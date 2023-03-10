package services

import (
	"github.com/ymho/qnin-sp/aperrors"
	"github.com/ymho/qnin-sp/models"
	"github.com/ymho/qnin-sp/repositories"
)

func (s *MyAppService) GetPDPListService(page int) ([]models.PatientDataProvider, error) {
	pdpList, err := repositories.SelectPDPList(s.db, page)
	if err != nil {
		err = aperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}
	if len(pdpList) == 0 {
		err := aperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return pdpList, nil

}

func (s *MyAppService) PostPDPService(pdp models.PatientDataProvider) (models.PatientDataProvider, error) {

	newPDP, err := repositories.InsertPDP(s.db, pdp)
	if err != nil {
		err = aperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.PatientDataProvider{}, err
	}
	return newPDP, nil
}

// TODO: ID指定
// TODO: UPDATEする関数
