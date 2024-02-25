package test2

import (
	"fmt"
	"os"
)

func RunTest2() {
	str := "Hello, world!"

	outputDestinations := []struct {
		writer *os.File
		target *os.File
	}{
		{writer: os.Stdout, target: os.Stdout},
		{writer: nil, target: nil}, // Placeholder for output.txt
	}

	for _, dest := range outputDestinations {
		var (
			writer *os.File = dest.writer
			target *os.File = dest.target
			err    error
		)

		if target == nil {
			target, err = os.Create("output.txt")
			if err != nil {
				fmt.Fprintf(os.Stderr, "error creating file: %v\n", err)
				return
			}
			defer target.Close()
			writer = target
		}

		_, err = fmt.Fprintf(writer, "%s\n", str)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error writing to file: %v\n", err)
			return
		}

		if dest.target == nil {
			fmt.Println("String successfully written to output.txt")
		}
	}
}
