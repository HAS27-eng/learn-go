package main

import "fmt"

func main() {
	var input string
	fmt.Println("Masukkan Nama Anda: ")
	i, r := fmt.Scanln(&input)
	fmt.Printf("Hasil Scanln: jumlah input = %v, error = %v", i, r)
	fmt.Printf("input: %v", input)
}
