package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// TODO serde_json rename_all alternative?
type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	// Data interface{} `json:"data"`
}

// type Rsp struct {
//     data struct {
//         user_id int
//     }
// }

/*
## init
go mod init go_nightly_integration_test

## how to run(multi methods)
1. go install go_nightly_integration_test && go_nightly_integration_test
2. go install . && go_nightly_integration_test
3. go build main.go && ./main
4. go run main.go
5. go run .
*/
func main() {
	base_url := "http://api.example.com"
	user_svc := "/0/api/v1/user"
	client := &http.Client{}

	// json account+password
	account := "w@go.dev"
	password_before_hash := "1234qwer"
	h := md5.New()
	io.WriteString(h, password_before_hash)
	password := fmt.Sprintf("%x", h.Sum(nil))
	// printf can format string
	req_body := make(map[string]string)
	req_body["account"] = account
	req_body["password"] = password
	fmt.Printf("%v+\n", req_body)

	req_body_str, err := json.Marshal(req_body)
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("%s%s/account/register", base_url, user_svc)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(req_body_str))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp %d %s\n%s\n\n", resp.StatusCode, url, resp_body)
	rsp := Resp{}
	if err := json.Unmarshal(resp_body, &rsp); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", rsp)
}
