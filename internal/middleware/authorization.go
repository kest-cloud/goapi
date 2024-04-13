package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/kest-cloud/goapi/api"
	"github.com/kest-cloud/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnauthorizationError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error
		if username == "" || token == "" {
			fmt.Println(err)
			log.Error(UnauthorizationError)
			api.RequestErrorHandler(w, UnauthorizationError)
			return
		}

		var database *tools.DatabaseInterface
		fmt.Println(database)
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnauthorizationError)
			api.RequestErrorHandler(w, UnauthorizationError)
			return
		}
		next.ServeHTTP(w, r)
	})

}
