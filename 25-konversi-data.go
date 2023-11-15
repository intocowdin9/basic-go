package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	/*
		strconv.Atoi() Fungsi ini digunakan untuk konversi data dari tipe
		nilai balik, yaitu hasil konversi dan nil error string ke int. Mengembalikan 2 buah
		(jika konversi sukses, maka error akan berisi).
	*/

	var str = "124"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num, reflect.TypeOf(num))
	}

	/*
		strconv.Itoa() Merupakan kebalikan dari strconv.Atoi, berguna untuk konversi
		int ke string.
	*/
	var num2 = 124
	var str2 = strconv.Itoa(num2)
	fmt.Println(str2, reflect.TypeOf(str2))

	/*
		strconv.ParseInt() Digunakan untuk konversi string berbentuk numerik dengan basis tertentu ke tipe numerik
		non-desimal dengan lebar data bisa ditentukan. Pada contoh berikut, string
		tipe data int64
	*/

	var str3 = "124"
	var num3, err2 = strconv.ParseInt(str3, 10, 64)

	if err2 == nil {
		fmt.Println(num3, reflect.TypeOf(num3))
	}

	/*
		strconv.FormatInt() Berguna untuk konversi data numerik
		int64 ke string dengan basis numerik bisa ditentukan sendiri.
	*/

	var num4 = int64(24)
	var str4 = strconv.FormatInt(num4, 8)

	fmt.Println(str4, reflect.TypeOf(str4))

	/*
		Casting string <-> byte
		String sebenarnya adalah slice/array byte. Di Golang sebuah karakter biasa (bukan
		unicode) direpresentasikan oleh sebuah elemen slice byte. Nilai slice tersebut adalah data
		int yang (default-nya) ber-basis desimal, yang merupakan kode ASCII dari karakter biasa
		tersebut.Cara mendapatkan slice byte dari sebuah data string adalah dengan meng-casting-nya ke
		tipe []byte. Tiap elemen byte isinya adalah data numerik dengan basis desimal.
	*/
	var text1 = "halo"
	var b = []byte(text1)
	fmt.Printf("%d %d %d %d \n\n", b[0], b[1], b[2], b[3])
	// 104 97 108 111

	/*
		Konversi Data interface{} Menggunakan Teknik Type Assertions
		Type assertions merupakan teknik casting data interface{}
		ke segala jenis tipe (dengan syarat data tersebut memang bisa di-casting).
		Berikut merupakan contoh penerapannya. Disiapkan variabel map[string]interface{}
		data bertipe dengan value berbeda beda tipe datanya.
	*/

	var data = map[string]interface{}{
		"name":    "john wick",
		"grade":   2,
		"height":  1.69,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	fmt.Println(data["name"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	/*
		Statement string data["nama"].(string) maksudnya adalah, nilai
		data["nama"] dicasting sebagai string.

		Tipe asli data pada variabel interface{} ke tipe type interface{} bisa diketahui dengan cara meng-casting.
		Namun casting ini hanya bisa dilakukan pada switch.
	*/
	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}
}
