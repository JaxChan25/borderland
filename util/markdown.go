package util

import (
	"bytes"
	"strings"

	"gopkg.in/russross/blackfriday.v2"
)

var buf bytes.Buffer

//NodeVisitor node.walk()的回调函数
func NodeVisitor(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {

	if entering {
		if node.Type == blackfriday.Image {
			return blackfriday.SkipChildren
		}

		//fmt.Println("我正在进入:  ", string(node.Literal))
		buf.Write(node.Literal)
	}

	return blackfriday.GoToNext
}

//Content2Intr 将article的content(makrdown格式)提取出纯文字作introduction
func Content2Intr(content string) string {

	render := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.CommonHTMLFlags | blackfriday.TOC,
	})

	parser := blackfriday.New(blackfriday.WithRenderer(render))
	node := parser.Parse([]byte(content))

	node.Walk(NodeVisitor)
	intr := string(buf.Bytes())
	intr = strings.Replace(strings.Replace(intr, "\r", "", -1), "\n", "", -1)

	return string([]rune(intr)[:197]) + "..."
}
