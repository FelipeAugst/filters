package picture

import (
	"fmt"
	"image"
	"image/png"
	_ "image/png"
	"os"

	"github.com/disintegration/gift"
)

type Picture struct {
	filtersList *gift.GIFT
	original    image.Image
	filtered    *image.NRGBA
	format      string
}

func NewPicture(source string, filters ...gift.Filter) Picture {
	var p Picture
	p.filtersList = gift.New(filters...)
	img, format, err := loadImage(source)
	if err != nil {
		fmt.Println("Failed to load source image \n ", err)
		return Picture{}
	}
	p.original = img
	p.format = format
	bounds := p.original.Bounds()
	p.filtered = image.NewNRGBA(bounds)
	return p

}

func loadImage(path string) (image.Image, string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}
	return img, format, nil
}

func (p *Picture) Apply() {

	p.filtersList.Draw(p.filtered, p.original)

}

func (p Picture) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	if err = png.Encode(file, p.filtered); err != nil {
		return err
	}
	return nil
}
