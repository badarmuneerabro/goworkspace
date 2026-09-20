package greetings

import "fmt"
type Person struct{
	Name string
}

func (p *Person) Greet(name string){
	fmt.Println(sayHelloTo(name))
	fmt.Println(p.introduceYourSelf())
}

func sayHelloTo(name string) string{
	return fmt.Sprintf("Hello %s.", name)
}

func (p *Person) introduceYourSelf() string{
	return fmt.Sprintf("I am %s.", p.Name)
}