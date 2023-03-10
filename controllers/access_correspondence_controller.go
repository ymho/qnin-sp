package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ymho/qnin-sp/controllers/services"
)

type AccessCorrespondenceController struct {
	service services.AccessCorrespondenceServicer
}

func NewAccessCorrespondenceController(s services.AccessCorrespondenceServicer) *AccessCorrespondenceController {
	return &AccessCorrespondenceController{service: s}
}

// /のハンドラ
func (c *AccessCorrespondenceController) DefaultHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: 実装
	io.WriteString(w, "Hello, world!\n")
}

// /helloのハンドラ
func (c *AccessCorrespondenceController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: 実装
	io.WriteString(w, "Hello, world!\n")
}

// /saml/のハンドラ
func (c *AccessCorrespondenceController) SAMLHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: 実装
	io.WriteString(w, "Hello, world!\n")
	// fmt.Fprintf(w, "Hello, %s, %s!", samlsp.AttributeFromContext(r.Context(), "urn:oid:0.9.2342.19200300.100.1.1"), samlsp.SessionFromContext(r.Context()))
}

// /listのハンドラ
func (c *AccessCorrespondenceController) GetACListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	acList, err := c.service.GetACListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(acList)
}
