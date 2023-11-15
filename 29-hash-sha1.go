package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {
	var text = "this is secret"
	var sha = sha1.New()
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println("original :", text)
	fmt.Println("hashed   :", encryptedString)

	// call do hash using salt
	mainDoHashUsingSalt()
}

// metode Salting Pada Hash
/*
	Salt dalam konteks kriptografi adalah data acak yang digabungkan pada data asli sebelum
	proses hash dilakukan.
	Hash merupakan enkripsi satu arah dengan lebar data yang sudah pasti, menjadikan sangat
	mungkin sekali kalau hasil hash untuk beberapa data adalah sama. Disinilah kegunaan salt,
	teknik ini berguna untuk mencegah serangan menggunakan metode pencocokan data-data
	yang hasil hash-nya adalah sama (dictionary attack).
	Langsung saja kita praktekan. Pertama import package yang dibutuhkan. Lalu buat fungsi
	untuk hash menggunakan salt dari waktu sekarang.
*/

func doHashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt: %s ", text, salt)
	fmt.Println(saltedText)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func mainDoHashUsingSalt() {
	var text = "this is secret"
	fmt.Printf("original : %s\n\n", text)

	var hashed1, salt1 = doHashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed1)

	var hashed2, salt2 = doHashUsingSalt(text)
	fmt.Printf("hashed 2 : %s\n\n", hashed2)

	var hashed3, salt3 = doHashUsingSalt(text)
	fmt.Printf("hashed 3 : %s\n\n", hashed3)

	_, _, _ = salt1, salt2, salt3
}
