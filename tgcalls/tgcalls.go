package tgcalls

import (
	"errors"
	"log"
	"os"
	"strconv"

	"GoMusicBot/queues"

	"github.com/gotgcalls/gotgcalls"
)

const CLIENT = "main"

var GoTGCalls *gotgcalls.GoTGCalls

var clientInitialized bool
var stringSession string
var apiId int
var apiHash string

func OnFinish(client string, chatId int64) bool {
	item := queues.Pull(chatId)
	if item == nil {
		stopResult, _ := GoTGCalls.Stop(client, chatId)
		return stopResult != gotgcalls.OK
	}

	Get().Stream(client, chatId, item.(string))
	return true
}

func Start() error {
	if GoTGCalls == nil {
		GoTGCalls = gotgcalls.NewGoTGCalls()
		GoTGCalls.OnFinish = func(client string, chatId int64) {
			OnFinish(client, chatId)
		}

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
