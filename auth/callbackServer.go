package auth

import (
	"io"
	"log"
	"net/http"
)

type CallbackServer struct {
	channel chan string
}

func (a *CallbackServer) HandleCallback(w http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	a.channel <- code
	_, _ = io.WriteString(w, "Successful. You can close the tab now.\n")
}

func (a *CallbackServer) GetAuthToken() string {
	return <-a.channel
}

func (a *CallbackServer) Init() {
	go func() {
		http.HandleFunc("/", a.HandleCallback)
		log.Fatal(http.ListenAndServe(":8787", nil))
	}()
}

func CreateAuthCallbackServer() CallbackServer {
	server := CallbackServer{make(chan string)}
	server.Init()
	return server
}
