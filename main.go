package main

import "fmt"

func tambah(a int, b int) int {
	return a + b
}

func main() {
	hasil := tambah(10, 20)
	fmt.Println("Hasil:", hasil)
}
