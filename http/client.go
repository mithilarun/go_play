package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Fetch URL from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter URL: ")
	url, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	/*
	 * Get rid of the trailing newline character
	 */
	url = strings.TrimSpace(url)
	/*
	 * Now that we have the URL, we will run it and display the response.
	 */

	 // Using only GET requests for now
	 response, err := http.Get(url)
	 if err != nil {
		 log.Fatal(err)
	 }

	 /*
	  * Close the response body when done reading it.
	 */
	 defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body[:]))
}
