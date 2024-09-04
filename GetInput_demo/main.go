package main

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// 综合cobra使用
func main() {
	var name string
	prompt := &survey.Input{
		Message: "What is your name",
	}
	survey.AskOne(prompt, &name)
	fmt.Printf("Hello, %s!\n", name)

	// menu
	var color string
	prompt2 := &survey.Select{
		Message: "Choose a color:",
		Options: []string{"red", "blue", "green", "yellow"},
	}
	survey.AskOne(prompt2, &color)
	fmt.Printf("You choose, %s!\n", color)
	// confirm
	var confirm bool
	prompt3 := &survey.Confirm{
		Message: "Do you want to proceed?",
	}
	survey.AskOne(prompt3, &confirm)

	if confirm {
		fmt.Println("Proceeding...")
	} else {
		fmt.Println("Operation canceled.")
	}
	// multi choose
	var languages []string
	prompt4 := &survey.MultiSelect{
		Message: "What programming languages do you know?",
		Options: []string{"Go", "Python", "JavaScript", "Rust"},
	}
	survey.AskOne(prompt4, &languages)

	fmt.Printf("You selected: %v\n", languages)

	// password
	var password string
	prompt5 := &survey.Password{
		Message: "Enter your password:",
	}
	survey.AskOne(prompt5, &password)

	fmt.Println("Password received.")
	// verify
	var email string
	prompt6 := &survey.Input{
		Message: "Enter your email:",
	}
	survey.AskOne(prompt6, &email, survey.WithValidator(survey.Required), survey.WithValidator(func(val interface{}) error {
		if str, ok := val.(string); ok {
			if !strings.Contains(str, "@") {
				return fmt.Errorf("invalid email address")
			}
		}
		return nil
	}))

	fmt.Printf("Email entered: %s\n", email)
}
