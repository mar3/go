package main

import "fmt"
import "io"
import "io/ioutil"
import "log"
import "os"
import "golang.org/x/crypto/ssh"


func _open() *ssh.Client {

	key, err := ioutil.ReadFile("/home/ec2-user/.ssh/id_rsa")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
		panic("")
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
		panic("")
	}

	conf := &ssh.ClientConfig{
		User: "ec2-user",
		Auth: []ssh.AuthMethod{
			// ssh.Password(password),
			ssh.PublicKeys(signer),
		},
	}

	host := "127.0.0.1:22"
	client, err := ssh.Dial("tcp", host, conf)
	if err != nil {
		fmt.Println(err)
		panic("")
	}

	return client
}

func _hello1() {

	client := _open()
	defer client.Close()
	session, _ := client.NewSession()
	defer session.Close()
	session.Stdout = os.Stdout

	// SUCCESS
	session.Run("id")
}

func _hello2() {

	client := _open()
	defer client.Close()
	session, _ := client.NewSession()
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stdin = os.Stdin

	length, err := session.Stdout.Write([]byte("id\n"))
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	fmt.Printf("[%d] bytes sent.\n", length)

	// s := bufio.NewScanner(session.Stdin)
	// for s.Scan() {
	// 	line := s.Text()
	// 	fmt.Println(line)
	// }
}

func _hello3() {

	client := _open()
	defer client.Close()
	session, _ := client.NewSession()
	defer session.Close()
	// session.Stdout = os.Stdout
	// session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		// ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err := session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		log.Fatalf("request for pseudo terminal failed: %s", err)
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatalf("Unable to setup stdin for session: %v\n", err)
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatalf("Unable to setup stdout for session: %v\n", err)
	}

	go io.Copy(os.Stdout, stdout)
	go io.Copy(stdin, os.Stdin)

	// length, err := session.Stdout.Write([]byte("id\n"))
	// if err != nil {
	// 	panic("Failed to create session: " + err.Error())
	// }
	// fmt.Printf("[%d] bytes sent.\n", length)

	session.Run("id")
	session.Run("cd /tmp")
	session.Run("pwd")

	// err := session.Shell()
	// if err != nil {
	// 	log.Fatalf("failed to start shell: %s", err)
	// }
	// if err := session.Run()
	fmt.Println("Ok.")
}

func main() {

	// _hello1()
	// _hello2()
	_hello3()
}
