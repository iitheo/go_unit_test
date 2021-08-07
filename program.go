package main

import "fmt"

func main(){
	output := SumNumbers(5,9,12)
	fmt.Println("Sum of numbers:",output)
}

func SumNumbers(input ...int) (result int){
	total := 0
	for _, v := range input {
		if v == 4 || v == 1{
			v += 2
		}
		total += v
	}
	return total
}
