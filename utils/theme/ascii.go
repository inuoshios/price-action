package theme

import (
	"log"

	f "github.com/mbndr/figlet4go"
)

func AsciiRender(str, fontName string, fontColour []f.Color, parser f.Parser) string {
	ascii := f.NewAsciiRender()

	opts := f.RenderOptions{
		FontName:  fontName,
		FontColor: fontColour,
		Parser:    parser,
	}

	renderText, err := ascii.RenderOpts(str, &opts)
	if err != nil {
		log.Println(err)
	}

	return renderText
}
