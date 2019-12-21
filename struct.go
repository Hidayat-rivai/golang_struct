package main

import "fmt"

func main(){
	var x2 = Pelajar{"Jack",11,16}
	fmt.Println(x2)
}

type Pelajar struct{
	nama string
	kelas int
	umur int
}