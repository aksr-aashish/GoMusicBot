package tgcalls

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/gotgcalls/gotgcalls"
)

const CLIENT = "main"

var GoTGCalls *gotgcalls.GoTGCalls

var clientInitialized bool
var stringSession string
var apiId int
var apiHash string

func Start() error {
	if GoTGCalls == nil {
		GoTGCalls = gotgcalls.NewGoTGCalls()

		stringSession, apiHash = os.Getenv("STRING_SESSION"), os.Getenv("API_HASH")
		apiId, _ = strconv.Atoi(os.Getenv("API_ID"))
		if apiId == 0 {
			return errors.New("invalid API_ID")
		}

		return GoTGCalls.Start()
	}

	return nil
}

func Get() *gotgcalls.GoTGCalls {
	if !clientInitialized {
		_, err := GoTGCalls.InitClient(CLIENT, stringSession, apiId, apiHash)
		if err != nil {
			log.Fatal("Error initializing client: " + err.Error())
		}

		clientInitialized = true
		return GoTGCalls
	}

	return GoTGCalls
}
