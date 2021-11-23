package main

import "fmt"

type Person struct {
	Name    string
	Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)

}
func (p Person) String2() string {
	return fmt.Sprintf("%v iss from %v", p.Name, p.Country)

}
func main() {
	rs:=Person{
		Name:    "John Doe",
		Country: "USA",
	}
	fmt.Printf("%s\n", rs)
}
