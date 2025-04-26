package code

/*
	原生图像库打水印，目前有注意到字体大小不是跟随图片大小的，大图片字体会很小
*/

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func RunWaterMark() {
	imgFile, err := os.Open("image/testimg.jpg")
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()
	img, err := jpeg.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	bgFile, err := os.Open("testbg.png")
	if err != nil {
		panic(err)
	}
	defer bgFile.Close()
	bg, err := png.Decode(bgFile)
	if err != nil {
		panic(err)
	}

	offset := image.Pt(bg.Bounds().Dx()-img.Bounds().Dx()-10, bg.Bounds().Dy()-img.Bounds().Dy()-10)
	b := bg.Bounds()
	m := image.NewRGBA(b)

	draw.Draw(m, b, bg, image.Point{}, draw.Src)
	draw.Draw(m, img.Bounds().Add(offset), img, image.Point{}, draw.Over)

	imgnew, err := os.Create("image/testimgnew.jpg")
	if err != nil {
		panic(err)
	}
	defer imgnew.Close()

	jpeg.Encode(imgnew, m, &jpeg.Options{Quality: 100})
}
