package services

import (
	"github.com/ymho/qnin-sp/models"
	"github.com/ymho/qnin-sp/repositories"
)

func (s *MyAppService) GetIdPListService(page int) ([]models.IdentityProvider, error) {
	idpList, err := repositories.SelectIdPList(s.db, page)
	if err != nil {
		return nil, err
	}

	return idpList, nil

}

func (s *MyAppService) PostIdPService(idp models.IdentityProvider) (models.IdentityProvider, error) {

	newIdP, err := repositories.InsertIdP(s.db, idp)
	if err != nil {
		return models.IdentityProvider{}, err
	}
	return newIdP, nil
}

// TODO: ID指定
// TODO: UPDATEする関数
