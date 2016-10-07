package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/hoisie/mustache"
)

func readJSON(context map[string]interface{}) {

	var parsed interface{}
	for i, arg := range os.Args {
		if arg != "--json" {
			continue
		}

		if i >= len(os.Args)-1 {
			fatalError(fmt.Errorf("The --json option requires a JSON string."))
		}

		if parsed != nil {
			fatalError(fmt.Errorf("The --json option can only be provided once."))
		}

		err := json.Unmarshal([]byte(os.Args[i+1]), &parsed)
		fatalError(err)
	}

	if parsed == nil {
		return
	}

	jMap, ok := parsed.(map[string]interface{})
	if !ok {
		fatalError(fmt.Errorf("JSON string must be an object."))
	}

	// Merge the JSON values into the model.

	for k, v := range jMap {
		context[k] = v
	}

}

func main() {
	// Read the template from stdin.

	source, err := ioutil.ReadAll(os.Stdin)
	fatalError(err)

	template, err := mustache.ParseString(string(source))
	fatalError(err)

	// Load the context/model from environment variables.

	context := make(map[string]interface{}, len(os.Environ()))
	for _, entry := range os.Environ() {
		split := strings.SplitN(entry, "=", 2)
		lines := strings.Split(split[1], "\n")
		if len(lines) > 1 {
			context[split[0]] = lines
		} else {
			context[split[0]] = split[1]
		}
	}

	// Read JSON model from the command line.
	readJSON(context)

	// Add key=value pairs from command-line arguments to the model.

	for _, arg := range os.Args[1:] {
		split := strings.SplitN(arg, "=", 2)
		if len(split) == 2 {
			fmt.Printf("\"%s\" => \"%s\"\n", split[0], split[1])
			context[split[0]] = split[1]
		}
	}

	// Render the template to stdout.

	fmt.Print(template.Render(context))
}

func fatalError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
