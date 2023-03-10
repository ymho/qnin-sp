package models

import (
	"net/url"
	"time"
)

type AccessCorrespondence struct {
	AccessID                 int       `json:"access_id"`
	SessionID                string    `json:"session_id"`
	Mail                     string    `json:"mail"`
	OrganizationName         string    `json:"organization_name"`
	JaOrganizationName       string    `json:"ja_organization_name"`
	OrganizationalUnitName   string    `json:"organizational_unit_name"`
	JaOrganizationalUnitName string    `json:"ja_organizational_unit_name"`
	SamlHash                 string    `json:"saml_hash"`
	Status                   string    `jspn:"status"`
	AuthorizedAt             time.Time `json:"authorized_at"`
	CreatedAt                time.Time `json:"created_at"`
}

type IdentityProvider struct {
	ID                 int       `json:"id"`
	OrganizationName   string    `json:"organization_name"`
	JaOrganizationName string    `json:"ja_organization_name"`
	MetadataURL        url.URL   `json:"metadata_url"`
	SSOURL             url.URL   `json:"sso_url"`
	LogoutURL          url.URL   `json:"logout_url"`
	ExpireAt           time.Time `json:"expire_at"`
	CreatedAt          time.Time `json:"created_at"`
}

type PatientBaseDataProvider struct {
	ID                 int       `json:"id"`
	OrganizationName   string    `json:"organization_name"`
	JaOrganizationName string    `json:"ja_organization_name"`
	QueryURL           url.URL   `json:"query_url"`
	MetadataURL        url.URL   `json:"metadata_url"`
	ExpireAt           time.Time `json:"expire_at"`
	CreatedAt          time.Time `json:"created_at"`
}

type PatientDataProvider struct {
	ID                 int       `json:"id"`
	OrganizationName   string    `json:"organization_name"`
	JaOrganizationName string    `json:"ja_organization_name"`
	QueryURL           url.URL   `json:"query_url"`
	MetadataURL        url.URL   `json:"metadata_url"`
	ExpireAt           time.Time `json:"expire_at"`
	CreatedAt          time.Time `json:"created_at"`
}
