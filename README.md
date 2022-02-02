# Beatnik

Beatnik is a **drum programming language**.

## What Does it Do?

With Beatnik you can type **text** which describes your beat, and Beatnik will interpret it and turn in into **MIDI**.

## Getting Started

1. Go to the (experimental) [demo page](https://beatnik-ihwgh.ondigitalocean.app/).
2. Read the [tutorial](https://github.com/fluhus/beatnik/blob/master/TUTORIAL.md).
3. Start making beats!

## Running Locally

1. Go to [releases](https://github.com/fluhus/beatnik/releases) and download the relevant file. It contains two executables: `btnk` and `gui`.
2. Read the [tutorial](https://github.com/fluhus/beatnik/blob/master/TUTORIAL.md).
3. Start making beats using `btnk` with the command line, or `gui` with your browser.

### btnk

`btnk` is a command line compiler that turns Beatnik code (typically with .btn suffix) into MIDI files.

From the command line run:

```
btnk file1.btn file2.btn
```

`btnk` will create file1.btn.mid and file2.btn.mid.

### gui

`gui` is a graphical web interface. Once you run it you can open your browser at the address it shows on the terminal.

This is an experimental thing I did for demonstration purposes. Would love to get feedback on where to take it further!

## Getting The Code

To download the code and build your own:

1. Install the [go](https://golang.org/) compiler (add it to your path).
2. Set the GOPATH environment variable to you code directory.
3. Run:  
   ```
   go get github.com/fluhus/beatnik
   go install github.com/fluhus/beatnik/btnk github.com/fluhus/beatnik/gui
   ```
4. Find the compiled binaries under `bin` and source code under `src`.

Happy coding!

## Samples

Isolated samples:
[Song 2](https://drive.google.com/file/d/1CVjNAYApnMNlhYOlAlGLJCB7WGvBDJO5/preview),
[50 Ways](https://drive.google.com/file/d/1qEw-5D6pLfflZBiCXrj60oeYwtHhJ1h_/preview)

Beatnik is used on my project [Out of Bounds](https://www.youtube.com/channel/UCAsR8ow5yv5dz4yZ6ZqsrTQ). Check it out for mixed samples.
