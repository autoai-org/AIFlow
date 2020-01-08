// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package entities

import (
	"github.com/autoai-org/aiflow/components/cmd/pkg/storage"
	"time"
)

// Event defines the basic structure of a spawned event
type Event struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	Data      string    `db:"data"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Status    string    `db:"status"`
	From      string    `db:"from"`
}

// TableName defines the tablename in database
func (e *Event) TableName() string {
	return "event"
}

// PK defines the primary key of Event
func (e *Event) PK() string {
	return "id"
}

// Save stores event into database
func (e *Event) Save() error {
	db := storage.GetDefaultDB()
	db.Connect()
	return db.Insert(e)
}
