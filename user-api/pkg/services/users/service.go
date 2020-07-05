package users

import (
	"encoding/json"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/database"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/database/redis"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/errors"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/models"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/structs"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/user-api/pkg/utils/general"
	"net/http"
)

var key = []byte("123-432-456-23-11-32-23415-1234!")
var orderApi = structs.OrderApi{StartOfUrl: "http://localhost:8001/order/", EndOfUrl: "/user"}

// List Makes the request in DB to list the users.
func List() (*[]structs.UserResponse, error) {
	users := []models.User{}
	reply, err := redis.Get("LISTUSER")
	if err != nil {
		err := database.DB.Find(&users).Error
		if err != nil {
			errors.Log(err)
			return nil, err
		}
		userBytes, err := json.Marshal(users)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
		err = redis.Set("LISTUSER", userBytes)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	} else {
		err = json.Unmarshal(reply, &users)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	}
	var response []structs.UserResponse

	for _, user := range users {
		response = append(response, structs.UserResponse{
			UserID:    user.ID,
			Cpf:       general.Decrypt(key, user.Cpf),
			Phone:     general.Decrypt(key, user.Phone),
			Name:      user.Name,
			Email:     general.Decrypt(key, user.Email),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return &response, nil
}

// Get Makes the request in DB to get the user information.
func Get(userID uint32) (*structs.UserResponse, error) {
	user := models.User{}
	reply, err := redis.Get("GETUSER" + general.Stringify(userID))
	if err != nil {
		err = database.DB.Where("id = ?", userID).Find(&user).Error
		if err != nil {
			errors.Log(err)
			return nil, err
		}
		userBytes, err := json.Marshal(user)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
		err = redis.Set("GETUSER" + general.Stringify(userID), userBytes)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	} else {
		err = json.Unmarshal(reply, &user)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	}

	return &structs.UserResponse{
		UserID:    user.ID,
		Cpf:       general.Decrypt(key, user.Cpf),
		Phone:     general.Decrypt(key, user.Phone),
		Name:      user.Name,
		Email:     general.Decrypt(key, user.Email),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Delete Makes the request in DB to delete the user information.
func Delete(userID uint32) (error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", orderApi.StartOfUrl + general.Stringify(userID) + orderApi.EndOfUrl, nil)
	if err != nil {
		errors.Log(err)
		return err
	}
	if req.Body != nil {
		defer func() { req.Body.Close() }()
	}

	// Fetch Request.
	resp, err := client.Do(req)
	if err != nil {
		return invalidDelete()
	}

	if resp.StatusCode != 200 {
		return invalidDelete()
	}

	user := models.User{}
	reply, err := redis.Get("GETUSER" + general.Stringify(userID))
	if err != nil {
		err = database.DB.Where("id = ?", userID).Find(&user).Error
		if err != nil {
			errors.Log(err)
			return err
		}
	} else {
		err = json.Unmarshal(reply, &user)
		if err != nil {
			errors.Log(err)
			return err
		}
	}
	if user.Name == "" {
		return nil
	}

	err = database.DB.Delete(user).Error

	if err != nil {
		errors.Log(err)
		return err
	}

	_, err = redis.Flush("GETUSER" + general.Stringify(userID))
	if err != nil {
		errors.Log(err)
		return err
	}
	_, err = redis.Flush("LISTUSER")
	if err != nil {
		errors.Log(err)
		return err
	}

	return nil
}

// Insert a new user in database.
func Insert(email string, cpf string, phone string, name string) (*structs.UserResponse, error) {
	user := models.User{
		Email: general.Encrypt(key, email),
		Cpf:   general.Encrypt(key, cpf),
		Phone: general.Encrypt(key, phone),
		Name:  name,
	}

	err := database.DB.Create(&user).Error
	if err != nil {
		errors.Log(err)
		return nil, err
	}

	_, err = redis.Flush("LISTUSER")
	if err != nil {
		errors.Log(err)
		return nil, err
	}

	return &structs.UserResponse{
		UserID:    user.ID,
		Cpf:       general.Decrypt(key, user.Cpf),
		Email:     general.Decrypt(key, user.Email),
		Name:      user.Name,
		Phone:     general.Decrypt(key, user.Phone),
		CreatedAt: user.CreatedAt,
	}, nil
}

// Update a new user in database.
func Update(userID uint32, u structs.PostUserRequest) (*structs.UserResponse, error) {
	user := models.User{}

	reply, err := redis.Get("GETUSER" + general.Stringify(userID))
	if err != nil {
		err := database.DB.Where("id = ?", userID).Find(&user).Error
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	} else {
		err = json.Unmarshal(reply, &user)
		if err != nil {
			errors.Log(err)
			return nil, err
		}
	}

	var cpfEncrypted string
	var emailEncrypted string
	var phoneEncrypted string

	if u.Cpf != "" {
		cpfEncrypted = general.Encrypt(key, u.Cpf)
	}
	if u.Email != "" {
		emailEncrypted = general.Encrypt(key, u.Email)
	}
	if u.Phone != "" {
		phoneEncrypted = general.Encrypt(key, u.Phone)
	}

	err = database.DB.Model(&user).Updates(models.User{
		Cpf:       cpfEncrypted,
		Email:     emailEncrypted,
		Name:      u.Name,
		Phone:     phoneEncrypted,
	}).Error

	if err != nil {
		errors.Log(err)
		return nil, err
	}

	_, err = redis.Flush("GETUSER" + general.Stringify(userID))
	if err != nil {
		errors.Log(err)
		return nil, err
	}
	_, err = redis.Flush("LISTUSER")
	if err != nil {
		errors.Log(err)
		return nil, err
	}

	return &structs.UserResponse{
		UserID:    user.ID,
		Cpf:       general.Decrypt(key, user.Cpf),
		Email:     general.Decrypt(key, user.Email),
		Name:      user.Name,
		Phone:     general.Decrypt(key, user.Phone),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}