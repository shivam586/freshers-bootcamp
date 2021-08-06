package main

import "fmt"

type Employee interface {
	Fulltime()
	Contractor()
	Freelancer()
}

type details struct {
	days,dailywage,workinghour int
}

func (t details) Fulltime() {
	fmt.Println(t.days*t.dailywage*t.workinghour)
}

func (t details) Contractor() {
	fmt.Println(t.days*t.dailywage*t.workinghour)
}

func (t details) Freelancer() {
	fmt.Println(t.days*t.dailywage*t.workinghour)
}
func main() {

	var employee string
	fmt.Println("please enter the type of employee")
	fmt.Scanln(&employee)
	var days,wage,hour int
     fmt.Scanln(&days)
	 fmt.Scanln(&wage)
	 fmt.Scanln(&hour)
	var i Employee = details{days,wage,hour}
	if employee == "fulltime" {
		i.Fulltime()
	}
	if employee == "contractor" {
		i.Contractor()
	}
	if employee == "freelancer" {
		i.Freelancer()
	}
}
