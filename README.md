# envstache

Command-line tool for rendering Mustache templates using environment variables.

## Usage

```
$ go get github.com/hoisie/mustache
$ go install github.com/tokenshift/envstache

$ echo "Test Value: {{TEST_VAL}}" | TEST_VAL="42" envstache
Test Value: 42

$ export TEST_VALUES="Test Value 1
> Test Value 2
> Test Value 3"
$ echo "{{#TEST_VALUES}}
>* {{.}}
>{{/TEST_VALUES}}" | envstache
* Test Value 1
* Test Value 2
* Test Value 3
```

**envstache** can also take `key=value` parameters on the command line; these
will be added (and override) environment variables:

```
$ export TEST_VALUE="testing"
$ echo "Test Value: {{TEST_VALUE}}" | envstache
Test Value: testing
$ echo "Test Value: {{TEST_VALUE}}" | envstache TEST_VALUE=hello
Test Value: hello
```

If you need to render structured/nested data, use the `--json` argument to
provide a string of JSON:

```
$ echo "Value: {{#some}}{{#nested}}{{value}}{{/nested}}{{/some}}" | \
  envstache --json '{"some": {"nested": {"value": "Got Here"}}}'
 Value: Got Here
 ```