package models

import "time"

type OptionDB struct {
	OptionID  int       `json:"option_id" db:"option_id"`
	OptionUrl string    `json:"option_url" db:"option_url"`
	MethodApi string    `json:"method_api" db:"method_api"`
	SP        string    `json:"sp" db:"sp"`
	LineNo    int       `json:"line_no" db:"line_no"`
	TableName string    `json:"table_name" db:"table_name"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	Updatedby string    `json:"updated_by" db:"updated_by"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type DefineColumn struct {
	ColumnField string `json:"column_field" db:"column_field"`
}

type OptionLookup struct {
	OptionLookUpCD string `json:"option_lookup_cd" db:"option_lookup_cd"`
	ColumnDB       string `json:"column_db" db:"column_db"`
	ViewName       string `json:"view_name" db:"view_name"`
	SourceField    string `json:"source_field" db:"source_field"`
	DisplayLookup  string `json:"display_lookup" db:"display_lookup"`
}
