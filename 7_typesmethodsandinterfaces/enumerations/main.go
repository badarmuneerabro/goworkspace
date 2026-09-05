package main
import "fmt"
type MailCategory int

const(
Uncategorized MailCategory = iota
Personal
Spam
Social
)

const(
Field1 = 0
Field2 = 1 + iota
Field3 = 20
Field4
Field5 = iota
Field6
Field7 = 21
Field8
Field9 = iota
)
func main(){
	fmt.Println(Uncategorized, " ", Personal, " ", Spam, " ", Social)
	fmt.Println(Field1, " ", Field2, " ", Field3, " ", Field4, " ", Field5, " ", Field6)
	fmt.Println(Field7, " ", Field8, " ", Field9)
}