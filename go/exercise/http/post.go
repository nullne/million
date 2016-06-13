package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
	// "net/url"
)
func main() {
	url := "http://127.0.0.1:8080/v1/formats/test"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"description": "MIS-BJ-5-3g6", "mip":"223.202.46.166"}
`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
