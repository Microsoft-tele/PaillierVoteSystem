package main

import (
	"SockGo/ImageUtils"
	"encoding/json"
	"fmt"
)

func ImageToJson(imagepath string) []byte {
	PixelSlice := ImageUtils.ReadPicToPixSlice(imagepath) // 生成像素数组，返回的是int切片
	for i, v := range PixelSlice {
		fmt.Printf("[%d : %v : %T]\n", i, v, v)
	}

	PixelSliceJson, err := json.Marshal(PixelSlice)
	if err != nil {
		fmt.Println("PixelSlice to json err:", err)
	}
	return PixelSliceJson
}
