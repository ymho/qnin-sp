package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ymho/qnin-sp/controllers"
	"github.com/ymho/qnin-sp/services"
)

func NewRouter(db *sql.DB) *mux.Router {

	ser := services.NewMyAppService(db)

	idpCon := controllers.NewIdentityProviderController(ser)
	pdpCon := controllers.NewPatientDataProviderController(ser)
	pbdpCon := controllers.NewPatientBaseDataProviderController(ser)
	aCon := controllers.NewAccessCorrespondenceController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/", aCon.DefaultHandler).Methods(http.MethodGet)
	r.HandleFunc("/hello", aCon.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/list", aCon.GetACListHandler).Methods(http.MethodGet)
	r.HandleFunc("/saml/", aCon.SAMLHandler)

	r.HandleFunc("/idp", idpCon.PostIdPHandler).Methods(http.MethodPost)
	r.HandleFunc("/idp/list", idpCon.GetIdPListHandler).Methods(http.MethodGet)

	r.HandleFunc("/pdp", pdpCon.PostPDPHandler).Methods(http.MethodPost)
	r.HandleFunc("/pdp/list", pdpCon.GetPDPListHandler).Methods(http.MethodGet)

	r.HandleFunc("/pbdp", pbdpCon.PostPBDPHandler).Methods(http.MethodPost)
	r.HandleFunc("/pbdp/list", pbdpCon.GetPBDPListHandler).Methods(http.MethodGet)

	return r
}
