package notelist

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"notepad/internal/database"
)

var ID int

func NewNoteList(e *widget.Entry, db database.DB) *widget.List {

	ss := database.TitleList

	notesList := widget.NewList(
		func() int { return len(ss) },
		func() fyne.CanvasObject {

			return widget.NewLabel("title")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			text, _ := db.GetNote(ss[id])
			obj.(*widget.Label).SetText(ss[id])
			e.SetText(text)
		},
	)

	notesList.OnSelected = func(id widget.ListItemID) {
		text, _ := db.GetNote(ss[id])
		ID = id
		e.SetText(text)
	}

	return notesList
}
