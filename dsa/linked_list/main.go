package main
import "fmt"
type Node struct{
	data int
	next *Node
}
type LinkedList struct{
	start *Node
}
func (l *LinkedList) insertAtStart(d int){
	l.start = &Node{
		data: d,
		next: l.start,
	}
	
	fmt.Println("insertAtStart() -> newStart=",l) 
}
func (l *LinkedList) PrintList(){
	for i := l.start; i != nil; i = i.next{
		fmt.Print(fmt.Sprintf("%d,", i.data))
	}
}
func main(){
	
	list := LinkedList{}
	list.insertAtStart(30)
	list.insertAtStart(20)
	list.insertAtStart(10)
	list.PrintList()
}