package db

import (
	"agora-vnf-manager/config"
	"agora-vnf-manager/core/log"
	"database/sql"
	"fmt"
)

type IDbConnection interface {
	NewSession() (session IDbSession)
	Sync(...any) (err error)
}

type IDbSession interface {
	Exec(string) (e error)
	Query(string, any) (e error)
	Create(any) (e error)
	Get(cond any) (has bool, e error)
	FindOne(data any, cond any) (e error)
	Find(data any, cond any) (e error)
	Update(data any) (e error)
	Delete(any, any) (e error)
	Model(any) IDbSession
	Begin() IDbSession
	Commit() IDbSession
	Rollback() IDbSession
}

type DbParams struct {
	Path string `json:"path"`
}

type Db struct {
	engine IDbConnection
}

func CreateDb(cfg *config.Config) *Db {
	output := Db{}
	output.engine = nil
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", cfg.Database.DbAddress, cfg.Database.DbPort, cfg.Database.DbUsername, cfg.Database.DbName)
	sqlDB, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Errorf("[CreateDb]: Could not initialize DB - %s", err.Error())
	}
	if db, err := NewDb(sqlDB); err != nil {
		log.Errorf("[CreateDb]: Could not initialize DB - %s", err.Error())
	} else {
		output.engine = db
	}
	return &output
}

func (d *Db) Sync(beans ...interface{}) error {
	if d.engine == nil {
		return fmt.Errorf(`DB engine is not initialized`)
	}
	return d.engine.Sync(beans...)
}
