package code

// Excelize 读写 Excel 的 demo

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/xuri/excelize/v2"
)

const (
	Endpoint        = "files.ozline.icu"
	AccesskeyID     = ""
	AccesskeySecret = ""
	BucketName      = ""
	MainDirectory   = "west2-online"
)

type BufferWriters struct {
	Buffer *bytes.Buffer
}

func (bw *BufferWriters) Write(p []byte) (n int, err error) {
	return bw.Buffer.Write(p)
}

func TestExcelizeWrite(t *testing.T) {
	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	sw, err := f.NewStreamWriter("Sheet1")

	if err != nil {
		panic(err)
	}

	for rowID := 1; rowID <= 1024; rowID++ {
		row := make([]interface{}, 1024)
		for colID := 0; colID < 50; colID++ {
			row[colID] = rand.Intn(68400)
		}

		cell, err := excelize.CoordinatesToCellName(1, rowID)

		if err != nil {
			panic(err)
		}

		if err := sw.SetRow(cell, row); err != nil {
			panic(err)
		}
	}

	// 结束流写
	if err := sw.Flush(); err != nil {
		panic(err)
	}

	// 保存

	// sum, err := f.WriteTo(bufferWiter)
	buffer, err := f.WriteToBuffer()

	if err != nil {
		panic(err)
	}

	// connect to oss
	client, err := oss.New(Endpoint, AccesskeyID, AccesskeySecret, oss.UseCname(true))

	if err != nil {
		panic(err)
	}

	bucket, err := client.Bucket(BucketName)

	if err != nil {
		panic(err)
	}

	xlsxFilename := MainDirectory + "/test.xlsx"

	if err := bucket.PutObject(xlsxFilename, buffer); err != nil {
		panic(err)
	}

	fmt.Printf("https://%s/%s/%s\n", Endpoint, MainDirectory, xlsxFilename)

}
