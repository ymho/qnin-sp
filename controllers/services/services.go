package services

import "github.com/ymho/qnin-sp/models"

type IdentityProviderServicer interface {
	GetIdPListService(page int) ([]models.IdentityProvider, error)
	PostIdPService(idp models.IdentityProvider) (models.IdentityProvider, error)
	// TODO サービスを定義
}

type PatientBaseDataProviderServicer interface {
	GetPBDPListService(page int) ([]models.PatientBaseDataProvider, error)
	PostPBDPService(pbdp models.PatientBaseDataProvider) (models.PatientBaseDataProvider, error)
	// TODO サービスを定義
}

type PatientDataProviderServicer interface {
	GetPDPListService(page int) ([]models.PatientDataProvider, error)
	PostPDPService(pdp models.PatientDataProvider) (models.PatientDataProvider, error)
	// TODO サービスを定義
}

type AccessCorrespondenceServicer interface {
	GetACListService(page int) ([]models.AccessCorrespondence, error)
	PostACService(ac models.AccessCorrespondence) (models.AccessCorrespondence, error)
	// TODO サービスを定義
}
