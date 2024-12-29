package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand/v2"
	"os"
)

var palette = []color.Color{color.White, color.RGBA{0x00, 0x00, 0xBB, 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer) {
	const (
		cycles  = 100
		res     = 0.001
		size    = 10
		nframes = 64
		delay   = 100
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, size, size)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main() {
	// 创建一个 GIF 文件
	file, err := os.Create("output.gif")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 将图形数据写入文件
	lissajous(file)
}
