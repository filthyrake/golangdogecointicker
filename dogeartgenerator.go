package main

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"image/color"
	"math/rand"
	"time"
	"os"
	"log"
//	"fmt"
)

// Create the canvas
var c = generativeart.NewCanva(500, 500)

// Let's add another art type!
func julia1(z complex128) complex128 {
 c := complex(-0.1, 0.651)

 z = z*z + c

 return z
}

// color map function for Domain Wrap art
func cmap(r, m1, m2 float64) color.RGBA {

        var randR = uint8(rand.Intn(255))
        var randG = uint8(rand.Intn(255))
        var randB = uint8(rand.Intn(255))
        var randA = uint8(rand.Intn(255))

	rgb := color.RGBA{
		randR,
		randG,
		randB,
		randA,
	}
	return rgb
}

// trying to make a single function to handle all random color generation
func randColor() color.RGBA {
        var randR = uint8(rand.Intn(255))
        var randG = uint8(rand.Intn(255))
        var randB = uint8(rand.Intn(255))
        var randA = uint8(rand.Intn(255))

        rgb := color.RGBA{
                randR,
                randG,
                randB,
                randA,
        }
        return rgb
}

// random color schema function
func randColorSchema() []color.RGBA{
        rgb := []color.RGBA{
                randColor(),
                randColor(),
                randColor(),
                randColor(),
                randColor(),
                randColor(),
                randColor(),
                randColor(),
                randColor(),
                randColor(),
        }
        return rgb
}

// Randomly pick what type of art we are making today and draw it
func randomArt() {
	min := 0
	max := 14
	var randNumArtType = rand.Intn(max - min + 1) + min

	if randNumArtType == 0 {
		//fmt.Println("NewRandomShape")
		max = 200
		var randNumMisc = rand.Intn(max - min + 1) + min
		c.Draw(arts.NewRandomShape(randNumMisc))
	} else if randNumArtType == 1 {
		//fmt.Println("NewColorCanve")
		max = 10
		var randNumMisc = rand.Intn(max - min + 1) + min
		c.SetLineWidth(8)
		c.Draw(arts.NewColorCanve(float64(randNumMisc)))
	} else if randNumArtType == 2 {
		//fmt.Println("NewColorCircle2")
		max = 40
		var randNumMisc = rand.Intn(max - min + 1) + min
		c.Draw(arts.NewColorCircle2(randNumMisc))
	} else if randNumArtType == 3 {
		//fmt.Println("NewCircleLoop2")
		max = 10
		var randNumMisc = rand.Intn(max - min + 1) + min
		c.Draw(arts.NewCircleLoop2(randNumMisc))
	} else if randNumArtType == 4 {
		//fmt.Println("NewPixelHole")
		max = 100
		var randNumMisc = rand.Intn(max - min + 1) + min
		c.SetIterations(1200)
		c.Draw(arts.NewPixelHole(randNumMisc))
	} else if randNumArtType == 5 {
		//fmt.Println("NewDotsWave")
		max = 500
		var randNumMisc = rand.Intn(max - min + 1) + min
		c.Draw(arts.NewDotsWave(randNumMisc))
	} else if randNumArtType == 6 {
		//fmt.Println("NewContourLine")
		max = 1000
		var randNumMisc = rand.Intn(max - min + 1) + min
		c.Draw(arts.NewContourLine(randNumMisc))
	} else if randNumArtType == 7 {
		//fmt.Println("NewNoiseLine")
		max = 2000
		var randNumMisc = rand.Intn(max - min + 1) + min
		c.Draw(arts.NewNoiseLine(randNumMisc))
	} else if randNumArtType == 8 {
		//fmt.Println("NewDotLine")
		max = 200
		var randNumMisc = rand.Intn(max - min + 1) + min
		c.SetLineWidth(10)
		c.SetIterations(15000)
		c.Draw(arts.NewDotLine(randNumMisc, 20, 50, false))
	} else if randNumArtType == 9 {
		//fmt.Println("NewCircleLoop")
		max = 200
		var randNumMisc = rand.Intn(max - min + 1) + min
		c.SetAlpha(30)
		c.SetIterations(1000)
	        c.SetLineColor(randColor())
		c.Draw(arts.NewCircleLoop(float64(randNumMisc)))
	} else if randNumArtType == 10 {
		//fmt.Println("Junas")
		c.SetForeground(randColor())
		max = 10
		var randNumMisc = rand.Intn(max - min + 1) + min
		c.Draw(arts.NewJanus(randNumMisc, 0.2))
	} else if randNumArtType == 11 {
		//fmt.Println("DomainWrap")
		c.Draw(arts.NewDomainWrap(0.01, 4,4, 20, cmap))
	} else if randNumArtType == 12 {
		//fmt.Println("CircleNoise")
		max = 4000
		var randNumMisc = rand.Intn(max - min + 1) + min
		c.SetAlpha(80)
		c.SetLineWidth(0.3)
		c.SetIterations(400)
		c.Draw(arts.NewCircleNoise(randNumMisc, 60, 80))
	} else if randNumArtType == 13 {
		//fmt.Println("Julia Set")
		 c.SetIterations(800)
		 c.Draw(arts.NewJulia(julia1, 40, 1.5, 1.5))
	} else if randNumArtType == 14 {
		//fmt.Println("New Silk Sky")
		c.SetAlpha(10)
		c.Draw(arts.NewSilkSky(15, 5))
	}

}

func main() {
	rand.Seed(time.Now().Unix())
	min := 0
 
	c.SetBackground(randColor())

	c.FillBackground()

	c.SetColorSchema(randColorSchema())

	// Now let's pick what kind of art we want to make and make some art!
	randomArt()


	// Do we want to add a second layer of art on top of the first layer?
	anotherMin := 0
	anotherMax := 1
	var randNumAdditional = rand.Intn(anotherMax - anotherMin + 1) + min
	if randNumAdditional == 0 {
		randomArt()
	}


	c.ToPNG("<replace>")

	// Let's get the randomly generated art file back open and ready for dogeifying
	image1,err := os.Open("<replace>")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	first, err := png.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()
 
	// Now for the dogecoin logo
	image2,err := os.Open("<replace>")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	second,err := png.Decode(image2)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image2.Close()

	// let's randomize where the logo goes!
	waterMarkMaxX := 300
	waterMarkMaxY := 300
	var randNumX = rand.Intn(waterMarkMaxX - min + 1) + min
	var randNumY = rand.Intn(waterMarkMaxY - min + 1) + min

	// Write the watermark
	offset := image.Pt(randNumX, randNumY)
	b := first.Bounds()
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, first, image.ZP, draw.Src)
	draw.Draw(image3, second.Bounds().Add(offset), second, image.ZP, draw.Over)

	third,err := os.Create("<replace>")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
	defer third.Close()
}
