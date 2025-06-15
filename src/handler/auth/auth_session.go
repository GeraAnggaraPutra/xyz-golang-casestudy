package auth

import (
	"context"
	"time"

	"kredit-plus/src/constant"
	"kredit-plus/src/query"
	"kredit-plus/toolkit/logger"
)

func (a *Auth) ValidateSession(ctx context.Context) (err error) {
	q := query.NewQuery(a.cache, a.db)

	session, err := q.ReadSessionQuery(ctx, a.claims.GUID)
	if err != nil {
		logger.PrintError(err, "error read session by guid : "+a.claims.GUID)
		return
	}

	if time.Now().After(session.AccessTokenExpiredAt) {
		err = constant.ErrTokenExpired
		return
	}

	return
}
