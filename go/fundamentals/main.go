package main

import (
	"fmt"
	"strconv" // untuk mengkonversi int to string
)

func main() {
	var negara string = "Indonesia" // deklarasi variabel dengan tipe data
	var namaDepan = "Dicky"         // deklarasi variabel tanpa tipe data
	var umur = 99                   // deklarasi variabel dengan nilai string

	hobi := "membaca" // := adalah svd (Short Variable Declaration)  hanya bisa digunakan di dalam fungsi dan tujuannya untuk mempersingkat deklarasi variabel saja

	fmt.Println("Hello world : " + namaDepan)
	fmt.Println("Umur : " + strconv.Itoa(umur))
	fmt.Println("Hobi : " + hobi)
	fmt.Println("Negara : " + negara)
}
