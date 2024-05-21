package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"

)

func Test_Http_Get(t *testing.T){
	
	url := fmt.Sprintf("https://www.cloudflare.com/ips-v4")

	response, err := http.Get(url)

	if err != nil {
		t.Fatalf(`ERR = %q`, err)
		return ;
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	str := string(body)

	t.Log(str)

}