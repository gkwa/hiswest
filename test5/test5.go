package test5

import (
	"bytes"
	"encoding/json"
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

	var bodyBytes []byte
	bodyBytes, err = io.ReadAll(resp.Body) // resp.Body is an io.ReadCloser
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading response body: %v\n", err)
		return
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading response body: %v\n", err)
		return
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, bodyBytes, "", "\t")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error formatting JSON: %v\n", err)
		return
	}

	_, err = fmt.Fprintf(os.Stdout, "%s\n", prettyJSON.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing response body to stdout: %v\n", err)
		return
	}
}
