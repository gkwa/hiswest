package test5

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func RunTest5() {
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading response body: %v\n", err)
		return
	}

	_, err = fmt.Fprintf(os.Stdout, "%s\n", body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing response body to stdout: %v\n", err)
		return
	}
}
