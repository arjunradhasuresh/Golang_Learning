package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//common way
	var n int
	fmt.Println("Enter the number")
	fmt.Scan(&n)
	fmt.Println("Number is", n)

	//read a file
	// This methord is not a good methord becouse it store all data in the memory
	file, err := os.ReadFile("raman.txt")
	if err != nil {
		fmt.Println("some error occured")
		return
	}
	fmt.Println(string(file))

	//methord2
	//this methord will avoide storage of all data in memory.
	file1, err := os.Open("raman.txt")
	if err != nil {
		fmt.Println("some erro occured")
		return
	}
	defer file1.Close()
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
