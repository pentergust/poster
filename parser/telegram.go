package parser

import (
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// Специальный хук ждя рендеринга Telegram сообщения
func renderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if _, ok := node.(*ast.Heading); ok {
		if entering {
			io.WriteString(w, "<strong>")

		} else {
			io.WriteString(w, "</strong>\n\n")
		}

		return  ast.GoToNext, true
	}

	if _, ok := node.(*ast.Paragraph); ok {
		if !entering {
			io.WriteString(w, "\n")
		}
		return  ast.GoToNext, true
	}
	
	return ast.GoToNext, false
}

// Преобразует сообщение в подходящий для Telegram формат
func ToTelegram(text string) string {
	p := parser.New()
	doc := p.Parse([]byte(text))
	render := html.NewRenderer(html.RendererOptions{RenderNodeHook: renderHook})
	html := markdown.Render(doc, render)
	return string(html)
}
