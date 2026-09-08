package main
import "fmt"
//1. Define a function type
type MyFunc func(string) string

//2. Put a method on that function type
func (mf MyFunc) sayHello(name string){
	fmt.Println("sayHello() -> Start...")
	result := mf(name)
	fmt.Println(result)
	fmt.Println("sayHello() -> End...")
}
func main(){
	var greet MyFunc = func(name string) string{
		return "Hello, " + name
	}
	
	greet.sayHello("Ahmed")
}