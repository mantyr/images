package images

import (
    "image"
    "errors"
    "image/color"
)

func (i *Image) negative() (img *image.NRGBA, err error) {
    width  := i.Width()
    height := i.Height()
    if width <= 0 || height <= 0 {
        err = errors.New("no image size")
    }

    img = image.NewNRGBA(image.Rect(0, 0, width, height))
    parallel(height, func(partStart, partEnd int) {
        for dstY := partStart; dstY < partEnd; dstY++ {
            for dstX := 0; dstX < width; dstX++ {
                color := i.Image.At(dstX, dstY).(color.NRGBA)
                color.R = 255 - color.R
                color.G = 255 - color.G
                color.B = 255 - color.B

                img.Set(dstX, dstY, color)
            }
        }
    })
    return
}