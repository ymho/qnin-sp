package services

import (
	"github.com/ymho/qnin-sp/aperrors"
	"github.com/ymho/qnin-sp/models"
	"github.com/ymho/qnin-sp/repositories"
)

func (s *MyAppService) GetIdPListService(page int) ([]models.IdentityProvider, error) {
	idpList, err := repositories.SelectIdPList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}
	if len(idpList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return idpList, nil

}

func (s *MyAppService) PostIdPService(idp models.IdentityProvider) (models.IdentityProvider, error) {

	newIdP, err := repositories.InsertIdP(s.db, idp)
	if err != nil {
		err = aperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.IdentityProvider{}, err
	}
	return newIdP, nil
}

// TODO: ID指定
// TODO: UPDATEする関数
