package main

import "fmt"

func main(){
	num := 5
	dividends := 0
	for i := 1; i <= 10; i++ {
		if dividends >= 2{
			break
		}
		if num % i == 0 {
			dividends++
		}
	}
	
	fmt.Println(dividends)
	if dividends <= 2 {
		fmt.Printf("%d is a prime number.", num)
	}
	
}