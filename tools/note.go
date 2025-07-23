package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	DOCS_DIR  = "./docs"
	TMPL_PATH = "./tools/README.md.tmpl"
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

// ToNoteLink 返回文件 "{DOCS_DIR}/{note}" URL 编码后的字符串
func ToNoteLink(note string) string {
	return url.QueryEscape(fmt.Sprintf("%s/%s", DOCS_DIR, note))
}

var funcMap = template.FuncMap{
	"ToNoteName": ToNoteName,
	"ToNoteLink": ToNoteLink,
}

func GetNoteData() Data {
	categories := make([]string, 0)
	notes := make(map[string][]string)

	entries, err := os.ReadDir(DOCS_DIR)
	CheckErr(err)

	for _, entry := range entries {
		name := entry.Name()
		if strings.HasPrefix(name, "_") {
			continue
		}

		currCategory := ToNoteCategory(name)
		if len(categories) == 0 || categories[len(categories)-1] != currCategory {
			categories = append(categories, currCategory)
			notes[currCategory] = make([]string, 0)
		}

		notes[currCategory] = append(notes[currCategory], name)
	}

	return Data{
		Categories: categories,
		Notes:      notes,
	}
}

func ListNotes() {
	data := GetNoteData()

	idLen := 2
	categoryLen := 0
	for _, category := range data.Categories {
		if len(category) > categoryLen {
			categoryLen = len(category)
		}
	}
	if categoryLen < 8 {
		categoryLen = 8
	}

	fmt.Println()
	fmt.Printf("     ID%s     CATEGORY%s     NAME\n", strings.Repeat(" ", idLen-2), strings.Repeat(" ", categoryLen-8))
	fmt.Printf("     ==%s     ========%s     ====\n", strings.Repeat(" ", idLen-2), strings.Repeat(" ", categoryLen-8))

	idx := 0
	for _, category := range data.Categories {
		for _, note := range data.Notes[category] {
			idx += 1
			fmt.Printf("     %-2d     %s%s     %s\n",
				idx, category, strings.Repeat(" ", categoryLen-len(category)), ToNoteName(note))
		}
	}
	fmt.Println()
}

func GenerateReadme() {
	tmpl, err := template.New(filepath.Base(TMPL_PATH)).Funcs(funcMap).ParseFiles(TMPL_PATH)
	CheckErr(err)

	outputFile, err := os.OpenFile("./README.md", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o666)
	CheckErr(err)
	defer func() {
		_ = outputFile.Close()
	}()

	err = tmpl.Execute(outputFile, GetNoteData())
	CheckErr(err)
}

func main() {
	args := os.Args
	if len(args) != 2 {
		CheckErr(fmt.Errorf("error parameter"))
	}

	switch args[1] {
	case "list":
		ListNotes()
	case "generate":
		GenerateReadme()
	}
}
