package main

import "fmt"

func printStudents(students []string) {
	fmt.Println("Students:", students)
}

func main() {
	fmt.Println("Slice operations")
	students := []string{"Aarav", "Diya", "Kabir"}
	printStudents(students)

	students = append(students, "Meera")
	fmt.Println("After add:")
	printStudents(students)

	removeIndex := 1
	students = append(students[:removeIndex], students[removeIndex+1:]...)
	fmt.Println("After remove by index 1:")
	printStudents(students)

	students[0] = "Anaya"
	fmt.Println("After update at index 0:")
	printStudents(students)

	fmt.Println("\nMap operations")
	marks := map[string]int{
		"Mathematics": 85,
		"Science":     91,
	}
	fmt.Println("Initial map:", marks)

	marks["English"] = 88
	fmt.Println("After insert English:", marks)

	delete(marks, "Science")
	fmt.Println("After delete Science:", marks)

	subject := "Mathematics"
	mark, found := marks[subject]
	if found {
		fmt.Printf("Lookup %s: %d\n", subject, mark)
	} else {
		fmt.Printf("Lookup %s: not found\n", subject)
	}
}
