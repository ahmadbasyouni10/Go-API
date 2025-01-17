package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ahmadbasyouni10/Go-API/api"
	"github.com/ahmadbasyouni10/Go-API/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	// decode the query params into the params struct
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// make a new database connection
	// here we wouldnt change code if we changed the database
	// bc we are using the interface which has the methods we need
	// and we just need the new db to implement the interface methods
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	// parantheses and * destruct the pointer and call the method
	tokenDetails = (*database).GetUserCoins(params.Username)

	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	// encodode the struct response into json and write it to the response writer
	// when we do json.NewEncoder(w).Encode(response) it writes the json to the response writer
	// the part that specicially makes the writer send info
	// to the client is when we do encode which writes the json to the writer
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
