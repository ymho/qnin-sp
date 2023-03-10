package services

import (
	"github.com/ymho/qnin-sp/models"
	"github.com/ymho/qnin-sp/repositories"
)

func (s *MyAppService) GetACListService(page int) ([]models.AccessCorrespondence, error) {
	acList, err := repositories.SelectACList(s.db, page)
	if err != nil {
		return nil, err
	}

	return acList, nil

}

func (s *MyAppService) PostACService(ac models.AccessCorrespondence) (models.AccessCorrespondence, error) {

	newAC, err := repositories.InsertAC(s.db, ac)
	if err != nil {
		return models.AccessCorrespondence{}, err
	}
	return newAC, nil
}
