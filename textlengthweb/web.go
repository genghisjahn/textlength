package main

import "github.com/genghisjahn/textlength"
import "fmt"
import "net/http"
import "html/template"
import "time"
import "strings"

var items []textlength.TextLengthItem

func main() {
	items, _ = textlength.BuildItems(10000)
	http.HandleFunc("/process/", processHandler)
	http.HandleFunc("/query/", queryHandler)
	http.ListenAndServe("localhost:8888", nil)
}

func processHandler(rw http.ResponseWriter, req *http.Request) {
	text := strings.TrimSpace(req.PostFormValue("text"))
	fmt.Printf("Processing text: {%v} at %v.\n", text, time.Now())
	result, _ := textlength.ProcessText(text, items)
	fmt.Fprintf(rw, "%v", result)
	fmt.Printf("Returned {%v} at %v.\n", result, time.Now())
}

func queryHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Request recieved for index.html at %v\n", time.Now())
	t, _ := template.ParseFiles("index.html")
	t.Execute(rw, t)
}
