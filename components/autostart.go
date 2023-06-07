package components

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/emersion/go-autostart"
)

func Autostart() {
	// Path to your Go project executable
	executablePath := "/home/nishantk/Desktop/code/quotes/main"

	// Create an autostart application
	app := &autostart.App{
		Name: "MyProject",
		Exec: []string{executablePath},
		//Description: "My Go Project",
	}

	// Check if autostart is enabled
	if app.IsEnabled() {
		log.Println("Autostart is already enabled for MyProject")
		return

	} else {
		// Enable autostart
		if err := app.Enable(); err != nil {
			log.Fatal("Failed to enable autostart:", err)
		}
		log.Println("Autostart is now enabled for MyProject")
	}

	// Run the Go project
	cmd := exec.Command(executablePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = filepath.Dir(executablePath)
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to run Go project:", err)
	}
}
