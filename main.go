package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hoisie/mustache"
)

func main() {
	// Read the template from stdin.

	source, err := ioutil.ReadAll(os.Stdin)
	fatalError(err)

	template, err := mustache.ParseString(string(source))
	fatalError(err)

	// Load the context/model from environment variables.

	context := make(map[string]interface{}, len(os.Environ()))
	for _, entry := range(os.Environ()) {
		split := strings.SplitN(entry, "=", 2)
		lines := strings.Split(split[1], "\n")
		if len(lines) > 1 {
			context[split[0]] = lines
		} else {
			context[split[0]] = split[1]
		}
	}

	// Add key=value pairs from command-line arguments to the model.

	for _, arg := range(os.Args) {
		split := strings.SplitN(arg, "=", 2)
		if len(split) == 2 {
			fmt.Printf("\"%s\" => \"%s\"\n", split[0], split[1])
			context[split[0]] = split[1]
		}
	}

	// Render the template to stdout.

	fmt.Println(template.Render(context))
}

func fatalError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}