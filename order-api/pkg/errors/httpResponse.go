package errors

import (
	"github.com/SelecaoSerasaConsumidor/BE_LucasCegielkowski/order-api/pkg/responses"
	"strings"
)

// GetHttpResponse Gets the correct Http response to the given error.
func GetHttpResponse(err error) responses.HttpResponse {
	if strings.Contains(err.Error(), "Duplicate entry") {
		return &responses.HttpConflict{}
	}

	if strings.Contains(err.Error(), "record not found") {
		return &responses.HttpNoContent{}
	}

	if strings.Contains(err.Error(), "the user does not exist") {
		return &responses.HttpNotFound{}
	}

	if strings.Contains(err.Error(), "a foreign key constraint fails") ||
		strings.Contains(err.Error(), "invalid amount given"){
		return &responses.HttpUnprocessableEntity{}
	}


	return &responses.HttpInternalError{}
}