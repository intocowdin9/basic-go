package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://depeloper.com:80/hello?name=john wick&age=27"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
	}

	fmt.Printf("url\t: %s\n", urlString)

	fmt.Printf("protocol\t: %s\n", u.Scheme)
	fmt.Printf("host\t: %s\n", u.Host)
	fmt.Printf("path\t: %s\n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name\t: %s, age: %s\n", name, age)
}
