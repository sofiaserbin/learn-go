package main

import "fmt"
import "unicode/utf8"

func main(){
	var intNum int = 32767
	fmt.Println(intNum)

	var floatNum float64 = 12345.78
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string = "Hello \nWorld"
	fmt.Println(myString)
	var longString string = `bhnjkkkkkkkk
	bnjmkm long string
	`

	fmt.Println(longString)

	//if the characters are outside the ASCII scope, the length will seem random
	//it will actually count the number of bytes
	fmt.Println(len(longString))
	fmt.Println(utf8.RuneCountInString("Å"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var boolean bool = false
	fmt.Println(boolean)

	myVar := "text"
	fmt.Println(myVar)
}