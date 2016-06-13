package main

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"fmt"
	// "log"
	"net"
	"os"
)

func main() {
	sock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		panic(err)
	}
	ag := agent.NewClient(sock)
	signers, err := ag.Signers()
	if err != nil {
		panic(err)
	}
	config := &ssh.ClientConfig{
		User: "le.yu",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signers...),
		},
	}

	fort, err := ssh.Dial("tcp", "192.168.10.171:22", config)
	if err != nil {
		panic(err)
	}
	target, err := fort.Dial("tcp", "219.146.210.23:22")
	if err != nil {
		panic(err)
	}

	c, chans, reqs, err := ssh.NewClientConn(target, "219.146.210.23:22", config)
	if err != nil {
		panic(err)
	}
	client := ssh.NewClient(c, chans, reqs)
	fmt.Println("*********")
	fmt.Println(client)
	session, err :=client.NewSession()
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run("uptime")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b.String())
}
