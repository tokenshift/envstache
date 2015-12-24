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