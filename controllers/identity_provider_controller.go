package controllers

import (
	"encoding/json"
	"github.com/ymho/qnin-sp/apperrors"
	"io"
	"net/http"
	"strconv"

	"github.com/ymho/qnin-sp/controllers/services"
)

// TODO: エラーハンドラの呼び出しによりエラー処理

type IdentityProviderController struct {
	service services.IdentityProviderServicer
}

func NewIdentityProviderController(s services.IdentityProviderServicer) *IdentityProviderController {
	return &IdentityProviderController{service: s}
}

// PostIdPHandler /idpのハンドラ
func (c *IdentityProviderController) PostIdPHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: 実装
	io.WriteString(w, "Hello, world!\n")
}

// GetIdPListHandler /idp/listのハンドラ
func (c *IdentityProviderController) GetIdPListHandler(w http.ResponseWriter, req *http.Request) {

	queryMap := req.URL.Query()
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "query param must be number")
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	idpList, err := c.service.GetIdPListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(idpList)
}

// TODO ハンドラメソッドの続きを実装
