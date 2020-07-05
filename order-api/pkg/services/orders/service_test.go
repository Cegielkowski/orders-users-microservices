package orders

import (
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/config"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/database"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/models"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/structs"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/utils/general"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const userId uint32 = 45
const itemDesc = "Item Test"
const itemQuant uint32 = 10993
const itemPric uint32 = 23
const totalVal uint32 = 5000

var existing *structs.OrderResponse
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

	order := models.Order{}

	if database.DB.Last(&order).RecordNotFound() {
		_, err := Insert(userId, itemDesc, itemQuant, itemPric, totalVal)
		if err != nil {
			log.Fatal(err)
		}
		database.DB.Last(&order)
	}

	existing = &structs.OrderResponse{
		OrderID:         order.ID,
		UserID:          order.UserID,
		ItemDescription: order.ItemDescription,
		ItemPrice:       order.ItemPrice,
		ItemQuantity:    order.ItemQuantity,
		TotalValue:      order.TotalValue,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}
	notExistingID = uint(order.ID + 900)
}

func TestGet(t *testing.T) {
	type args struct {
		orderID uint
	}
	tests := []struct {
		name    string
		args    args
		want    *structs.OrderResponse
		wantErr bool
	}{
		{
			name:    "Exists",
			args:    args{orderID: uint(existing.OrderID)},
			want:    existing,
			wantErr: false,
		},
		{
			name:    "NotExists",
			args:    args{orderID: notExistingID},
			want:    nil,
			wantErr: true,
		},
	}
	defer func() {
		database.DB.Unscoped().Delete(models.Order{}, "item_description LIKE ?", itemDesc)
	}()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(uint32(tt.args.orderID))
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
