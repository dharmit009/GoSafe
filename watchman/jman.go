package watchman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dharmit009/gopass/passutil"
)

const DbFilename = "passwords.json"

type PasswordEntry struct {
	Website  string `json:"website"`
	Username string `json:"username"`
	Password []byte `json:"password"`
}

type PasswordDB struct {
	Entries []PasswordEntry `json:"entries"`
}

func LoadPasswordDB(filename string) PasswordDB {
	var db PasswordDB
	file, err := os.Open(filename)
	if err == nil {
		defer file.Close()
		bytes, err := ioutil.ReadAll(file)
		if err == nil {
			json.Unmarshal(bytes, &db)
		}
	}
	return db
}

func savePasswordDB(db PasswordDB, filename string) error {
	bytes, _ := json.MarshalIndent(db, "", "    ")
	err := ioutil.WriteFile(filename, bytes, 0600)
	return err
}

func SearchPasswordDBByWebsite(db PasswordDB, websiteName string) []PasswordEntry {
	matchingEntries := make([]PasswordEntry, 0)
	for _, entry := range db.Entries {
		if entry.Website == websiteName {
			matchingEntries = append(matchingEntries, entry)
		}
	}
	return matchingEntries
}

func PrintPasswordEntries(entries []PasswordEntry, mpass string) string {
	var decpasswd []byte
	for _, entry := range entries {
		fmt.Printf("Website: %s\n", entry.Website)
		fmt.Printf("Username: %s\n", entry.Username)
		decpasswd, _ = passutil.Decrypt(entry.Password, mpass)
		fmt.Printf("Password: %s\n", decpasswd)
		fmt.Println()
	}
	return string(decpasswd)
}

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
	err := savePasswordDB(db, DbFilename)
	return err

}

//	func RemoveEntry(website, mpass string) error {
//		db := LoadPasswordDB(DbFilename)
//		// Example usage
//
//		db.Entries = append(db.Entries, newEntry)
//		err := savePasswordDB(db, DbFilename)
//		return err
//
// }
func GetEntry(website, mpass string) string {
	db := LoadPasswordDB(DbFilename)
	// Retrieve the entry by website name
	matchingEntries := SearchPasswordDBByWebsite(db, website)
	out := PrintPasswordEntries(matchingEntries, mpass)
	return out
}
