package appNew

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func RunNewApp() (w fyne.Window, a fyne.App) {

	a = app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w = a.NewWindow("Примеры кода")
	w.Resize(fyne.NewSize(800, 800))
	w.SetIcon(theme.DocumentCreateIcon())

	return w, a
}
