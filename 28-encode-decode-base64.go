package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "john wick"

	// penerapan fungsi EndodeToString() & DecodeString()
	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("decoded:", decodedString)

	// penerapan fungsi Encode() & Decode()
	var data2 = "bruce lee"
	var encoded2 = make([]byte, base64.StdEncoding.EncodedLen(len(data2)))
	base64.StdEncoding.Encode(encoded2, []byte(data))
	var encodedString2 = string(encoded2)
	fmt.Println(encodedString2)

	var decoded2 = make([]byte, base64.StdEncoding.DecodedLen(len(encoded2)))
	var _, err = base64.StdEncoding.Decode(decoded2, encoded2)
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodedString2 = string(decoded2)
	fmt.Println(decodedString2)

	// Encode & Decode Data URL
	var data3 = "http://depeloper.com/"

	var encodedString3 = base64.URLEncoding.EncodeToString([]byte(data3))
	fmt.Println(encodedString3)
}
