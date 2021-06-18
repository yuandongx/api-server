package mysql

import (
	"testing"
)

type TestCaseStruct struct {
	Tname    string `sql:"name" len:"257"`
	Tinfo    string `sql:"info"`
	Tage     int    `sql:"age"`
	Tworking bool   `sql:"working"`
}

func TestCreate(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"2343",
			args{&TestCaseStruct{}},
			0,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery(t *testing.T) {
	result := Query("SELECT * FROM `xuyuandong`.`app-accesscredentials-v1` ORDER BY `id`;", 8)
	t.Log(result)
}
