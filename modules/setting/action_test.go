package setting

import "testing"

func TestCreateAccessCredentialTable(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"create access credential table test",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateAccessCredentialTable(); (err != nil) != tt.wantErr {
				t.Errorf("CreateAccessCredentialTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccessCredentials_add(t *testing.T) {
	type fields struct {
		Id             int64
		Name           string
		UserName       string
		Password       string
		BecomeMethod   string
		BecomeUser     string
		BecomePassword string
		PublicKey      string
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"insert a oject",
			fields{
				Name:           "abcd",
				UserName:       "abcd",
				Password:       "abcd@12345",
				BecomeUser:     "anbd",
				BecomeMethod:   "su",
				BecomePassword: "abcd@12345",
				PublicKey:      "qwer@12345",
			},
			0,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AccessCredentials{
				Id:             tt.fields.Id,
				Name:           tt.fields.Name,
				UserName:       tt.fields.UserName,
				Password:       tt.fields.Password,
				BecomeMethod:   tt.fields.BecomeMethod,
				BecomeUser:     tt.fields.BecomeUser,
				BecomePassword: tt.fields.BecomePassword,
				PublicKey:      tt.fields.PublicKey,
			}
			got, err := a.add()
			if (err != nil) != tt.wantErr {
				t.Errorf("add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.want {
				t.Errorf("add() got = %v, want %v", got, tt.want)
			}
		})
	}
}
