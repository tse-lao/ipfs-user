package ipfs

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/tse-lao/ether-user/wallet"
)

type FileStatus struct {
	Name           string `json:"name"`
	Cid            string `json:"cat"`
	Size           int    `json:"size"`
	CumulativeSize int    `json:"cumlativeSize"`
	ChildBlocks    int    `json:"childBlocks"`
	FileType       string `json:"fileType"`
}

type FileData struct {
	Owner string `json:"owner"`
	Data  []byte `json:"data"`
	Time  string `json:"data_added"`
	Type  string `json:"type_doc"`
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
		fmt.Println("First Character needs tot start wtih /")
		return false, "The first character should be equal to /"
	}

	cmdStruct := exec.Command("ipfs", "files", "mkdir", "-p", folder_name)
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

func UploadToIPFS(currentPath string) (bool, string) {
	cmdStruct := exec.Command("ipfs", "add", currentPath)
	out, err := cmdStruct.Output()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(out))

	return true, string(out)
}

func DownloadFromIPFS(path string) (bool, string) {
	cmdStruct := exec.Command("ipfs", "get", path)
	out, err := cmdStruct.Output()

	if err != nil {
		return false, err.Error()
	}

	return true, string(out)
}

func AddFile(currentPath, newPath string) (bool, string) {
	//read the data before we can add itr

	cmdStruct := exec.Command("ipfs", "files", "write", "-p", "--create", newPath, currentPath)

	out, err := cmdStruct.Output()

	if err != nil {
		fmt.Println("Error occured when moving it to the right path. ")
		fmt.Println(err)
	}

	fmt.Println(out)

	return true, "succesfully implemented"
}

//addEncr
func PrivateFileAdd(currentPath, newPath, password string) (bool, string) {
	//read the data before we can add it
	result := ReadFile(currentPath)

	// publicKey *ecdsa.PublicKey

	encryptedData := wallet.DecryptWithPrivateKey(password, result)

	//returned the encrypted data, now need to write this byte format into something with teh right stats.
	//create metadata for this.
	fmt.Println(encryptedData)

	cmdStruct := exec.Command("ipfs", "files", "write", "-p", "--create", newPath, currentPath)

	out, err := cmdStruct.Output()

	if err != nil {
		fmt.Println("Error occured when moving it to the right path. ")
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
		//still an error that needs to be tackled. when reading file.
		fmt.Println("error occured while reading ")
	}
	//	cmdStruct := exec.Command

	fmt.Println(out)
	return out
}

/*
README:
*/
func RemoveFile(path string) (bool, string) {
	cdmStruct := exec.Command("ipfs", "files", "rm", path)
	out, err := cdmStruct.Output()
	fmt.Println(out)
	if err != nil {
		fmt.Println("unable to remove the file", err)
		return false, "unable to remove the file, file does not exists"
	}
	fmt.Println(out)

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
		filestatus.Name = scanner.Text()
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
			status.Cid = scanner.Text()
		case 2:
			integer, _ := strconv.Atoi(scanner.Text())
			status.Size = integer
		case 4:
			integer, _ := strconv.Atoi(scanner.Text())
			status.CumulativeSize = integer
		case 6:
			integer, _ := strconv.Atoi(scanner.Text())
			status.ChildBlocks = integer
		case 8:
			status.FileType = scanner.Text()
		}
		counter++
	}
	return status
}

func WriteToIPFS(data string) {
	cdmStruct := exec.Command("ipfs", "dag", "put", data)

	out, err := cdmStruct.Output()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(out)
}
