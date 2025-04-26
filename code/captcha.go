package code

/*
	关于 dchest/captcha 的使用 demo
*/

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/dchest/captcha"
)

var code string
var buf bytes.Buffer
var buf2 bytes.Buffer

func RunCaptcha() {

	captcha.SetCustomStore(captcha.NewMemoryStore(100, 10*time.Minute))

	id := captcha.New()
	// digits := captcha.RandomDigits(6)
	// img := captcha.NewImage(id, digits, captcha.StdWidth, captcha.StdHeight)
	// img.WriteTo(file)
	if err := captcha.WriteImage(&buf, id, captcha.StdWidth, captcha.StdHeight); err != nil {
		panic(err)
	}

	captchaBytes := buf.Bytes()
	filePath := "img.png"
	err := os.WriteFile(filePath, captchaBytes, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("id: %v\n", id)

	// verify
	fmt.Printf("input code:")
	fmt.Scanf("%v", &code)

	fmt.Printf("your input: [%v]\n", code)

	fmt.Printf("Check: %v\n", captcha.VerifyString(id, code))

	// test reload
	captcha.Reload(id)
	if err := captcha.WriteImage(&buf2, id, captcha.StdWidth, captcha.StdHeight); err != nil {
		panic(err)
	}

	captchaBytes = buf2.Bytes()
	filePath = "img2.png"
	if err := os.WriteFile(filePath, captchaBytes, 0644); err != nil {
		panic(err)
	}

	fmt.Printf("input code:")
	fmt.Scanf("%v", &code)

	fmt.Printf("your input: [%v]\n", code)

	fmt.Printf("Check: %v\n", captcha.VerifyString(id, code))
}
