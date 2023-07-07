package project

import "fmt"

func PrintPointer(){
	hello := "Pointer"

	fmt.Println(hello) // Pointer
	fmt.Println(&hello) // 0xc000010200

	// Alamat memori disimpan di variabel yang disebut pointer
	// \* dikenal sebagai operator pointer, yang berfungsi untuk mengambil nilai dari alamat memori/deklarasi variabel
	// & dikenal sebagai operator alamat, yang berfungsi untuk mengambil alamat memori dari sebuah variabel

	/* 
	Di dalam Go, pointer adalah sebuah variabel yang menyimpan alamat memori dari suatu nilai atau variabel. Pointer memungkinkan Anda untuk secara langsung mengakses dan memanipulasi nilai yang disimpan di alamat memori tersebut.

	Dalam Go, tanda "*" digunakan untuk mendeklarasikan variabel pointer dan juga untuk mengambil nilai yang disimpan di alamat memori yang ditunjuk oleh pointer (dereference operator). Operator "&" digunakan untuk mengambil alamat memori dari suatu variabel.
	*/
	number := 4
	var numberA *int = &number //tipe data pointer int = *int
	fmt.Println("number", numberA) // 0xc000010208
	fmt.Println("number", *numberA) // 4 dereference operator, mengambil nilai dari alamat memori
	fmt.Println("number", &numberA) //
	fmt.Println("number", *&numberA) // ini jadinya numberA, karena &* = *&, jadi *&numberA = numberA
	fmt.Println("number", &*numberA) // kenapa hasilnya malah alamat karena yang diassign adalah * which is pointer
	fmt.Println("number", *&number) // 4
	// fmt.Println("number", *number) // Kalau sudah seperti tidak bisa diakases lagi, karena sudah mengeluarkan isi dari alamat, error
}