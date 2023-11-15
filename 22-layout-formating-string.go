package main

import "fmt"

func main() {
	// Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 2 (biner).
	fmt.Printf("%b\n", 26)
	// Digunakan untuk memformat data numerik yang merupakan kode unicode, menjadi bentuk string karakter unicode-nya.
	fmt.Printf("%c\n", 1400)
	fmt.Printf("%c\n", 1235)
	// Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 10 (basis bilangan yang kita gunakan).
	fmt.Printf("%d\n", 26)
	// Digunakan untuk memformat data numerik desimal ke dalam bentuk notasi standar
	fmt.Printf("%e\n", 182.5)
	// Berfungsi untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan. Secara default lebar digit desimal adalah 6 digit.
	fmt.Printf("%f %.9f %.2f \n", 182.5, 182.5, 182.5)
	/*
		Berfungsi untuk memformat data numerik desimal, dengan lebar desimal bisa ditentukan.
		Lebar kapasitasnya sangat besar, pas digunakan untuk data yang jumlah digit desimalnya
		cukup banyak. Bisa dilihat pada kode berikut perbandingan antara %e, %f, dan %g.
	*/
	fmt.Printf("%e\n", 0.123123123123)
	fmt.Printf("%f\n", 0.123123123123)
	fmt.Printf("%g\n", 0.123123123123)
	/*
		Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 8
		(oktal).
	*/
	fmt.Printf("%o\n", 26)
	/*
		Digunakan untuk memformat data pointer, mengembalikan alamat pointer referensi variabel-
		nya. Alamat pointer dituliskan dalam bentuk numerik berbasis 16 dengan prefix 0x
	*/
	var name = "rafli"
	fmt.Printf("%p\n", &name)
	/*
		Digunakan untuk escape string. Meskipun string yang dipakai menggunakan literal \ akan
		tetap di-escape.
	*/
	fmt.Printf("%q\n", `"name \ height "`)
	// Digunakan untuk memformat data string.
	fmt.Printf("%s\n", "rafli")
	// Digunakan untuk memformat data boolean, menampilkan nilai bool-nya.
	var isGraduated bool = true
	fmt.Printf("%t\n", isGraduated)
	// Berfungsi untuk mengambil tipe variabel yang akan diformat.
	var hobbies = []string{}
	fmt.Printf("%T, %T, %T, %T, %T \n",
		12.20, isGraduated, "rafli", 2, hobbies,
	)
	/*
		Digunakan untuk memformat data apa saja (termasuk data bertipe
		interface{} ). Hasil kembaliannya adalah string nilai data aslinya.
		Jika data adalah objek cetakan struct, maka akan ditampilkan semua secara property
		berurutan.
	*/
	var data = student{
		name:        "rafli",
		height:      1.69,
		age:         26,
		isGraduated: true,
		hobbies:     []string{"jndakdakn"},
	}

	fmt.Printf("%v\n", data)
	/*
		 %+v Digunakan untuk memformat struct, mengembalikan nama tiap property dan nilainya
		berurutan sesuai dengan struktur struct.
	*/
	fmt.Printf("%+v\n", data)
	/*
		%#v Digunakan untuk memformat struct, mengembalikan nama dan nilai tiap property sesuai
		dengan struktur struct dan juga bagaimana objek tersebut dideklarasikan.
	*/
	fmt.Printf("%#v\n", data)
	/*
		Ketika menampilkan objek yang deklarasinya adalah menggunakan teknik anonymous
		struct, maka akan muncul juga struktur anonymous struct nya.
	*/
	var data2 = struct {
		name        string
		isGraduated bool
	}{
		name:        "wick",
		isGraduated: true,
	}

	fmt.Printf("%#v\n", data2)

	/*
		Digunakan untuk memformat data numerik, menjadi bentuk string numerik berbasis 16
		(heksadesimal).
	*/
	fmt.Printf("%x\n", 26)
	// Jika digunakan pada tipe data string, maka akan mengembalikan kode heksadesimal tiap karakter.

	var d = data.name
	fmt.Printf("%x%x%x%x\n", d[0], d[1], d[2], d[3])
	/*
		%x dan %X memiliki fungsi yang sama. Perbedaannya adalah
		%X akan mengembalikan string dalam bentuk uppercase atau huruf kapital.
	*/
	fmt.Printf("%X\n", d)
	// Cara untuk menulis karakter % pada string format.
	fmt.Printf("%%\n")
}

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}
