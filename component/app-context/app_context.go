package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetMainDB() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}

func (ctx *appCtx) GetMainDB() *gorm.DB {
	return ctx.db
}
