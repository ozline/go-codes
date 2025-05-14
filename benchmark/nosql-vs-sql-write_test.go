package benchmark

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func BenchmarkWriteMongoDB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		start := time.Now()

		bulk := make([]any, 0, 1000)
		for i := 0; i < numDocuments; i++ {
			doc := bson.M{
				"name":  fmt.Sprintf("Name_%d", i),
				"value": rand.Intn(10000000),
			}
			bulk = append(bulk, doc)

			// 每 1000 条执行一次批量插入
			if len(bulk) == 1000 {
				_, err := collection.InsertMany(context.TODO(), bulk)
				if err != nil {
					b.Fatalf("Failed to insert documents: %v", err)
				}
				bulk = bulk[:0]
			}
		}

		// 插入剩余的文档
		if len(bulk) > 0 {
			_, err := collection.InsertMany(context.TODO(), bulk)
			if err != nil {
				b.Fatalf("Failed to insert remaining documents: %v", err)
			}
		}

		duration := time.Since(start)
		b.ReportMetric(float64(numDocuments)/duration.Seconds(), "docs/sec")
	}
}

func BenchmarkWriteMySQL(b *testing.B) {
	for n := 0; n < b.N; n++ {
		start := time.Now()

		tx, err := mysqlDB.Begin()
		if err != nil {
			b.Fatalf("Failed to start transaction: %v", err)
		}

		// 每次批量插入 1000 条
		batchSize := 1000
		values := make([]any, 0, batchSize*2) // 每条记录有两个值 (name, value)
		queryPrefix := fmt.Sprintf("INSERT INTO %s (name, value) VALUES ", collectionName)
		query := queryPrefix

		for i := 0; i < numDocuments; i++ {
			query += "(?, ?),"
			values = append(values, fmt.Sprintf("Name_%d", i), rand.Intn(10000000))

			// 每 1000 条执行一次批量插入
			if (i+1)%batchSize == 0 {
				query = query[:len(query)-1] // 去掉最后的逗号
				_, err := tx.Exec(query, values...)
				if err != nil {
					b.Fatalf("Failed to execute batch insert: %v", err)
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
				b.Fatalf("Failed to execute final batch insert: %v", err)
			}
		}

		if err := tx.Commit(); err != nil {
			b.Fatalf("Failed to commit transaction: %v", err)
		}

		duration := time.Since(start)
		b.ReportMetric(float64(numDocuments)/duration.Seconds(), "docs/sec")
	}
}
