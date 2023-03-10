package services

import (
	"github.com/ymho/qnin-sp/aperrors"
	"github.com/ymho/qnin-sp/models"
	"github.com/ymho/qnin-sp/repositories"
)

func (s *MyAppService) GetACListService(page int) ([]models.AccessCorrespondence, error) {
	acList, err := repositories.SelectACList(s.db, page)
	if err != nil {
		err = aperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}
	if len(acList) == 0 {
		err := aperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return acList, nil

}

func (s *MyAppService) PostACService(ac models.AccessCorrespondence) (models.AccessCorrespondence, error) {

	newAC, err := repositories.InsertAC(s.db, ac)
	if err != nil {
		err = aperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.AccessCorrespondence{}, err
	}
	return newAC, nil
}
