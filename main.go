package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile() {
	// Read the file

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	/*
		data, err := ioutil.ReadFile("example.txt")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println(string(data))
	*/
}

func main() {

	// Create a file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	// Write to the file
	writer := bufio.NewWriter(file)

	// read the complete string from the console
	reader := bufio.NewReader(os.Stdin)

	// read the string from the console until the delimiter is found
	input, _ := reader.ReadString('\n')

	_, err = writer.WriteString(input)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	writer.Flush()

	fmt.Println("File written successfully")
	ReadFile()
}
