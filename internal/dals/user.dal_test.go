package dals

import (
	"context"
	"os"
	"testing"

	"github.com/mainlycricket/go-mongo/internal/database"
	"github.com/mainlycricket/go-mongo/internal/database/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestUserDal(t *testing.T) {
	conn, err := database.Connect(os.Getenv("CLUSTER_URL"))
	if err != nil {
		t.Errorf("db connection failed: %v", err)
		return
	}
	userDal := NewUserDal(conn.Database(os.Getenv("DB_NAME") + "_test"))

	type args struct {
		ctx  context.Context
		user *models.User
	}

	tests := []struct {
		name    string
		args    args
		got     bson.ObjectID
		wantErr bool
	}{
		{
			name: "simple user",
			args: args{
				ctx: context.Background(),
				user: &models.User{
					Name:     "tushar",
					Email:    "tushar@gmail.com",
					Password: "secret",
				},
			},
			wantErr: false,
		},
		{
			name: "duplicate user",
			args: args{
				ctx: context.Background(),
				user: &models.User{
					Name:     "tushar",
					Email:    "tushar@gmail.com",
					Password: "secret",
				},
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			id, err := userDal.InsertOne(test.args.ctx, test.args.user)
			if err != nil != test.wantErr {
				t.Logf("InsertOne Error: %v", err)
				t.FailNow()
			}

			if err == nil {
				if _, err := userDal.ReadById(test.args.ctx, id); err != nil {
					t.Logf("ReadById Error: %v", err)
					t.Fail()
				}
			}
		})
	}

	if _, err := userDal.ReadAll(context.Background()); err != nil {
		t.Logf("ReadAll Error: %v", err)
		t.Fail()
	}
}
