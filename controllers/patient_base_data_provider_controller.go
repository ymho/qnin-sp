package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/ymho/qnin-sp/controllers/services"
)

type PatientBaseDataProviderController struct {
	service services.PatientBaseDataProviderServicer
}

func NewPatientBaseDataProviderController(s services.PatientBaseDataProviderServicer) *PatientBaseDataProviderController {
	return &PatientBaseDataProviderController{service: s}
}

// /pbdpのハンドラ
func (c *PatientBaseDataProviderController) PostPBDPHandler(w http.ResponseWriter, req *http.Request) {
	// TODO: 実装
	io.WriteString(w, "Hello, world!\n")
}

// /pbdp/listのハンドラ
func (c *PatientBaseDataProviderController) GetPBDPListHandler(w http.ResponseWriter, req *http.Request) {

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

	pbdpList, err := c.service.GetPBDPListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(pbdpList)
}

// TODO ハンドラメソッドを実装
