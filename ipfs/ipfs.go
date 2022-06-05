package ipfs

import (
	"fmt"
	"os/exec"
)

func Init() {
	//initialize ipfs locally'
	out, err := exec.Command("ipfs", "init").Output()
	if err != nil {
		//TODO: Figure out a better way to display this error.
		fmt.Println("Error with IPFS INIT", err)
	}
	//the outcome of the init, we would like to store in the same setup files in tmp.

	fmt.Println(string(out))

}

func GenKey(address string) {
	//initialize ipfs locally'
	cmdStruct := exec.Command("ipfs", "key", "gen", address)
	out, stderr := cmdStruct.Output()
	if stderr != nil {
		fmt.Println(stderr)
	}
	fmt.Println(string(out))
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

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	return "the path should be returned here. "
}

func ReadFile(path string) []byte {
	result := []byte("Something like a byte will be displayed here. ")
	//we need to add this functionality in of course.

	//	cmdStruct := exec.Command
	return result
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
