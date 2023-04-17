package jman

import (
	"fmt"
)

const jsfile = "passwords.json"

type PasswordEntry struct {
	CredId   int    `json:"cred_id"`
	Website  string `json:"website"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type PasswordManager struct {
	Entries []PasswordEntry `json:"entries"`
}

// func LoadJsonFile(jsfile string) PasswordManager{
// 	var manager PasswordManager
// 	// Load existing entries from file, if it exists
// 	if _, err := os.Stat(jsfile); err == nil {
// 		data, _ := ioutil.ReadFile(jsfile)
// 		err = json.Unmarshal(data, &manager)
// 	}
//   return manager
// }

func (m *PasswordManager) GetEntry(id int) (PasswordEntry, error) {
	for _, entry := range m.Entries {
		if entry.CredId == id {
			return entry, nil
		}
	}
	return PasswordEntry{}, fmt.Errorf("entry with CredId %d not found", id)
}

func (m *PasswordManager) AddEntry(entry PasswordEntry) {
	// Find the highest existing CredId, and increment by 1 for the new entry

  fmt.Println(jsfile)
  fmt.Println(entry.Website)
  fmt.Println(entry.Username)
  fmt.Println(entry.Password)
	highestId := 0
	for _, existing := range m.Entries {
		if existing.CredId > highestId {
			highestId = existing.CredId
		}
	}
	entry.CredId = highestId + 1
	m.Entries = append(m.Entries, entry)
}

func (m *PasswordManager) RemoveEntry(id int) error {
	for i, entry := range m.Entries {
		if entry.CredId == id {
			m.Entries = append(m.Entries[:i], m.Entries[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("entry with CredId %d not found", id)
}

func (m *PasswordManager) UpdateEntry(updatedEntry PasswordEntry) error {
	for i, entry := range m.Entries {
		if entry.CredId == updatedEntry.CredId {
			m.Entries[i] = updatedEntry
			return nil
		}
	}
	return fmt.Errorf("entry with CredId %d not found", updatedEntry.CredId)
}

func GetNewEntry(w, u, p string) (PasswordEntry, error) {
	var entry PasswordEntry

	entry.Website = w
	entry.Username = u
	entry.Password = p

	return entry, nil
}


// FOR REFERENCE
// 
	// for {
	// 	fmt.Println("Select an action:")
	// 	fmt.Println("1. Retrieve an entry")
	// 	fmt.Println("2. Add an entry")
	// 	fmt.Println("3. Remove an entry")
	// 	fmt.Println("4. Update an entry")
	// 	fmt.Println("5. Save and quit")

	// 	var choice int
	// 	fmt.Scanln(&choice)

	// 	switch choice {
	// 	case 1:
	// 		fmt.Println("Enter CredId:")
	// 		var id int
	// 		fmt.Scanln(&id)
	// 		entry, err := manager.GetEntry(id)
	// 		if err != nil {
	// 			fmt.Println("Error retrieving entry:", err)
	// 		} else {
	// 			fmt.Println(entry)
	// 		}
	// 	case 2:
	// 		entry, err := GetNewEntry()
	// 		if err != nil {
	// 			fmt.Println("Error creating new entry:", err)
	// 		} else {
	// 			manager.AddEntry(entry)
	// 		}
	// 	case 3:
	// 		fmt.Println("Enter CredId:")
	// 		var id int
	// 		fmt.Scanln(&id)
	// 		err := manager.RemoveEntry(id)
	// 		if err != nil {
	// 			fmt.Println("Error removing entry:", err)
	// 		}
	// 	case 4:
	// 		fmt.Println("Enter CredId:")
	// 		var id int
	// 		fmt.Scanln(&id)
	// 		entry, err := manager.GetEntry(id)
	// 		if err != nil {
	// 			fmt.Println("Error retrieving entry:", err)
	// 		} else {
	// 			fmt.Println("Enter attribute to update:")
	// 			fmt.Println("1. Website")
	// 			fmt.Println("2. Username")
	// 			fmt.Println("3. Password")
	// 			var attr int
	// 			fmt.Scanln(&attr)
	// 			switch attr {
	// 			case 1:
	// 				fmt.Println("Enter new website:")
	// 				var website string
	// 				fmt.Scanln(&website)
	// 				entry.Website = website
	// 			case 2:
	// 				fmt.Println("Enter new username:")
	// 				var username string
	// 				fmt.Scanln(&username)
	// 				entry.Username = username
	// 			case 3:
	// 				fmt.Println("Enter new password:")
	// 				var password string
	// 				fmt.Scanln(&password)
	// 				entry.Password = password
	// 			default:
	// 				fmt.Println("Invalid attribute")
	// 				continue
	// 			}
	// 			err = manager.UpdateEntry(entry)
	// 			if err != nil {
	// 				fmt.Println("Error updating entry:", err)
	// 			}
	// 		}
	// 	case 5:
	// 		data, err := json.MarshalIndent(manager, "", "    ")
	// 		if err != nil {
	// 			fmt.Println("Error saving passwords:", err)
	// 			return
	// 		}
	// 		err = ioutil.WriteFile("passwords.json", data, 0644)
	// 		if err != nil {
	// 			fmt.Println("Error saving passwords:", err)
	// 			return
	// 		}
	// 		fmt.Println("Passwords saved. Goodbye!")
	// 		return
	// 	default:
	// 		fmt.Println("Invalid choice")
	// 	}
	// }

