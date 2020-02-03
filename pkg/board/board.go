package board

import (
	"github.com/peterhellberg/giphy"
	"github.com/unrolled/render"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	NoGifFoundURL          = "https://media.giphy.com/media/9J7tdYltWyXIY/giphy.mp4"
	InternalServerErrorURL = "https://media.giphy.com/media/OJBrW6nM5hoNW/giphy.mp4"
	HelloWorldURL          = "https://media.giphy.com/media/PnUatAYWMEMvmiwsyx/giphy.mp4"
	boardAssetName         = "templates/board.tmpl"
	formAssetName          = "templates/form.tmpl"
)

var DefaultBoard = &Board{
	URL:       HelloWorldURL,
	Mood:      []string{"Hello", "World"},
	CreatedAt: time.Now().Unix(),
}

// language=HTML
var boardHTMLTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Moodboard</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <style>
        body, html {
            height: 100%;
            margin: 0;

            overflow: hidden;
        }

        .video {
            display: block;

            /* Full height */
            height: 100vh;
            width: 100vw;

            /* Center and scale the image nicely */
            background-position: center;
            background-repeat: no-repeat;
            background-size: contain;
            background-color: black;
        }

        .content {
            color: lightgray;
            background-color: black;

            position: absolute;
            bottom: 0;
            left: 0;
            margin-left: 3vw;
            margin-bottom: 3vh;
            padding: 10px;

            border: ridge lightgray;

            font-size: x-large;
            font-weight: bold;
        }

        @media screen and (max-width: 600px) {
            .content {
                font-size: small;
            }
        }
    </style>
</head>
<body>

<video class="video" autoplay muted loop>
    <source src="{{.URL}}" type="video/mp4">
</video>

<div class="bg">
    <p class="content">MOOD: {{ StringJoin .Mood " " | ToUpper }}</p>
</div>

<script type="text/javascript">
    let initial = null;
    let reload = null;

    setInterval(async () => {
        const response = await fetch('//' + location.host + '/mood');
        const {mood, timestamp} = await response.json();
        if (initial != null && (JSON.stringify(mood) !== JSON.stringify(initial) || JSON.stringify(reload) !== JSON.stringify(timestamp))) {
            location.reload();
        } else {
            initial = mood;
            reload = timestamp;
        }
    }, 1000);
</script>

</body>
</html>
`

// language=HTML
var formHTMLTemplate = `
<!DOCTYPE html>
<html> 
<h3>What is your mood today?</h3> 
<body> 
	<form action="/mood/form" method="post" id="mood"> 
		<label for="mood">Mood:</label> 
		<input type="text" name="mood" id="MoodForm" autofocus> 
		<input type="submit" value="Submit"> 
		<input type="reset" value="Reset"> 
	</form> 
	<p class="content">Current Mood: {{ StringJoin . " " | ToUpper }}</p>
</body> 
</html> 

`

type Board struct {
	URL       string
	Mood      []string
	CreatedAt int64
	mutex     sync.Mutex
}

var templateFunctions = map[string]interface{}{
	"StringJoin": strings.Join,
	"ToUpper":    strings.ToUpper,
}

var renderTemplate = render.New(render.Options{
	Funcs: []template.FuncMap{templateFunctions},
	Asset: func(name string) ([]byte, error) {
		if name == formAssetName {
			return []byte(formHTMLTemplate), nil
		}
		return []byte(boardHTMLTemplate), nil
	},
	AssetNames: func() []string {
		return []string{boardAssetName, formAssetName}
	},
}).HTML

func (b *Board) RenderIndex(writer io.Writer) {
	renderTemplate(writer, http.StatusOK, "board", b)
}

func (b *Board) RenderForm(writer io.Writer) {
	renderTemplate(writer, http.StatusOK, "form", b.Mood)
}

func (b *Board) UpdateMood(mood []string) {
	b.mutex.Lock()
	b.Mood = mood
	b.CreatedAt = time.Now().Unix()
	b.URL = findGifForMood(mood)
	b.mutex.Unlock()
}

func (b *Board) GetMoodWithTime() (mood []string, timestamp int64) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.Mood, b.CreatedAt
}

func findGifForMood(mood []string) string {
	g := giphy.DefaultClient
	gif, err := g.Search(mood)
	if err != nil {
		return InternalServerErrorURL
	}

	if len(gif.Data) < 1 {
		return NoGifFoundURL
	}
	gifId := gif.Data[rand.Intn(len(gif.Data))].ID

	return toMP4URL(gifId)
}

func toMP4URL(id string) string {
	return "https://media.giphy.com/media/" + id + "/giphy.mp4"
}

func toGifURL(id string) string {
	return "https://media.giphy.com/media/" + id + "/giphy.gif"
}

func toThumbnailURL(id string) string {
	return "https://media.giphy.com/media/" + id + "/200w_d.gif"
}
