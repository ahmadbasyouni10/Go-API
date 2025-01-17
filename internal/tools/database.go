package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

// bc these funcs are in the interface, they have to be implemented
// by a var db that is a pointer to the mockdb struct
// and has databaseinterface as its type
// doing (*database).func will destruct the pointer and call a method
// that the db will have bc its type of databaseinterface
// and this interface has the methods GetUserLoginDetails and GetUserCoins
// we use db interface so we can change the database to a real one later
// without changing the handlers or middleware code
// since the handlers and middleware only know about the interface
// a connection isnt needed in the interface so we can use a mockdb
// also in main we can use the real db and in tests we can use the mockdb
type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

// database has interface which allows it to do methods, but then is an instance
// of the mockDB struct which is empty right now cuz mock but the & is so we
// can pass the pointer of the database we made to the functions getusercoins
// and getuserlogindetails and setupdatabase
func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil

}
