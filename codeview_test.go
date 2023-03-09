package codeview

import (
	"os"
	"testing"
)

func TestCodeView(t *testing.T) {
	text := `package codeview

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
		title: "CatBin",
	}
}`

	lg, _ := os.ReadFile("assets/logo.png")

	SetProjectName("MyProject")
	SetLogo(lg)

	cv := NewCodeView()
	cv.SetText(text)
	cv.SetTitle("Some Example Code")

	cv.SetScale(1.0)

	err := cv.RenderToPng("test.png")
	if err != nil {
		t.Fatal(err)
	}
}
