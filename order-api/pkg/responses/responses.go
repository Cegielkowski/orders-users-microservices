package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type HttpResponse interface{
	GetMessage() []byte
	GetStatusCode() int
}

type SimpleMessage struct{
	Msg string `json:"msg"`
}

// Finish Finishes the request by putting the status code and writing a response.
func Finish(w http.ResponseWriter, h HttpResponse){
	// Not necessary if 200.
	if h.GetStatusCode() != 200 {
		w.WriteHeader(h.GetStatusCode())
	}
	_, _ = w.Write(h.GetMessage())
}

// HttpSuccessContent 200.
type HttpSuccessContent struct{
	Content interface{}
}

func (h *HttpSuccessContent) GetMessage() []byte {
	b, err := json.Marshal(h.Content)
	if err != nil {
		log.Println(err)
	}

	return b
}

func  (h *HttpSuccessContent) GetStatusCode() int {
	return 200
}

// HttpCreated 201.
type HttpCreated struct{
	Content interface{}
}

func (h *HttpCreated) GetMessage() []byte {
	b, err := json.Marshal(h.Content)
	if err != nil {
		log.Println(err)
	}

	return b
}

func  (h *HttpCreated) GetStatusCode() int {
	return http.StatusCreated
}

// HttpNoContent 204.
type HttpNoContent struct{}

func (h *HttpNoContent) GetMessage() []byte {
	return nil
}

func  (h *HttpNoContent) GetStatusCode() int {
	return http.StatusNoContent
}

// HttpBadRequest 400.
type HttpBadRequest struct{}

func (h *HttpBadRequest) GetMessage() []byte {
	b, err := json.Marshal(SimpleMessage{"The request sent doesn't match it's requirements"})
	if err != nil {
		log.Println(err)
	}

	return b
}

func  (h *HttpBadRequest) GetStatusCode() int {
	return http.StatusBadRequest
}

// HttpNotFound 404.
type HttpNotFound struct{}

func (h *HttpNotFound) GetMessage() []byte {
	b, err := json.Marshal(SimpleMessage{"the user does not exist"})
	if err != nil {
		log.Println(err)
	}

	return b
}

func  (h *HttpNotFound) GetStatusCode() int {
	return http.StatusNotFound
}


// HttpConflict 409.
type HttpConflict struct{}

func (h *HttpConflict) GetMessage() []byte {
	b, err := json.Marshal(SimpleMessage{"There was a conflict in the sent request"})
	if err != nil {
		log.Println(err)
	}

	return b
}

func  (h *HttpConflict) GetStatusCode() int {
	return http.StatusConflict
}

// HttpUnprocessableEntity 422.
type HttpUnprocessableEntity struct{}

func (h *HttpUnprocessableEntity) GetMessage() []byte {
	b, err := json.Marshal(SimpleMessage{"There is a problem with the content of the entity requested"})
	if err != nil {
		log.Println(err)
	}

	return b
}

func  (h *HttpUnprocessableEntity) GetStatusCode() int {
	return http.StatusUnprocessableEntity
}

// HttpInternalError 500.
type HttpInternalError struct{}

func (h *HttpInternalError) GetMessage() []byte {
	b, err := json.Marshal(SimpleMessage{"There was an intern error"})
	if err != nil {
		log.Println(err)
	}

	return b
}

func  (h *HttpInternalError) GetStatusCode() int {
	return http.StatusInternalServerError
}