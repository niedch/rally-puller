package markdown

import (
	"testing"

	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"github.com/stretchr/testify/assert"
)

func TestConvertToMarkdown(t *testing.T) {
	tests := []struct {
		name     string
		htmlInput string
		expectedMarkdown string
		expectError bool
	}{
		{
			name:     "Simple HTML to Markdown",
			htmlInput: "<h1>Hello World</h1>",
			expectedMarkdown: "# Hello World",
			expectError: false,
		},
		{
			name:     "HTML with paragraph",
			htmlInput: "<p>This is a paragraph.</p>",
			expectedMarkdown: "This is a paragraph.",
			expectError: false,
		},
		{
			name:     "HTML with link",
			htmlInput: `<a href="https://example.com">Example</a>`,
			expectedMarkdown: "[Example](https://example.com)",
			expectError: false,
		},
		{
			name:     "HTML with custom tag 'info'",
			htmlInput: `<info>Some info text</info>`,
			expectedMarkdown: "", // The rendererForTable currently renders nothing for 'info'
			expectError: false,
		},
		{
			name:     "Empty HTML",
			htmlInput: "",
			expectedMarkdown: "",
			expectError: false,
		},
		{
			name:     "Complex HTML with multiple tags",
			htmlInput: `<div><h1>Title</h1><p>Paragraph with <strong>bold</strong> text.</p><ul><li>Item 1</li><li>Item 2</li></ul></div>`,
			expectedMarkdown: "# Title\n\nParagraph with **bold** text.\n\n- Item 1\n- Item 2",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			markdown, err := ConvertToMarkdown(tt.htmlInput)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedMarkdown, markdown)
			}
		})
	}
}

func TestRendererForTable(t *testing.T) {
	t.Run("RendererForTable returns RenderSuccess", func(t *testing.T) {
		status := rendererForTable(nil, nil, nil)
		assert.Equal(t, converter.RenderSuccess, status)
	})
}
