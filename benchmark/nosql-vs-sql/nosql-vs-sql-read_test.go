package benchmark

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	numDocumentsForRead = 1000000
)

func PreInsertDocumentsMongoDB() {
	bulk := make([]any, 0, 1000)
	for i := 0; i < numDocumentsForRead; i++ {
		doc := bson.M{
			"name":  fmt.Sprintf("Name_%d", i),
			"value": rand.Intn(10000000),
		}
		bulk = append(bulk, doc)

		if len(bulk) == 1000 {
			_, err := collection.InsertMany(context.TODO(), bulk)
			if err != nil {
				panic(err)
			}
			bulk = bulk[:0]
		}
	}

	if len(bulk) > 0 {
		_, err := collection.InsertMany(context.TODO(), bulk)
		if err != nil {
			panic(err)
		}
	}
}

func PreInsertDocumentsMySQL() {
	tx, err := mysqlDB.Begin()
	if err != nil {
		panic(fmt.Sprintf("Failed to start transaction: %v", err))
	}

	// 每次批量插入 1000 条
	batchSize := 1000
	values := make([]any, 0, batchSize*2) // 每条记录有两个值 (name, value)
	queryPrefix := fmt.Sprintf("INSERT INTO %s (name, value) VALUES ", collectionName)
	query := queryPrefix

	for i := 0; i < numDocumentsForRead; i++ {
		query += "(?, ?),"
		values = append(values, fmt.Sprintf("Name_%d", i), rand.Intn(10000000))

		// 每 1000 条执行一次批量插入
		if (i+1)%batchSize == 0 {
			query = query[:len(query)-1] // 去掉最后的逗号
			_, err := tx.Exec(query, values...)
			if err != nil {
				panic(fmt.Sprintf("Failed to execute batch insert: %v", err))
			}

			// 重置 query 和 values
			query = queryPrefix
			values = values[:0]
		}
	}

	// 插入剩余的记录
	if len(values) > 0 {
		query = query[:len(query)-1] // 去掉最后的逗号
		_, err := tx.Exec(query, values...)
		if err != nil {
			panic(fmt.Sprintf("Failed to execute final batch insert: %v", err))
		}
	}

	if err := tx.Commit(); err != nil {
		panic(fmt.Sprintf("Failed to commit transaction: %v", err))
	}
}

func BenchmarkReadMongoDB(b *testing.B) {
	// 预先插入数据
	PreInsertDocumentsMongoDB()

	for n := 0; n < b.N; n++ {
		start := time.Now()

		cursor, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {
			b.Fatalf("Failed to find documents: %v", err)
		}
		defer cursor.Close(context.TODO())

		count := 0
		for cursor.Next(context.TODO()) {
			count++
		}

		if err := cursor.Err(); err != nil {
			b.Fatalf("Cursor error: %v", err)
		}

		duration := time.Since(start)
		b.ReportMetric(float64(count)/duration.Seconds(), "docs/sec")
	}
}

func BenchmarkReadMySQL(b *testing.B) {
	// 预先插入数据
	PreInsertDocumentsMySQL()

	for n := 0; n < b.N; n++ {
		start := time.Now()

		rows, err := mysqlDB.Query(fmt.Sprintf("SELECT id, name, value FROM %s", collectionName))
		if err != nil {
			b.Fatalf("Failed to query rows: %v", err)
		}
		defer rows.Close()

		count := 0
		for rows.Next() {
			var id int
			var name string
			var value int

			if err := rows.Scan(&id, &name, &value); err != nil {
				b.Fatalf("Failed to scan row: %v", err)
			}
			count++
		}

		duration := time.Since(start)
		b.ReportMetric(float64(count)/duration.Seconds(), "docs/sec")
	}
}
