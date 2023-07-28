package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/stretchr/testify/require"
)

const URL = "http://127.0.0.1:8888/apigateway/studentservice/Register"

var httpCli = &http.Client{Timeout: 3 * time.Second}

type requestBody struct {
	Name string	`json:"name"`
	ID int64 `json:"id"`
}

func TestMain(t *testing.T){
	reqData := requestBody{
		Name: "tom",
		ID : 1,
	}

	req, err := generateReq(reqData, t)
	require.NoError(t, err)

	dump, err := httputil.DumpRequest(req, true)
    require.NoError(t, err)
    fmt.Println(string(dump))

	resp, err := httpCli.Do(req)
	require.NoError(t, err)
	require.Equal(t, consts.StatusOK, resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        fmt.Println("Error occurred while reading the response body:", err)
        return
    }

	fmt.Println(string(bodyBytes))
}

func BenchmarkMain(b *testing.B) {
	b.SetParallelism(8)
	b.RunParallel(func(pb *testing.PB){
		for pb.Next() {
			reqBody := `{"name":"tom", "id":1}`
			req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer([]byte(reqBody)))
			req.Header.Set("Content-Type", "application/json")
			require.NoError(b, err)

			client := &http.Client{Timeout: 3 * time.Second}
			resp, err := client.Do(req)
			require.NoError(b, err)
			require.Equal(b, consts.StatusOK, resp.StatusCode)
		}
	})
}

func generateReq(reqData requestBody, t *testing.T) (*http.Request, error){
	reqJSON, err := json.Marshal(reqData)
	require.NoError(t, err)

	reqBody := string(reqJSON)
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("Content-Type", "application/json")

	return req, err
}