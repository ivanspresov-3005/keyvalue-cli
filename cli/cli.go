package main

import (
	"fmt"
	"os"
	"strings"
	"testtask.com/kvlogic"
)

// doOperation performs the specified operation (set, get, delete, exit) based on user input.
// It interacts with the user to gather necessary information for each operation and calls
// the corresponding methods from the kvlogic.KeyValue instance.
func doOperation(keyvalue *kvlogic.KeyValue, operation string) {
	var key, value string
	var ttl int
	if operation != "exit" {
		fmt.Print("Please enter key: ")
		fmt.Scanln(&key)
		key = strings.TrimSpace(key)
		fmt.Println("") // Separate input prompt from subsequent output
	}

	switch operation {
	case "set":
		// For 'set' operation, prompt user to enter value and ttl
		fmt.Print("Please enter value: ")
		fmt.Scanln(&value)
		value = strings.TrimSpace(value)
		fmt.Println("")

		fmt.Print("Please enter ttl: ")
		fmt.Scanln(&ttl)

		// Validate input and call the set method from kvlogic.KeyValue
		if key == "" || value == "" {
			fmt.Println("Both key and value are required for set command")
			return
		}

		if ttl < 0 {
			fmt.Println("TTL cannot be negative")
			return
		}

		keyvalue.Set(key, value, ttl)
		fmt.Printf("The new entity with key=%s,value=%s and ttl=%d was added to db\n", key, value, ttl)

	case "get":
		// For 'get' operation, prompt user to enter key, retrieve value, and print it
		if key == "" {
			fmt.Println("Key is required for get command")
			return
		}

		value := keyvalue.Get(key)
		fmt.Printf("Value for key %s: %s\n", key, value)

	case "delete":
		// For 'delete' operation, prompt user to enter key and call the delete method
		if key == "" {
			fmt.Println("Key is required for delete command")
			return
		}

		keyvalue.Delete(key)
		fmt.Printf("Value for key %s was deleted\n", key)

	case "exit":
		// Exit the program if the user selects 'exit'
		fmt.Println("Exiting the program.")
		os.Exit(0)

	default:
		// Handle invalid operation input
		fmt.Println("Invalid operation. Please use set, get, delete, or exit.")
	}
	fmt.Println() // Add a newline after each operation for better separation
}

// main function is the entry point of the program.
// It initializes a kvlogic.KeyValue instance and repeatedly prompts the user
// to select an operation (set, get, delete, exit) and performs the chosen operation.
func main() {
	// initiate keyvalue instance
	keyvalue := kvlogic.NewKeyValue()
	var operation string

	// Continuously prompt user for operation until 'exit' is chosen
	for {
		fmt.Print("Please select operation (set, get, delete, exit): ")
		fmt.Scanln(&operation)
		fmt.Println("") // Separate input prompt from subsequent output
		operation = strings.TrimSpace(operation)
		doOperation(keyvalue, operation)
	}
}
