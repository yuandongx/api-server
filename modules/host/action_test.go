package host

import (
	"ping/com/mysql"
	"reflect"
	"testing"
)

func Test_addHost(t *testing.T) {
	type args struct {
		h Switch
	}
	tests := []struct {
		name    string
		args    args
		wantId  int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "addhosts",
			args: args{Switch{
				Hostname:   "abc",
				Address:    "1.1.1.14",
				Port:       22,
				Group:      "abcd",
				Comment:    "abcd",
				UserName:   "abcd",
				Password:   "abcd@1234",
				StatusCode: 24}},
			wantId:  0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := addHost(&tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("addHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId == tt.wantId {
				t.Errorf("addHost() gotId = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func Test_create_table(t *testing.T) {
	h := Switch{}
	mysql.Create(&h)
}

func Test_getAllHost(t *testing.T) {
	type args struct {
		obj mysql.Object
	}
	tests := []struct {
		name string
		args args
		want []Device
	}{
		// TODO: Add test cases.
		{
			name: "all hosts",
			args: args{obj: &Switch{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAll(tt.args.obj); reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllHost() = %v, want %v", got, tt.want)
			}
		})
	}
}
