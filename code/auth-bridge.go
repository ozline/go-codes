package code

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAuthBridgeAnswer(url string, token string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s?token=%s", url, token), nil)
	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}

type DockingResponse struct {
	MsgCode         int64  `json:"msgCode"`
	ErrMsg          string `json:"errMsg"`
	ReceiptDateTime string `json:"receiptDateTime"`
	ReturnDateTime  string `json:"returnDateTime"`
	Item            struct {
		Info struct {
			TimeStamp string `json:"timeStamp"`
			UserId    string `json:"userId"`
			Username  string `json:"username"`
		}
	}
	Success bool `json:"success"`
}

func RunAuthTest() {
	res, err := GetAuthBridgeAnswer("http://skydog.ltd:11111/api/fjyz/info", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0aW1lU3RhbXAiOjE3MDEwMTc5NDQ2MzIsInVzZXJJZCI6InRlc3RhZG1pbiIsInVzZXJuYW1lIjoi5pWZ6IGM5bel5rWL6K-V6LSm5Y-3In0.cXtlG6i3iDN380-dlozdRBn9s5KgK1qC7LowRI13g7E#/")
	if err != nil {
		panic(err)
	}

	var resp DockingResponse

	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Printf("%+v\n", resp)
}
