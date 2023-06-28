package model

import (
	"database/sql/driver"
	"errors"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Method struct {
	gorm.Model      `json:"-"`
	Name            string          `gorm:"not null;index" json:"name" validate:"camelCase,required" example:"MethodName"`
	Params          Params          `gorm:"not null" validate:"required" json:"params" example:"string, string, int"`
	Description     string          `gorm:"not null" validate:"required" json:"description" example:"This method does an operation"`
	ResultStructure ResultStructure `gorm:"not null" validate:"required" json:"result_structure" example:"{ \"key\": \"value\" }"`
}

type ResultStructure map[string]any

func (r *ResultStructure) GormDataType() string { return "JSONB" }

func (r *ResultStructure) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "mysql", "sqlite":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}

type Params []string

func (p *Params) GormDataType() string {
	return "VARCHAR(255)"
}

func (p *Params) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value cannot cast to []byte")
	}
	*p = strings.Split(string(bytes), ",")
	return nil
}

func (p Params) Value() (driver.Value, error) {
	if len(p) == 0 {
		return nil, nil
	}
	return strings.Join(p, ","), nil
}
