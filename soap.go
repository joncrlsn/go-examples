package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	soapuiorg()
}

func soapSecretServer() {

}

func soapuiorg() {

	resp, err := http.Post("http://www.soapui.org", " text/xml;charset=UTF-8", strings.NewReader("the body"))

	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}
