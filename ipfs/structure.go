package ipfs

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func CreateCid() {
	//push the level of CID in different structure.

	//[] Here we put object data into it. and we get back a VID.
	// [] CID

	out, err := exec.Command("ipfs", "dag", "put").Output()

	//this givees me back a result. and will be stored in a different row of lists. like meta data.

	//we need to ipfs init
	if err != nil {
		//TODO: Figure out a better way to display this error, because it is already initialized.
		log.Println("Error at IPFS.go, func Init()", err)
	}

	fmt.Println(string(out))
}

func ReadAllSubDirectories(path string) []byte {
	//[] Read the current path and grab it
	// [] Read all the files
	cmd := exec.Command("ls", "-lR")

	if len(path) > 1 {
		cmd.Dir = path
	}

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
	}

	return out
}

func DagPut(data string) {
	cmd := exec.Command("ipfs", "dag", "put")

	out, _ := cmd.Output()

	fmt.Println("not working yet,,,,", out)
}
func CreateMetaDataFile(allFiles []byte, name string, fileName string) {
	reader := bytes.NewReader(allFiles)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	//here we create the dag and the values for later on, now e leave it like this,

	// CreateDag(name)

	directory := fileName

	//[] Check if there is a directory called FileName
	CreateFolder(fileName)

	var allPaths []string

	for scanner.Scan() {
		//fmt.Println(scanner.Text())

		lineReader := bytes.NewReader(scanner.Bytes())
		lineScanner := bufio.NewScanner(lineReader)
		lineScanner.Split(bufio.ScanWords)

		if strings.HasPrefix(scanner.Text(), "./") {
			directory = trimDirectory(scanner.Text())
			//check of directory is there.

		}

		if len(lineScanner.Text()) > 0 || lineScanner.Text() == scanner.Text() {
			if (lineScanner.Text()) == "." {
				fmt.Println("Directory:" + trimDirectory(lineScanner.Text()))
			}
		}

		counter := 0
		folder := false
		for lineScanner.Scan() {
			switch counter {
			case 1:
				i, _ := strconv.Atoi(lineScanner.Text())
				if i > 1 {
					folder = true
				}
			case 8:

				allPaths = append(allPaths, directory+"/"+lineScanner.Text())
				if folder {
					CreateFolder(directory + "/" + lineScanner.Text())
					fmt.Println(directory + "/" + lineScanner.Text())
				}
			}
			counter++
			folder = false
		}
	}

	//now we want to only get the files and collect the data where it belongs to.
	//we create directories and dag. we put the dag above the rest.
	writePaths, err := os.Create("paths.txt")

	if err != nil {
		fmt.Println(err)
	}

	w := bufio.NewWriter(writePaths)
	for _, line := range allPaths {
		fmt.Fprintln(w, line)
	}

	//[] we want to move all the files into the ipfs based on their

	write1, _ := os.Create("metadata.txt")
	write1.Write(allFiles)

	moveArrayOfFiles(allPaths, name)

}

func trimDirectory(directory string) string {
	if len(directory) <= 2 {
		log.Fatal("this is too short to even be a directory")
	}
	length := len(directory) - 1
	newDirectory := directory[1:length]

	return newDirectory
}

// now we want to put it in the ipfs as simple as tthat.
func moveArrayOfFiles(paths []string, name string) string {
	for _, path := range paths {
		//	AddFile((name + "/" + path), path)
		var namePath string = name + path
		var newName string = path
		fmt.Println(namePath)
		fmt.Println(newName)

		//TODO:[] here we need to add the encryption for the file.. one by one.
		//AddFile(namePath, newName, )
	}

	return "we have uploaded all to ipfs"
}

//make sure the new directory is there.

func checkDirectory(path string) {
	fmt.Println("We check if directory exsilen(fmt ")

}
