package b

import (
	"fmt"

	_ "github.com/ozline/go-codes/language/init_test/multi_package_init/b/c"
)

func init() {
	fmt.Println("run init 1 from b.go")
}

func init() {
	fmt.Println("run init 2 from b.go")
}
