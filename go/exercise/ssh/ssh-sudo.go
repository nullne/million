package main

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"fmt"
	// "strings"
)

// password implements the ClientPassword interface
type password string

func (p password) Password(user string) (string, error) {
	return string(p), nil
}

var Password string = "ubuntu"

func main() {
	// An SSH client is represented with a ClientConn. Currently only
	// the "password" authentication method is supported.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of ClientAuth via the Auth field in ClientConfig.
	config := &ssh.ClientConfig{
		User: "le.yu",
		Auth: []ssh.AuthMethod{
			ssh.Password("ubuntu"),
		},
	}
	client, err := ssh.Dial("tcp", "192.168.25.177:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()
	fmt.Println(session)

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	stdin, err := session.StdinPipe()
	err = session.Run("sudo -u yule whoami")
	stdin.Write([]byte("ubuntu\r\n"))
	if err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

}
