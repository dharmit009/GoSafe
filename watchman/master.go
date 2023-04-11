package passutil

import (
	"os"

	"github.com/dharmit009/gopass/passutil"
)

const fileName = "./test.enc"

// This function is used to check if master key exists or not ...
func CheckMasterkey(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Passcheck is used to confirm if any two passwords are ssame or not ...
func PassEqual(pass1, pass2 string) bool {
	return pass1 == pass2
}

// If password which is passed is equal to MP then return true else false
func PassEqualToMP(pass1 string) bool {
	comparer, err := string(passutil.DecryptFile(fileName, pass1))
	if err != nil {
		return false
	}

	if pass1 != comparer {
		return false
	} else {
		return true
	}
}
