package main

import (
	"flag"
	"io/ioutil"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

var data = flag.String("data", "data", "path to data folder")
var file = flag.String("file", "file", "file to edit")

func appMain(driver gxui.Driver) {
	t := dark.CreateTheme(driver)
	w := t.CreateWindow(800, 600, "Editor")

	e := t.CreateCodeEditor()
	w.AddChild(e)

	text, err := ioutil.ReadFile(*file)
	if err != nil {
		e := displayError(t, err.Error())
		e.OnClose(driver.Terminate)
		gxui.EventLoop(driver)
	}
	e.SetText(string(text))

	w.OnClose(driver.Terminate)
	gxui.EventLoop(driver)
}

func displayError(theme gxui.Theme, m string) gxui.Window {
	w := theme.CreateWindow(300, 100, "Error")

	ll := theme.CreateLinearLayout()
	ll.SetOrientation(gxui.Vertical)
	ll.SetHorizontalAlignment(gxui.AlignCenter)
	w.AddChild(ll)

	l := theme.CreateLabel()
	l.SetText(m)
	ll.AddChild(l)

	b := theme.CreateButton()
	b.SetText("Okay")
	b.OnClick(func(e gxui.MouseEvent) {
		w.Close()
	})
	ll.AddChild(b)

	return w
}

func main() {
	flag.Parse()
	gl.StartDriver(*data, appMain)
}
