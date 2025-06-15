package auth

import (
	"kredit-plus/src/model"
	"kredit-plus/toolkit/logger"
)

func (a *Auth) User() (data model.User, err error) {
	data = model.User{GUID: a.claims.UserGUID}

	if err = a.db.First(&data).Error; err != nil {
		logger.PrintError(err, "error find user", "id", a.claims.UserGUID)
		return
	}

	return
}
