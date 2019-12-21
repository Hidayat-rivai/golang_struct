package main

import "fmt"

func main(){
	var x2 Pelajar
	
	
	x2.nama = "Jack"
	x2.kelas = 11
	x2.umur = 16
	
	fmt.Println(x2.nama)
	fmt.Println(x2.kelas)
	fmt.Println(x2.umur)
}

type Pelajar struct{
	nama string
	kelas int
	umur int
}