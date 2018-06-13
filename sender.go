package main

import (
	"os"
	"os/exec"
	"strings"
	"fmt"
	em "fileshare/email"
)

func main() {

	toEmail := os.Args[1]
	filepath := os.Args[2]
	UploadToIPFS(filepath, toEmail)
}

func UploadToIPFS(filepath string, email string) {

	gatewayURL := "localhost"

	// Add to IPFS
	cmd := exec.Command("ipfs", "add", filepath)
	output, err := cmd.Output()
	fmt.Println(string(output))
	if err != nil {
		fmt.Println(err)
	}

	// Create a File URL and return
	lines := strings.Split(string(output[:]), "\n")
	words := strings.Split(lines[len(lines) - 2], " ")
	hash := words[1]
	url := gatewayURL + "/ipfs/" + hash
	response := map[string]interface{}{ "url": url }

	fmt.Println(response)
	em.SendMail(email, "yuva670@gmail.com", url)
}
