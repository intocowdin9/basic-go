package main

import (
	"fmt"
	"io"
	"os"
)

// Membuat File Baru
/*
	Pembuatan file di Golang sangatlah mudah, cukup dengan memanggil fungsi
	os.Create() lalu memasukkan path file ingin dibuat sebagai parameter fungsi tersebut.
	Jika file yang akan dibuat sudah ada, maka akan ditimpa. Bisa memanfaatkan
	os.IsNotExist() untuk mendeteksi apakah file sudah dibuat atau belum.
	Berikut merupakan contoh pembuatan file.
*/

var path = "/home/return/go/src/self-try/file-test.txt"

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func createFile() {
	// deteksi apakah file sudah ada
	var _, err = os.Stat(path)

	// buat file baru jika belum ada
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkError(err)
		defer file.Close()
	}
}

func writeFile() {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString("halo\n")
	checkError(err)
	_, err = file.WriteString("mari belajar golang\n")
	checkError(err)

	// simpan perubahan
	err = file.Sync()
	checkError(err)
}

func readFile() {
	// buka file
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	checkError(err)
	defer file.Close()

	// baca file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			checkError(err)
		}
		if n == 0 {
			break
		}
	}

	fmt.Println(string(text))
	checkError(err)
}

func deleteFile() {
	var err = os.Remove(path)
	checkError(err)
}

func main() {
	createFile()
	writeFile()
	readFile()
	// deleteFile()
}
