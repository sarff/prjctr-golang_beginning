/*
Треба написати програму для умовної  обробки зображення різного типу, з використанням поліморфізму,
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

func (i *PNGImageProcessor) Process() error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("We processed a panic, result = %v. Image: %s\n", r, i.path)
		} else {
			fmt.Println("Everything finished successfully")
		}
	}()

	fmt.Println("Trying to process png image")

	var s *string
	fmt.Println(*s)
	return nil
}

func (i *BMPImageProcessor) Process() error {
	fmt.Printf("We processed bmp image: %s\n", i.path)
	return nil
}

func main() {
	var bmp = BMPImageProcessor{
		path: "./image/image.bmp",
	}
	var png = PNGImageProcessor{
		path: "./image/image.png",
	}

	err := error(nil)
	processors := []ImageProcessor{&bmp, &png}
	for _, p := range processors {
		err = p.Process()
		if err != nil {
			return
		}
	}
}
