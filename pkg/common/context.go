package common

import (
	"context"
	"xfd-backend/database/db/model"
	"xfd-backend/pkg/consts"
)

func GetUserID(ctx context.Context) string {
	return ctx.Value(consts.CONTEXT_HEADER_USER_ID).(string)
}

func GetUserRole(ctx context.Context) model.UserRole {
	return ctx.Value(consts.CONTEXT_HEADER_USER_ROLE).(model.UserRole)
}
