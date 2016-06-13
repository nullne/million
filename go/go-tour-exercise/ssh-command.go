package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	// "reflect"
)

// An SSH client is represented with a ClientConn. Currently only
// the "password" authentication method is supported.
//
// To authenticate with the remote server you must pass at least one
// implementation of AuthMethod via the Auth field in ClientConfig.
func main() {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("ubuntu"),
		},
	}
	client, err := ssh.Dial("tcp", "192.168.107.137:22", config)
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

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		switch err.(type) {
		default:
			fmt.Println("Unknow")
		case *ssh.ExitError:
			fmt.Println(err.(*ssh.ExitError).ExitStatus())
		}
        //
		// fmt.Println(err.(type))
		// if (reflect.TypeOf(err) == reflect.TypeOf(&ssh.ExitError{})) {
		// 	fmt.Println(err.(*ssh.ExitError).ExitStatus())
		// 	fmt.Println(111)
		// }
		// fmt.Println(reflect.TypeOf(err))
		// fmt.Println(err.ExitStatus())
		// panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
