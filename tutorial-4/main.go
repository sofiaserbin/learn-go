package main

import "fmt"

func main(){
	var intArr [3]int32 = [...]int32{1,2,3}
	intArr[1] = 132
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3]) //elements 1 and 2

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])


	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap["Jason"]
	delete(myMap2, "Adam")
	if ok{
		fmt.Printf("The age is %v", age)
	}else{
		fmt.Println("Invalid name")
	}

	for name:= range myMap{ //no order preserved
		fmt.Printf("Name: %v", name)
	}

	for i, v := range intArr{
		fmt.Printf("\nI is %v, value is %v", i, v)
	}

}