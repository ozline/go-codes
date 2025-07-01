module github.com/ozline/go-codes

go 1.24.4 // 这个版本直接人工替换就可以了

// algorithm 使用
require (
	github.com/bits-and-blooms/bitset v1.22.0 // alogirhtm
	github.com/twmb/murmur3 v1.1.8 // algorithm
)

// benchmark 使用
require (
	github.com/go-sql-driver/mysql v1.9.2 // benchmark
	go.mongodb.org/mongo-driver v1.17.3 // benchmark
)

// implement 使用
require (
	github.com/aliyun/aliyun-oss-go-sdk v3.0.1+incompatible // implement
	github.com/dchest/captcha v1.0.0 // implement
	github.com/evanphx/json-patch v0.5.2 // implement
	github.com/fsnotify/fsnotify v1.6.0 // implement
	github.com/xuri/excelize/v2 v2.8.0 // implement
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.3 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/xuri/efp v0.0.0-20230802181842-ad255f2331ca // indirect
	github.com/xuri/nfp v0.0.0-20230819163627-dc951e3ffe1a // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/time v0.6.0 // indirect
)
