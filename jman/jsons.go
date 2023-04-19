package jman

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/dharmit009/gopass/passutil"
)

const jsfile = "password.json"

type Entry struct {
	ID       int    `json:"id"`
	Website  string `json:"website"`
	Username string `json:"username"`
	Password []byte `json:"password"`
}

type Jman struct {
	Entries []Entry `json:"entries"`
}

func NewJman() (*Jman, error) {
	j := &Jman{
		Entries: []Entry{},
	}
	if fileExists(jsfile) {
		err := j.Load()
		if err != nil {
			return nil, err
		}
	} else {
		err := createFile(jsfile)
		if err != nil {
			return nil, err
		}
	}
	return j, nil
}

func (j *Jman) Load() error {
	data, err := ioutil.ReadFile(jsfile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &j)
	if err != nil {
		return err
	}
	return nil
}

func (j *Jman) Save() error {
	data, err := json.Marshal(j)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(jsfile, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (j Jman) GetEntries() ([]Entry, error) {
	entries := []Entry{}
	for _, e := range j.Entries {
		entries = append(entries, e)
	}
	return entries, nil
}

func (j *Jman) AddEntry(website, username, password, mpass string) error {
	id := 0
	for _, e := range j.Entries {
		if e.ID > id {
			id = e.ID
		}
	}
	encpasswd, _ := passutil.Encrypt([]byte(password), mpass)
	entry := Entry{
		ID:       id + 1,
		Website:  website,
		Username: username,
		Password: encpasswd,
	}
	j.Entries = append(j.Entries, entry)
	return j.Save()
}

func (j *Jman) RemoveEntry(id int) error {
	found := false
	for i, e := range j.Entries {
		if e.ID == id {
			copy(j.Entries[i:], j.Entries[i+1:])
			j.Entries = j.Entries[:len(j.Entries)-1]
			found = true
			break
		}
	}
	if !found {
		return errors.New("entry not found")
	}
	return j.Save()
}

// func (j *Jman) UpdateEntry(id int, website, username, password, mpass string) error {
// 	for _, e := range j.Entries {
// 		if e.ID == id {
// 			encpasswd, _ := passutil.Encrypt([]byte(password), mpass)
// 			e.Website = website
// 			e.Username = username
// 			e.Password = encpasswd
// 			return j.Save()
// 		}
// 	}
// 	return errors.New("entry not found")
// }
func (j *Jman) UpdateEntry(id int, website, username, password, mpass string) error {
	for i, e := range j.Entries {
		if e.ID == id {
			encpasswd, _ := passutil.Encrypt([]byte(password), mpass)
			j.Entries[i].Website = website
			j.Entries[i].Username = username
			j.Entries[i].Password = encpasswd
			return j.Save()
		}
	}
	return errors.New("entry not found")
}


func fileExists(jsfile string) bool {
	_, err := os.Stat(jsfile)
	return err == nil
}

func createFile(jsfile string) error {
	file, err := os.Create(jsfile)
	if err != nil {
		return err
	}
	defer file.Close()
	j := &Jman{
		Entries: []Entry{},
	}
	data, err := json.Marshal(j)
	if err != nil {
		return err

	}
	err = ioutil.WriteFile(jsfile, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (j *Jman) GetEntry(id int) (Entry, error) {
	for _, e := range j.Entries {
		if e.ID == id {
			return e, nil
		}
	}
	return Entry{}, errors.New("entry not found")
}

func (j *Jman) GetEntryPassword(id int, mpass string) ([]byte, error) {
	for _, e := range j.Entries {
		if e.ID == id {
			password, err := passutil.Decrypt(e.Password, mpass)
			if err != nil {
				return nil, err
			}
			return password, nil
		}
	}
	return nil, errors.New("entry not found")
}

func (j *Jman) GetEntryByUsername(username string) (Entry, error) {
	for _, e := range j.Entries {
		if e.Username == username {
			return e, nil
		}
	}
	return Entry{}, errors.New("entry not found")
}

func (j *Jman) GetEntryByWebsite(website string) (Entry, error) {
	for _, e := range j.Entries {
		if e.Website == website {
			return e, nil
		}
	}
	return Entry{}, errors.New("entry not found")
}

func (j *Jman) GetEntryById(id int) (Entry, error) {
	for _, e := range j.Entries {
		if e.ID == id {
			return e, nil
		}
	}
	return Entry{}, errors.New("entry not found")
}
