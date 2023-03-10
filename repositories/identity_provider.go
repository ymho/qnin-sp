package repositories

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ymho/qnin-sp/models"
)

const (
	idpNumPerPage = 10
)

func InsertIdP(db *sql.DB, idp models.IdentityProvider) (models.IdentityProvider, error) {
	const sqlStr = `
	insert into identity_provider (organization_name, ja_organization_name, metadata_url, sso_url, logout_url, expire_at, created_at) values
	(?,?,?,?,?,now(),now());
	`

	var newIdP models.IdentityProvider
	newIdP.OrganizationName, newIdP.JaOrganizationName, newIdP.MetadataURL, newIdP.SSOURL, newIdP.LogoutURL = idp.OrganizationName, idp.JaOrganizationName, idp.MetadataURL, idp.SSOURL, idp.LogoutURL

	result, err := db.Exec(sqlStr, idp.OrganizationName, idp.JaOrganizationName, idp.MetadataURL, idp.SSOURL, idp.LogoutURL)
	if err != nil {
		return models.IdentityProvider{}, err
	}

	id, _ := result.LastInsertId()

	newIdP.ID = int(id)

	return newIdP, nil
}

func SelectIdPList(db *sql.DB, page int) ([]models.IdentityProvider, error) {
	const sqlStr = `
	select id, organization_name, ja_organization_name, metadata_url, sso_url, logout_url
	from identity_provider
	limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, idpNumPerPage, ((page - 1) * idpNumPerPage))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	idpArray := make([]models.IdentityProvider, 0)
	for rows.Next() {
		var idp models.IdentityProvider
		rows.Scan(&idp.ID, &idp.OrganizationName, &idp.JaOrganizationName, &idp.MetadataURL, &idp.SSOURL, &idp.LogoutURL)
		idpArray = append(idpArray, idp)
	}

	return idpArray, nil
}

// TODO: IDを指定する関数
// TODO: UPDATEする関数
