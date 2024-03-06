package gonebot

import (
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func UpgradeWebsocket(w http.ResponseWriter, r *http.Request) error {
	xSelfId := r.Header.Get("x-self-id")
	botId, err := strconv.ParseInt(xSelfId, 10, 64)
	if err != nil {
		return err
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	NewBot(botId, c)
	return nil
}
