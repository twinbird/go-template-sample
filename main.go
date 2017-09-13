package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderData はテンプレートへ引き渡す表示用のデータです
type RenderData struct {
	Title string
	Entry *Entry
}

// Entry は1記事に相当します
type Entry struct {
	Text string
}

var (
	// サンプルデータ
	sampleEntry *Entry      = &Entry{Text: "This is a sample entry"}
	sampleData  *RenderData = &RenderData{Title: "sample", Entry: sampleEntry}
	// テンプレートディレクトリ
	templatesDir string = "templates"
)

func main() {
	// admin:編集ページ
	http.HandleFunc("/admin/edit/someEntry", editAdminSomeEntry)
	// admin:参照ページ
	http.HandleFunc("/admin/refer/someEntry", referAdminSomeEntry)
	// public:参照ページ
	http.HandleFunc("/refer/someEntry", referSomeEntry)

	http.ListenAndServe(":8080", nil)
}

// admin用:編集ページ
func editAdminSomeEntry(w http.ResponseWriter, r *http.Request) {
	execTemplate(w, sampleData, "layout", "admin_menu", "entry_editor")
}

// admin用:参照ページ
func referAdminSomeEntry(w http.ResponseWriter, r *http.Request) {
	execTemplate(w, sampleData, "layout", "admin_menu", "entry_view")
}

// public用:参照ページ
func referSomeEntry(w http.ResponseWriter, r *http.Request) {
	execTemplate(w, sampleData, "layout", "public_menu", "entry_view")
}

// execTemplate はfilesのテンプレートからHTMLを構築して,
// wに対して書き込みます.
// HTML構築の際にはdataを利用します.
// filesで指定するテンプレートには必ず{{define "layout"}}された
// ものを1つだけ含む必要があります.
func execTemplate(w http.ResponseWriter, data interface{}, files ...string) {
	// 渡された引数からテンプレートパスの集合を作る
	var pathes []string
	for _, f := range files {
		p := fmt.Sprintf("%s/%s.html", templatesDir, f)
		pathes = append(pathes, p)
	}
	//上記で作ったパスの一覧を使ってテンプレートを作る
	template := template.Must(template.ParseFiles(pathes...))
	// layoutが必ず基点になるという事にする
	template.ExecuteTemplate(w, "layout", data)
}
