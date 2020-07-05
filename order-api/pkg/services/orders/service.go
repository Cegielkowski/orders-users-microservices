package orders

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/database"
	dbelastic "github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/database/elastic"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/database/redis"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/errors"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/models"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/structs"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/utils/general"
	"io"
	"net/http"
	"strconv"
)

var userApi = structs.UserApi{Url: "http://localhost:8000/user/"}

// List Makes the request in DB to list the orders.
func List() (*[]structs.OrderResponse, error) {
	orders := []models.Order{}

	reply, err := redis.Get("LISTORDER")
	if err != nil {
		err := database.DB.Find(&orders).Error
		if err != nil {
			errors.Log(err)
			return nil, err
		}
		orderBytes, err := json.Marshal(orders)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
		err = redis.Set("LISTORDER", orderBytes)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	} else {
		err = json.Unmarshal(reply, &orders)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	}

	var response []structs.OrderResponse

	for _, order := range orders {
		err, userStruct := CheckUser(order.UserID)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
		response = append(response, structs.OrderResponse{
			OrderID:         order.ID,
			UserID:          order.UserID,
			ItemDescription: order.ItemDescription,
			ItemQuantity:    order.ItemQuantity,
			ItemPrice:       order.ItemPrice,
			TotalValue:      order.TotalValue,
			User:            userStruct,
			CreatedAt:       order.CreatedAt,
			UpdatedAt:       order.UpdatedAt,
		})
	}
	return &response, nil
}

// Get Makes the request in DB to get the order information by the userID.
func GetByUser(id uint32) (*[]structs.OrderResponse, error) {
	orders := []models.Order{}
	reply, err := redis.Get("GETBYUSER" + general.Stringify(id))
	if err != nil {
		err := database.DB.Where("user_id = ?", id).Find(&orders).Error
		if err != nil {
			errors.Log(err)
			return nil, err
		}
		orderBytes, err := json.Marshal(orders)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
		err = redis.Set("GETBYUSER"+general.Stringify(id), orderBytes)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	} else {
		err = json.Unmarshal(reply, &orders)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	}

	var response []structs.OrderResponse

	for _, order := range orders {
		err, userStruct := CheckUser(order.UserID)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
		response = append(response, structs.OrderResponse{
			OrderID:         order.ID,
			UserID:          order.UserID,
			ItemDescription: order.ItemDescription,
			ItemQuantity:    order.ItemQuantity,
			ItemPrice:       order.ItemPrice,
			TotalValue:      order.TotalValue,
			User:            userStruct,
			CreatedAt:       order.CreatedAt,
			UpdatedAt:       order.UpdatedAt,
		})
	}
	return &response, nil
}

// Get Makes the request in DB to get the order information.
func Get(id uint32) (*structs.OrderResponse, error) {
	order := models.Order{}
	reply, err := redis.Get("GETORDER" + general.Stringify(id))
	if err != nil {
		err := database.DB.Where("id = ?", id).Find(&order).Error
		if err != nil {
			errors.Log(err)
			return nil, err
		}
		orderBytes, err := json.Marshal(order)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
		err = redis.Set("GETORDER"+general.Stringify(id), orderBytes)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	} else {
		err = json.Unmarshal(reply, &order)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	}
	err, userStruct := CheckUser(order.UserID)
	if err != nil {
		errors.Log(err)
		return nil, err
	}
	return &structs.OrderResponse{
		OrderID:         order.ID,
		UserID:          order.UserID,
		ItemDescription: order.ItemDescription,
		ItemQuantity:    order.ItemQuantity,
		ItemPrice:       order.ItemPrice,
		TotalValue:      order.TotalValue,
		User:            userStruct,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}, nil
}

// Delete Makes the request in DB to delete the order information.
func Delete(orderID uint32) error {
	order := models.Order{}
	reply, err := redis.Get("GETORDER" + general.Stringify(orderID))
	if err != nil {
		err := database.DB.Where("id = ?", orderID).Find(&order).Error
		if err != nil {
			errors.Log(err)
			return err
		}
	} else {
		err = json.Unmarshal(reply, &order)
		if err != nil {
			errors.Log(err)
			return err
		}
	}

	err = database.DB.Delete(order).Error
	if err != nil {
		errors.Log(err)
		return err
	}
	_, err = redis.Flush("GETORDER" + general.Stringify(order.ID))
	if err != nil {
		errors.Log(err)
		return err
	}
	_, err = redis.Flush("GETBYUSER" + general.Stringify(order.UserID))
	if err != nil {
		errors.Log(err)
		return err
	}
	_, err = redis.Flush("LISTORDER")
	if err != nil {
		errors.Log(err)
		return err
	}

	return nil
}

// Delete Makes the request in DB to delete the order information.
func DeleteByUser(userId uint64) error {
	order := models.Order{}
	orders, err := GetByUser(uint32(userId))
	if err != nil {
		errors.Log(err)
		return err
	}

	err = database.DB.Where("user_id = " + strconv.FormatUint(userId, 10)).Delete(order).Error
	if err != nil {
		errors.Log(err)
		return err
	}
	for _, ord := range *orders {
		_, err = redis.Flush("GETBYORDER" + general.Stringify(ord.OrderID))
		if err != nil {
			errors.Log(err)
			return err
		}
	}

	_, err = redis.Flush("GETBYUSER" + general.Stringify(uint32(userId)))
	if err != nil {
		errors.Log(err)
		return err
	}
	_, err = redis.Flush("LISTORDER")
	if err != nil {
		errors.Log(err)
		return err
	}

	return nil
}

// Insert a new order in database.
func Insert(userId uint32, iDescription string, iQuantity uint32, iPrice uint32, total uint32) (*structs.OrderResponse, error) {
	err, user := CheckUser(userId)
	if err != nil && user.Name == "" {
		return nil, err
	}

	ctx := context.Background()
	esclient, err := dbelastic.GetESClient()
	if err != nil {
		errors.Log(err)
		return nil, err
	}

	order := models.Order{
		UserID:          userId,
		ItemDescription: iDescription,
		ItemQuantity:    iQuantity,
		ItemPrice:       iPrice,
		TotalValue:      total,
	}

	err = database.DB.Create(&order).Error
	if err != nil {
		errors.Log(err)
		return nil, err
	}

	_, err = redis.Flush("GETBYUSER" + general.Stringify(userId))
	if err != nil {
		errors.Log(err)
		return nil, err
	}
	_, err = redis.Flush("LISTORDER")
	if err != nil {
		errors.Log(err)
		return nil, err
	}

	newOrder := structs.OrderResponse{
		OrderID:         order.ID,
		UserID:          order.UserID,
		ItemDescription: order.ItemDescription,
		ItemQuantity:    order.ItemQuantity,
		ItemPrice:       order.ItemPrice,
		TotalValue:      order.TotalValue,
		User:            user,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}
	dataJSON, err := json.Marshal(newOrder)
	if err != nil {
		errors.Log(err)
		return nil, err
	}

	js := string(dataJSON)
	_, err = esclient.Index().Index("orders").Type("order").BodyJson(js).Do(ctx)
	fmt.Println(err)

	if err != nil {
		errors.Log(err)
		return nil, err
	}

	return &newOrder, nil
}

// Update a new order in database.
func Update(orderID uint32, order structs.PostOrderRequest) (*structs.OrderResponse, error) {
	orderModel := models.Order{}
	reply, err := redis.Get("GETORDER" + general.Stringify(orderID))
	if err != nil {
		err := database.DB.Where("id = ?", orderID).Find(&orderModel).Error
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	} else {
		err = json.Unmarshal(reply, &orderModel)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	}

	err = database.DB.Model(&orderModel).Updates(models.Order{
		ItemDescription: order.ItemDescription,
		ItemQuantity:    order.ItemQuantity,
		ItemPrice:       order.ItemPrice,
		TotalValue:      order.TotalValue,
	}).Error

	if err != nil {
		errors.Log(err)
		return nil, err
	}

	_, err = redis.Flush("GETORDER" + general.Stringify(orderModel.ID))
	if err != nil {
		errors.Log(err)
		return nil, err
	}
	_, err = redis.Flush("GETBYUSER" + general.Stringify(orderModel.UserID))
	if err != nil {
		errors.Log(err)
		return nil, err
	}
	_, err = redis.Flush("LISTORDER")
	if err != nil {
		errors.Log(err)
		return nil, err
	}

	err, userStruct := CheckUser(orderModel.UserID)
	if err != nil {
		errors.Log(err)
		return nil, err
	}
	return &structs.OrderResponse{
		OrderID:         orderModel.ID,
		UserID:          orderModel.UserID,
		ItemDescription: orderModel.ItemDescription,
		ItemQuantity:    orderModel.ItemQuantity,
		ItemPrice:       orderModel.ItemPrice,
		TotalValue:      orderModel.TotalValue,
		User:            userStruct,
		CreatedAt:       orderModel.CreatedAt,
		UpdatedAt:       orderModel.UpdatedAt,
	}, nil
}

func CheckUser(userId uint32) (error, structs.UserResponse) {
	emptyUser := structs.UserResponse{}
	var user structs.UserResponse
	resp, err := http.Get(userApi.Url + strconv.FormatUint(uint64(userId), 10))
	if err != nil {
		errors.Log(err)
		return err, emptyUser
	}

	if resp.Body == http.NoBody {
		return invalidUser(), emptyUser
	} else {
		defer func() { resp.Body.Close() }()
	}

	err = json.NewDecoder(resp.Body).Decode(&user)
	switch {
	case err == io.EOF:
		return invalidUser(), emptyUser
	case err != nil:
		return err, emptyUser
	}

	if user == emptyUser {
		return invalidUser(), emptyUser
	}

	return nil, user
}
