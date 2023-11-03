package repo

import (
	"context"
	"gorm.io/gorm"
	"xfd-backend/pkg/xerr"
)

type contextKeyTransaction struct{}

var transactionKey = contextKeyTransaction{}

type IRepo interface {
	GetDB(context.Context) *gorm.DB
}

type TransactionHandler interface {
	WithTransaction(ctx context.Context, f func(context.Context) xerr.XErr) xerr.XErr
}

var _ IRepo = repo{}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) IRepo {
	return repo{db: db}
}

func NewTransactionHandler(db *gorm.DB) TransactionHandler {
	return repo{db: db}
}

func (r repo) WithTransaction(ctx context.Context, f func(ctx context.Context) xerr.XErr) xerr.XErr {
	tx := r.db.Begin()
	if tx.Error != nil {
		return xerr.WithCode(xerr.ErrorDatabase, tx.Error)
	}
	ctxWithTx := context.WithValue(ctx, transactionKey, tx)

	var xErr xerr.XErr
	defer func() {
		if xErr != nil {
			tx.Rollback()
		}
		if recoverInfo := recover(); recoverInfo != nil {
			tx.Rollback()
			panic(recoverInfo)
		}
	}()

	if xErr = f(ctxWithTx); xErr != nil {
		return xErr
	}

	return xerr.WithCode(xerr.ErrorDatabase, tx.Commit().Error)
}

func (r repo) GetDB(ctx context.Context) *gorm.DB {
	itx := ctx.Value(transactionKey)
	if itx == nil {
		return r.db.Unscoped()
	}
	tx, _ := itx.(*gorm.DB)
	return tx.Unscoped()
}
