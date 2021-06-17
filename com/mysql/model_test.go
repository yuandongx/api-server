package mysql

import (
	"reflect"
	"testing"
)

func TestModel_CloneRow(t *testing.T) {
	type fields struct {
		name       string
		_type      string
		indexs     []int
		id         string
		fieldKinds []reflect.Kind
		columns    []string
		values     []interface{}
		table      string
		count      int
	}
	tests := []struct {
		name    string
		fields  fields
		wantRow []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				name:       tt.fields.name,
				_type:      tt.fields._type,
				indexs:     tt.fields.indexs,
				id:         tt.fields.id,
				fieldKinds: tt.fields.fieldKinds,
				columns:    tt.fields.columns,
				values:     tt.fields.values,
				table:      tt.fields.table,
				count:      tt.fields.count,
			}
			if gotRow := m.CloneRow(); !reflect.DeepEqual(gotRow, tt.wantRow) {
				t.Errorf("CloneRow() = %v, want %v", gotRow, tt.wantRow)
			}
		})
	}
}

func TestModel_SetColumn(t *testing.T) {
	type fields struct {
		name       string
		_type      string
		indexs     []int
		id         string
		fieldKinds []reflect.Kind
		columns    []string
		values     []interface{}
		table      string
		count      int
	}
	type args struct {
		v      interface{}
		values []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				name:       tt.fields.name,
				_type:      tt.fields._type,
				indexs:     tt.fields.indexs,
				id:         tt.fields.id,
				fieldKinds: tt.fields.fieldKinds,
				columns:    tt.fields.columns,
				values:     tt.fields.values,
				table:      tt.fields.table,
				count:      tt.fields.count,
			}
			m.init(tt.args.v)
		})
	}
}

//func Test_convert(t *testing.T) {
//	type args struct {
//		kind  reflect.Kind
//		field reflect.Value
//	}
//	tests := []struct {
//		name      string
//		args      args
//		wantValue interface{}
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if gotValue := convert(tt.args.kind, tt.args.field); !reflect.DeepEqual(gotValue, tt.wantValue) {
//				t.Errorf("convert() = %v, want %v", gotValue, tt.wantValue)
//			}
//		})
//	}
//}

func Test_tableName(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"tableName",
			args{s: "abc"},
			"app-abc-v1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tableName(tt.args.s); got != tt.want {
				t.Errorf("tableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
