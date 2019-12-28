// Copyright (c) 2019 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package storage

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"github.com/ilibs/gosql/v2"
)

var logger = utilities.NewDefaultLogger("./logs/system.log")

// Database is the primary class for persisting contents
type Database struct {
	configs     map[string]*gosql.Config
	isConnected bool
}

// NewDB returns a new Database Instance
func NewDB(driver string, uri string) Database {
	configs := make(map[string]*gosql.Config)
	configs["default"] = &gosql.Config{
		Enable:  true,
		Driver:  driver,
		Dsn:     uri,
		ShowSql: true,
	}
	return Database{
		configs:     configs,
		isConnected: false,
	}
}

// CreateTables Create all Tables
func (db *Database) CreateTables() {
	// Read SQL
	var sqlString = utilities.GetRemoteFile("https://raw.githubusercontent.com/autoai-org/CVPM/master/components/cmd/pkg/storage/db.sql")
	gosql.Exec(sqlString)
}

// Connect tries to connect the database
func (db *Database) Connect() {
	gosql.Connect(db.configs)
	db.isConnected = true
}