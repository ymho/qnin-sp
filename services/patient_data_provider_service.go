package services

import (
	"github.com/ymho/qnin-sp/models"
	"github.com/ymho/qnin-sp/repositories"
)

func (s *MyAppService) GetPDPListService(page int) ([]models.PatientDataProvider, error) {
	pdpList, err := repositories.SelectPDPList(s.db, page)
	if err != nil {
		return nil, err
	}

	return pdpList, nil

}

func (s *MyAppService) PostPDPService(pdp models.PatientDataProvider) (models.PatientDataProvider, error) {

	newPDP, err := repositories.InsertPDP(s.db, pdp)
	if err != nil {
		return models.PatientDataProvider{}, err
	}
	return newPDP, nil
}

// TODO: ID指定
// TODO: UPDATEする関数
