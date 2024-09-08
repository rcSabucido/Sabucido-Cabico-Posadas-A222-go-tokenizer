# Tokenizer in Go

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

## Collaborators
This project is developed by:

- Tracy Angelo Posadas
- Karsten Gabriel Cabico
- Ryz Clyd Sabucido 
