package main

import (
	"./employee"
)

func main() {
	// 1.
	// e := employee.Employee{
	// 	FirstName:   "Sam",
	// 	LastName:    "Adolf",
	// 	TotalLeaves: 30,
	// 	LeavesTaken: 20,
	// }

	// 2. new
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
