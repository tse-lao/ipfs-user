package main

import (
	"fmt"
	"os"
	"time"

	ipns "github.com/ipfs/go-ipns"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/tse.lao/ipfs-user/ipfs"
)

func main() {
	// createUserProfile("0xc94d737b36A32BbC4eaf545832C05420fa11B919")
	fmt.Println("=== INITIALIZING THE IPFS ====")
	ipfs.Init()
	ipfs.GenKey("0xc94d737b36A32BbC4eaf545832C05420fa11B916")
	//returned this: k51qzi5uqu5dk4ap2b5ufp6qxibilq5wrj5omngnoezwyir6qvbzd6onxp9f91
	ipfs.AllKeys()

	// TODO: Check if the profile is working and running.
	fmt.Println("=== FIND THE NEW PROFILES ====")
	ipfs.FindItems("/newProfile")

	//TODO: Not implemented yet..
	ipfs.CheckFolder("")

	//TODO: This is not implementated yet.
	ipfs.CreateFolder("/newProfile", "notimplementedyet!")

	// generatePrivates()

}

func generatePrivates() {
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

func createUserProfile(user_address string) {
	//first we need to create a user profile
	//which will be build based on user_address

	//lets try this then
	d1 := []byte("main_address:" + user_address)
	err := os.WriteFile("/tmp/profile", d1, 0655)

	if err != nil {
		fmt.Println("Something wrong with the creation of the profile:", err)
	}

	fmt.Println("We wrote the profile file, together with the addres")

	fmt.Println("\n printing the init now")
	//ipfs.Init()

}

//we need to make sure that we can run the ipfs daemon in the go language.

//we need to make a function to add files to the ipfs locally.

//function to distribute the IPFS.

//make sure that there is a mechanism that someone is able to retrieve the hosting/sharing of the data,
