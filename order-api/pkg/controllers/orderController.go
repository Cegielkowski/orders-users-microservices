package controllers

import (
	"encoding/json"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/errors"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/responses"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/services/orders"
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/structs"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetByUserOrderController Get the order information of the requested id.
func GetByUserOrderController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["userID"], 10, 32)
	if err != nil {
		errors.Log(err)
		responses.Finish(w, &responses.HttpBadRequest{})
		return
	}

	ret, err := orders.GetByUser(uint32(userID))
	if err != nil {
		responses.Finish(w, errors.GetHttpResponse(err))
		return
	}

	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		responses.Finish(w, errors.GetHttpResponse(err))
	}
}

// GetOrderController Get the order information of the requested id.
func GetOrderController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	orderID, err := strconv.ParseUint(vars["orderID"], 10, 32)
	if err != nil {
		errors.Log(err)
		responses.Finish(w, &responses.HttpBadRequest{})
		return
	}

	ret, err := orders.Get(uint32(orderID))
	if err != nil {
		responses.Finish(w, errors.GetHttpResponse(err))
		return
	}

	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		responses.Finish(w, errors.GetHttpResponse(err))
	}
}

// ListOrderController List the orders.
func ListOrderController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ret, err := orders.List()
	if err != nil {
		responses.Finish(w, errors.GetHttpResponse(err))
		return
	}

	err = json.NewEncoder(w).Encode(ret)
	if err != nil {
		responses.Finish(w, errors.GetHttpResponse(err))
	}
}

// PostOrderController Creates a new order.
func PostOrderController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order structs.PostOrderRequest
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		errors.Log(err)
		responses.Finish(w, &responses.HttpBadRequest{})
		return
	}
	defer func() { r.Body.Close() }()

	created, err := orders.Insert(order.UserID, order.ItemDescription, order.ItemQuantity, order.ItemPrice, order.TotalValue)
	if err != nil {
		errors.Log(err)
		responses.Finish(w, errors.GetHttpResponse(err))
		return
	}

	responses.Finish(w, &responses.HttpCreated{
		Content: created,
	})
}

// PutOrderController Updates an order.
func PutOrderController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	orderID, err := strconv.ParseUint(vars["orderID"], 10, 32)
	if err != nil {
		errors.Log(err)
		responses.Finish(w, &responses.HttpBadRequest{})
		return
	}

	var order structs.PostOrderRequest
	err = json.NewDecoder(r.Body).Decode(&order)

	defer func() { r.Body.Close() }()

	if err != nil {
		errors.Log(err)
		responses.Finish(w, &responses.HttpBadRequest{})
		return
	}

	updated, err := orders.Update(uint32(orderID), order)
	if err != nil {
		errors.Log(err)
		responses.Finish(w, errors.GetHttpResponse(err))
		return
	}

	responses.Finish(w, &responses.HttpSuccessContent{
		Content: updated,
	})
}

// DeleteOrderController Delete a order.
func DeleteOrderController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	orderID, err := strconv.ParseUint(vars["orderID"], 10, 32)
	if err != nil {
		errors.Log(err)
		responses.Finish(w, &responses.HttpBadRequest{})
		return
	}

	err = orders.Delete(uint32(orderID))
	if err != nil {
		responses.Finish(w, errors.GetHttpResponse(err))
		return
	}

	responses.Finish(w, &responses.HttpSuccessContent{})
}

// DeleteOrderController Delete a order.
func DeleteByUserOrderController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["userID"], 10, 32)
	if err != nil {
		errors.Log(err)
		responses.Finish(w, &responses.HttpBadRequest{})
		return
	}

	err = orders.DeleteByUser(userID)
	if err != nil {
		responses.Finish(w, errors.GetHttpResponse(err))
		return
	}

	responses.Finish(w, &responses.HttpSuccessContent{})
}