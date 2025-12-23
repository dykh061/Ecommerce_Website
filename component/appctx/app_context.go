package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
}
type appContext struct {
	db        *gorm.DB
	secretKey string
}

func NewAppContext(db *gorm.DB, secretKey string) *appContext {
	return &appContext{db: db, secretKey: secretKey}
}

func (ctx *appContext) GetMainDBConnection() *gorm.DB { return ctx.db }

func (ctx *appContext) SecretKey() string { return ctx.secretKey }
