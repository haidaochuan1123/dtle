package common

import (
	"reflect"
	"testing"

	test "github.com/outbrain/golib/tests"
)

func TestParseColumnList(t *testing.T) {
	names := "id,category,max_len"

	columnList := ParseColumnList(names)
	test.S(t).ExpectEquals(columnList.Len(), 3)
	test.S(t).ExpectTrue(reflect.DeepEqual(columnList.Names(), []string{"id", "category", "max_len"}))
	test.S(t).ExpectEquals(columnList.Ordinals["id"], 0)
	test.S(t).ExpectEquals(columnList.Ordinals["category"], 1)
	test.S(t).ExpectEquals(columnList.Ordinals["max_len"], 2)
}

func TestGetColumn(t *testing.T) {
	names := "id,category,max_len"
	columnList := ParseColumnList(names)
	{
		column := columnList.GetColumn("category")
		test.S(t).ExpectTrue(column != nil)
		test.S(t).ExpectEquals(column.RawName, "category")
	}
	{
		column := columnList.GetColumn("no_such_column")
		test.S(t).ExpectTrue(column == nil)
	}
}
