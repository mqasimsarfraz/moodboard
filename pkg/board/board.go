package board

import (
	"github.com/peterhellberg/giphy"
	"github.com/unrolled/render"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

const NoGifFoundURL = "https://media.giphy.com/media/9J7tdYltWyXIY/giphy.gif"
const InternalServerErrorURL = "https://media.giphy.com/media/OJBrW6nM5hoNW/giphy.gif"

type Board struct {
}

type Gif struct {
	URL  string
	Mood []string
}

func NewBoard() *Board {
	return &Board{}
}

var templateFunctions = map[string]interface{}{
	"StringJoin": strings.Join,
	"ToUpper":    strings.ToUpper,
}

var renderTemplate = render.New(render.Options{
	Funcs: []template.FuncMap{templateFunctions},
}).HTML

func (b *Board) Render(writer io.Writer, mood []string) {
	renderTemplate(writer, http.StatusOK, "board.html", Gif{URL: b.getGifURL(mood), Mood: mood})
}

func (b *Board) getGifURL(mood []string) string {
	g := giphy.DefaultClient
	gif, err := g.Search(mood)
	if err != nil {
		return InternalServerErrorURL
	}

	if len(gif.Data) < 1 {
		return NoGifFoundURL
	}
	return gif.Data[rand.Intn(len(gif.Data))].MediaURL()
}
