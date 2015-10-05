package rest

import (
	"fmt"
	"net/http"

	"github.com/bitraf/overlord/model"
)

func errorJson(w http.ResponseWriter, code int, message string, args ...interface{}) {
	err := model.Error{Code: code, Message: message}

	if len(args) > 0 {
		err.Message = fmt.Sprintf(message, args)
	}

	serveJson(w, code, err)
}
