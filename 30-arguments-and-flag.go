package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Penggunaan Arguments
	/*
		Arguments adalah data opsional yang disisipkan ketika eksekusi program. Sedangkan flag
		merupakan ekstensi dari argument. Dengan flag, penulisan argument menjadi lebih rapi dan
		terstruktur

		Data arguments bisa didapat lewat variabel os.Args (package
		osperlu di-import terlebih dahulu). Data tersebut tersimpan dalam bentuk array dengan pemisah adalah tanda spasi.
		Berikut merupakan contoh penggunaannya.
	*/
	var argsRaw = os.Args
	fmt.Printf("->%#v\n", argsRaw)

	var args = argsRaw[1:]
	fmt.Printf("->%#v\n", args)

	// Penggunaan Flag
	/*
		Flag memiliki kegunaan yang sama seperti arguments, yaitu untuk parameterize eksekusi
		program, dengan penulisan dalam bentuk key-value. Berikut merupakan contoh
		penerapannya.
	*/

	var name = flag.String("name", "anonymous", "type your name")
	var age = flag.Int64("age", 26, "type your age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)
}
