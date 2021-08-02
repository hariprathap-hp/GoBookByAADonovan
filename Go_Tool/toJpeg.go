package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	_ "image/png" // register PNG decoder
	"io"
	"os"
)

var im = flag.String("image", "jpeg", "enter your required output image type")
var myimg string

func main() {
	flag.Parse()
	myimg = *im
	str := "i.jpeg"
	f, err := os.Open(str)
	if err != nil {
		fmt.Println(err)
	}
	if err := toJPEG(f, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}
func toJPEG(in io.Reader, out io.Writer) error {
	//fmt.Println(img)
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
L:
	switch myimg {
	case "jpeg":
		fmt.Println("Case Jpeg")
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	default:
		break L
	}
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
