package watchman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dharmit009/gopass/passutil"
)

// Default file where files are saved
const DbFilename = "passwords.json"

// A simple structure which is designed to store passwords
// website name, username, passsword.
type PasswordEntry struct {
	Website  string `json:"website"`
	Username string `json:"username"`
	Password []byte `json:"password"`
}

// Passwords Database Structure in Json.
type PasswordDB struct {
	Entries []PasswordEntry `json:"entries"`
}

// A function to load passwords from a json file.
func LoadPasswordDB(filename string) PasswordDB {
	var db PasswordDB
	file, err := os.Open(filename)
  defer file.Close()
	if err == nil {
		defer file.Close()
		bytes, err := ioutil.ReadAll(file)
		if err == nil {
			json.Unmarshal(bytes, &db)
		}
	}
	return db
}

// SavePasswordDB() is a funciton which is used to store passwords 
// in a json file.
func SavePasswordDB(db PasswordDB, filename string) error {
	bytes, _ := json.MarshalIndent(db, "", "    ")
	err := ioutil.WriteFile(filename, bytes, 0600)
	return err
}

// SearchPasswordDBByWebsite() is a funciton which is used to search passwords 
// with the help of website name from a json file.
func SearchPasswordDBByWebsite(db PasswordDB, websiteName string) []PasswordEntry {
	matchingEntries := make([]PasswordEntry, 0)
	for _, entry := range db.Entries {
		if entry.Website == websiteName {
			matchingEntries = append(matchingEntries, entry)
		}
	}
	return matchingEntries
}

// PrintPasswordEntries() is a function which can used to print all passwords 
// entries after decrypting.
func PrintPasswordEntries(entries []PasswordEntry, mpass string) (string, string, string) {
	var wname string
	var uname string
	var decpasswd []byte
	for _, entry := range entries {
		wname = entry.Website
		uname = entry.Username
		decpasswd, _ = passutil.Decrypt(entry.Password, mpass)
		fmt.Printf("Website: %s\n", entry.Website)
		fmt.Printf("Username: %s\n", entry.Username)
		fmt.Printf("Password: %s\n", decpasswd)
	}
	return wname, uname, string(decpasswd)
}

// PrintPairEntries() is a function used to print website name and username in pair
func PrintPairEntries(entries []PasswordEntry, mpass string) (string, string) {
	var wname string
	var uname string
	for _, entry := range entries {
		wname = entry.Website
		uname = entry.Username
		fmt.Printf("Website: %s\n", entry.Website)
		fmt.Printf("Username: %s\n", entry.Username)
	}
	return wname, uname
}

// AddEntry is a function which is used to make entry in passwords.json file.
func AddEntry(website, username, password, mpass string) error {
	db := LoadPasswordDB(DbFilename)
	// Example usage
	encpasswd, _ := passutil.Encrypt([]byte(password), mpass)
	newEntry := PasswordEntry{
		Website:  website,
		Username: username,
		Password: encpasswd,
	}

	db.Entries = append(db.Entries, newEntry)
	err := SavePasswordDB(db, DbFilename)
	return err

}

//	func RemoveEntry(website, mpass string) error {
//		db := LoadPasswordDB(DbFilename)
//		// Example usage
//
//		db.Entries = append(db.Entries, newEntry)
//		err := SavePasswordDB(db, DbFilename)
//		return err
//
// }

// GetEntry() is a function which you can use to get any specific entry by using website name.
func GetEntry(website, mpass string) (string, string, string) {
	db := LoadPasswordDB(DbFilename)
	// Retrieve the entry by website name
	matchingEntries := SearchPasswordDBByWebsite(db, website)
	wname, uname, pass := PrintPasswordEntries(matchingEntries, mpass)
	return wname, uname, pass
}
