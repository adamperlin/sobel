package sobel

import (
//  "image/jpeg"
  "image"
  "image/color"
  "math"
//  "fmt"
  "grayscale"
)

var (
   FilterX = [3][3]int8{
    {-1, 0, 1},
    {-2, 0, 2},
    {-1, 0, 1},
  }

  FilterY = [3][3]int8{
    {-1, -2, -1},
    {0, 0, 0},
    {1, 2, 1},
  }
)

func Filter(img image.Image) image.Image {
  img = grayscale.ToGrayscale(img)
  max := img.Bounds().Max
  min := img.Bounds().Min
  /* filtered image must be two pixels shorter, because
  there must be a row of pixels on each side of a pixel for the sobel operator
  to work*/
  var filtered = image.NewGray(image.Rect(max.X - 2, max.Y - 2, min.X, min.Y))
  width := max.X
  height := max.Y
  var pixel color.Color
  for x := 1; x < width - 1; x++ {
    for y := 1; y < height - 1; y++ {
      fX, fY := applyFilter(img, x, y)
      v := uint32(math.Ceil(math.Sqrt(float64((fX * fX) + (fY * fY)))))
    //  fmt.Println("v: ", v)
      pixel = color.Gray{Y: uint8(v)}
      filtered.SetGray(x, y, pixel.(color.Gray))
    }
  }
  return filtered
}

func applyFilter(img image.Image, x int, y int) (uint32, uint32) {
    var fX, fY int
    curX := x - 1
    curY := y - 1
    for i := 0; i < 3; i++ {
      for j := 0; j < 3; j++ {
        pixel := getGrayPixel(img.At(curX, curY))
        fX += int(FilterX[i][j]) * int(pixel)
        fY += int(FilterY[i][j]) * int(pixel)
        curX++
      }
      curX = x - 1
      curY++
    }
    uFX, uFY := uint32(math.Abs(float64(fX))), uint32(math.Abs(float64(fY)))
  //  fmt.Printf("uFX: %d, uFY: %d", uFX, uFY)
    return uFX, uFY
}

func getGrayPixel(c color.Color) uint8 {
  p, _, _, _ := c.RGBA()
//  fmt.Printf("pixel is: %d", pixel)

  ret :=  uint8(p)
  //fmt.Println("ret: ", ret)
  return ret
}
