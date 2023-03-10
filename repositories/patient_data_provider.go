package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ymho/qnin-sp/models"
)

const (
	pdpNumPerPage = 10
)

func InsertPDP(db *sql.DB, pdp models.PatientDataProvider) (models.PatientDataProvider, error) {
	const sqlStr = `
	insert into patient_data_provider (organization_name, ja_organization_name, metadata_url, query_url, expire_at, created_at) values
	(?,?,?,?,now(),now());
	`

	var newPDP models.PatientDataProvider
	newPDP.OrganizationName, newPDP.JaOrganizationName, newPDP.MetadataURL, newPDP.QueryURL = pdp.OrganizationName, pdp.JaOrganizationName, pdp.MetadataURL, pdp.QueryURL

	result, err := db.Exec(sqlStr, pdp.OrganizationName, pdp.JaOrganizationName, pdp.MetadataURL, pdp.QueryURL)
	if err != nil {
		return models.PatientDataProvider{}, err
	}

	id, _ := result.LastInsertId()

	newPDP.ID = int(id)

	return newPDP, nil
}

func SelectPDPList(db *sql.DB, page int) ([]models.PatientDataProvider, error) {
	const sqlStr = `
	select id, organization_name, ja_organization_name, metadata_url, query_url
	from patient_data_provider
	limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, pdpNumPerPage, ((page - 1) * pdpNumPerPage))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pdpArray := make([]models.PatientDataProvider, 0)
	for rows.Next() {
		var pdp models.PatientDataProvider
		rows.Scan(&pdp.ID, &pdp.OrganizationName, &pdp.JaOrganizationName, &pdp.MetadataURL, &pdp.QueryURL)
		pdpArray = append(pdpArray, pdp)
	}

	return pdpArray, nil
}

// TODO: IDを指定する関数
// TODO: UPDATEする関数
