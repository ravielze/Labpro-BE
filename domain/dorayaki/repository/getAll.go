package repository

import (
	"github.com/ravielze/Labpro-BE/constants"
	"github.com/ravielze/Labpro-BE/model/dao"
	"github.com/ravielze/oculi/common/functions"
	"github.com/ravielze/oculi/request"
)

func (r *repository) GetAll(req request.Context, page int) ([]dao.Dorayaki, int, error) {
	var totalItems int64
	if err := req.Transaction().Model(dao.Dorayaki{}).Count(&totalItems).Error(); err != nil {
		return nil, 0, err
	}

	totalPage := functions.CalculateTotalPages(int(totalItems), 10)

	offset := ((page) - 1) * 10
	if offset > int(totalItems) {
		return nil, 0, constants.ErrPageOutOfRange
	}

	var result []dao.Dorayaki
	if err := req.Transaction().Model(dao.Dorayaki{}).Offset(int(offset)).Limit(10).Find(&result).Error(); err != nil {
		return nil, 0, err
	}
	return result, totalPage, nil
}
