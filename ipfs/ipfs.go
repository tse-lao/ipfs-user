package ipfs

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

func Init() {
	//initialize ipfs locally'
	out, err := exec.Command("ipfs", "init").Output()
	if err != nil {
		//TODO: Figure out a better way to display this error, because it is already initialized.
		log.Println("Error at IPFS.go, func Init()", err)
	}

	fmt.Println(string(out))

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

func RunDaemon() Notification {
	result := Notification{}

	//cmdStruct := exec.Command("ipfs", "daemon")
	cmd := exec.Command("ipfs", "daemon")

	//out, err := cmdStruct.CombinedOutput()

	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	stdout := NewCapturingPassThroughWriter(os.Stdout)
	stderr := NewCapturingPassThroughWriter(os.Stderr)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		wg.Done()
	}()

	_, errStderr = io.Copy(stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatalf("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	return result

}

type CapturingPassThroughWriter struct {
	buf bytes.Buffer
	w   io.Writer
}

// NewCapturingPassThroughWriter creates new CapturingPassThroughWriter
func NewCapturingPassThroughWriter(w io.Writer) *CapturingPassThroughWriter {
	return &CapturingPassThroughWriter{
		w: w,
	}
}

func (w *CapturingPassThroughWriter) Write(d []byte) (int, error) {
	w.buf.Write(d)
	return w.w.Write(d)
}

// Bytes returns bytes written to the writer
func (w *CapturingPassThroughWriter) Bytes() []byte {
	return w.buf.Bytes()
}

func copyAndCapute(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)

	for {
		n, err := r.Read(buf[:])

		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)

			if err != nil {
				return out, err
			}
		}
		if err != nil {
			//Read retinr io.OEF at the end of the file, which is not an error
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
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
