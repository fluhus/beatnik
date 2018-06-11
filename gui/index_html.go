// Auto-generated with gen.go from index.html.

package main

import (
	"html/template"
	"io/ioutil"
)

// Template is safe to use because it was tested during generation.
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

var indexPageTemplate = template.Must(template.New("index").Parse(indexPageContent))

const indexPageContent = `<!DOCTYPE html>
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
42,36. 42. 42,38. 42.
42,36. 42. 42,38. 42.
42,36. 42. 42,38. 42.
38+.. 38.. 38.. 38.. 43+.. 43.. 41.. 41..
36,57~~</textarea>
    </div>
    <button type="submit" class="btn btn-primary">Get MIDI!</button>
  </form>

  <hr>

  <button class="btn btn-success" data-toggle="collapse" data-target="#tutorial">Tutorial</button>
  <div id="tutorial" class="collapse">
    <h3>Beatnik Language Tutorial</h3>
	<hr>
    <h4>Notes</h4>
    <p>A note is a single hit on multiple drums at the same time.
	A note has drum numbers, velocities and duration.</p>
    <p>Example:</p>
    <h4>
      <samp>
        <span class="text-primary">38</span>,<span class="text-primary">46</span><span
        class="text-danger">+</span><span class="text-success">..</span>
      </samp>
    </h4>
    <ol>
      <li>
        <span class="text-primary"><b>Drum numbers:</b></span>
        Drum numbers (midi standard) seperated by commas.
      </li>
      <li>
        <span class="text-danger"><b>Drum velocity:</b></span>
        Each drum number can be followed by +'s or -'s for velocity.
        <ul>
          <li><b>[nothing]</b> Forte</li>
          <li><b>++</b> Fortississimo</li>
          <li><b>+</b> Fortissimo</li>
          <li><b>-</b> Mezzo-forte</li>
          <li><b>--</b> Mezzo-piano</li>
          <li><b>---</b> Piano</li>
          <li><b>----</b> Pianissimo</li>
          <li><b>-----</b> Pianississimo</li>
        </ul>
      </li>
	  <li>
        <span class="text-success"><b>Note duration:</b></span>
        Each note can be followed by .'s or ~'s for duration.
        <ul>
          <li><b>[nothing]</b> 1/4 bar</li>
          <li><b>~~</b> 1 bar</li>
          <li><b>~</b> 1/2 bar</li>
          <li><b>.</b> 1/8 bar</li>
          <li><b>..</b> 1/16 bar</li>
          <li><b>...</b> 1/32 bar</li>
          <li><b>....</b> 1/64 bar</li>
          <li><b>.....</b> 1/128 bar</li>
        </ul>
      </li>
    </ol>
	<hr>
	<h4>Comments</h4>
	<p>When a hash sign (#) appears, everything on it's right is ignored.
	Use this to annotate your beat!</p>
	<p><pre>aasdf</pre></p>
  </div>
</div>

</body>
</html>
`
