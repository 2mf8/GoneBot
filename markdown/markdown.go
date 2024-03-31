package markdown

import (
	"fmt"
	"net/url"
	"strings"
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
	str := fmt.Sprintf("~~%s~~", content)
	r.Str += str
	return r
}

func (r *MarkDown) Bold(content string) *MarkDown {
	str := fmt.Sprintf("**%s** ", content)
	r.Str += str
	return r
}

func (r *MarkDown) Italic(content string) *MarkDown {
	str := fmt.Sprintf("*%s* ", content)
	r.Str += str
	return r
}

func (r *MarkDown) ItalicBold(content string) *MarkDown {
	str := fmt.Sprintf("***%s*** ", content)
	r.Str += str
	return r
}

func (r *MarkDown) BlockReference(content string) *MarkDown {
	str := fmt.Sprintf("\\n> %s\\n", content)
	r.Str += str
	return r
}

func (r *MarkDown) Image(text, url string, width, height int) *MarkDown {
	str := fmt.Sprintf("\\n![%s #%vpx #%vpx](%s)", text, width, height, url)
	r.Str += str
	return r
}

func (r *MarkDown) DividerLine() *MarkDown {
	str := fmt.Sprintln("\\n ---\\n")
	r.Str += str
	r.Str = strings.TrimSpace(r.Str)
	return r
}

func (r *MarkDown) Text(content string) *MarkDown {
	r.Str += content
	return r
}

func (r *MarkDown) NewLine() *MarkDown {
	r.Str += "\\n"
	return r
}

func (r *MarkDown) Json(content string) *MarkDown {
	c := strings.ReplaceAll(strings.ReplaceAll(content, "\t", "\\t"), "\n", "\\n")
	str := fmt.Sprintf("\\n```json\\n%s\\n```\\n", c)
	r.Str += str
	return r
}