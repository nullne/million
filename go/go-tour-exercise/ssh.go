package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"io/ioutil"
	"os"
	"os/user"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"

	"github.com/howeyc/gopass"
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
	defer client.Close()

	// this sends keepalive packets every 2 seconds
	// there's no useful response from these, so we can just abort if there's an error
	go func() {
		t := time.NewTicker(2 * time.Second)
		defer t.Stop()
		for {
			<-t.C
			res, suc, err := client.Conn.SendRequest("keepalive@golang.org", true, nil)
			if err != nil {
				fmt.Println(err.Error())
				return
			} else {
				fmt.Println(res)
				fmt.Println(suc)
			}
		}
	}()
	time.Sleep(100 * time.Second)
	return client, nil
}

func authMethod(t string) ssh.AuthMethod {
	switch t {
	case "password":
		fmt.Printf("Password: ")
		pass := gopass.GetPasswdMasked()
		return ssh.Password(string(pass))
	case "key":
		usr, _ := user.Current()
		file := usr.HomeDir + "/.ssh/id_rsa"
		buf, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		key, err := ssh.ParsePrivateKey(buf)
		if err != nil {
			panic(err)
		}
		return ssh.PublicKeys(key)
	case "forwarding":
		sock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
		if err != nil {
			panic(err)
		}
		agent := agent.NewClient(sock)
		signers, err := agent.Signers()
		if err != nil {
			log.Fatal(err)
		}
		return ssh.PublicKeys(signers...)
	}
	return nil
}

func main() {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			authMethod("key"),
		},
	}
	_, err := SSHDialTimeout("tcp", "192.168.25.177:22", config, 5*time.Second)
	if err != nil {
		panic(err)
	}

	/*
		//ssh config
		if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
			sshConfig = getSSHConfig(opt.username, sshAgent)
			defer sshAgent.Close()
		} else {
			log.Fatal(err.Error())
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
		if err != nil {
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
	*/

	// if sttyOutput := buf.String(); !strings.Contains(sttyOutput, "-echo ") {
	// 	fmt.Println(err)
	// 	// t.Fatalf("terminal mode failure: expected -echo in stty output, got %s", sttyOutput)
	// }else{
	// 	fmt.Println(sttyOutput)
	// }

}
