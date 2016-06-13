package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"sync"

	"code.google.com/p/go.crypto/ssh"
)

const PS1 = "|>"

func MuxShell(w io.Writer, r, e io.Reader) (chan<- string, <-chan string) {
	in := make(chan string, 1)
	out := make(chan string, 100)
	var wg sync.WaitGroup
	go func() {
		for cmd := range in {
			wg.Add(1)
			w.Write([]byte(cmd + "\n"))
			wg.Wait()
		}
	}()

	go func() {
		init := false
		var (
			buf [65 * 1024]byte
			t   int = 0
		)
		for {
			n, err := r.Read(buf[t:])
			if err != nil {
				close(in)
				close(out)
				return
			}
			if s := string(buf[t:]); strings.Contains(s, "[sudo]") {
				w.Write([]byte("untu\n"))
			} else {
				t += n
			}

			if s := string(buf[:t]); strings.HasSuffix(s, PS1) { //assuming the $PS1 == 'sh-4.3$ '
				if init == false {
					init = true
				} else {
					out <- strings.Trim(s, PS1+"\n\r ")
				}
				t = 0
				wg.Done()
			}
		}
	}()
	return in, out
}

func main() {
	config := &ssh.ClientConfig{
		User: "le.yu",
		Auth: []ssh.AuthMethod{
			ssh.Password("ubuntu"),
		},
	}
	client, err := ssh.Dial("tcp", "192.168.25.177:22", config)
	if err != nil {
		panic(err)
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	defer client.Close()
	// Create a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("unable to create session: %s", err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Fatal(err)
	}

	w, err := session.StdinPipe()
	if err != nil {
		panic(err)
	}
	r, err := session.StdoutPipe()
	if err != nil {
		panic(err)
	}

	e, err := session.StderrPipe()
	if err != nil {
		panic(err)
	}
	in, out := MuxShell(w, r, e)
	// if err := session.Start("echo 'ubuntu'|sudo -S/bin/sh"); err != nil {
	if err := session.Shell(); err != nil {
		log.Fatal(err)
	}

	in <- "export PS1='|>'"
	in <- "whoami"
	in <- "sudo whoami"
	in <- "whoami"
	in <- "sudo whoami"
	in <- "whoami"
	in <- "exit"
	for o := range out {
		fmt.Printf("1:\t%s\n", o)
	}
	// in <- "sudo whoami"
	// fmt.Printf("whoami: %s\n", <-out)

	session.Wait()
}
