package main

import (
	"database/sql"
	"time"

	"context"

	"github.com/jinzhu/gorm"
)

type SampleStruct struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string
	Mobile    string
}

func QueryWithTimeout(ctx context.Context, db *gorm.DB) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var sample SampleStruct
	tx := db.BeginTx(ctx, &sql.TxOptions{})
	err = tx.Set("gorm:query_option", "FOR UPDATE").Find(&sample, "name = Ankit").Error
	return
}
