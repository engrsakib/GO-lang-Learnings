package main

type Person struct {
	Name string
	Age  int
	DOB  string
}

type Student struct {
	Person
	University string
	Major	  string
	RollNo     int
	Semester    int
	Grade string
}


type Address struct {
	State    string
	Division string
	District string
	Street   string
	City     string
	Zip      string
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	a := Address{State: "California", Division: "West", District: "Los Angeles", Street: "123 Main St", City: "Los Angeles", Zip: "90001"}

	println("Person:", p.Name, p.Age)
	println("Address:", a.Street, a.City, a.State, a.Zip, a.District, a.Division)


	s := Student{
		Person:     Person{Name: "Bob", Age: 20, DOB: "2004-01-15"},
		University: "XYZ University",
		Major:      "Computer Science",
		RollNo:     101,
		Semester:   4,
		Grade:      "A",
	}
	println("Student:", s.Name, s.Age, s.University, s.Major, s.RollNo, s.Semester, s.Grade, "DOB:", s.DOB)
}

func init() {
	println("Struct.go initialized")
}