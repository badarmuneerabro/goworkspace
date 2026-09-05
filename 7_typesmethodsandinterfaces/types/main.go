package main

import "fmt"
// User-defined type named Person, having underlying type of struct.
type Person struct{
	FirstName string
	LastName string
	Age int
}

type Score int
type Converter func(string) Score

type HighScore Score

func main(){
	var i int = 10
	var s Score = 100
	var hs HighScore = 1000
	
	//hs = s
	//s = i
	s = Score(i)
	hs = HighScore(s)
	
	fmt.Println(s)
	fmt.Println(hs)
}