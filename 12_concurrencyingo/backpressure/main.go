package main

import "fmt"
import "time"

type PressureGauge struct{
	ch chan struct{}
}

func New(limit int) *PressureGauge{
	return &PressureGauge{
		ch: make(chan struct{}, limit),
	}
}

func (pg *PressureGauge) Process(f func()){
	select{
		case pg.ch <- struct{}{}:
		f()
		<-pg.ch
		default:
		fmt.Println("No more capacity")
	}
}

func doThingsThatShouldBeLimited(){
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
func main(){
	pg := New(10)
	for i := 0; i <= 20; i++{
		go pg.Process(func(){
			doThingsThatShouldBeLimited()
		})
	}
	
	time.Sleep(30 * time.Second)
}