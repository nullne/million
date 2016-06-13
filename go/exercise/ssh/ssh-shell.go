package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	// "log"
	"net"
	"time"
	// "bufio"
	// "strings"
	"bytes"
	// "os"
	"io"
)

// Conn wraps a net.Conn, and sets a deadline for every read
// and write operation.
type Conn struct {
	net.Conn
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (c *Conn) Read(b []byte) (int, error) {
	err := c.Conn.SetReadDeadline(time.Now().Add(c.ReadTimeout))
	if err != nil {
		return 0, err
	}
	return c.Conn.Read(b)
}

func (c *Conn) Write(b []byte) (int, error) {
	err := c.Conn.SetWriteDeadline(time.Now().Add(c.WriteTimeout))
	if err != nil {
		return 0, err
	}
	return c.Conn.Write(b)
}

func SSHDialTimeout(network, addr string, config *ssh.ClientConfig, timeout time.Duration) (*ssh.Client, error) {
	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		fmt.Println("error here")
		return nil, err
	}

	timeoutConn := &Conn{conn, timeout, timeout}
	c, chans, reqs, err := ssh.NewClientConn(timeoutConn, addr, config)
	if err != nil {
		return nil, err
	}
	client := ssh.NewClient(c, chans, reqs)

	// this sends keepalive packets every 2 seconds
	// there's no useful response from these, so we can just abort if there's an error
	// go func() {
	// 	t := time.NewTicker(2 * time.Second)
	// 	defer t.Stop()
	// 	for {
	// 		<-t.C
	// 		_, _, err := client.Conn.SendRequest("keepalive@golang.org", true, nil)
	// 		if err != nil {
	// 			return
	// 		}
	// 	}
	// }()
	// time.Sleep(100 * time.Second)
	return client, nil
}
func main() {
	config := &ssh.ClientConfig{
		User: "le.yu",
		Auth: []ssh.AuthMethod{
			ssh.Password("ubuntu"),
		},
	}
	client, err := SSHDialTimeout("tcp", "192.168.25.177:22", config, 10*time.Second)
	if err != nil {
		panic(err)
	}
	session, _ := client.NewSession()
	defer session.Close()
	// err = session.Shell()
	// fmt.Println(err)

	stdout, err := session.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		// t.Fatalf("unable to acquire stdout pipe: %s", err)
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		fmt.Println(err)
		// t.Fatalf("unable to acquire stdin pipe: %s", err)
	}

	stderr, err := session.StderrPipe()
	if err != nil{
		fmt.Println(err)
	}

	tm := ssh.TerminalModes{ssh.ECHO: 0}
	if err = session.RequestPty("xterm", 80, 40, tm); err != nil {
		fmt.Println(err)
		// t.Fatalf("req-pty failed: %s", err)
	}

	err = session.Shell()
	if err != nil {
		fmt.Println(err)
		// t.Fatalf("session failed: %s", err)
	}

	stdin.Write([]byte("fuck;exit\n"))

	var buf, errbuf bytes.Buffer
	if _, err := io.Copy(&buf, stdout); err != nil {
		fmt.Println(err)
		// t.Fatalf("reading failed: %s", err)
	}

	if _, err = io.Copy(&errbuf, stderr); err != nil {
		fmt.Println(err)
	}

	fmt.Println(errbuf.String())
	fmt.Println("----------------")
	fmt.Println(buf.String())

	// if sttyOutput := buf.String(); !strings.Contains(sttyOutput, "-echo ") {
	// 	fmt.Println(err)
	// 	// t.Fatalf("terminal mode failure: expected -echo in stty output, got %s", sttyOutput)
	// }else{
	// 	fmt.Println(sttyOutput)
	// }

}
