package sonicwall

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

// RunSSHCommand : what is this?
// Greet : describe what this function does
// CreateMessage : describe what this function does
func RunSSHCommand(username string, password string, hostname string, port string, command []string) (string, error) {

	//port := "22"
	if len(port) == 0 {
		port = "22"
	}

	// SSH client config
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		// Non-production only
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	results, err := executeCommand(command, hostname, port, config)

	return results, err
}

func executeCommand(commands []string, hostname string, port string, config *ssh.ClientConfig) (string, error) {

	// Connect to host
	client, err := ssh.Dial("tcp", hostname+":"+port, config)
	if err != nil {
		//log.Print(err)
		fmt.Println(err.Error())
	}
	defer client.Close()

	// Create sesssion
	sess, err := client.NewSession()
	if err != nil {
		//log.Fatal("Failed to create session: ", err)
		//log.Print("Failed to create session: ", err)
		fmt.Println("Failed to create session: " + err.Error())
	}
	defer sess.Close()

	// StdinPipe for commands
	stdin, err := sess.StdinPipe()
	if err != nil {
		log.Print(err)
	}

	// with a terminal config
	if err := sess.RequestPty("xterm", 40, 80, ssh.TerminalModes{ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400}); err != nil {
		return "", err
	}

	// Uncomment to store output in variable
	//var b bytes.Buffer
	//var errOut bytes.Buffer
	//sess.Stdout = &b
	//sess.Stderr = &b

	// Enable system stdout
	// Comment these if you uncomment to store in variable
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	// Start remote shell
	err = sess.Shell()
	if err != nil {
		//log.Fatal(err)
		//log.Print(err)
		fmt.Println(err.Error())
	}

	// send the commands
	for _, cmd := range commands {
		_, err = fmt.Fprintf(stdin, "%s\n", cmd)
		if err != nil {
			//log.Print(err)
			fmt.Println(err.Error())
		}
	}

	// Wait for sess to finish
	err = sess.Wait()
	if err != nil {

		//log.Print(err)
		fmt.Println(err.Error())

	}

	// Uncomment to store in variable
	//fmt.Println(b.String())

	//return b.String(), err
	return "Success", err

}
