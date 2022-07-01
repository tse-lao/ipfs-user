package main

import (
	"fmt"
	"os"
	"time"

	ipns "github.com/ipfs/go-ipns"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/tse-lao/ipfs-user/ipfs"
)

func main() {
	//var pathlink string = "/Users/koenschuite/Desktop/facebook-koenschuite/comments_and_reactions/comments.json"

	//ipfs.AddFile(pathlink, "newlycreated/newtry.json")

	const pathName = "/Users/koenschuite/Desktop/facebook-koenschuite"
	const fileName = "/facebook-data/activity_messages"

	ipfs.CreateFolder(fileName)

	//result := ipfs.ReadAllSubDirectories(pathName)

	//ipfs.CreateMetaDataFile(result, pathName, fileName)

}

func TestingAddFile() {
	// createUserProfile("0xc94d737b36A32BbC4eaf545832C05420fa11B919")

	/*
		[x] Find files
		[x] display as an array.
	*/
	const movefile = "/Users/koenschuite/Documents/pasfoto.png"

	ipfs.AddFile(movefile, "/newimage.png")

	result := ipfs.ReadDirectory("")
	fmt.Println(result)
	/* file := ipfs.ReadFile("/pasfoto.png")
	fmt.Println(file)
	fmt.Println("Result : " + result)

	//remove v=file
	ipfs.RemoveFile("/newPassfoto.png") */
}

func creationProcess() {
	fmt.Println("=== INITIALIZING THE IPFS ====")
	ipfs.Init()

	ipfs.GenKey("0xc94d737b36A32BbC4eaf545832C05420fa11B919")
	//returned this: k51qzi5uqu5dk4ap2b5ufp6qxibilq5wrj5omngnoezwyir6qvbzd6onxp9f91
	ipfs.AllKeys()

	// TODO: Check if the profile is working and running.
	fmt.Println("=== FIND THE NEW PROFILES ====")
	ipfs.ReadDirectory("/newProfile")

	//TODO: This is not implementated yet.
	ipfs.CreateFolder("/newProfile")

	// generatePrivates()

	/*
		DAEMON RUNNING
		[]fix error so it won't crash
		result := ipfs.RunDaemon()
		fmt.Println(result)
	*/
}

func GeneratePrivates() {
	privateKey, publicKey, err := crypto.GenerateKeyPair(crypto.RSA, 2048)

	if err != nil {
		panic(err)
	}

	fmt.Println("Private key: \n", privateKey)
	fmt.Println("\n\nPublic key:\n", publicKey)

	ipnsRecord, err := ipns.Create(privateKey, []byte("/ipfs/Qme1knMqwt1hKZbc1BmQFmnm9f36nyQGwXxPGVpVJ9rMK5"), 0, time.Now().Add(1*time.Hour), 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(ipnsRecord, publicKey)
}

func CreateUserProfile(user_address string) {
	//first we need to create a user profile
	//which will be build based on user_address

	//lets try this then
	d1 := []byte("main_address:" + user_address)
	err := os.WriteFile("/tmp/profile", d1, 0655)

	if err != nil {
		fmt.Println("Something wrong with the creation of the profile:", err)
	}

	//ipfs.Init()

}

//we need to make sure that we can run the ipfs daemon in the go language.

//we need to make a function to add files to the ipfs locally.

//function to distribute the IPFS.

//make sure that there is a mechanism that someone is able to retrieve the hosting/sharing of the data,
