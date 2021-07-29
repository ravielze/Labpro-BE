package constants

import (
	"errors"
	"net/http"

	consts "github.com/ravielze/oculi/constant/errors"
	errorUtil "github.com/ravielze/oculi/errors"
)

var (
	ErrResetUnauthorized = errors.New("reset unauthorized")

	HealthMappers = errorUtil.Mappers{
		{Code: http.StatusUnauthorized, Err: ErrResetUnauthorized},
	}

	ErrDorayakiNotFound = errors.New("dorayaki not found")
	ErrShopNotFound     = errors.New("shop not found")

	ErrFailedToParseID = errors.New("id should be more than 0 and in base-10 format")
	ErrPageOutOfRange  = errors.New("page out of range")

	ShopMappers = errorUtil.Mappers{
		{Code: http.StatusBadRequest, Err: ErrFailedToParseID},
		{Code: http.StatusNotFound, Err: ErrShopNotFound},
		{Code: http.StatusBadRequest, Err: consts.ErrRecordNotFound},
		{Code: http.StatusBadRequest, Err: ErrPageOutOfRange},
	}

	ErrStockNotEnough       = errors.New("there is not enough stock to transfer")
	TransferDorayakiMappers = errorUtil.Mappers{
		{Code: http.StatusBadRequest, Err: ErrFailedToParseID},
		{Code: http.StatusBadRequest, Err: ErrStockNotEnough},
		{Code: http.StatusNotFound, Err: consts.ErrRecordNotFound},
		{Code: http.StatusNotFound, Err: ErrShopNotFound},
		{Code: http.StatusNotFound, Err: ErrDorayakiNotFound},
	}

	StandardDorayakiMappers = errorUtil.Mappers{
		{Code: http.StatusBadRequest, Err: ErrFailedToParseID},
		{Code: http.StatusNotFound, Err: consts.ErrRecordNotFound},
		{Code: http.StatusNotFound, Err: ErrDorayakiNotFound},
		{Code: http.StatusNotFound, Err: ErrShopNotFound},
	}
)
