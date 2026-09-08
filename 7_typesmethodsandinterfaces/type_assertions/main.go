package main
import "fmt"
type MyInt int

func main(){
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt)
	fmt.Println(i2 + 1)
	
	i3, ok := i.(int)
	if !ok {
		fmt.Printf("Unexpected type for %v", i)
	}else{
		fmt.Printf("%d", i3)
	}
}