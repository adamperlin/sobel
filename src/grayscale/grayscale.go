package grayscale

import (
  "image/color"
  "image"
)

func ToGrayscale(img image.Image) image.Image {
  max := img.Bounds().Max
  min := img.Bounds().Min

  var filtered = image.NewGray(image.Rect(max.X, max.Y, min.X, min.Y))
  for x := 0; x < max.X; x++ {
    for y := 0; y < max.Y; y++ {
      grayColor := color.GrayModel.Convert(img.At(x, y))
      filtered.Set(x, y, grayColor)
    }
  }
  return filtered
}
