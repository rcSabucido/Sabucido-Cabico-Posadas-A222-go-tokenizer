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

type User struct {
	Input string
}

func main() {
	var user User
	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title(lipgloss.NewStyle().Background(lipgloss.Color("212")).Foreground(lipgloss.Color("#ffffff")).Render("Go Tokenizer")).
			Description("Welcome to _Go Tokenizer_!\n\nThis project was made by Sabucido, Cabico, and Posadas\n\n").
			Next(true).
			NextLabel("Next"),
		),

		huh.NewGroup(huh.NewInput().
			Value(&user.Input).
			Title("Enter your string:").
			Placeholder("this-is-an-example.").
			Description(lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render("Please use a hyphen(-) as a delimiter.")),
		),
	)

	err := form.Run()

	if err != nil {
		fmt.Println("Uh oh: ", err)
		os.Exit(1)
	}

	prepareString := func() {
		time.Sleep(2 * time.Second)
	}

	_ = spinner.New().Title("Preparing your tokenized string...").Action(prepareString).Run()

	{
		var sb strings.Builder
		keyword := func(s string) string {
			return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212")).Render(s)
		}
		tokens := Tokenize(user.Input)

		_, _ = fmt.Fprintf(&sb, "%s\n\n%s", keyword("TOKENIZER RESULTS:"), output(tokens))

		fmt.Println(
			lipgloss.NewStyle().
				Width(60).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render(sb.String()),
		)
	}
}
