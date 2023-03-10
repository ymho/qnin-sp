package repositories

import (
	"database/sql"

	"github.com/ymho/qnin-sp/models"
)

const (
	acNumPerPage = 10
)

func InsertAC(db *sql.DB, ac models.AccessCorrespondence) (models.AccessCorrespondence, error) {
	const sqlStr = `
	insert into access_correspondence (session_id, mail, organization_name, ja_organization_name, organizational_unit_name, ja_organizational_unit_name, saml_hash, status, authorized_at, created_at) values
	(?,?,?,?,?,?,?,?,now(),now());
	`

	var newAC models.AccessCorrespondence
	newAC.SessionID, newAC.Mail, newAC.OrganizationName, newAC.JaOrganizationName, newAC.OrganizationalUnitName, newAC.JaOrganizationalUnitName, newAC.SamlHash, newAC.Status = ac.SessionID, ac.Mail, ac.OrganizationName, ac.JaOrganizationName, ac.OrganizationalUnitName, ac.JaOrganizationalUnitName, ac.SamlHash, ac.Status

	result, err := db.Exec(sqlStr, ac.SessionID, ac.Mail, ac.OrganizationName, ac.JaOrganizationName, ac.OrganizationalUnitName, ac.JaOrganizationalUnitName, ac.SamlHash, ac.Status)
	if err != nil {
		return models.AccessCorrespondence{}, err
	}

	id, _ := result.LastInsertId()

	newAC.AccessID = int(id)

	return newAC, nil
}

func SelectACList(db *sql.DB, page int) ([]models.AccessCorrespondence, error) {
	const sqlStr = `
	select access_id, session_id, mail, organization_name, ja_organization_name, organizational_unit_name, ja_organizational_unit_name, saml_hash, status
	from access_correspondence
	limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, acNumPerPage, ((page - 1) * acNumPerPage))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	acArray := make([]models.AccessCorrespondence, 0)
	for rows.Next() {
		var ac models.AccessCorrespondence
		rows.Scan(&ac.AccessID, &ac.SessionID, &ac.Mail, &ac.OrganizationName, &ac.JaOrganizationName, &ac.OrganizationalUnitName, &ac.JaOrganizationalUnitName, &ac.SamlHash, &ac.Status)
		acArray = append(acArray, ac)
	}

	return acArray, nil
}

// TODO: IDを指定する関数
// TODO: UPDATEする関数
