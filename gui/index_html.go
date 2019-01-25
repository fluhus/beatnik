// Auto-generated with gen.go from index.html.

package main

import (
	"html/template"
	"io/ioutil"
)

// Default template is safe to use because it was tested during generation.
func indexPage(f string) (*template.Template, error) {
	if f == "" {
		return indexPageTemplate, nil
	}
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	return template.New("index").Parse(string(data))
}

var indexPageTemplate = template.Must(template.New("index").Parse(indexPageSrc))

const indexPageSrc = `<!DOCTYPE html>
<html lang="en">
<head>
  <title>Beatnik</title>
  <meta charset="utf-8">
  
  <!-- Bootstrap stuff. -->
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>
<body>

<div class="container">
  <h1>Beatnik <small>text your drum track</small></h1>
  <p></p>

  <h3>Source</h3>
  <form action="/midi" method="post" target="_blank">
    <div class="form-group">
      <textarea class="form-control" rows="10" name="src" style="font-family: monospace">bpm:120

# Try this example!
HC,K. HC.   HC,S. HC.
HC,K. HC,K. HC,S. HC.
HC,K. HC.   HC,S. HC.
S+.. S.. S.. S.. T3+.. T3.. T4.. T4..
K,C3~~</textarea>
    </div>
    <button type="submit" class="btn btn-primary">Get MIDI!</button>
  </form>

</div>

</body>
</html>
`
