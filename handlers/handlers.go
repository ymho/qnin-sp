package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ymho/qnin-sp/models"
)

// /のハンドラ
func DefaultHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// /helloのハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// /samlのハンドラ
// TODO: パスパラメータ、クエリパラメータ
func PostSAMLHandler(w http.ResponseWriter, req *http.Request) {
	access := models.Access1
	jsonData, err := json.Marshal(access)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
