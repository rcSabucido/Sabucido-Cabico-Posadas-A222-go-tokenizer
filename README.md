# go-tokenizer 

## Overview
This project is a **Tokenizer** written in Go, designed to process and analyze input text by breaking it into tokens such as words, punctuation, and special characters. The tokenizer identifies alphanumeric tokens, punctuation, and end-of-line markers, providing an efficient solution for text analysis.

## Features
- **Accurate Tokenization**: Accurately separates the text into tokens for easy parsing.
- **Token Classification**: Correctly identifies and classifies the type of each token, including:
  - **Word**: Alphabetic sequences (e.g., `hello`, `GoLang`)
  - **Punctuation**: Characters such as commas, periods, and others (e.g., `,`, `.`)
  - **Delimeter**: Uses a Hyphen symbol as the Delimeter (e.g., `-`, `Go-Lang`)
  - **Number**: Numeric sequences (e.g., `123`, `456.78`)
  - **Alphanumeric**: Sequences containing both letters and numbers (e.g., `abc123`)
  - **End of Line**: Handles end-of-line markers and returns appropriate tokens for line breaks.
- **TUI Implementation**: The tokenizer features a simple and intuitive Text User Interface (TUI) implemented in Go for ease of use and interaction.

## Dependencies Used
- [**Huh?**](https://github.com/charmbracelet/huh) - is a simple and powerful TUI library for Go
- [**lipgloss**](https://github.com/charmbracelet/lipgloss) - is a styling library for Go

### How to install
<p>Inside your Go directory, run these commands:</p>
<p><code>go get "github.com/charmbracelet/huh"</code></p>
<code>go get "github.com/charmbracelet/huh/spinner"</code>


## Collaborators
This project is developed by:

- [Tracy Angelo Posadas](https://github.com/y0b1)
- [Karsten Gabriel Cabico](https://github.com/POKNUTDONUT)
- [Ryz Clyd Sabucido](https://github.com/rcSabucido)
