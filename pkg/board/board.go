package board

import (
	"github.com/unrolled/render"
	"io"
	"net/http"
)

type Board struct {
}

func NewBoard() *Board {
	return &Board{}
}

var renderTemplate = render.New().HTML

func (b *Board) Render(writer io.Writer, mood string) {
	renderTemplate(writer, http.StatusOK, "board.html", mood)
}
