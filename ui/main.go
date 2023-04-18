package main

import (
	"fmt"
	"log"

	"github.com/dharmit009/gopass/ui/jman"
)

func main() {
	j, err := jman.NewJman()
	if err != nil {
		log.Fatal(err)
	}

	entries, err := j.GetEntries()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Entries: %v\n", entries)

	err = j.AddEntry("example.com", "user1", "password1")
	if err != nil {
		log.Fatal(err)
	}

	entries, err = j.GetEntries()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Entries: %v\n", entries)

	err = j.UpdateEntry(1, "example.com", "user2", "password2")
	if err != nil {
		log.Fatal(err)
	}

	entries, err = j.GetEntries()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Entries: %v\n", entries)

	err = j.RemoveEntry(1)
	if err != nil {
		log.Fatal(err)
	}

	entries, err = j.GetEntries()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Entries: %v\n", entries)
}

