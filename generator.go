package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Data struct {
	Categories []string
	Notes      map[string][]string
}

func ToNoteName(note, category string) string {
	var s string
	s, _ = strings.CutPrefix(note, fmt.Sprintf("【%s】", category))
	s, _ = strings.CutSuffix(s, ".md")
	return s
}

func ToNoteLink(note string) string {
	return url.QueryEscape(fmt.Sprintf("./docs/%s", note))
}

var (
	data = Data{
		Categories: []string{"Linux", "Go"},
		Notes: map[string][]string{
			"Linux": {
				"【Linux】命令速查表.md",
			},
			"Go": {
				"【Go】slice.md",
			},
		},
	}
	funcMap = template.FuncMap{
		"ToNoteName": ToNoteName,
		"ToNoteLink": ToNoteLink,
	}
)

func main() {
	tmplPath := "./README.md.tmpl"
	tmpl, err := template.New(filepath.Base(tmplPath)).Funcs(funcMap).ParseFiles(tmplPath)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.OpenFile("./README.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	tmpl.Execute(outputFile, data)
}
