package image

import (
    "image"
    "image/color"
)

func Clone(img image.Image) *image.NRGBA {
    srcBounds := img.Bounds()
    srcMinX := srcBounds.Min.X
    srcMinY := srcBounds.Min.Y

    dstBounds := srcBounds.Sub(srcBounds.Min)
    dstW := dstBounds.Dx()
    dstH := dstBounds.Dy()
    dst := image.NewNRGBA(dstBounds)

    switch src := img.(type) {

        case *image.NRGBA:
            rowSize := srcBounds.Dx() * 4
            parallel(dstH, func(partStart, partEnd int) {
                for dstY := partStart; dstY < partEnd; dstY++ {
                    di := dst.PixOffset(0, dstY)
                    si := src.PixOffset(srcMinX, srcMinY+dstY)
                    copy(dst.Pix[di:di+rowSize], src.Pix[si:si+rowSize])
                }
            })

        case *image.NRGBA64:
		parallel(dstH, func(partStart, partEnd int) {
			for dstY := partStart; dstY < partEnd; dstY++ {
				di := dst.PixOffset(0, dstY)
				si := src.PixOffset(srcMinX, srcMinY+dstY)
				for dstX := 0; dstX < dstW; dstX++ {

					dst.Pix[di+0] = src.Pix[si+0]
					dst.Pix[di+1] = src.Pix[si+2]
					dst.Pix[di+2] = src.Pix[si+4]
					dst.Pix[di+3] = src.Pix[si+6]

					di += 4
					si += 8

				}
			}
		})

	case *image.RGBA:
		parallel(dstH, func(partStart, partEnd int) {
			for dstY := partStart; dstY < partEnd; dstY++ {
				di := dst.PixOffset(0, dstY)
				si := src.PixOffset(srcMinX, srcMinY+dstY)
				for dstX := 0; dstX < dstW; dstX++ {

					a := src.Pix[si+3]
					dst.Pix[di+3] = a
					switch a {
					case 0:
						dst.Pix[di+0] = 0
						dst.Pix[di+1] = 0
						dst.Pix[di+2] = 0
					case 0xff:
						dst.Pix[di+0] = src.Pix[si+0]
						dst.Pix[di+1] = src.Pix[si+1]
						dst.Pix[di+2] = src.Pix[si+2]
					default:
						dst.Pix[di+0] = uint8(uint16(src.Pix[si+0]) * 0xff / uint16(a))
						dst.Pix[di+1] = uint8(uint16(src.Pix[si+1]) * 0xff / uint16(a))
						dst.Pix[di+2] = uint8(uint16(src.Pix[si+2]) * 0xff / uint16(a))
					}

					di += 4
					si += 4

				}
			}
		})

	case *image.RGBA64:
		parallel(dstH, func(partStart, partEnd int) {
			for dstY := partStart; dstY < partEnd; dstY++ {
				di := dst.PixOffset(0, dstY)
				si := src.PixOffset(srcMinX, srcMinY+dstY)
				for dstX := 0; dstX < dstW; dstX++ {

					a := src.Pix[si+6]
					dst.Pix[di+3] = a
					switch a {
					case 0:
						dst.Pix[di+0] = 0
						dst.Pix[di+1] = 0
						dst.Pix[di+2] = 0
					case 0xff:
						dst.Pix[di+0] = src.Pix[si+0]
						dst.Pix[di+1] = src.Pix[si+2]
						dst.Pix[di+2] = src.Pix[si+4]
					default:
						dst.Pix[di+0] = uint8(uint16(src.Pix[si+0]) * 0xff / uint16(a))
						dst.Pix[di+1] = uint8(uint16(src.Pix[si+2]) * 0xff / uint16(a))
						dst.Pix[di+2] = uint8(uint16(src.Pix[si+4]) * 0xff / uint16(a))
					}

					di += 4
					si += 8

				}
			}
		})

	case *image.Gray:
		parallel(dstH, func(partStart, partEnd int) {
			for dstY := partStart; dstY < partEnd; dstY++ {
				di := dst.PixOffset(0, dstY)
				si := src.PixOffset(srcMinX, srcMinY+dstY)
				for dstX := 0; dstX < dstW; dstX++ {

					c := src.Pix[si]
					dst.Pix[di+0] = c
					dst.Pix[di+1] = c
					dst.Pix[di+2] = c
					dst.Pix[di+3] = 0xff

					di += 4
					si += 1

				}
			}
		})

	case *image.Gray16:
		parallel(dstH, func(partStart, partEnd int) {
			for dstY := partStart; dstY < partEnd; dstY++ {
				di := dst.PixOffset(0, dstY)
				si := src.PixOffset(srcMinX, srcMinY+dstY)
				for dstX := 0; dstX < dstW; dstX++ {

					c := src.Pix[si]
					dst.Pix[di+0] = c
					dst.Pix[di+1] = c
					dst.Pix[di+2] = c
					dst.Pix[di+3] = 0xff

					di += 4
					si += 2

				}
			}
		})

	case *image.YCbCr:
		parallel(dstH, func(partStart, partEnd int) {
			for dstY := partStart; dstY < partEnd; dstY++ {
				di := dst.PixOffset(0, dstY)
				switch src.SubsampleRatio {
				case image.YCbCrSubsampleRatio422:
					siy0 := dstY * src.YStride
					sic0 := dstY * src.CStride
					for dstX := 0; dstX < dstW; dstX = dstX + 1 {
						siy := siy0 + dstX
						sic := sic0 + ((srcMinX+dstX)/2 - srcMinX/2)
						r, g, b := color.YCbCrToRGB(src.Y[siy], src.Cb[sic], src.Cr[sic])
						dst.Pix[di+0] = r
						dst.Pix[di+1] = g
						dst.Pix[di+2] = b
						dst.Pix[di+3] = 0xff
						di += 4
					}
				case image.YCbCrSubsampleRatio420:
					siy0 := dstY * src.YStride
					sic0 := ((srcMinY+dstY)/2 - srcMinY/2) * src.CStride
					for dstX := 0; dstX < dstW; dstX = dstX + 1 {
						siy := siy0 + dstX
						sic := sic0 + ((srcMinX+dstX)/2 - srcMinX/2)
						r, g, b := color.YCbCrToRGB(src.Y[siy], src.Cb[sic], src.Cr[sic])
						dst.Pix[di+0] = r
						dst.Pix[di+1] = g
						dst.Pix[di+2] = b
						dst.Pix[di+3] = 0xff
						di += 4
					}
				case image.YCbCrSubsampleRatio440:
					siy0 := dstY * src.YStride
					sic0 := ((srcMinY+dstY)/2 - srcMinY/2) * src.CStride
					for dstX := 0; dstX < dstW; dstX = dstX + 1 {
						siy := siy0 + dstX
						sic := sic0 + dstX
						r, g, b := color.YCbCrToRGB(src.Y[siy], src.Cb[sic], src.Cr[sic])
						dst.Pix[di+0] = r
						dst.Pix[di+1] = g
						dst.Pix[di+2] = b
						dst.Pix[di+3] = 0xff
						di += 4
					}
				default:
					siy0 := dstY * src.YStride
					sic0 := dstY * src.CStride
					for dstX := 0; dstX < dstW; dstX++ {
						siy := siy0 + dstX
						sic := sic0 + dstX
						r, g, b := color.YCbCrToRGB(src.Y[siy], src.Cb[sic], src.Cr[sic])
						dst.Pix[di+0] = r
						dst.Pix[di+1] = g
						dst.Pix[di+2] = b
						dst.Pix[di+3] = 0xff
						di += 4
					}
				}
			}
		})

	case *image.Paletted:
		plen := len(src.Palette)
		pnew := make([]color.NRGBA, plen)
		for i := 0; i < plen; i++ {
			pnew[i] = color.NRGBAModel.Convert(src.Palette[i]).(color.NRGBA)
		}

		parallel(dstH, func(partStart, partEnd int) {
			for dstY := partStart; dstY < partEnd; dstY++ {
				di := dst.PixOffset(0, dstY)
				si := src.PixOffset(srcMinX, srcMinY+dstY)
				for dstX := 0; dstX < dstW; dstX++ {

					c := pnew[src.Pix[si]]
					dst.Pix[di+0] = c.R
					dst.Pix[di+1] = c.G
					dst.Pix[di+2] = c.B
					dst.Pix[di+3] = c.A

					di += 4
					si += 1

				}
			}
		})

	default:
		parallel(dstH, func(partStart, partEnd int) {
			for dstY := partStart; dstY < partEnd; dstY++ {
				di := dst.PixOffset(0, dstY)
				for dstX := 0; dstX < dstW; dstX++ {

					c := color.NRGBAModel.Convert(img.At(srcMinX+dstX, srcMinY+dstY)).(color.NRGBA)
					dst.Pix[di+0] = c.R
					dst.Pix[di+1] = c.G
					dst.Pix[di+2] = c.B
					dst.Pix[di+3] = c.A

					di += 4

				}
			}
		})

    }

    return dst
}
