package main

import "fmt"
import "sync"

func main() {
	result := Parse("hello")

	fmt.Println(result)

	result2 := Parse("goodbye")

	fmt.Println(result2)
}

type SlowComplicatedParser interface {
	Parse(string) string
}

var parser SlowComplicatedParser
var once sync.Once

func Parse(dataToParse string) string {
	once.Do(func() {
		parser = initParser()
	})

	return parser.Parse(dataToParse)
}

func initParser() SlowComplicatedParser {
	fmt.Println("inializing!")

	return SCPI{}
}

type SCPI struct {
}

func (s SCPI) Parse(in string) string {
	if len(in) > 1 {
		return in[0:1]
	}

	return ""
}
