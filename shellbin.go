package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	input := getDataFromStdin()
	response, err := makeHTTPRequest(input)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("\x1B[36m ShellBin\033[0m")
	fmt.Printf("Here is your bin: https://shellbin.nextblu.com/#/bin/%s\n", response)
}

func getDataFromStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("There was an error reading your data: %v\n", err)
		panic(err)
	}
	return strings.Join(lines, "\n")
}

func makeHTTPRequest(data string) (string, error) {
	requestBody, err := json.Marshal(data)
	if err != nil {
		log.Printf("There was an error preparing the data %v\n", err)
		return "", errors.New("Error preparing data")
	}
	req, err := http.NewRequest("POST", "https://shellbin-api.nextblu.com/api/v1/bin/new", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Printf("There was an error creating a request to the ShellBin server %v\n", err)
		return "", errors.New("Error creating a request to the ShellBin server")
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("There was an error contacting the ShellBin server %v\n", err)
		return "", errors.New("Error contacting the ShellBin server")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("There was an error interpreting the ShellBin server response %v\n", err)
		return "", errors.New("Error interpreting ShellBin server response")
	}
	stringedBody := string(body)
	stringedBody = strings.ReplaceAll(stringedBody, "\"", "")
	return stringedBody, nil
}
