package board

import (
	"github.com/peterhellberg/giphy"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
	"io"
	"math/rand"
	"net/http"
)

var NoGifFoundURL = "https://media.giphy.com/media/9J7tdYltWyXIY/giphy.gif"

type Board struct {
}

type Gif struct {
	URL  string
	Mood []string
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

	var gifURL string
	if (len(gif.Data)) < 1 {
		gifURL = NoGifFoundURL
	} else {
		gifURL = gif.Data[rand.Intn(len(gif.Data))].MediaURL()
	}

	err = renderTemplate(writer, http.StatusOK, "board.html", Gif{URL: gifURL, Mood: mood})
	if err != nil {
		return errors.WithMessage(err, "rendering template")
	}
	return nil
}
