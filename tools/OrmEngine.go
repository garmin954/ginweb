package tools

import (
	"ginweb/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngin *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm, error) {
	dbconfig := cfg.Database
	conn := dbconfig.User + ":" + dbconfig.Password + "@tcp(" + dbconfig.Host + ":" + dbconfig.Port + ")/" + dbconfig.DbName + "?charset=" + dbconfig.Charset

	engine, err := xorm.NewEngine(dbconfig.Driver, conn)
	if err != nil {
		return nil, err
	}

	engine.ShowSQL(dbconfig.ShowSql)

	ersr := engine.Sync2(
		new(model.SmsCode),
		new(model.User),
	)
	if ersr != nil {
		return nil, ersr
	}
	orm := new(Orm)
	orm.Engine = engine
	DbEngin = orm
	return orm, nil
}
