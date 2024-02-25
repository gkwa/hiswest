package test4

import (
	"fmt"
	"net/http"
	"os"
)

func RunTest4() {
	url := "https://httpbin.org/headers"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error making GET request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "unexpected status code: %v\n", resp.StatusCode)
		return
	}

	_, err = fmt.Fprintln(os.Stdout, "Response from", url+":")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing to stdout: %v\n", err)
		return
	}

	_, err = fmt.Fprintf(os.Stdout, "%s\n", resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing response body to stdout: %v\n", err)
		return
	}
}
