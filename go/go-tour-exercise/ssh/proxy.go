package main

import (
	"bytes"
	"code.google.com/p/go.crypto/ssh"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Passwordless authentication
	sshsock := os.ExpandEnv("$SSH_AUTH_SOCK")
	addr, _ := net.ResolveUnixAddr("unix", sshsock)
	agentConn, _ := net.DialUnix("unix", nil, addr)
	agent := ssh.NewAgentClient(agentConn)
	config := &ssh.ClientConfig{
		User: "example-user",
		Auth: []ssh.ClientAuth{
			ssh.ClientAuthAgent(agent),
		},
	}

	fmt.Println("Dialing server B")
	bastion, err := ssh.Dial("tcp", "server-b:22", config)
	if err != nil {
		panic("Failed to dial server B: " + err.Error())
	}
	defer bastion.Close()

	fmt.Println("Dialing server C from server B")
	bastionClient, err := bastion.Dial("tcp", "server-c:22")
	if err != nil {
		panic("Failed to dial server C from server B: " + err.Error())
	}
	defer bastionClient.Close()

	fmt.Println("Handshaking server C from server B")
	client, err := ssh.Client(bastionClient, config)
	if err != nil {
		panic("Failed to handshake server C from server B: " + err.Error())
	}
	defer client.Close()

	// Create a session
	fmt.Println("Session on server C")
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("unable to create session: %s", err)
	}
	defer session.Close()

	var stdErr, stdOut bytes.Buffer
	session.Stdout = &stdOut
	session.Stdout = &stdErr
	fmt.Println("Running uptime on server C")
	if err := session.Run("/usr/bin/uptime"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(stdOut.String())
	fmt.Println(stdErr.String())
}
