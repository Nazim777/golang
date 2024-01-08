package main

import (
	"fmt";
	// "strings"
)

func main() {
    // fmt.Println("hello world")

	
	// variable
	// 1st
	// var email string
	// var phone int
	// email = "test@gmail.com"
	// phone = 123456789
	// fmt.Println(email)
	// fmt.Println(phone)


	// 2nd
	// name := "rahim"
	// age := 20
	// fmt.Println(name)
	// fmt.Println(age)

	// loops
	// for i:=0; i<=5; i++{
	// 	fmt.Println(i)
	// }


	// numbers :=[]int{1,2,3,4}

	// for index, value:=range numbers{
	// 	fmt.Printf("index: %d, value: %d\n",index, value)

	// }

	// for _, value:= range numbers{
	// 	fmt.Printf("value %d\n",value)

	// }

	// for i := 0; i < 6; i++ {
	// 	if i==4{
	// 		break
	// 	}
	// 	fmt.Println(i)
		
	// }
	
	// for i := 0; i < 6; i++ {
	// 	if i==4{
	// 		continue
	// 	}
	// 	fmt.Println(i)
		
	// }


	
	// conditions
	// age:=25
	// if age<20{
	// 	fmt.Println("you can not drink")
	// }else if age>20 && age<22 {
	// 	fmt.Println("you can drink but can not drive")
		
	// }else{
	// 	fmt.Println("you can drink and drive")
	// }


	// function
	// value := sum(4,6)
	// fmt.Println(value)
	// sum, multiply :=calculate(2,5)
	// fmt.Printf("sum :%d, multiply: %d",sum,multiply)
	
	// array

	// var arrayName [size]Type
	// var numbers  [5]int
	// numbers:= [5]int{1,2,3,4,5}
	// arrayLength := len(numbers)
	// fmt.Println(arrayLength)

	// for i:=0; i<len(numbers); i++{
	// 	fmt.Println(numbers[i])
	// }

	// slice
	// var sliceName []Type
	// numbers:=[]int{1,2,3,4}
	// for i := 0; i <len(numbers); i++ {
	// 	fmt.Println(numbers[i])
		
	// }

	// stuct

	// type StructName struct {
	// 	Field1 Type1
	// 	Field2 Type2
	// 	// ... additional fields
	// }


	// type Person struct{
	// 	name string
	// 	email string
	// 	phone int
	// 	isAdmin bool
	// }

	// person:=Person{
	// 	name: "Rahim Islam",
	// 	email: "rahim@gmail.com",
	// 	phone: 123456789,
	// 	isAdmin: true,
	// }

	// fmt.Println(person.name)


	// take user input
	// var name string
	// var age int
	// fmt.Println("Enter your name!")
	// fmt.Scan(&name)
	// fmt.Println("Enter your age!")
	// fmt.Scan(&age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)

	// string
	// text := "Go Programming"
	// substring1 := "Go"
	// substring2 := "Golang"
  
	// check if Go is present in Go Programming
	// result := strings.Contains(text, substring1)
	// fmt.Println(result)
  
	// check if Golang is present in Go Programming
	// result = strings.Contains(text, substring2)
	// fmt.Println(result)

	// text1 := "go is fun to learn"
	// convert to uppercase
	// text1 = strings.ToUpper(text1)
	// fmt.Println(text1)

	// text2 := "I LOVE GOLANG"
	// convert to lowercase
	// text2 = strings.ToLower(text2)
	// fmt.Println(text2)



}


// func sum(a int, b int) int{
// 	return a+b
// }

// func calculate (a int,b int)(int, int){
// 	sum:=a+b
// 	mulitply := a*b
// 	return sum, mulitply

// }