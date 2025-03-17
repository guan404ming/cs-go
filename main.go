package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/guan404ming/cs-go/cmd"
)

func main() {
	// If there are command line arguments, use cmd.Execute
	if len(os.Args) > 1 {
		if err := cmd.Execute(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	// Read commands from standard input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Parse command and arguments
		parts := strings.SplitN(line, " ", 2)
		command := parts[0]

		var args []string
		if len(parts) > 1 {
			// Process argument part
			argPart := parts[1]
			var currentArg strings.Builder
			inQuote := false

			for i := 0; i < len(argPart); i++ {
				char := argPart[i]

				if char == '\'' {
					if inQuote {
						// End quote
						args = append(args, currentArg.String())
						currentArg.Reset()
						inQuote = false
					} else {
						// Start quote
						inQuote = true
					}
				} else if char == ' ' && !inQuote {
					// Space outside quotes, separate arguments
					if currentArg.Len() > 0 {
						args = append(args, currentArg.String())
						currentArg.Reset()
					}
				} else {
					// Normal character
					currentArg.WriteByte(char)
				}
			}

			// Process the last argument
			if currentArg.Len() > 0 {
				args = append(args, currentArg.String())
			}
		}

		// Set os.Args for cmd.Execute to use
		os.Args = []string{os.Args[0], command}
		os.Args = append(os.Args, args...)

		if err := cmd.Execute(); err != nil {
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
