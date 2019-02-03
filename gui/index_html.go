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

  <!-- Meta information. -->
  <meta property="og:title" content="Beatnik" />
  <meta property="og:description" content="Text your drum track" />
  <meta property="og:type" content="website" />
  <meta property="og:image" content="https://www.clipartmax.com/png/middle/173-1731477_drum-kit-sprite-005-cartoon-drum-kit-png.png" />

  <!-- Bootstrap stuff. -->
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>

  <!-- MIDI player. -->
  <script type='text/javascript' src='//www.midijs.net/lib/midi.js'></script>

  <script>
    function compile() {
	  let src = $("#src").val();
	  $.ajax({
	    url: "/compile",
		data: {src: src},
		success: function(result) {
          MIDIjs.play("/midi/" + result);
        },
	  });
	}

	function download() {
	  let src = $("#src").val();
	  $.ajax({
	    url: "/compile",
		data: {src: src},
		success: function(result) {
          top.location.href = "/midi/" + result;
        },
	  });
	}
  </script>
</head>
<body>

<div class="container">
  <h1>Beatnik <small>text your drum track</small></h1>
  <p></p>

  <h3>Source</h3>
    <div class="form-group">
      <textarea id="src" class="form-control" rows="10"
	   style="font-family: monospace">bpm:120

# Try this example!
HC,K. HC.   HC,S. HC.
HC,K. HC,K. HC,S. HC.
HC,K. HC.   HC,S. HC.
S+.. S.. S.. S.. T3+.. T3.. T4.. T4..
K,C2~~</textarea>
    </div>
	<button class="btn btn-default" type="button"
	    data-toggle="tooltip" title="Play" onclick="compile()">
	  <span class="glyphicon glyphicon-play" aria-hidden="true"></span>
	</button>
	<button class="btn btn-default" type="button"
	    data-toggle="tooltip" title="Stop" onclick="MIDIjs.stop()">
	  <span class="glyphicon glyphicon-stop" aria-hidden="true"></span>
	</button>
    <button class="btn btn-default" type="button"
	    data-toggle="tooltip" title="Download" onclick="download()">
	  <span class="glyphicon glyphicon-download-alt" aria-hidden="true"></span>
	</button>

  <hr>
  <p><a href="https://github.com/fluhus/beatnik/blob/master/TUTORIAL.md"
      target="_blank">Language tutorial</a></p>
  <p>Beatnik is <a href="https://github.com/fluhus/beatnik"
                 target="_blank">open source</a></p>

</div>

</body>
</html>
`
