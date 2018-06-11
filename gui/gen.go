//+build ignore

// Command gen generates go source from HTML pages.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func main() {
	flag.Parse()
	files := filesWithSuffix(".html")
	t := template.Must(template.New("").Funcs(template.FuncMap{
		"stringLiteral":   stringLiteral,
		"templateVarName": templateVarName,
		"constName":       constName,
		"funcName":        funcName,
	}).Parse(outputTemplate))

	for _, f := range files {
		fmt.Fprintln(os.Stderr, f)

		// Load file.
		data, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Fprintln(os.Stderr, "  failed to read:", err)
			os.Exit(2)
		}

		// Validate template.
		_, err = template.New("").Parse(string(data))
		if err != nil {
			fmt.Fprintln(os.Stderr, "  invalid template in file:", err)
			os.Exit(2)
		}

		// Create go source.
		base := f[:len(f)-len(".html")]
		b := bytes.NewBuffer(nil)
		t.Execute(b, map[string]string{
			"name":    base,
			"content": string(data),
		})

		// Go fmt.
		src, err := format.Source(b.Bytes())
		if err != nil {
			fmt.Fprintln(os.Stderr, "  generated source is invalid:", err)
			os.Exit(2)
		}

		ioutil.WriteFile(base+"_html.go", src, 0600)
		fmt.Fprintln(os.Stderr, "  success")
	}
}

// filesWithSuffix returns a list of file names in the current directory that
// have the given suffix.
func filesWithSuffix(s string) []string {
	var result []string
	files, _ := ioutil.ReadDir(".")
	for _, f := range files {
		if strings.HasSuffix(f.Name(), s) {
			result = append(result, f.Name())
		}
	}
	return result
}

// stringLiteral returns a go string literal for the given string.
func stringLiteral(s string) string {
	return "`" + strings.Replace(s, "`", "`+\"`\"+`", -1) + "`"
}

// templateVarName returns a variable name for the given HTML file prefix.
func templateVarName(s string) string {
	return camelCase(append(strings.Split(s, "_"), "page", "template"))
}

// constName returns a const name for the content of the given HTML file
// prefix.
func constName(s string) string {
	return camelCase(append(strings.Split(s, "_"), "page", "content"))
}

// funcName returns a getter function name for the given HTML file prefix.
func funcName(s string) string {
	return camelCase(append(strings.Split(s, "_"), "page"))
}

// camelCase returns a camel-case string made from concatenating the given parts.
func camelCase(parts []string) string {
	if len(parts) == 0 {
		return ""
	}
	result := strings.ToLower(parts[0])
	for _, s := range parts[1:] {
		if s == "" {
			continue
		}
		result += strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
	}
	return result
}

// outputTemplate is the source template for the output go source.
const outputTemplate = `// Auto-generated with gen.go from {{.name}}.html.

package main

import (
	"html/template"
	"io/ioutil"
)

// Template is safe to use because it was tested during generation.
func {{funcName .name}}(f string) (*template.Template, error) {
	if f == "" {
		return {{templateVarName .name}}, nil
	}
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return template.New("{{.name}}").Parse(string(data))
}

var {{templateVarName .name}} = template.Must(template.New("{{.name}}").Parse({{constName .name}}))

const {{constName .name}} = {{stringLiteral .content}}
`
