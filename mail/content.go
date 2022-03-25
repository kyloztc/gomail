package mail

type EmailContent interface {
	GetContent() string
	GetContentType() MailType
}

type TextContent struct {
	Content string
}

func NewTextContent(text string) *TextContent {
	return &TextContent{
		Content: text,
	}
}

func (c *TextContent) GetContent() string {
	return c.Content
}

func (c *TextContent) GetContentType() MailType {
	return TextMailType
}

type HtmlContext struct {
	Content string
}

func NewHtmlContent(text string) *HtmlContext {
	return &HtmlContext{
		Content: text,
	}
}

func (c *HtmlContext) GetContent() string {
	return c.Content
}

func (c *HtmlContext) GetContentType() MailType {
	return HtmlMailType
}

