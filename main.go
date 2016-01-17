package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

func AddPhotoFrame() {
	//TODO:user the res/gophergala.jpg to generate a image and write to res/m.jpg which is similar to like the logo in the README.md
	p1, err := os.Open("res/gophergala.jpg")
	if err != nil {
		return
	}
	defer p1.Close()
	m1, _, err := image.Decode(p1)

	p2, err := os.Open("res/m.jpg")

	if err != nil {
		return
	}
	defer p2.Close()

	m2, _, err := image.Decode(p2)

	if err != nil {
		return
	}

	// fill the core pic
	r := image.Rect(117, 132, 663, 645)
	m := image.NewRGBA(image.Rect(0, 0, 779, 779))

	draw.Draw(m, r, m1, m1.Bounds().Min.Add(image.Pt(117, 131)), draw.Src)

	// fill the upper

	rupper := image.Rect(0, 0, 779, 132)
	draw.Draw(m, rupper, m2, m2.Bounds().Min, draw.Src)

	rleft := image.Rect(0, 132, 117, 779)
	draw.Draw(m, rleft, m2, m2.Bounds().Min.Add(image.Pt(0, 132)), draw.Src)

	rright := image.Rect(663, 132, 779, 779)
	draw.Draw(m, rright, m2, m2.Bounds().Min.Add(image.Pt(663, 132)), draw.Src)

	rdown := image.Rect(116, 645, 663, 779)
	draw.Draw(m, rdown, m2, m2.Bounds().Min.Add(image.Pt(116, 645)), draw.Src)

	toImg, _ := os.Create("new.jpg")
	defer toImg.Close()

	jpeg.Encode(toImg, m, &jpeg.Options{jpeg.DefaultQuality})
}

func main() {
	AddPhotoFrame()
	println(`This final exercise,let's add a photo frame for gala logo!
You should use image package to generate a new iamge from the gala logo(which is stored in res/gophergala.jpg,and makes it like the res/m.jpg.
Now edit main.go to complete 'AddPhotoFrame' function,this task has no test,enjoy your trip!`)
}
