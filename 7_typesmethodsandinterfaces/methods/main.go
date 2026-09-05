package main
import "fmt"
import "time"
type Person struct{
	FirstName string
	LastName string
	Age int
}

func (p Person) String() string{
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

type Incrementer interface{
	Increment()
}

type Counter struct{
	total int
	lastUpdated time.Time
}

func (c *Counter) Increment(){
	if c == nil{
		fmt.Println("Provided data is nil.")
		return
	}
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string{
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main(){
	p := Person{
		FirstName: "Badar",
		LastName: "Muneer",
		Age: 24,
	}
	
	fmt.Println(p.String())
	
	var pointerCounter *Counter
	
	var incrementer Incrementer
	incrementer = pointerCounter
	
	incrementer.Increment()
}