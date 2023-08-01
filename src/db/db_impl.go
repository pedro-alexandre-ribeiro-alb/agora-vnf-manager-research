package db

import (
	vnf_log "agora-vnf-manager/core/log"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type dbConnectionGorm struct {
	engine *gorm.DB
}

type dbSessionGorm struct {
	session *gorm.DB
}

var gorm_logger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		ParameterizedQueries:      true,
		Colorful:                  true,
	},
)

var gorm_silent_logger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		LogLevel: logger.Silent,
	},
)

func NewDb(sqlConn *sql.DB) (db IDbConnection, err error) {
	// &gorm.Config{Logger: gorm_logger}
	if db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlConn}), &gorm.Config{Logger: gorm_silent_logger}); err == nil {
		return &dbConnectionGorm{engine: db}, err
	} else {
		vnf_log.Errorf("[NewDb]: Could not initialize database - %s", err.Error())
		return nil, err
	}
}

func (c *dbConnectionGorm) NewSession() (session IDbSession) {
	output := dbSessionGorm{session: c.engine.Session(&gorm.Session{Logger: gorm_logger})}
	return &output
}

func (c *dbConnectionGorm) Sync(data ...any) (err error) {
	return c.engine.AutoMigrate(data...)
}

func (s *dbSessionGorm) Exec(rawQuery string) error {
	result := s.session.Exec(rawQuery)
	return result.Error
}

func (s *dbSessionGorm) Query(rawQuery string, dest any) error {
	result := s.session.Raw(rawQuery).Scan(&dest)
	return result.Error
}

func (s *dbSessionGorm) Create(beans any) error {
	result := s.session.Create(beans)
	return result.Error
}

func (s *dbSessionGorm) Get(beans any) (bool, error) {
	result := s.session.Preload(clause.Associations).Take(&beans)
	return true, result.Error
}

func (s *dbSessionGorm) FindOne(beans any, cond any) error {
	result := s.session.Preload(clause.Associations).Limit(1).Find(beans, cond)
	return result.Error
}

func (s *dbSessionGorm) Find(beans any, cond any) error {
	result := s.session.Preload(clause.Associations).Find(beans, cond)
	return result.Error
}

func (s *dbSessionGorm) Update(beans any) error {
	result := s.session.Model(beans).Clauses(clause.Returning{}).Updates(beans)
	return result.Error
}

func (s *dbSessionGorm) Delete(beans any, conds any) error {
	result := s.session.Unscoped().Delete(beans, conds)
	if result.RowsAffected <= 0 {
		return fmt.Errorf("could not delete record: %s", result.Error.Error())
	}
	return result.Error
}

func (s *dbSessionGorm) Model(data any) IDbSession {
	s.session = s.session.Model(data)
	return s
}

func (s *dbSessionGorm) Begin() IDbSession {
	s.session = s.session.Begin()
	return s
}

func (s *dbSessionGorm) Commit() IDbSession {
	s.session = s.session.Commit()
	return s
}

func (s *dbSessionGorm) Rollback() IDbSession {
	s.session = s.session.Rollback()
	return s
}

func (s *dbSessionGorm) NativeImpl() *gorm.DB {
	return s.session
}

func (db *Db) GetEngine() IDbConnection {
	return db.engine
}

func (d *Db) SessionGorm() (transaction IDbSession, err error) {
	if d.engine == nil {
		return nil, fmt.Errorf(`DB engine is not initialized`)
	}
	return d.engine.NewSession(), nil
}
