package main

import (
	"fmt"
	"flag"
	"os"
	"log"
	"bufio"
	"strings"
	"net/http"
	"io"
)

func main() {
	log.SetFlags(2 | 3)
	wordlist := flag.String("w", "", "wordlist")
	url := flag.String("u", "", "url")
	method := flag.String("m", "GET", "method")
	flag.Parse()
	fmt.Println("wordlist:", *wordlist)

	// read wordlist
	file, err := os.Open(*wordlist)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// iterate over wordlist
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		requestUrl := strings.Replace(*url, "{}", scanner.Text(), 1)

		// send request
		switch *method {
			case "GET":
				resp, _ := http.Get(requestUrl)

				if resp.StatusCode == 200 {
					defer resp.Body.Close()
					body, _ := io.ReadAll(resp.Body)
					fmt.Println(string(body))
					os.Exit(0)
				}

			case "POST":
				resp, _ := http.Post(requestUrl, "text", nil)
				if resp.StatusCode == 200 {
					defer resp.Body.Close()
					body, _ := io.ReadAll(resp.Body)
					fmt.Println(string(body))
					os.Exit(0)
				}
		}

	}
}