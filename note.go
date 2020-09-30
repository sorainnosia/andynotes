package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne"
)

const (
	noteCountKey  = "noteCount"
	noteKeyFormat = "note%d"
)

type note struct {
	content string
}

func (n *note) title() string {
	if n.content == "" {
		return "Untitled"
	}
	return strings.SplitN(n.content, "\n", 2)[0]
}

type noteList struct {
	list  []*note
	prefs fyne.Preferences
}

func (l *noteList) add() *note {
	n := &note{}

	l.list = append([]*note{n}, l.list...)

	return n
}

func (l *noteList) loadNotes() {
	total := l.prefs.Int(noteCountKey)

	if total == 0 {
		return
	}

	for i := 0; i < total; i++ {
		key := fmt.Sprintf(noteKeyFormat, i)
		content := l.prefs.String(key)

		l.list = append(l.list, &note{content: content})
	}
}

func (l *noteList) remove(n *note) {
	if len(l.list) == 0 {
		return
	}

	for i, note := range l.list {
		if note != n {
			continue
		}

		if i == len(l.list)-1 {
			l.list = l.list[:i]
		} else {
			l.list = append(l.list[:i], l.list[i+1:]...)
		}
	}
}

func (l *noteList) saveNotes() {
	for i, n := range l.list {
		key := fmt.Sprintf(noteKeyFormat, i)

		l.prefs.SetString(key, n.content)
	}
	l.prefs.SetInt(noteCountKey, len(l.list))
}
