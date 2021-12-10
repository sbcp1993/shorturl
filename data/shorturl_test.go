package data

import (
	"os"
	"shorturl/db"
	"strings"
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
				url: "https://docs.microsoft.com/en-us/azure/devops/report/powerbi/data-connector-connect?view=azure-devops",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		url, err := AddShorturl(tt.args.con, tt.args.url)
		assert.NoError(t, err)
		assert.NotEmpty(t, url)
		assert.Len(t, url.UrlString, 30)
	}
}

func TestGetRealurl(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	postgres, err := db.NewDBConnection()
	assert.NoError(t, err)
	assert.NotNil(t, postgres)

	testurl := "https://docs.microsoft.com/en-us/azure/devops/report/powerbi/data-connector-connect?view=azure-devops"
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
				url: strings.Split(surl.UrlString, baseurl)[1],
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
