package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings.Contains()
	var isExists = strings.Contains("john wick", "wick")
	fmt.Println(isExists) // Variabel isExists "john wick" akan bernilai true, karena string "wick" merupakan bagian dari "john wick".

	// strings.HasPrefix()
	/*
		Digunakan untuk deteksi apakah sebuah string (parameter pertama) diawali string tertentu
		(parameter kedua).
	*/
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1)

	var isPrefix2 = strings.HasPrefix("john wick", "wick")
	fmt.Println(isPrefix2)

	// strings.HasSuffix()
	/*
		Digunakan untuk deteksi apakah sebuah string (parameter pertama) diakhiri string tertentu
		(parameter kedua).
	*/
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1)
	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2)

	// strings.Count()
	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany)

	// strings.Index()
	/*
		String "ha"berada pada posisi ke2 dalam string "ethan hunt"
		(indeks dimulai dari 0). Jika diketemukan dua substring, maka yang diambil adalah yang pertama, contoh:
	*/

	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2)

	// strings.Replace()
	var text = "apple"
	var find = "p"
	var replaceWith = "q"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1)

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2)

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3)

	// strings.Repeat()
	var str = strings.Repeat("ha", 3)
	fmt.Println(str)

	// strings.Split()
	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1)

	var string2 = strings.Split("batman", "")
	fmt.Println(string2)

	// strings.Join()
	var data = []string{"apple", "manggo", "kurma"}
	var str2 = strings.Join(data, "-")
	fmt.Println(str2)

	// strings.ToLower()
	var str3 = strings.ToLower("AlaY")
	fmt.Println(str3)

	var str4 = strings.ToUpper("eat!")
	fmt.Println(str4)
}
