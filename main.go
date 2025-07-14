package main

import (
	"fmt"
	"mypasswords/auth"
	"mypasswords/commands"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if err := auth.Authenticate("mypasswords"); err != nil {
			fmt.Fprintln(os.Stderr, "Authentication failed:", err)
			os.Exit(1)
		}
		// Tryb CLI z argumentami
		switch os.Args[1] {
		case "add":
			// Dodaj hasło: password-manager add <service> <user> <password>
			if len(os.Args) != 5 {
				fmt.Println("Usage: add <service> <user> <password>")
				return
			}
			fmt.Printf("Dodano hasło dla %s użytkownika %s\n", os.Args[2], os.Args[3])
			// Tu dodaj logikę zapisu
		case "list":
			fmt.Println("Wyświetlam listę haseł...")
			// Tu dodaj logikę wyświetlania
		default:
			fmt.Println("Nieznana komenda:", os.Args[1])
		}
	} else {
		// Tryb interaktywny (TUI)
		commands.Execute()
	}
}
