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
}
func (l *LinkedList) insertAtEnd(d int){
	if l.start == nil{
		l.insertAtStart(d)
		return
	}
	
	for i := l.start; i != nil; i = i.next {
		if i.next == nil {
			i.next = &Node{
				data: d,
				next: nil,
			}
			return
		}
	}
}

func (l *LinkedList) insertAt(index, data int){
	if l.start == nil || index <= 0 {
		l.insertAtStart(data)
		return
	}
	var pos int = 0
	for i := l.start; i != nil; i = i.next{
		if pos == (index - 1) {
			i.next = &Node{
				data: data,
				next: i.next,
			}
			return
		}
		pos++
		fmt.Println("pos=", pos)
	}
	l.insertAtEnd(data)
}

func (l *LinkedList) find(target int) int {
	index := 0
	for i := l.start; i != nil; i = i.next {
		if i.data == target {
			return index
		}
		index++
	}
	return -1
}
func (l *LinkedList) length() int{
	length := 0
	for i := l.start; i != nil; i = i.next {
		length++
	}
	
	return length
}
func (l *LinkedList) PrintList(){
	for i := l.start; i != nil; i = i.next{
		fmt.Print(fmt.Sprintf("%d,", i.data))
	}
	
	fmt.Println()
}
func main(){
	
	list := LinkedList{}
	list.insertAtStart(30) //0
	list.insertAtStart(20) //1
	list.insertAtStart(10) //2
	list.insertAtEnd(40) //3
	list.insertAtEnd(50) //4
	list.insertAt(2, 25)
	list.insertAt(-1, 5)
	list.PrintList()
	
	fmt.Println(list.find(10))
	fmt.Println("length=", list.length())
}