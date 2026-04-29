package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

const (
	width  = 1200
	height = 630
)

var (
	paper = color.RGBA{247, 242, 232, 255}
	ink   = color.RGBA{22, 20, 15, 255}
	white = color.RGBA{255, 253, 248, 255}
	blue  = color.RGBA{35, 86, 216, 255}
	green = color.RGBA{22, 128, 75, 255}
	pink  = color.RGBA{228, 86, 122, 255}
)

func main() {
	mustWrite("static/images/default-og.jpg", drawDefault())
	mustWrite("static/images/posts/seo-static-blog-cover.jpg", drawCover())
	mustWrite("static/images/posts/cloudflare-pages-hugo-cover.jpg", drawCloudflareCover())
}

func mustWrite(path string, img image.Image) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err := jpeg.Encode(file, img, &jpeg.Options{Quality: 92}); err != nil {
		panic(err)
	}
}

func canvas() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	fill(img, img.Bounds(), paper)
	stroke(img, image.Rect(32, 32, 1168, 598), ink, 20)
	return img
}

func drawDefault() image.Image {
	img := canvas()
	fill(img, image.Rect(112, 160, 690, 220), ink)
	fill(img, image.Rect(112, 252, 780, 294), ink)
	fill(img, image.Rect(116, 338, 620, 372), blue)
	fill(img, image.Rect(116, 424, 720, 450), ink)
	fill(img, image.Rect(116, 466, 650, 492), ink)
	drawBadge(img, 830, 110)
	return img
}

func drawCover() image.Image {
	img := canvas()
	fill(img, image.Rect(112, 152, 690, 214), ink)
	fill(img, image.Rect(112, 250, 828, 312), ink)
	fill(img, image.Rect(116, 400, 850, 435), blue)
	fill(img, image.Rect(116, 455, 580, 486), pink)
	fill(img, image.Rect(790, 132, 1035, 322), white)
	stroke(img, image.Rect(790, 132, 1035, 322), ink, 8)
	fill(img, image.Rect(830, 174, 866, 210), green)
	fill(img, image.Rect(884, 174, 920, 210), green)
	fill(img, image.Rect(938, 174, 974, 210), white)
	fill(img, image.Rect(830, 244, 985, 262), blue)
	fill(img, image.Rect(830, 280, 944, 298), pink)
	return img
}

func drawCloudflareCover() image.Image {
	img := canvas()
	fill(img, image.Rect(112, 152, 740, 214), ink)
	fill(img, image.Rect(112, 250, 930, 312), ink)
	fill(img, image.Rect(116, 400, 860, 435), blue)
	fill(img, image.Rect(116, 455, 640, 486), green)
	fill(img, image.Rect(805, 122, 1025, 342), white)
	stroke(img, image.Rect(805, 122, 1025, 342), ink, 8)
	fill(img, image.Rect(845, 165, 985, 205), pink)
	fill(img, image.Rect(845, 232, 985, 252), ink)
	fill(img, image.Rect(845, 276, 945, 296), blue)
	return img
}

func drawBadge(img *image.RGBA, x, y int) {
	fill(img, image.Rect(x, y, x+170, y+140), white)
	stroke(img, image.Rect(x, y, x+170, y+140), ink, 8)
	for _, r := range []image.Rectangle{
		image.Rect(x+28, y+28, x+54, y+54),
		image.Rect(x+64, y+28, x+90, y+54),
		image.Rect(x+136, y+28, x+162, y+54),
		image.Rect(x+28, y+64, x+54, y+90),
		image.Rect(x+100, y+64, x+126, y+90),
		image.Rect(x+136, y+64, x+162, y+90),
	} {
		fill(img, r, green)
		stroke(img, r, ink, 4)
	}
}

func fill(img *image.RGBA, rect image.Rectangle, c color.Color) {
	bounds := img.Bounds()
	rect = rect.Intersect(bounds)
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			img.SetRGBA(x, y, rgba)
		}
	}
}

func stroke(img *image.RGBA, rect image.Rectangle, c color.Color, size int) {
	fill(img, image.Rect(rect.Min.X, rect.Min.Y, rect.Max.X, rect.Min.Y+size), c)
	fill(img, image.Rect(rect.Min.X, rect.Max.Y-size, rect.Max.X, rect.Max.Y), c)
	fill(img, image.Rect(rect.Min.X, rect.Min.Y, rect.Min.X+size, rect.Max.Y), c)
	fill(img, image.Rect(rect.Max.X-size, rect.Min.Y, rect.Max.X, rect.Max.Y), c)
}
