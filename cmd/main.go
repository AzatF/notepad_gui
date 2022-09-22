package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	appNew "notepad/app"
	"notepad/config"
	"notepad/internal/create"
	"notepad/internal/database"
	"notepad/internal/notelist"
	"notepad/pkg/logging"
)

var (
	e            *widget.Entry
	createWindow fyne.Window
	btn          *widget.Button
	buttonCreate *widget.Button
	list         *widget.List
)

func main() {

	logger := logging.GetLogger("trace")
	logger.Info("start")

	cfg := config.GetConfig("./etc/.env")
	db, _ := database.NewDataBase(cfg, logger)

	w, a := appNew.RunNewApp()

	database.TitleList, _ = db.GetTitles()

	e = widget.NewMultiLineEntry()
	e.SetMinRowsVisible(25)
	e.SetPlaceHolder("Новая запись...")

	selectList := widget.NewSelect(database.TitleList, func(s string) {
		text, _ := db.GetNote(s)
		e.SetText(text)
		e.Refresh()
	})
	selectList.PlaceHolder = "Выбрать запись"

	list = notelist.NewNoteList(e, db)

	createWindow = create.NewNoteCreate(a, db)

	buttonCreate = widget.NewButton("Создать новую запись", func() {
		createWindow.SetOnClosed(func() {
			createWindow = create.NewNoteCreate(a, db)
		})
		createWindow.Show()
		logger.Info("title list: ", database.TitleList)
	})

	btn = widget.NewButton("Сохранить", func() {

		err := db.WriteNote(database.TitleList[notelist.ID], e.Text)
		if err != nil {
			logger.Error(err)
		}
	})

	cont2 := container.NewVScroll(list)

	splitCont := container.NewHSplit(
		cont2,
		container.NewVBox(e, buttonCreate, btn),
	)
	splitCont.Offset = 0.2

	w.SetContent(splitCont)

	w.Show()
	w.SetMaster()
	a.Run()
}
