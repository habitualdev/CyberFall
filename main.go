package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"math/rand"
	"time"
)

var lineContent = []string{}
var style = tcell.StyleDefault

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghi                jklmnopqrs              tuvwxyzABCD  EFGHIJKLMNOPQR           STUVWXYZ1234567890     ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func DrawDoNothing(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	if len(lineContent) == 0 {
		for i := 0; i < height; i++ {
			lineContent = append(lineContent, RandStringRunes(width*2))
		}
	} else {
		lineContent = append(lineContent[1:], RandStringRunes(width*2))
	}
	lines := height
	for i := 0; i < lines; i++ {
		lineRand := lineContent[i]
		for j := 0; j < width; j++ {
			screen.SetContent(j, y+i, rune(lineRand[j]), nil, style)
		}
	}
	return x, y, width, height
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			println(r)
			main()
		}
	}()
	style = style.Foreground(tcell.ColorDarkOliveGreen)
	style = style.Background(tcell.ColorDarkGreen)
	app := tview.NewApplication()
	box := tview.NewBox().SetBorder(true)
	box.SetBackgroundColor(tcell.ColorGreen)
	box.SetDrawFunc(DrawDoNothing)

	go func() {
		for {
			app.Draw()
			time.Sleep(200 * time.Millisecond)
		}
	}()

	box.GetRect()
	if err := app.SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
