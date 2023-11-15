package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Penggunaan Exec
	/*
		Golang menyediakan package exec	berisikan banyak fungsi untuk keperluan eksekusi
		perintah CLI.
		Cara untuk eksekusi command cukup mudah, yaitu dengan menuliskan command dalam
		bentuk string, diikuti arguments-nya (jika ada) sebagai parameter variadic pada fungsi
		exec.Command()
		. Contoh:
	*/

	var output1, _ = exec.Command("ls").Output()
	fmt.Printf(" -> ls\n%s\n", string(output1))

	var output2, _ = exec.Command("pwd").Output()
	fmt.Printf(" -> pwd\n%s\n", string(output2))

	var output3, _ = exec.Command("git", "config", "user.name").Output()
	fmt.Printf(" -> git config user.name\n%s\n", string(output3))
}
