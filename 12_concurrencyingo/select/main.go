package main
import "fmt"
func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go func(){
		inGoroutine := 1
		ch1 <- inGoroutine
		fromMain := <- ch2
		
		fmt.Println("inGoroutine:", fromMain, inGoroutine)
	}()
	inMain := 2
	var fromGoroutine int
	
	select{
		case ch2 <- inMain:
		fmt.Println("Send data from main.")
		case fromGoroutine = <- ch1:
		fmt.Println("Read data in main.")
	}
	
	fmt.Println("main:", inMain, fromGoroutine)
}