package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Method struct {
	gorm.Model      `json:"-"`
	Name            string          `gorm:"not null;index" json:"name" validate:"camelCase,required" example:"MethodName"`
	Params          Params          `gorm:"not null" validate:"required" json:"params" example:"string, string, int"`
	Description     string          `gorm:"not null" validate:"required" json:"description" example:"This method does an operation"`
	ResultStructure ResultStructure `gorm:"not null" validate:"required" json:"result_structure" example:"{ \"key\": \"value\" }"`
	Kind            MethodKind      `gorm:"not null;default:concurrent" validate:"required" json:"kind" example:"concurrent"`
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

type MethodKind byte

const (
	Broadcast MethodKind = iota
	Concurrent
	Indepotent
	Exchange
)

var (
	methodKindToString = map[MethodKind]string{
		Broadcast:  "broadcast",
		Concurrent: "concurrent",
		Indepotent: "indepotent",
		Exchange:   "exchange",
	}
	stringToMethodKind = map[string]MethodKind{
		"broadcast":  Broadcast,
		"concurrent": Concurrent,
		"indepotent": Indepotent,
		"exchange":   Exchange,
	}
)

func (m *MethodKind) GormDataType() string {
	keys := lo.Map(lo.Keys(stringToMethodKind), func(k string, _ int) string { return fmt.Sprintf("'%s'", k) })
	return fmt.Sprintf("enum(%s)", strings.Join(keys, ","))
}

func (m *MethodKind) Scan(src any) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("src value cannot cast to []byte")
	}
	*m = stringToMethodKind[string(bytes)]
	return nil
}

func (m MethodKind) Value() (driver.Value, error) {
	return methodKindToString[m], nil
}

func (m MethodKind) String() string {
	return methodKindToString[m]
}

func (m MethodKind) MarshalJSON() ([]byte, error) {
	return []byte(m.String()), nil
}

func (m *MethodKind) UnmarshalJSON(b []byte) error {
	*m = stringToMethodKind[string(b)]
	return nil
}
