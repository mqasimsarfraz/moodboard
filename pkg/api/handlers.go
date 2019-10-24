package api

import (
	"encoding/json"
	"fmt"
	"github.com/MQasimSarfraz/moodboard/pkg/board"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

var mood = []string{"hello", "world"}

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
	router.OPTIONS("/mood/:mood", api.handleCorsPreflight)
	router.GET("/mood", api.handleMoodGet)

	return router
}

func (api *Api) ping(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	fmt.Fprint(writer, "pong")
}

func (api *Api) handleIndex(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/html")
	api.Board.Render(writer, mood)
}

func (api *Api) handleMoodUpdate(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	api.setCorsHeaders(writer)
	mood = strings.Split(params.ByName("mood"), " ")
	resp, err := json.Marshal(&Response{Mood: mood})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(resp)
}

func (api *Api) handleMoodGet(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	resp, err := json.Marshal(&Response{Mood: mood})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(resp)
}

func (api *Api) handleCorsPreflight(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	api.setCorsHeaders(writer)
}

func (api *Api) setCorsHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, PUT")
}

type Response struct {
	Mood []string `json:"mood"`
}
