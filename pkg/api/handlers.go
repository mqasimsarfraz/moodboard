package api

import (
	"fmt"
	"github.com/MQasimSarfraz/moodboard/pkg/board"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var mood = "hello"

type Api struct {
	Board *board.Board
}

func NewApi(board *board.Board) *Api {
	return &Api{board}
}

func register(api *Api) http.Handler {
	router := httprouter.New()

	// register router handlers here
	router.GET("/ping", api.ping)
	router.GET("/", api.handleIndex)
	router.PUT("/mood/:mood", api.handleMoodUpdate)

	return router
}

func (api *Api) ping(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprint(writer, "pong")
}

func (api *Api) handleIndex(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	api.Board.Render(writer, mood)
}

func (api *Api) handleMoodUpdate(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	mood = params.ByName("mood")
}
