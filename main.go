package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"os"
	"strings"
	"time"
)

// User struct to store user input from the form
type User struct {
	Input string
}

func main() {
	var user User

	// Define a form using the 'huh' package with two groups: a title note and an input field
	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title(lipgloss.NewStyle().Background(lipgloss.Color("212")).Foreground(lipgloss.Color("#ffffff")).Render("Go Tokenizer")). // Style for title
			Description("Welcome to _Go Tokenizer_!\n\nThis project was made by Sabucido, Cabico, and Posadas\n\n").                   // Description
			Next(true).                                                                                                                // Enable the "Next" button
			NextLabel("Next"),                                                                                                         // Label for the "Next" button
		),

		// Group to capture input from the user
		huh.NewGroup(huh.NewInput().
			Value(&user.Input).                                                                                                  // Bind the user input to the Input field in the User struct
			Title("Enter your string:").                                                                                         // Title for the input
			Placeholder("this-is-an-example.").                                                                                  // Input placeholder hint
			Description(lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render("Please use a hyphen(-) as a delimiter.")), // Description for user guidance
		),
	)

	// Run the form and capture any error
	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh: ", err) // Print error if form fails
		os.Exit(1)                  // Exit if error occurs
	}

	// Function to simulate a delay for preparing the string
	prepareString := func() {
		time.Sleep(2 * time.Second) // Simulate a 2-second processing delay
	}

	// Create a spinner to show progress while the string is being "prepared"
	_ = spinner.New().Title("Preparing your tokenized string...").Action(prepareString).Run()

	// Tokenize and display the result
	{
		var sb strings.Builder // String builder to efficiently build the output
		keyword := func(s string) string {
			// Styling function for keywords (bold and colored)
			return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212")).Render(s)
		}

		// Tokenize the user input string (assuming Tokenize is defined elsewhere)
		tokens := Tokenize(user.Input)

		// Format the output string and write to the string builder
		_, _ = fmt.Fprintf(&sb, "%s\n\n%s", keyword("TOKENIZER RESULTS:"), output(tokens))

		// Display the final result with border styling
		fmt.Println(
			lipgloss.NewStyle().
				Width(60).                              // Set output box width
				BorderStyle(lipgloss.RoundedBorder()).  // Apply a rounded border
				BorderForeground(lipgloss.Color("63")). // Set border color
				Padding(1, 2).                          // Set padding inside the box
				Render(sb.String()),                    // Render the built string inside the styled box
		)
	}
}
