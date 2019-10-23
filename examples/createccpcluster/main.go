package main

import (
	"fmt"
	"os"

	"github.com/swiftdiaries/ccp-clientlibrary-go/v3/ccp"
)

func main() {
	password := os.Getenv("CCP_PASSWORD")
	username := os.Getenv("CCP_USERNAME")
	baseurl := os.Getenv("CCP_BASEURL")

	client := ccp.NewClient(username, password, baseurl)
	err := client.Login(client)
	if err != nil {
		fmt.Printf("Error logging into client: %v", err)
	}
}
