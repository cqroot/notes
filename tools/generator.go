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

// ToNoteCategory 从类似 `01【Linux】笔记名.md` 这样形式的字符串中提取 `Linux` 关键字
// 如何传入的字符串不是该种形式，返回源字符串
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

// ToNoteName 从类似 `01【Linux】笔记名.md` 这样形式的字符串中提取 `笔记名` 关键字
// 如何传入的字符串不是该种形式，返回源字符串
func ToNoteName(note string) string {
	if len(note) < 2 {
		return note
	}
	note = note[2:]

	endIdx := strings.Index(note, "】")
	if endIdx == -1 {
		return note
	}

	note = note[endIdx+len("】"):]
	note, _ = strings.CutSuffix(note, ".md")
	return note
}

// ToNoteLink 返回文件 "./docs/{note}" URL 编码后的字符串
func ToNoteLink(note string) string {
	return url.QueryEscape(fmt.Sprintf("./docs/%s", note))
}

var funcMap = template.FuncMap{
	"ToNoteName": ToNoteName,
	"ToNoteLink": ToNoteLink,
}

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

	outputFile, err := os.OpenFile("./README.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o666)
	CheckErr(err)
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, GetNoteData())
	CheckErr(err)
}
