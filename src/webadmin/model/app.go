package model

import "github.com/jmoiron/sqlx"

type AppInfo struct {
	AppId       int    `db:"app_id"`
	AppName     string `db:"app_name"`
	AppType     string `db:"app_type"`
	CreateTime  string `db:"create_time"`
	DevelopPath string `db:"develop_path"`
	IP          []string
}

var (
	Db *sqlx.DB
)

func InitDb(db *sqlx.DB) {
	Db = db
}
