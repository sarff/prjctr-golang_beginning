/*
Треба написати програму для умовної обробки зображення різного типу, з використанням поліморфізму,
контрактів і panic/recover
*/
package main

import "fmt"

type ImageProcessor interface {
	Process() error
}

type BMPImageProcessor struct {
	path string
}

type PNGImageProcessor struct {
	path string
}

func (p *PNGImageProcessor) Process() error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("We processed a panic, result = %v. Image: %s\n", r, p.path)
		} else {
			fmt.Println("Everything finished successfully")
		}
	}()

	fmt.Println("Trying to process png image")

	// fake panic
	var s *string
	fmt.Println(*s)

	return nil
}

func (p *BMPImageProcessor) Process() error {
	fmt.Printf("We processed bmp image: %s\n", p.path)
	return nil
}

func main() {
	bmp := BMPImageProcessor{
		path: "./image/image.bmp",
	}
	png := PNGImageProcessor{
		path: "./image/image.png",
	}

	processors := []ImageProcessor{&bmp, &png}
	for _, p := range processors {
		err := p.Process()
		if err != nil {
			fmt.Printf("Error processing image: %s\n", err)
			return
		}
	}
}
