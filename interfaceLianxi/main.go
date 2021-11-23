package main

import (
	"errors"
	"fmt"
	"os"
)

type Account struct {
	FirstName string
	LastName  string
}
type Employee struct {
	Account
	Credits float64
}

func (a *Account) ChangeName(newName string) {
	a.FirstName = newName
	fmt.Println(*a)

}
func (e Employee) String() string {
	return fmt.Sprintf("Name:%s %s\nCredits:%.2f\n", e.FirstName, e.LastName, e.Credits)

}
func CreateEmployee(firstName string, lastName string, credits float64) (*Employee, error) {
	return &Employee{Account{firstName, lastName}, credits}, nil

}
func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		e.Credits = e.Credits + amount
		return e.Credits, nil
	}
	return 0.0, errors.New("Invalid credit amount")

}
func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= e.Credits {
			e.Credits -= amount
			return e.Credits, nil
		}
		return 0.0, errors.New("You can't remove more credits than the account has")
	}
	return 0.0, errors.New("You can't remove negative numbers")

}
func (e *Employee) CheckCredits() float64 {
	return e.Credits
}

func main() {
	bruce,_:=CreateEmployee("maxin","xin",500)
	fmt.Println(bruce)
	fmt.Println(bruce.CheckCredits())
	_,err:=bruce.AddCredits(250)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(bruce.CheckCredits())
	bruce.ChangeName("maxin2")
	fmt.Println(bruce)


}
