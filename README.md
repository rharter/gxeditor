# Golang Editor

This is some (very) simple experimentation with the [GXUI](https://github.com/google/gxui/) library. I've done exactly nothing new here, just some code diving and playing around.

As noted in the GXUI readme, this is a project in active development so things will change and break.

This project will let you open a text file and edit it. It's not particularly useful, since file saving isn't a think, and it probably can't even open really really big files because of the way I'm using `ioutils`. This is just a fun little project to play with GXUI.

![Screenshot](/images/screenshot.png?raw=true)

## Usage

I use a Mac, so I had to install glew.  I don't know how to do this on other systems.

    brew install glew

Get the libraries using `go get` (all may not be required, but what the hell).

    go get github.com/google/gxui
    go get github.com/google/gxui/drivers/gl
    go get github.com/google/gxui/math
    go get github.com/google/gxui/themes/dark

Get this project using `go get`

    go get github.com/rharter/gxeditor

Run the app using `go run`

    go run main.go -data="data" -file="main.go"

I even made an error dialog, playing around with (only slightly) more complex layouts.

    go run main.go -data="data"

