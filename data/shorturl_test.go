package data

import (
	"os"
	"shorturl/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddShorturl(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	postgres, err := db.NewDBConnection()
	assert.NoError(t, err)
	assert.NotNil(t, postgres)

	type args struct {
		con *db.DBConnection
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid case",
			args: args{
				con: postgres,
				url: "https://go.dev/tour/welcome/1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		url, err := AddShorturl(tt.args.con, tt.args.url)
		assert.NoError(t, err)
		assert.NotEmpty(t, url)
		assert.Len(t, url, 8)
	}
}

func TestGetRealurl(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	postgres, err := db.NewDBConnection()
	assert.NoError(t, err)
	assert.NotNil(t, postgres)

	testurl := "https://gobyexample.com/channels"
	surl, err := AddShorturl(postgres, testurl)
	assert.NoError(t, err)

	type args struct {
		con *db.DBConnection
		url string
	}
	tests := []struct {
		name         string
		args         args
		wantReal_url string
		wantErr      bool
	}{
		{
			name: "valid case",
			args: args{
				con: postgres,
				url: surl.UrlString,
			},
		},
	}
	for _, tt := range tests {
		url, err := GetRealurl(tt.args.con, tt.args.url)
		assert.NoError(t, err)
		assert.NotEmpty(t, url)
		assert.Equal(t, testurl, url)
	}
}
