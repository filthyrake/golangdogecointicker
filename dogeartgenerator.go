package main

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
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


// Randomly pick what type of art we are making today and draw it
func randomArt() {
	min := 0
	max := 12
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
		//fmt.Println(randNumMisc)
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
		c.SetColorSchema(common.DarkRed)
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
	}

}

// Some art needs a special schema.  Here we randomly pick and set that
func specialSchema() {
	min := 0
	max := 3
	var randSpecialSchema = rand.Intn(max - min + 1) + min
	if randSpecialSchema == 0 {
		c.SetColorSchema(common.Outdoors)
	} else if randSpecialSchema == 1 {
		c.SetColorSchema(common.Reddery)
	} else if randSpecialSchema == 2 {
		c.SetColorSchema(common.DarkPink)
	} else if randSpecialSchema == 3 {
		c.SetColorSchema(common.DarkRed)
	}
}

// Now let's randomize and set our Color Schema
func setColorSchema() {
	min := 0
	max := 9
	var randNumSchema = rand.Intn(max - min + 1) + min
	if randNumSchema == 0 {
		var colors = []color.RGBA{
			{0xED, 0x34, 0x41, 0xFF},
			{0xFF, 0xD6, 0x30, 0xFF},
			{0x32, 0x9F, 0xE3, 0xFF},
			{0x15, 0x42, 0x96, 0xFF},
			{0x00, 0x00, 0x00, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 1 {
		var colors = []color.RGBA{
			{0xCF, 0x2B, 0x34, 0xFF},
			{0xF0, 0x8F, 0x46, 0xFF},
			{0xF0, 0xC1, 0x29, 0xFF},
			{0x19, 0x6E, 0x94, 0xFF},
			{0x35, 0x3A, 0x57, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 2 {
		var colors = []color.RGBA{
			{0xF9, 0xC8, 0x0E, 0xFF},
			{0xF8, 0x66, 0x24, 0xFF},
			{0xEA, 0x35, 0x46, 0xFF},
			{0x66, 0x2E, 0x9B, 0xFF},
			{0x43, 0xBC, 0xCD, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 3 {
		var colors = []color.RGBA{
			{0x05, 0x1F, 0x34, 0xFF},
			{0x02, 0x74, 0x95, 0xFF},
			{0x01, 0xA9, 0xC1, 0xFF},
			{0xBA, 0xD6, 0xDB, 0xFF},
			{0xF4, 0xF5, 0xF5, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 4 {
		var colors = []color.RGBA{
			{0xFF, 0xC6, 0x18, 0xFF},
			{0xF4, 0x25, 0x39, 0xFF},
			{0xFE, 0x84, 0xFE, 0xFF},
			{0xFF, 0x81, 0x19, 0xFF},
			{0x98, 0x19, 0xFA, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 5 {
		var colors = []color.RGBA{
			{0x11, 0x60, 0xC6, 0xFF},
			{0xFD, 0xD9, 0x00, 0xFF},
			{0xF5, 0xB4, 0xF8, 0xFF},
			{0xEF, 0x13, 0x55, 0xFF},
			{0xF4, 0x9F, 0x0A, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 6 {
		var colors = []color.RGBA{
			{0x58, 0x18, 0x45, 0xFF},
			{0x90, 0x0C, 0x3F, 0xFF},
			{0xC7, 0x00, 0x39, 0xFF},
			{0xFF, 0x57, 0x33, 0xFF},
			{0xFF, 0xC3, 0x0F, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 7 {
		var colors = []color.RGBA{
			{0xFF, 0xBE, 0x0B, 0xFF},
			{0xFB, 0x56, 0x07, 0xFF},
			{0xFF, 0x00, 0x6E, 0xFF},
			{0x83, 0x38, 0xEC, 0xFF},
			{0x3A, 0x86, 0xFF, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 8 {
		var colors = []color.RGBA{
			{0x06, 0x7B, 0xC2, 0xFF},
			{0x84, 0xBC, 0xDA, 0xFF},
			{0xEC, 0xC3, 0x0B, 0xFF},
			{0xF3, 0x77, 0x48, 0xFF},
			{0xD5, 0x60, 0x62, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 9 {
		var colors = []color.RGBA{
			{0x67, 0xD2, 0xE9, 0xFF},
			{0xA7, 0xDB, 0xDA, 0xFF},
			{0xE0, 0xE4, 0xCC, 0xFF},
			{0xF3, 0x86, 0x30, 0xFF},
			{0xFA, 0x69, 0x00, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 10 {
		var colors = []color.RGBA{
			{0x2C, 0x35, 0x31, 0xFF},
			{0x11, 0x64, 0x66, 0xFF},
			{0xD9, 0xB0, 0x8C, 0xFF},
			{0xFF, 0xCB, 0x9A, 0xFF},
			{0xD1, 0xE8, 0xE2, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 11 {
		var colors = []color.RGBA{
			{0x56, 0x80, 0xE9, 0xFF},
			{0x84, 0xCE, 0xEB, 0xFF},
			{0x5A, 0x89, 0xEA, 0xFF},
			{0xC1, 0xC8, 0xEA, 0xFF},
			{0x88, 0x60, 0xD0, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 12 {
		var colors = []color.RGBA{
			{0xA6, 0x4A, 0xC9, 0xFF},
			{0xFC, 0xCD, 0x04, 0xFF},
			{0xFF, 0xB4, 0x8F, 0xFF},
			{0xF5, 0xE6, 0xCC, 0xFF},
			{0x17, 0xE9, 0xE0, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 13 {
		var colors = []color.RGBA{
			{0xA1, 0xC3, 0xD1, 0xFF},
			{0xB3, 0x9B, 0xC8, 0xFF},
			{0xF0, 0xEB, 0xF4, 0xFF},
			{0xF1, 0x72, 0xA1, 0xFF},
			{0xE6, 0x43, 0x98, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 14 {
		var colors = []color.RGBA{
			{0x1F, 0x26, 0x05, 0xFF},
			{0x1F, 0x65, 0x21, 0xFF},
			{0x53, 0x90, 0x0F, 0xFF},
			{0xA4, 0xA7, 0x1E, 0xFF},
			{0xD6, 0xCE, 0x15, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 15 {
		var colors = []color.RGBA{
			{0xFB, 0xE8, 0xA6, 0xFF},
			{0xF4, 0x97, 0x6C, 0xFF},
			{0x30, 0x3C, 0x6C, 0xFF},
			{0xB4, 0xDF, 0xE5, 0xFF},
			{0xD2, 0xFD, 0xFF, 0xFF},
		}
		c.SetColorSchema(colors)
	} else if randNumSchema == 16 {
		var colors = []color.RGBA{
			{0x8D, 0x87, 0x41, 0xFF},
			{0xF4, 0x97, 0x6C, 0xFF},
			{0x30, 0x3C, 0x6C, 0xFF},
			{0xB4, 0xDF, 0xE5, 0xFF},
			{0xD2, 0xFD, 0xFF, 0xFF},
		}
		c.SetColorSchema(colors)
	}



}

func main() {
	rand.Seed(time.Now().Unix())
	min := 0
 
	c.SetBackground(randColor())

	c.FillBackground()

	setColorSchema()

	// Now let's pick what kind of art we want to make and make some art!
	randomArt()


	// Do we want to add a second layer of art on top of the first layer?
	anotherMin := 0
	anotherMax := 1
	var randNumAdditional = rand.Intn(anotherMax - anotherMin + 1) + min
	if randNumAdditional == 0 {
		randomArt()
	}


	c.ToPNG("/root/dogeart/randomart.png")

	// Let's get the randomly generated art file back open and ready for dogeifying
	image1,err := os.Open("/root/dogeart/randomart.png")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	first, err := png.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()
 
	// Now for the dogecoin logo
	image2,err := os.Open("/root/dogeart/dogecoinlogo.png")
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

	third,err := os.Create("/root/dogeart/result.jpg")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
	defer third.Close()
}
