package entity

import (
	"fmt"
)

type FilterFields interface {
	Filter() string
}

type FilterField struct {
	TableName  string
	FieldName  string
	Operator   string
	FieldValue string
}

func NewFilterField(tn, fn, o, fv string) *FilterField {
	return &FilterField{
		TableName:  tn,
		FieldName:  fn,
		Operator:   o,
		FieldValue: fv,
	}
}

func (ff FilterField) Filter() string {
	return fmt.Sprintf("%s.%s %s '%s'", ff.TableName, ff.FieldName, ff.Operator, ff.FieldValue)
}

type FilterGroup struct {
	Fields   []FilterFields
	Operator string
}

func NewFilterGroup(ff []FilterFields, o string) *FilterGroup {
	nff := make([]FilterFields, len(ff))
	copy(nff, ff)

	return &FilterGroup{
		Fields:   nff,
		Operator: o,
	}
}

func (fg FilterGroup) Filter() string {
	s := ""
	if len(fg.Fields) > 0 {
		s += "("
	}

	for k, v := range fg.Fields {
		s += v.Filter()
		if k < len(fg.Fields)-1 {
			s += fg.Operator
		}
	}

	if len(fg.Fields) > 0 {
		s += ")"
	}

	return s
}
