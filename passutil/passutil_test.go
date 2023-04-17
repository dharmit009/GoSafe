package passutil_test

import (
	"testing"
  "fmt"

	"github.com/dharmit009/gopass/passutil"
)

func TestGeneratePassword(t *testing.T) {
	// Call the GeneratePassword function
	password := passutil.GeneratePassword()

  fmt.Println("Generated Password: ", password)
	// Assert that the generated password is not empty
	if password == "" {
		t.Errorf("Generated password is empty")
	} else{
    fmt.Println("Password Generated Successfully")
  }

	// Assert that the length of the generated password is 16
	if len(password) != 16 {
		t.Errorf("Expected password length: %d, actual length: %d", 16, len(password))
	}

	// Add more assertions to check the password format, allowed characters, etc.
}
