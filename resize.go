package image

import (
    "image"
    "math"
)

type iwpair struct {
    i int
    w int32
}

type pweights struct {
    iwpairs []iwpair
    wsum    int32
}

func precomputeWeights(dstSize, srcSize int, filter ResampleFilter) []pweights {
    du := float64(srcSize) / float64(dstSize)
    scale := du
    if scale < 1.0 {
        scale = 1.0
    }
    ru := math.Ceil(scale * filter.Support)

    out := make([]pweights, dstSize)

    for v := 0; v < dstSize; v++ {
        fu := (float64(v)+0.5)*du - 0.5

        startu := int(math.Ceil(fu - ru))
        if startu < 0 {
            startu = 0
        }
        endu := int(math.Floor(fu + ru))
        if endu > srcSize-1 {
            endu = srcSize - 1
        }

        wsum := int32(0)
        for u := startu; u <= endu; u++ {
            w := int32(0xff * filter.Kernel((float64(u)-fu)/scale))
            if w != 0 {
                wsum += w
                out[v].iwpairs = append(out[v].iwpairs, iwpair{u, w})
            }
        }
        out[v].wsum = wsum
    }
    return out
}

func Resize(i image.Image, width, height int, filter ResampleFilter) *image.NRGBA {
    return nil
}



func (i *Image) resize(width, height int) (img *image.NRGBA, err error) {
    img = image.NewNRGBA(image.Rect(0, 0, width, height))

    w_ratio := i.Width64() / float64(width)
    h_ratio := i.Height64() / float64(height)

    parallel(height, func(partStart, partEnd int) {
        for dstY := partStart; dstY < partEnd; dstY++ {
            fy := (float64(dstY)+0.5)*h_ratio - 0.5

            for dstX := 0; dstX < width; dstX++ {
                fx := (float64(dstX)+0.5)*w_ratio - 0.5

                srcX := int(math.Min(math.Max(math.Floor(fx+0.5), 0.0), i.Width64()))
                srcY := int(math.Min(math.Max(math.Floor(fy+0.5), 0.0), i.Height64()))

                srcOff := srcY*i.Image.Stride + srcX*4
                dstOff := dstY*img.Stride + dstX*4

                copy(img.Pix[dstOff:dstOff+4], i.Image.Pix[srcOff:srcOff+4])
            }
        }
    })
    return
}

func (i *Image) resizeW(width int, filter ResampleFilter) (img *image.NRGBA, err error) {
    img = image.NewNRGBA(image.Rect(0, 0, width, i.Height()))

    weights := precomputeWeights(width, i.Width(), filter)

    parallel(i.Height(), func(partStart, partEnd int) {
        for dstY := partStart; dstY < partEnd; dstY++ {
            for dstX := 0; dstX < width; dstX++ {
                var c [4]int32
                for _, iw := range weights[dstX].iwpairs {
                    i_position := dstY*i.Image.Stride + iw.i*4
                    c[0] += int32(i.Image.Pix[i_position+0]) * iw.w
                    c[1] += int32(i.Image.Pix[i_position+1]) * iw.w
                    c[2] += int32(i.Image.Pix[i_position+2]) * iw.w
                    c[3] += int32(i.Image.Pix[i_position+3]) * iw.w
                }
                j := dstY*img.Stride + dstX*4
                sum := weights[dstX].wsum
                img.Pix[j+0] = clampint32(int32(float32(c[0])/float32(sum) + 0.5))
                img.Pix[j+1] = clampint32(int32(float32(c[1])/float32(sum) + 0.5))
                img.Pix[j+2] = clampint32(int32(float32(c[2])/float32(sum) + 0.5))
                img.Pix[j+3] = clampint32(int32(float32(c[3])/float32(sum) + 0.5))
            }
        }
    })
    return
}

func (i *Image) resizeH(height int, filter ResampleFilter) (img *image.NRGBA, err error) {
    img = image.NewNRGBA(image.Rect(0, 0, i.Width(), height))

    weights := precomputeWeights(height, i.Height(), filter)

    parallel(i.Width(), func(partStart, partEnd int) {
        for dstX := partStart; dstX < partEnd; dstX++ {
            for dstY := 0; dstY < height; dstY++ {
                var c [4]int32
                for _, iw := range weights[dstY].iwpairs {
                    i_position := iw.i*i.Image.Stride + dstX*4
                    c[0] += int32(i.Image.Pix[i_position+0]) * iw.w
                    c[1] += int32(i.Image.Pix[i_position+1]) * iw.w
                    c[2] += int32(i.Image.Pix[i_position+2]) * iw.w
                    c[3] += int32(i.Image.Pix[i_position+3]) * iw.w
                }
                j := dstY*img.Stride + dstX*4
                sum := weights[dstY].wsum
                img.Pix[j+0] = clampint32(int32(float32(c[0])/float32(sum) + 0.5))
                img.Pix[j+1] = clampint32(int32(float32(c[1])/float32(sum) + 0.5))
                img.Pix[j+2] = clampint32(int32(float32(c[2])/float32(sum) + 0.5))
                img.Pix[j+3] = clampint32(int32(float32(c[3])/float32(sum) + 0.5))
            }
        }
    })
    return
}
