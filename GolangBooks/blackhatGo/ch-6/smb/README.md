# SMB
A Go package for communicating over SMB. Currently only minimal funcationality exists for client-side functions.

This is a book-friendly version. If you wish to contribute to the official package, see the repository at github.com/stacktitan/smb.

Here is a sample client that establishes a session with a server:

```go
package main

import (
	"log"

	"blackhat/ch-6/smb/smb"
)

func main() {

	host := "172.16.248.192"
	options := smb.Options{
		Host:        host,
		Port:        445,
		User:        "alice",
		Domain:      "corp",
		Workstation: "",
		Password:    "Password123!",
	}
	debug := false
	session, err := smb.NewSession(options, debug)
	if err != nil {
		log.Fatalln("[!]", err)
	}
	defer session.Close()

	if session.IsSigningRequired {
		log.Println("[-] Signing is required")
	} else {
		log.Println("[+] Signing is NOT required")
	}

	if session.IsAuthenticated {
		log.Println("[+] Login successful")
	} else {
		log.Println("[-] Login failed")
	}

	if err != nil {
		log.Fatalln("[!]", err)
	}
}

```
