package employee

import "fmt"

// 雇员
type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

// New
func New(FirstName string, LastName string, TotalLeaves int, LeavesTaken int) Employee {
	e := Employee{FirstName, LastName, TotalLeaves, LeavesTaken}
	return e
}

// 信息
func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
