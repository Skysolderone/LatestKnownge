package main

type Worker struct {
	ID   int
	Name string
}

func (w Worker) GetID() int {
	return w.ID
}

func (w Worker) GetName() string {
	return w.Name
}

type Supervisor struct {
	ID   int
	Name string
}

func (s Supervisor) GetID() int {
	return s.ID
}

func (s Supervisor) GetName() string {
	return s.Name
}

type Department struct {
	Workers     []Worker
	Supervisors []Supervisor
}
type Employee interface {
	GetID() int
	GetName() string
}
type Departments struct {
	Employees []Employee
}

func main() {

}