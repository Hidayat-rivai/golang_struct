package main

import "fmt"

func main(){
	var x1 Pelajar
	
	x1.nama = "Johan"
	x1.kelas = 12
	x1.umur = 17
	
	fmt.Println(x1)
	
}

type Pelajar struct{
	nama string
	kelas int
	umur int
}