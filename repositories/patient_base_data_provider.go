package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ymho/qnin-sp/models"
)

const (
	pbdpNumPerPage = 10
)

func InsertPBDP(db *sql.DB, pbdp models.PatientBaseDataProvider) (models.PatientBaseDataProvider, error) {
	const sqlStr = `
	insert into patient_base_data_provider (organization_name, ja_organization_name, metadata_url, query_url, expire_at, created_at) values
	(?,?,?,?,now(),now());
	`

	var newPBDP models.PatientBaseDataProvider
	newPBDP.OrganizationName, newPBDP.JaOrganizationName, newPBDP.MetadataURL, newPBDP.QueryURL = pbdp.OrganizationName, pbdp.JaOrganizationName, pbdp.MetadataURL, pbdp.QueryURL

	result, err := db.Exec(sqlStr, pbdp.OrganizationName, pbdp.JaOrganizationName, pbdp.MetadataURL, pbdp.QueryURL)
	if err != nil {
		return models.PatientBaseDataProvider{}, err
	}

	id, _ := result.LastInsertId()

	newPBDP.ID = int(id)

	return newPBDP, nil
}

func SelectPBDPList(db *sql.DB, page int) ([]models.PatientBaseDataProvider, error) {
	const sqlStr = `
	select id, organization_name, ja_organization_name, metadata_url, query_url
	from patient_base_data_provider
	limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, pbdpNumPerPage, ((page - 1) * pbdpNumPerPage))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pbdpArray := make([]models.PatientBaseDataProvider, 0)
	for rows.Next() {
		var pbdp models.PatientBaseDataProvider
		rows.Scan(&pbdp.ID, &pbdp.OrganizationName, &pbdp.JaOrganizationName, &pbdp.MetadataURL, &pbdp.QueryURL)
		pbdpArray = append(pbdpArray, pbdp)
	}

	return pbdpArray, nil
}

// TODO: IDを指定する関数
// TODO: UPDATEする関数
