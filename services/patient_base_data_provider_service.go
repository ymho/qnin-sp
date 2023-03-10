package services

import (
	"github.com/ymho/qnin-sp/apperrors"
	"github.com/ymho/qnin-sp/models"
	"github.com/ymho/qnin-sp/repositories"
)

func (s *MyAppService) GetPBDPListService(page int) ([]models.PatientBaseDataProvider, error) {
	pbdpList, err := repositories.SelectPBDPList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(pbdpList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return pbdpList, nil

}

func (s *MyAppService) PostPBDPService(pbdp models.PatientBaseDataProvider) (models.PatientBaseDataProvider, error) {

	newPBDP, err := repositories.InsertPBDP(s.db, pbdp)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.PatientBaseDataProvider{}, err
	}
	return newPBDP, nil
}

// TODO: ID指定
// TODO: UPDATEする関数
