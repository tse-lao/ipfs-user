package ipfs

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

type FileStatus struct {
	name           string
	cid            string
	size           int
	cumulativeSize int
	childBlocks    int
	fileType       string
	dir            bool
}

func errorhandling(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

/*create folders
Make sure you implement the currect path to create a directory.
*/
func CreateFolder(folder_name string) (bool, string) {
	//check if the path string starts with a slash
	if folder_name[0:1] != "/" {
		return false, "The first character should be equal to /"
	}

	cmdStruct := exec.Command("ipfs", "files", "mkdir", folder_name)
	out, err := cmdStruct.Output()

	if err != nil {
		fmt.Println(err)
		return false, "An error occured with creating the a folder"
	}
	fmt.Println(string(out))

	return true, "succesfully created the directory"
}

//moves the path.
func MoveFile(currentPath, newPath string) (bool, string) {
	//moves the file to different path.
	cmdStruct := exec.Command("ipfs", "files", "mv", currentPath, newPath)
	out, err := cmdStruct.Output()

	if err != nil {
		return false, err.Error()
	}
	fmt.Println(out)

	return false, string(out)
}

//moves the path.
func AddFile(currentPath, newPath string) (bool, string) {
	//read the data before we can add it

	//moves the file to different path.
	// we cannot do this.
	cmdStruct := exec.Command("ipfs", "files", "write", "--create", newPath, currentPath)

	out, err := cmdStruct.Output()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(out)

	return true, "succesfully implemented"
}

func CopyAndMoveFile(currentPath, newPath string) (bool, string) {
	return true, "succesfully copied the file and moved it to the IPFS path."
}

func UploadFile(file []byte, path string) string {
	/*
		[ ] Make sure that a file can be added.
	*/
	cdmStruct := exec.Command("ipfs", "file", "add", "/newProfile")
	out, _ := cdmStruct.Output()

	fmt.Println(string(out))
	return "the path"
}

func ReadFile(path string) []byte {

	out, err := exec.Command("ipfs", "files", "read", path).Output()

	if err != nil {
		fmt.Println("error occured while reading ")
	}
	//	cmdStruct := exec.Command
	return out
}

/*
README:
*/
func RemoveFile(path string) (bool, string) {
	cdmStruct := exec.Command("ipfs", "files", "rm", path)
	out, err := cdmStruct.Output()

	if err != nil {
		fmt.Println("unable to remove the file", err)
		return false, "unable to remove the file, file does not exists"
	}

	fmt.Println("succesfully removed the file", out)
	return true, "succesfully removed the file"

}

/*
README: Read the directory should start with /
*/
func ReadDirectory(path string) []FileStatus {
	if len(path) < 1 {
		out, err := exec.Command("ipfs", "files", "ls").Output()
		errorhandling(err)

		result := ReadDirectoryAsList(path, out)
		return result
	}

	if string(path[0]) != "/" {
		log.Fatal("not possible to do this, since it needs to start with /")
	}

	out, err := exec.Command("ipfs", "files", "ls", path).Output()
	errorhandling(err)
	result := ReadDirectoryAsList(path, out)
	return result
}

func ReadDirectoryAsList(path string, files []byte) []FileStatus {
	reader := bytes.NewReader(files)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	var directory []FileStatus
	count := 0
	for scanner.Scan() {
		count++
		filestatus := FileStat(path + "/" + scanner.Text())
		filestatus.name = scanner.Text()
		directory = append(directory, filestatus)
	}
	return directory
}

func FileStat(path string) FileStatus {
	status := FileStatus{}

	out, err := exec.Command("ipfs", "files", "stat", path).Output()

	if err != nil {
		fmt.Println("error occured in FileStat", err)
		return status
	}

	reader := bytes.NewReader(out)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	counter := 0
	for scanner.Scan() {
		switch counter {
		case 0:
			status.cid = scanner.Text()
		case 2:
			integer, _ := strconv.Atoi(scanner.Text())
			status.size = integer
		case 4:
			integer, _ := strconv.Atoi(scanner.Text())
			status.cumulativeSize = integer
		case 6:
			integer, _ := strconv.Atoi(scanner.Text())
			status.childBlocks = integer
		case 8:
			status.fileType = scanner.Text()
		}
		counter++
	}
	return status
}
