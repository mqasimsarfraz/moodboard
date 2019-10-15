package board

import (
	"github.com/peterhellberg/giphy"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
	"io"
	"math/rand"
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
	gif, err := g.Search(mood)
	if err != nil {
		return errors.WithMessage(err, "getting gif")
	}

	if (len(gif.Data)) < 1 {
		return ErrGifNotFound
	}

	err = renderTemplate(writer, http.StatusOK, "board.html", gif.Data[rand.Intn(g.Limit)].MediaURL())
	if err != nil {
		return errors.WithMessage(err, "rendering template")
	}
	return nil
}
