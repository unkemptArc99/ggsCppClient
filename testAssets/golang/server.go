package main

import (
	"fmt"
	"io"
	"net/http"

	gopher "github.com/hewiefreeman/GopherGameServer"
	"github.com/hewiefreeman/GopherGameServer/core"
)

func main() {
	// Set the core server settings
	settings := gopher.ServerSettings{
		ServerName:     "!s!",
		MaxConnections: 0,

		HostName:  "localhost",
		HostAlias: "localhost",
		IP:        "localhost",
		Port:      8080,

		AdminLogin:    "admin",
		AdminPassword: "admin",
	}

	gopher.SetClientConnectCallback(onClientConnect)
	// Make a Room type and set broadcasts and callbacks
	chatRoomType := core.NewRoomType("chat", true)
	chatRoomType.EnableBroadcastUserEnter().EnableBroadcastUserLeave().
		SetUserEnterCallback(onEnterChat).SetUserLeaveCallback(onLeaveChat)

	// Open a Room
	_, roomErr := core.NewRoom("chatExample", "chat", false, 0, "")
	if roomErr != nil {
		fmt.Println("Error while opening Room:", roomErr)
		return
	}

	gopher.Start(&settings)
}

func onClientConnect(writer *http.ResponseWriter, request *http.Request) bool {
	message := "Welcome to the server"
	_, e := io.WriteString(*writer, message)
	return e == nil
}

func onEnterChat(room *core.Room, user *core.RoomUser) {
	// Example of using parameters to send a welcome message to the entering User
	message := "Welcome! Please read the chat room rules, and have fun!"
	messageErr := room.ServerMessage(message, core.ServerMessageNotice, []string{user.User().Name()})
	if messageErr != nil {
		fmt.Println("Error while messaging User:", messageErr)
	}
}

func onLeaveChat(room *core.Room, user *core.RoomUser) {
	// ...

	// To convert RoomUser to User:
	// u := user.User()
}
