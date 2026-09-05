package main
import "fmt"
type Node struct{
	data int
	next *Node
}
type LinkedList struct{
	Node
}
func (l *LinkedList) insertAtStart(d int){
	l.Node = Node{
		data: d,
		next: &l.Node,
	}
}

func (l *LinkedList) PrintList(){
	if l == nil{
		fmt.Println("[]");
	}
	fmt.Print(fmt.Sprintf("[%d,", l.Node.data))
	
	for i := l.Node.next; i != nil; i = i.next{
		fmt.Print(fmt.Sprintf("%d,", l.Node.data))
	}
	
	fmt.Print(fmt.Sprintf("]\n"))
}
func (l *LinkedList) insert(d int){
	
}
func main(){
	
	list := LinkedList{}
	list.insertAtStart(30)
	list.insertAtStart(20)
	list.insertAtStart(10)
	list.PrintList()
}