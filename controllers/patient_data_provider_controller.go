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

type PatientDataProviderController struct {
	service services.PatientDataProviderServicer
}

func NewPatientDataProviderController(s services.PatientDataProviderServicer) *PatientDataProviderController {
	return &PatientDataProviderController{service: s}
}

// PostPDPHandler /pdpのハンドラ
func (c *PatientDataProviderController) PostPDPHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: 実装
	io.WriteString(w, "Hello, world!\n")
}

// GetPDPListHandler /pdp/listのハンドラ
func (c *PatientDataProviderController) GetPDPListHandler(w http.ResponseWriter, req *http.Request) {

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

	pdpList, err := c.service.GetPDPListService(page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(pdpList)
}

// TODO ハンドラメソッドを実装
