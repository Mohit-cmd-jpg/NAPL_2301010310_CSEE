package main

import (
	"fmt"
	"sync"
)

func calculateSquares(numbers []int, results chan<- string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	fmt.Println("Squares goroutine started")
	for _, number := range numbers {
		results <- fmt.Sprintf("Square of %d = %d", number, number*number)
	}
	fmt.Println("Squares goroutine completed")
}

func calculateCubes(numbers []int, results chan<- string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	fmt.Println("Cubes goroutine started")
	for _, number := range numbers {
		results <- fmt.Sprintf("Cube of %d = %d", number, number*number*number)
	}
	fmt.Println("Cubes goroutine completed")
}

func calculateFibonacci(count int, results chan<- string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	fmt.Println("Fibonacci goroutine started")
	first, second := 0, 1
	for index := 0; index < count; index++ {
		results <- fmt.Sprintf("Fibonacci term %d = %d", index+1, first)
		first, second = second, first+second
	}
	fmt.Println("Fibonacci goroutine completed")
}

func main() {
	numbers := []int{2, 3, 4, 5}
	results := make(chan string)

	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	go calculateSquares(numbers, results, &waitGroup)
	go calculateCubes(numbers, results, &waitGroup)
	go calculateFibonacci(len(numbers), results, &waitGroup)

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}

	fmt.Println("All concurrent tasks completed")
}
