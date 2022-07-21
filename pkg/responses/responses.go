package responses

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"test/pkg/customerror"
	errorCustomStatus "test/pkg/error"
)

type rest struct {
	Error   bool        `json:"error"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success Serve Data With Success Status
func Success(w http.ResponseWriter, data interface{}) {
	result := rest{
		Error:   false,
		Status:  200,
		Message: "SUCCESS",
		Data:    data,
	}

	responses, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		Error(w, customerror.ErrInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(result.Status)
	w.Write(responses)
	return
}

// Error Serve Data With Error Status
func Error(w http.ResponseWriter, err error) {
	var result = rest{}
	customErrors, _ := err.(*errorCustomStatus.Error)

	if customErrors == nil {
		result = rest{
			Error:   true,
			Status:  500,
			Message: customerror.ErrInternalServerError.Title,
			Data:    customerror.ErrInternalServerError.Detail,
		}
	} else {
		status, err := strconv.Atoi(customErrors.Status)
		if err != nil {
			log.Println(err)
			Error(w, customerror.ErrInternalServerError)
			return
		}
		result = rest{
			Error:   true,
			Status:  status,
			Message: customErrors.Title,
			Data:    customErrors.Detail,
		}
	}

	responses, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		Error(w, customerror.ErrInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(result.Status)
	w.Write(responses)
}
