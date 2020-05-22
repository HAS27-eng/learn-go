package main

import "fmt"

func main() {
	var input *int
	fmt.Println("Masukkan Nama Anda: ")
	i, r := fmt.Scanln(input)
	fmt.Printf("Hasil Scanln: length = %v, error = %v", i, r)
	fmt.Printf("input: %v", input)
}
