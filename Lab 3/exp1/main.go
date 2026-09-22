package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (person *Person) ReadData() {
	fmt.Print("Enter name: ")
	fmt.Scan(&person.Name)
	fmt.Print("Enter age: ")
	fmt.Scan(&person.Age)
	fmt.Print("Enter job: ")
	fmt.Scan(&person.Job)
	fmt.Print("Enter salary: ")
	fmt.Scan(&person.Salary)
}

func (person Person) Display() {
	fmt.Println("--------------------")
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Job: %s\n", person.Job)
	fmt.Printf("Salary: %.2f\n", person.Salary)
	fmt.Println("--------------------")
}

func main() {
	var firstPerson Person
	var secondPerson Person

	fmt.Println("Enter details for Person 1")
	firstPerson.ReadData()

	fmt.Println("\nEnter details for Person 2")
	secondPerson.ReadData()

	fmt.Println("\nPerson 1")
	firstPerson.Display()

	fmt.Println("Person 2")
	secondPerson.Display()
}
