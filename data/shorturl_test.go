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
		assert.Len(t, url, 6)
	}
}
