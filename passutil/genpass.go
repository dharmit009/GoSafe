package passutil

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const bold = "\033[1m"
const normal = "\033[0m"

func header(mess string) {
	fmt.Println()
	fmt.Println(bold + mess + normal)
	fmt.Println()
}

func GeneratePassword() string {
	// header("### Pass Generator ###")

	var uppercase_letters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var lowercase_letters string = strings.ToLower(uppercase_letters)
	var digits string = "0123456789"
	var symbols string = "(){}[].,_"
	var upps, lows, digs, syms bool = true, true, true, true

	// fmt.Println(bold+"uppercase:\t"+normal, uppercase_letters)
	// fmt.Println(bold+"lowercase:\t"+normal, lowercase_letters)
	// fmt.Println(bold+"digits:   \t"+normal, digits)
	// fmt.Println(bold+"symbols:  \t"+normal, symbols)

	// header("### Boolean Values: ###")

	// fmt.Println(bold+"upps:\t "+normal, upps)
	// fmt.Println(bold+"lows:\t "+normal, lows)
	// fmt.Println(bold+"digs:\t "+normal, digs)
	// fmt.Println(bold+"syms:\t "+normal, syms)

	// header("### Password Generation ### ")

	var length int32 = 16
	var allowedCharacters string = ""

	if upps == true {
		allowedCharacters += uppercase_letters
	}
	if lows == true {
		allowedCharacters += lowercase_letters
	}
	if digs == true {
		allowedCharacters += digits
	}
	if syms == true {
		allowedCharacters += symbols
	}

	// fmt.Println(bold+"Length: "+normal, length)
	// fmt.Println(bold+"All: "+normal, allowedCharacters)

	rand.Seed(time.Now().UnixNano())
	passwordBytes := make([]byte, length)
	for i := range passwordBytes {
		passwordBytes[i] = allowedCharacters[rand.Intn(len(allowedCharacters))]
	}
	var genPassword string = string(passwordBytes)

	// fmt.Println(bold + "Generated Password: " + normal + genPassword)
	return genPassword

}
