package main

import "fmt"
type Employee struct{
	Name string
	ID string
}

func (e Employee) Description() string{
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct{
	Employee
	Reports []Employee
}
func (m Manager) FindNewEmployee(){
	
	fmt.Println("FindNewEmployee() -> Hello.")
	
}
func main(){
	m := Manager{
		Employee: Employee {
			Name: "Badar",
			ID: "20",
		},
		Reports: []Employee{},
	}
	
	fmt.Println(m.Description())
	m.FindNewEmployee()
}