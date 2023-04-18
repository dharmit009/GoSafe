package jman

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

const jsfile = "password.json"

type Entry struct {
	ID       int    `json:"id"`
	Website  string `json:"website"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Jman struct {
	Entries []*Entry `json:"entries"`
}

func NewJman() (*Jman, error) {
	j := &Jman{
		Entries: []*Entry{},
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

func (j *Jman) GetEntries() ([]Entry, error) {
	entries := []Entry{}
	for _, e := range j.Entries {
		entries = append(entries, *e)
	}
	return entries, nil
}

func (j *Jman) AddEntry(website, username, password string) error {
	id := 0
	for _, e := range j.Entries {
		if e.ID > id {
			id = e.ID
		}
	}
	entry := Entry{
		ID:       id + 1,
		Website:  website,
		Username: username,
		Password: password,
	}
	j.Entries = append(j.Entries, &entry)
	return j.Save()
}

func (j *Jman) RemoveEntry(id int) error {
	found := false
	for i, e := range j.Entries {
		if e.ID == id {
			copy(j.Entries[i:], j.Entries[i+1:])
			j.Entries[len(j.Entries)-1] = nil
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

func (j *Jman) UpdateEntry(id int, website, username, password string) error {
	for _, e := range j.Entries {
		if e.ID == id {
			e.Website = website
			e.Username = username
			e.Password = password
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
		Entries: []*Entry{},
	}
	data, err := json.Marshal(j)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
