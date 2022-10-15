package ImageUtils

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
	"strings"
)

func ReadPicToPixSlice(PicturePath string) (PixelSlice [][]int64) {
	f, err := os.Open(PicturePath)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err2 := f.Close()
		if err2 != nil {
			fmt.Println("Close picture err!")
		}
	}(f)

	img, fmtName, err1 := image.Decode(f)
	if err1 != nil {
		panic(err)
	}
	Dx := img.Bounds().Dx()
	Dy := img.Bounds().Dy()

	fmt.Printf("Name: %v, Bounds: %v, Color: %v\n", fmtName, []int{Dx, Dy}, img.ColorModel())

	PixelSlice = make([][]int64, 0)
	for i := 0; i < Dx; i++ {
		for j := 0; j < Dy; j++ {
			str := fmt.Sprintf("%v\n", img.At(i, j))
			_, af, _ := strings.Cut(str, "{")
			be, _, _ := strings.Cut(af, "}")
			res := strings.Split(be, " ")
			resInt := make([]int64, 0)
			for _, v := range res {
				tmp, err3 := strconv.ParseInt(v, 10, 64)
				if err3 != nil {
					fmt.Println("Strconv err is", err3)
				}
				resInt = append(resInt, tmp)
			}
			PixelSlice = append(PixelSlice, resInt)
		}
	}
	return PixelSlice
}
