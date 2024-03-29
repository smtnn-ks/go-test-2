// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type usrTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *usrTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("usr").
func (v *usrTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *usrTableType) Columns() []string {
	return []string{
		"id",
		"usr_name",
		"pass",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *usrTableType) NewStruct() reform.Struct {
	return new(Usr)
}

// NewRecord makes a new record for that table.
func (v *usrTableType) NewRecord() reform.Record {
	return new(Usr)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *usrTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// UsrTable represents usr view or table in SQL database.
var UsrTable = &usrTableType{
	s: parse.StructInfo{
		Type:    "Usr",
		SQLName: "usr",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "int64", Column: "id"},
			{Name: "UsrName", Type: "string", Column: "usr_name"},
			{Name: "Pass", Type: "string", Column: "pass"},
		},
		PKFieldIndex: 0,
	},
	z: new(Usr).Values(),
}

// String returns a string representation of this struct or record.
func (s Usr) String() string {
	res := make([]string, 3)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "UsrName: " + reform.Inspect(s.UsrName, true)
	res[2] = "Pass: " + reform.Inspect(s.Pass, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Usr) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.UsrName,
		s.Pass,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Usr) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.UsrName,
		&s.Pass,
	}
}

// View returns View object for that struct.
func (s *Usr) View() reform.View {
	return UsrTable
}

// Table returns Table object for that record.
func (s *Usr) Table() reform.Table {
	return UsrTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Usr) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Usr) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Usr) HasPK() bool {
	return s.ID != UsrTable.z[UsrTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *Usr) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = UsrTable
	_ reform.Struct = (*Usr)(nil)
	_ reform.Table  = UsrTable
	_ reform.Record = (*Usr)(nil)
	_ fmt.Stringer  = (*Usr)(nil)
)

func init() {
	parse.AssertUpToDate(&UsrTable.s, new(Usr))
}
