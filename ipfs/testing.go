package ipfs

import (
	"fmt"
	"testing"

	"github.com/tse-lao/ether-user/wallet"
)

func TestEncryptDescrypt(t *testing.T) {

	//GET PUBLIC KEY
	publicKey := wallet.GetPublicKey("somepassword")

	fmt.Println(publicKey)
	data := []byte("encrypt this data please")

	fmt.Println("We encrypted data below with public key: \n\n ")
	encrypted := wallet.EncryptWithPublicKey(publicKey, data)

	fmt.Println("We descrypted with private key: \n\n ")
	wallet.DecryptWithPrivateKey("somepassword", encrypted)

}
