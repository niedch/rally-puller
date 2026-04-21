package markdown

import (
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/base"
	"github.com/JohannesKaufmann/html-to-markdown/v2/plugin/commonmark"
	"golang.org/x/net/html"
)

func ConvertToMarkdown(html string) (string, error) {
	conv := converter.NewConverter(
		converter.WithPlugins(
			base.NewBasePlugin(),
			commonmark.NewCommonmarkPlugin(),
		),
	)

	conv.Register.RendererFor("info", converter.TagTypeInline, rendererForTable, converter.PriorityStandard)

	return conv.ConvertString(html)
}

func rendererForTable(ctx converter.Context, w converter.Writer, node *html.Node) converter.RenderStatus {
	return converter.RenderSuccess
}

