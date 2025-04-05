package main

import "fmt"
import "errors"

func main(){
	var printValue string = "Hello"
	printMe(printValue)
	var result, remainder, err = intDivistion(11, 0)
	if err!=nil{
		fmt.Printf(err.Error())
	}else if remainder == 0{
		fmt.Printf("The result of the integer division is %v", result)
	}else{
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
}
	}
	

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivistion(numerator int, denominator int) (int, int, error){
	var err error
	if denominator==0{
		err = errors.New("Cannot Divide by Zero")
		return 0,0,err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}