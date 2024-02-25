package test1

import (
	"fmt"
	"os"
)

func RunTest1() {
	str := "Hello, world!"

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf("error creating file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s\n", str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing to file: %v\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "%s\n", str)
}
