package markdown

import (
	"fmt"
	"net/url"
)

type MarkDown struct {
	Str string
}

func NewMarkDown() *MarkDown {
	return &MarkDown{
		Str: "",
	}
}

func (r *MarkDown) MqqApi(content string) *MarkDown {
	str := fmt.Sprintf("\\n[%s](mqqapi://aio/inlinecmd?command=%s&reply=false&enter=false)", content, url.PathEscape(content))
	r.Str += str
	return r
}

func (r *MarkDown) MqqApiAuto(content string) *MarkDown {
	str := fmt.Sprintf("\\n[%s](mqqapi://aio/inlinecmd?command=%s&reply=false&enter=true)", content, url.PathEscape(content))
	r.Str += str
	return r
}

func (r *MarkDown) Url(name, webUrl string) *MarkDown {
	str := fmt.Sprintf("\\n[🔗%s](%s)", name, webUrl)
	r.Str += str
	return r
}

func (r *MarkDown) H1(content string) *MarkDown {
	str := fmt.Sprintf("\\n# %s", content)
	r.Str += str
	return r
}

func (r *MarkDown) H2(content string) *MarkDown {
	str := fmt.Sprintf("\\n## %s", content)
	r.Str += str
	return r
}

func (r *MarkDown) H3(content string) *MarkDown {
	str := fmt.Sprintf("\\n### %s", content)
	r.Str += str
	return r
}

func (r *MarkDown) DeleteLine(content string) *MarkDown {
	str := fmt.Sprintf("\\n~~%s~~", content)
	r.Str += str
	return r
}

func (r *MarkDown) Bold(content string) *MarkDown {
	str := fmt.Sprintf("\\n**%s**", content)
	r.Str += str
	return r
}

func (r *MarkDown) Italic(content string) *MarkDown {
	str := fmt.Sprintf("\\n*%s*", content)
	r.Str += str
	return r
}

func (r *MarkDown) ItalicBold(content string) *MarkDown {
	str := fmt.Sprintf("\\n***%s***", content)
	r.Str += str
	return r
}

func (r *MarkDown) BlockReference(content string) *MarkDown {
	str := fmt.Sprintf("\\n\\n> %s\\n", content)
	r.Str += str
	return r
}
