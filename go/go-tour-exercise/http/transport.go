package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	values := url.Values{}
	values.Set("id", "you")
	values.Set("ip", "me")
	values.Set("stage", "1")
	values.Set("err", "err")
	resp, err := client.PostForm("http://127.0.0.1:8080/task/status/update", values)
	// url.Values{"taskid": {"568b6b8a464f18c7f68ff0be"}}
	fmt.Println(resp.Status, "|||", resp.StatusCode == 200)
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
	fmt.Println(err)
}
