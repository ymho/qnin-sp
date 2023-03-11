package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/crewjam/saml/samlsp"
	"github.com/ymho/qnin-sp/apperrors"
	"io"
	"net/http"
	"strconv"

	"github.com/ymho/qnin-sp/controllers/services"
)

// TODO: エラーハンドラの呼び出しによりエラー処理

type AccessCorrespondenceController struct {
	service services.AccessCorrespondenceServicer
}

func NewAccessCorrespondenceController(s services.AccessCorrespondenceServicer) *AccessCorrespondenceController {
	return &AccessCorrespondenceController{service: s}
}

// DefaultHandler /のハンドラ
func (c *AccessCorrespondenceController) DefaultHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: 実装
	io.WriteString(w, "Hello, world!\n")
}

// HelloHandler /helloのハンドラ
func (c *AccessCorrespondenceController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: 実装
	fmt.Fprintf(w, "Hello, %s, %s!", samlsp.AttributeFromContext(req.Context(), "urn:oid:0.9.2342.19200300.100.1.1"), samlsp.SessionFromContext(req.Context()))
	//io.WriteString(w, "Hello, world!\n")
}

// SAMLHandler /saml/のハンドラ
func (c *AccessCorrespondenceController) SAMLHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: 実装
	io.WriteString(w, "Hello, world!\n")
	// fmt.Fprintf(w, "Hello, %s, %s!", samlsp.AttributeFromContext(r.Context(), "urn:oid:0.9.2342.19200300.100.1.1"), samlsp.SessionFromContext(r.Context()))
}

// GetACListHandler /listのハンドラ
func (c *AccessCorrespondenceController) GetACListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "query param must be number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		page = 1
	}

	acList, err := c.service.GetACListService(page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(acList)
}
