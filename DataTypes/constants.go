package main

import (
	"fmt"
)

func main() {
	fmt.Println("constant")
	fmt.Println("Typed Constants .....")
	const indentifier string = "hello"
	fmt.Println("const indentifier string = hello")
	fmt.Println("identifier : ", indentifier)
	fmt.Println("Compilations Constants .....")
	const sample1 = 2 / 3
	fmt.Println("const sample1 = 2/3 ")
	fmt.Println("sample1 : ", sample1)
	fmt.Println("Multiple Constants .....")
	const EGG, TWO, RICE = `protien`, 2, `carb`
	fmt.Println("const EGG, TWO, RICE = `protien` , 2 , `carb`")
	fmt.Println("EGG :", EGG)
	fmt.Println("EGG :", TWO)
	fmt.Println("EGG :", RICE)
	fmt.Println("Multiple Constants with Type.....")
	const ONE, THREE, Four int = 1, 3, 4
	fmt.Println("const ONE, THREE, Four int = 1 , 3 , 4")
	fmt.Println("ONE :", ONE)
	fmt.Println("THREE :", THREE)
	fmt.Println("FOUR :", Four)
	fmt.Println("Enumerations in Constants.....")
	const (
		OTHERS = iota
		MALE   = iota
		FEMALE = iota
	)
	fmt.Println(`	
	   const (
		OTHERS = iota
		 MALE = iota
		 FEMALE = iota
	 )
	 `)
	fmt.Println("The first use of iota gives 0. Whenever iota is used again on a new line, its value is incremented by 1")
	fmt.Println("OTHERS : ", OTHERS)
	fmt.Println("MALE : ", MALE)
	fmt.Println("FEMALE : ", FEMALE)

	fmt.Println("Typed Enumerations in Constants.....")
	type Gender int
	const (
		NEW Gender = iota
		UNKNOWN
	)
	fmt.Println(`
	 type Gender int 
	 const (
		NEW Gender = iota
		UNKNOWN
	 )
	 `)
	fmt.Println("New", NEW)
	fmt.Println("UNKNOWN", UNKNOWN)

}
