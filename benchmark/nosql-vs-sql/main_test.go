package benchmark

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURI       = "mongodb://root:example@localhost:27017"
	mysqlDSN       = "root:root@tcp(127.0.0.1:3306)/ozline?charset=utf8mb4&parseTime=True&loc=Local"
	databaseName   = "performance_test"
	collectionName = "benchmark" // 同样也是 TableName
	numDocuments   = 1000000
)

var (
	collection *mongo.Collection
	mysqlDB    *sql.DB
)

func TestMain(m *testing.M) {
	// MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Connected to MongoDB")

	collection = client.Database(databaseName).Collection(collectionName)
	if err := collection.Drop(context.TODO()); err != nil {
		panic(err)
	}

	fmt.Println("Dropped existing collection")

	// MySQL
	mysqlDB, err = sql.Open("mysql", mysqlDSN)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MySQL")

	_, err = mysqlDB.Exec(fmt.Sprintf("TRUNCATE TABLE %s", collectionName))
	if err != nil {
		panic(err)
	}
	fmt.Println("Truncated MySQL table, start for benchmarking")

	exitCode := m.Run()
	os.Exit(exitCode)
}
