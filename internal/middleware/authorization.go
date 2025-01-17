package middleware

import (
	"errors"
	"net/http"

	"github.com/ahmadbasyouni10/Go-API/api"
	"github.com/ahmadbasyouni10/Go-API/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Unauthorized request username or token is wrong/misssing")

// Authorization middleware takes a next http.handler which is the next handler in the chain from api.go in our case would be router.get(/coins)
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		// middleware checks if the username and token are empty
		// so then in the handler that comes after we are guaranteed to have a username and token
		// bc the middleware checks for it
		// we can return the coin balance for the user

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			// uses predefined requestErrorhandler which takes the writer and the error
			// we use the request error handler and the internal error handler
			// bc we want to return the error to the client
			// but if its internal they dont need to know the exact error
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})

}
