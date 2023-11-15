package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	// Decode JSON Ke Variabel Objek Cetakan Struct
	/*
	   Data json tipenya adalah []byte casting). Dengan menggunakan, bisa didapat dari file ataupun string (dengan hasil json.Unmarshal
	   bentuk objek, entah itu dalam bentuk struct, data tersebut bisa dikonversi menjadi
	   map[string]interface{} ataupun variabel objek hasil. Program berikut ini adalah contoh cara decoding json ke bentuk objek. Pertama import
	   package yang dibutuhkan, dan siapkan struct User.
	*/
	var jsonString = `{"Name": "john wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)

	// Decode JSON ke map[string]interface{} & interface{}
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])

	// Decode Array JSON ke Array Objek
	var jsonString3 = `[
		{"Name": "john wick", "Age": 27},
		{"Name": "ethan hunt", "Age": 32}
	]`

	var data3 []*User

	var err3 = json.Unmarshal([]byte(jsonString3), &data3)
	if err3 != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("user 1:", data3[0].FullName)
	fmt.Println("user 2:", data3[1].FullName)

	// Encode Objek ke JSON
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData4, err4 = json.Marshal(object)
	if err4 != nil {
		fmt.Println(err4.Error())
	}

	var jsonString4 = string(jsonData4)
	fmt.Println(jsonString4)
}
