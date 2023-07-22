package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/bhg/chapter6_SMB_and_NTLM_A_Peek_Down_the_Rabbit_Hole/smb/ntlmssp"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatalln("Usage: main <dictionary/file> <user> <domain> <hash>")
	}

	hash := make([]byte, len(os.Args[4])/2)
	_, err := hex.Decode(hash, []byte(os.Args[4]))
	if err != nil {
		log.Fatalln(err)
	}

	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	var found string
	passwords := bytes.Split(f, []byte{'\n'})
	for _, password := range passwords {
		h := ntlmssp.Ntowfv2(string(password), os.Args[2], os.Args[3])
		if bytes.Equal(hash, h) {
			found = string(password)
			break
		}
	}
	if found != "" {
		fmt.Printf("[+] Recovered password: %s\n", found)
	} else {
		fmt.Println("[-] Failed to recover password")
	}
}
