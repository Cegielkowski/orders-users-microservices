package users

import (
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/config"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/database"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/models"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/structs"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/utils/general"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const nameTest = "TESTINSERT"
const cpfTest = "41865586898"
const phoneTest = "40028922"
const emailTest = "testee@emai.com"

var existing *structs.UserResponse
var notExistingID uint

func init() {
	testDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(filepath.Dir(filepath.Dir(testDir)))

	err = general.ReadConfigJson(&config.CONFIG, dir+"/config/tests_config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = database.CreateDB()
	if err != nil {
		log.Fatal(err)
	}
	database.DB.LogMode(false)

	user := models.User{}

	if database.DB.Last(&user).RecordNotFound() {
		_, err := Insert(emailTest, cpfTest, phoneTest, nameTest)
		if err != nil {
			log.Fatal(err)
		}
		database.DB.Last(&user)
	}

	existing = &structs.UserResponse{
		UserID:    user.ID,
		Cpf:       general.Decrypt(key, user.Cpf),
		Phone:     general.Decrypt(key, user.Phone),
		Email:     general.Decrypt(key, user.Email),
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	notExistingID = uint(user.ID + 900)
}

func TestGet(t *testing.T) {
	type args struct {
		userID uint
	}
	tests := []struct {
		name    string
		args    args
		want    *structs.UserResponse
		wantErr bool
	}{
		{
			name:    "Exists",
			args:    args{userID: uint(existing.UserID)},
			want:    existing,
			wantErr: false,
		},
		{
			name:    "NotExists",
			args:    args{userID: notExistingID},
			want:    nil,
			wantErr: true,
		},
	}
	defer func() {
		database.DB.Unscoped().Delete(models.User{}, "name LIKE ?", nameTest)
	}()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(uint32(tt.args.userID))
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name    string
		cpf     string
		wantErr bool
	}{
		{
			name:    "First insert Unique cpf num",
			cpf:     cpfTest,
			wantErr: false,
		},
		{
			name:    "Second insert non-Unique cpf num",
			cpf:     cpfTest,
			wantErr: true,
		},
	}
	defer func() {
		database.DB.Unscoped().Delete(models.User{}, "name LIKE ?", nameTest)
	}()
	for _, tt := range tests {
		t.Run(tt.cpf, func(t *testing.T) {
			if _, err := Insert(emailTest, tt.cpf, phoneTest, nameTest); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v, cpf %v", err, tt.wantErr, tt.cpf)
			}
		})
	}
}
