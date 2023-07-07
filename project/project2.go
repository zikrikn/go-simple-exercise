package project

import (
	"fmt"
	"strings"
)

func Project2() {
	var text = "Hello, Project 2! 1"
	fmt.Println(text)

	text2 := "Hello, Project 2! 2"
	fmt.Println(text2)

	var postiveNumber uint8 = 89
	fmt.Println(postiveNumber)

	var negativeNumber string = "Negative Number"
	fmt.Println(negativeNumber)

	var decimalNumber = 2.62
	fmt.Printf("Bilang desimal: %.3f\n", decimalNumber) // %f = float

	var exist bool = true
	fmt.Printf("exist? %t \n", exist) // %t = boolean

	var message = `Nama saya "John Wick"`
	fmt.Printf("message: %s \n", message) // %s = string

	// Teknik Casting (Konversi tipe data)
	var b int32 = int32(decimalNumber)
	fmt.Printf("b: %d \n", b) // %d = integer

	// Package Strings
	var index1 = strings.Index("ethan hunt", "ha") // mencari index dari string tapi dari katanya, dihitung dari kata ethan karena ethan e = 0, t = 1, h = 2, a = 3, n = 4, yang dikembalikan adalah yang paling kecil.
	fmt.Println(index1) // 2
}