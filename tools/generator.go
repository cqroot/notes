package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Data struct {
	Categories []string
	Notes      map[string][]string
}

func ToNoteCategory(note string) string {
	if len(note) < 2 {
		return note
	}
	note = note[2:]

	if !strings.HasPrefix(note, "【") {
		return note
	}

	endIdx := strings.Index(note, "】")
	if endIdx == -1 {
		return note
	}

	return note[len("【"):endIdx]
}

func ToNoteName(note, category string) string {
	if len(note) < 2 {
		return note
	}
	note = note[2:]

	var s string
	s, _ = strings.CutPrefix(note, fmt.Sprintf("【%s】", category))
	s, _ = strings.CutSuffix(s, ".md")
	return s
}

func ToNoteLink(note string) string {
	return url.QueryEscape(fmt.Sprintf("./docs/%s", note))
}

var (
	funcMap = template.FuncMap{
		"ToNoteName": ToNoteName,
		"ToNoteLink": ToNoteLink,
	}
)

func GetNoteData() Data {
	data := Data{
		Categories: []string{"Linux", "Go", "Tools", "Misc"},
		Notes: map[string][]string{
			"Linux": {},
			"Go":    {},
			"Tools": {},
			"Misc":  {},
		},
	}

	entries, err := os.ReadDir("./docs/")
	CheckErr(err)

	for _, entry := range entries {
		name := entry.Name()
		if strings.HasPrefix(name, "_") {
			continue
		}

		data.Notes[ToNoteCategory(name)] = append(data.Notes[ToNoteCategory(name)], name)
	}

	return data
}

func main() {
	tmplPath := "./tools/README.md.tmpl"
	tmpl, err := template.New(filepath.Base(tmplPath)).Funcs(funcMap).ParseFiles(tmplPath)
	CheckErr(err)

	outputFile, err := os.OpenFile("./README.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	CheckErr(err)
	defer outputFile.Close()

	tmpl.Execute(outputFile, GetNoteData())
}
