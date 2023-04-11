package watchman

import (
	"fmt"
	"os"

	"github.com/dharmit009/gopass/passutil"
)

const fileName = "./test.enc"

// This function is used to check if master key exists or not ...
func CheckMasterKey(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Passcheck is used to confirm if any two passwords are ssame or not ...
func CheckPassEqual(pass1, pass2 string) bool {
	return pass1 == pass2
}

// If password which is passed is equal to MP then return true else false
func CheckPassEqualToMP(pass1 string) bool {
	comparer, err := passutil.DecryptFile(fileName, pass1)
	if err != nil {
		return false
	}

	if pass1 != string(comparer) {
		return false
	} else {
		return true
	}
}

func CreateMasterKey(fileName, pass1, pass2 string) bool{
	if CheckMasterKey(fileName) == true {
		return true
	} else {
		if CheckPassEqual(pass1, pass2) == true {
			if err := passutil.EncryptFile(fileName, []byte(pass2), pass1); err != nil {
				fmt.Println("Error while creating masterkey")
				return false
			}
		} else {
			fmt.Println("Error: passwords are not equal")
			return false
		}
		return true
	}

}

// func main() {
// 	fmt.Println(CheckMasterKey(fileName))
// 	fmt.Println(CheckPassEqual("testing", "testing "))
// 	fmt.Println(CheckPassEqual("tsting", "testing "))
// 	fmt.Println(CheckPassEqual("testing", "testing"))
// 	fmt.Println(CreateMasterKey(fileName, "testing", "tester"))
// 	fmt.Println(CreateMasterKey(fileName, "testing", "testttt"))
// 	fmt.Println(CreateMasterKey(fileName, "testing", "testing"))
// 	fmt.Println(CheckPassEqualToMP("testing"))
// 	fmt.Println(CheckPassEqualToMP("test"))
// 	fmt.Println(CheckPassEqualToMP("asd;fj"))
//
// }
