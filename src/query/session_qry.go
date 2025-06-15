package query

import (
	"context"
	"encoding/json"
	"fmt"

	"kredit-plus/src/constant"
	"kredit-plus/src/model"
	"kredit-plus/toolkit/logger"
)

const sessionKey = "session-%s"

func (q *Query) ReadSessionQuery(
	ctx context.Context,
	sessionGUID string,
) (data model.Session, err error) {
	cmd := q.cache.Get(ctx, fmt.Sprintf(sessionKey, sessionGUID))
	if cmd.Err() == nil {
		var value []byte

		value, err = cmd.Bytes()
		if err != nil {
			logger.WithContext(ctx).Error(err, "error cache get session")
			return
		}

		if err = json.Unmarshal(value, &data); err != nil {
			logger.WithContext(ctx).Error(err, "error unmarshal session model")
			return
		}

		return
	}

	data = model.Session{GUID: sessionGUID}

	if err = q.db.First(&data).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error find session by guid : "+sessionGUID)
		return
	}

	q.setSessionCache(ctx, data)

	return
}

func (q *Query) CreateSessionQuery(
	ctx context.Context,
	arg model.Session,
) (err error) {
	if err = q.db.Create(&arg).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error create session", "session model", arg)
		return
	}

	q.setSessionCache(ctx, arg)

	return
}

func (q *Query) UpdateSessionQuery(
	ctx context.Context,
	arg model.Session,
) (err error) {
	if err = q.db.Updates(&arg).Error; err != nil {
		logger.WithContext(ctx).Error(err, "error update session", "session model", arg)
		return
	}

	q.setSessionCache(ctx, arg)

	return
}

func (q *Query) setSessionCache(
	ctx context.Context,
	arg model.Session,
) {
	value, err := json.Marshal(arg)
	if err != nil {
		logger.WithContext(ctx).Error(err, "error marshal session model", "session model", arg)
		return
	}

	err = q.cache.Set(ctx, fmt.Sprintf(sessionKey, arg.GUID), value, constant.DefaultCacheExpireDuration).Err()
	if err != nil {
		logger.WithContext(ctx).Error(err, "error cache set session")
		return
	}
}
