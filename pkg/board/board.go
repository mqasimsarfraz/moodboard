package board

import (
	"github.com/peterhellberg/giphy"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
	"io"
	"net/http"
)

var ErrGifNotFound = errors.New("No gif found")

type Board struct {
}

func NewBoard() *Board {
	return &Board{}
}

var renderTemplate = render.New().HTML

func (b *Board) Render(writer io.Writer, mood []string) error {

	g := giphy.DefaultClient
	g.Limit = 1

	gif, err := g.Search(mood)
	if err != nil {
		return errors.WithMessage(err, "getting gif")
	}

	if (len(gif.Data)) < 1 {
		return ErrGifNotFound
	}

	err = renderTemplate(writer, http.StatusOK, "board.html", gif.Data[0].MediaURL())
	if err != nil {
		return errors.WithMessage(err, "rendering template")
	}
	return nil
}
