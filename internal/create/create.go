package create

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"notepad/internal/database"
)

func NewNoteCreate(a fyne.App, db database.DB) fyne.Window {

	ee := widget.NewEntry()
	createWindow := a.NewWindow("Создать новую запись")
	createWindow.Resize(fyne.NewSize(500, 150))
	createWindow.SetIcon(theme.FileIcon())
	createWindow.CenterOnScreen()
	createWindow.SetContent(
		container.NewVBox(
			widget.NewLabel("Введите название записи"),
			ee,
			widget.NewButton("Создать", func() {
				_ = db.CreateNewTable(ee.Text)
				createWindow.Hide()
			}),
		),
	)

	return createWindow
}
