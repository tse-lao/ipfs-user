package ipfs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Init() {
	//initialize ipfs locally'
	out, err := exec.Command("ipfs", "init").Output()

	//we need to ipfs init
	if err != nil {
		//TODO: Figure out a better way to display this error, because it is already initialized.
		log.Println("Error at IPFS.go, func Init()", err)
	}
	fmt.Println(string(out))

}

func RunDaemon() interface{} {
	daemon := make(chan string)
	go StartDaemon(daemon)

	//for now this is fine to create and start the daemon  on hte backed.

	results := []interface{}{}
	for {
		msg := <-daemon
		fmt.Println(msg)

		results = append(results, msg)

		if msg == "Daemon is ready" {
			return results
		}

	}

	results = append(results, "Error has been occured")
	return results
}

func GenKey(address string) {
	//initialize ipfs locally'
	cmdStruct := exec.Command("ipfs", "key", "gen", address)
	out, err := cmdStruct.Output()

	fmt.Println("== CONFIGURE AND UPLOAD THE FILE ==")

	if err != nil {
		//TODO: need to log this in an error file, and command that this is an error
		log.Println("Error occured in generating key at GenKey", err)
	}

	fmt.Println(string(out))

	//we nneed to check if the string out is something new or already existing/
	fmt.Println("== UPLOAD THE FILE ==")

	os.WriteFile("/tmp/ipfs", out, 0644)

	//now lets tore this part for the address which is linked to the other file. this also needs to to be confirmed.
}

func AllKeys() {
	cmdStruct := exec.Command("ipfs", "key", "list")
	out, err := cmdStruct.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

//create folders
func CreateFolder(folder string, path string) {
	//check if the path string starts with a slash
	if folder[0:1] != "/" {
		fmt.Println("The first character should be equal to /")

		return
	}

	cmdStruct := exec.Command("ipfs", "files", "mkdir", folder)
	out, err := cmdStruct.Output()
	if err != nil {
		fmt.Println("Error creating folder", err)
	}
	fmt.Println(string(out))
}

func CheckFolder(name string) {
	out, err := exec.Command("ipfs", "files", "ls").Output()

	if err != nil {
		fmt.Println("Error with checking this")
	}

	fmt.Println(string(out))
}

func Run() {
	//this public function need to make sure that we can run the daemon
	cmdStruct := exec.Command("ipfs", "daemon")
	out, err := cmdStruct.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func AddFile(path string, data string) string {
	//create based on data provided/

	cmdStruct := exec.Command("ipfs", "files", "write", path, data)

	out, err := cmdStruct.Output()

	cmdStruct.StdoutPipe()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	return "the path should be returned here. "
}

func UploadFile(file []byte, path string) string {

	cdmStruct := exec.Command("ipfs", "file", "add", "")
	out, _ := cdmStruct.Output()

	fmt.Println(string(out))
	return "the path"
}

func ReadFile(path string) []byte {
	result := []byte("Something like a byte will be displayed here. ")
	//we need to add this functionality in of course.

	//	cmdStruct := exec.Command
	return result
}

type Notification struct {
	status  bool
	message string
}

//Shutdown the IPFS command for Daemon
func IpfsShutdown() Notification {
	result := Notification{}
	cmdStruct := exec.Command("ipfs", "shutdown")
	out, err := cmdStruct.Output()

	if err != nil {
		result.message = "Error shutting down the Daemon"
		result.status = false
	} else {
		result.message = string(out)
		result.status = true
	}

	return result
}

func StartDaemon(out chan string) {

	/*
		[x] Create connecting and constantely report the changes made.
		[x] Make sure the process is running on the background.
		[x] Make sure that is all up and running.
		[] Create a timeout, that can be triggered by the owner of the application.
	*/

	defer close(out)
	cmd := exec.Command("ipfs", "daemon")

	cmdReader, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)

	done := make(chan bool)

	go func() {
		for scanner.Scan() {
			//This does work. But there is still something going on.
			out <- scanner.Text()
		}
		done <- true
	}()

	cmd.Start()
	<-done
	err := cmd.Wait()

	if err != nil {
		fmt.Println(err)
		out <- "We have closed the connection, because an error occured"
		IpfsShutdown()
		close(out)
	}
}

func FindItems(path string) []byte {
	if len(path) < 1 {
		cmdStruct := exec.Command("ipfs", "files", "ls")
		out, err := cmdStruct.Output()

		message, errors := cmdStruct.CombinedOutput()

		fmt.Println(message)

		if errors != nil {
			fmt.Println("Combined Output gives me some errors", errors)
		}

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(out))

		return out
	}

	cmdStruct := exec.Command("ipfs", "files", "ls", path)
	out, err := cmdStruct.Output()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	return out
}
