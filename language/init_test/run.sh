echo ">>>>>>>>>>运行 single_init 测试<<<<<<<<<<"
go run ./single_init/single_init.go

echo ">>>>>>>>>>运行 multi_file_init 测试<<<<<<<<<<"
go run ./multi_file_init/*.go

echo ">>>>>>>>>>运行 multi_package_init 测试<<<<<<<<<<"
cd multi_package_init && go run . && cd ..