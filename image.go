package codeview

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
)

type CodeView struct {
	text  string
	title string
	scale float64
}

func NewCodeView() *CodeView {
	return &CodeView{
		scale: 1,
		text:  "",
		title: projectName,
	}
}

func (cv *CodeView) SetText(text string) {
	cv.text = text
}

func (cv *CodeView) SetTitle(title string) {
	max := 31 - len(projectName)

	if len(title) > max {
		title = title[:max] + "..."
	}

	cv.title = title + " - " + projectName
}

func (cv *CodeView) SetScale(scale float64) {
	cv.scale = scale
}

func (cv *CodeView) Render() (image.Image, error) {
	return _render(cv.text, cv.title, cv.scale)
}

func (cv *CodeView) RenderToPng(path string) error {
	img, err := _render(cv.text, cv.title, cv.scale)
	if err != nil {
		return err
	}

	return draw2dimg.SaveToPngFile(path, img)
}
