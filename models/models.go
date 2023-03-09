package models

import "time"

type Access struct {
	AccessID     string    `json:"access_id"`
	SessionID    string    `json:"session_id"`
	UserName     string    `json:"user_name"`
	Organization string    `json:"organization"`
	SAMLHash     string    `json:"saml_hash"`
	Status       string    `jspn:"status"`
	AuthorizedAt time.Time `json:"authorized_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type IdP struct {
	ID        string    `json:"idp_id"`
	Name      string    `json:"idp_name"`
	Url       string    `json:"url"`
	ExpireAt  time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
}

type PIdP struct {
	ID        string    `json:"idp_id"`
	Name      string    `json:"idp_name"`
	Url       string    `json:"url"`
	ExpireAt  time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
}

type PDP struct {
	ID        string    `json:"idp_id"`
	Name      string    `json:"idp_name"`
	Url       string    `json:"url"`
	ExpireAt  time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
}
