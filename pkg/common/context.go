package common

import (
	"context"
	"xfd-backend/pkg/consts"
)

func GetUserID(ctx context.Context) string {
	return ctx.Value(consts.CONTEXT_HEADER_USER_ID).(string)
}
