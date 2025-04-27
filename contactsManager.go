package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  uint8
}

const CONTACTS_FILE = "contacts.json"

var isRunning bool = true
var contacts []Person

func printMainMenu() {
	fmt.Println(`
Main Menu

1. Add new contact
2. List all contacts
3. Exit
	`)
}

func main() {
	contacts = TryLoadContactsFromFile()
	for isRunning {
		printMainMenu()
		var userChoice uint8
		if _, err := fmt.Scan(&userChoice); err != nil {
			fmt.Println("ERROR! ", err)
		}
		handleUserChoice(userChoice)
	}
}

func TryLoadContactsFromFile() []Person {
	bytes, err := os.ReadFile("contacts.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File %v doesn't exist\n", CONTACTS_FILE)
		} else {
			fmt.Println("Error loading contacts from file: ", CONTACTS_FILE, err)
		}
	} else {
		if err := json.Unmarshal(bytes, &contacts); err != nil {
			fmt.Println("Failed to deserialize contacts: ", err)
		}
	}
	return contacts
}

func handleUserChoice(userChoice uint8) {
	switch userChoice {
	case 1:
		HandleAddNewContact()
	case 2:
		HandleListAllContacts()
	case 3:
		HandleExit()
	default:
		fmt.Println("You entered an invalid selection. Try again!")
	}
}

func HandleExit() {
	fmt.Println("Exiting...")
	isRunning = false
}

func HandleListAllContacts() {
	if len(contacts) == 0 {
		fmt.Println("No contacts to list!")
	} else {
		for _, contact := range contacts {
			fmt.Printf("Name: %v, Age: %d\n", contact.Name, contact.Age)
		}
	}
}

func HandleAddNewContact() {
	var name string
	var age uint8

	fmt.Println("Enter a name: ")
	fmt.Scan(&name)

	fmt.Println("Enter an age: ")
	if _, err := fmt.Scan(&age); err != nil {
		fmt.Println("Invalid age entered", err)
	} else {
		contacts = append(contacts, Person{
			Name: name, Age: age,
		})
		SaveContacts(contacts)
	}
}

func SaveContacts(contacts []Person) {
	contactsJson, _ := json.MarshalIndent(contacts, "", "  ")
	file, err := os.OpenFile(CONTACTS_FILE, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}

	defer file.Close()

	if _, err := file.Write(contactsJson); err != nil {
		fmt.Println("Error writing contact to file", err)
	}
}
