package image

import (
    "fmt"
    "errors"
    "strings"
    "os"
    "image"
    "image/jpeg"
    "image/png"
    "image/gif"
//    _ "image/draw"
    "golang.org/x/image/tiff"
    "golang.org/x/image/bmp"
)

type Image struct {
    Image *image.NRGBA
    Address string
    Format string
    Quality int
    width int
    height int
    Error error
}

const (
    FormatTIFF string = "tiff"
    FormatJPEG string = "jpeg"
    FormatGIF  string = "gif"
    FormatPNG  string = "png"
    FormatBMP  string = "bmp"
)

var Formats = [5]string{"tiff", "jpeg", "gif", "png", "bmp"}

func Open(address string) (i *Image, err error){
    i = new(Image)
    i.Address = address
    i.Quality = 100

    file, err := os.Open(address)
    if err != nil {
        i.Error = err
        return
    }
    defer file.Close()
    var img image.Image
    img, i.Format, err = image.Decode(file)
    if err != nil {
        i.Error = err
        return
    }
    i.Image = toNRGBA(img)

    err = i.GetConfig()
    if err != nil {
        i.Error = err
        return
    }
    return
}

func (i *Image) GetConfig() (err error) {
    file, err := os.Open(i.Address)
    if err != nil {
        return
    }
    defer file.Close()
    conf, _, err := image.DecodeConfig(file)
    if err != nil {
        return
    }
    i.width  = conf.Width
    i.height = conf.Height
    return
}

func (i *Image) Save(params ...string) (err error) {
    if i.Image == nil {
        if i.Error != nil {
            return i.Error
        }
        return errors.New("no image")
    }

    address := i.Address
    if len(params) > 0 && strings.Trim(params[0], " \n\t\r") != "" {
        address = params[0]
    }

    file, err := os.Create(address)
    defer file.Close()

    switch i.Format {
        case FormatJPEG:
            var rgba *image.RGBA
            if i.Image.Opaque() {
                rgba = &image.RGBA{
                    Pix:    i.Image.Pix,
                    Stride: i.Image.Stride,
                    Rect:   i.Image.Rect,
                }
            }

            if rgba != nil {
                err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: i.Quality})
            } else {
                err = jpeg.Encode(file, i.Image, &jpeg.Options{Quality: i.Quality})
            }
        case FormatPNG:
            encoder := &png.Encoder{CompressionLevel: png.NoCompression}
            err = encoder.Encode(file, i.Image)
        case FormatGIF:
            err = gif.Encode(file, i.Image, &gif.Options{NumColors: 256})
        case FormatTIFF:
            err = tiff.Encode(file, i.Image, &tiff.Options{Compression: tiff.Uncompressed, Predictor: true})
        case FormatBMP:
            err = bmp.Encode(file, i.Image)
        default:
            err = errors.New("unsupported image format "+i.Format)
    }
    return
}

func (i *Image) ResizeMax(width, height int) (image *Image){
    var w_ratio float64 = float64(width) / i.Width64()
    var h_ratio float64 = float64(height) / i.Height64()

    if w_ratio < h_ratio {
        height = int(i.Height64() * w_ratio)
    } else {
        width  = int(i.Width64() * h_ratio)
    }
    return i.Resize(width, height)
}

func (i *Image) ResizeMin(width, height int) *Image{
    var w_ratio float64 = float64(width) / i.Width64()
    var h_ratio float64 = float64(height) / i.Height64()

    if w_ratio == h_ratio {
        return i.Resize(width, height)
    }
    var new_width int  = width
    var new_height int = height

    if w_ratio > h_ratio {
        new_height = int(i.Height64() * w_ratio)
    } else {
        new_width = int(i.Width64() * h_ratio)
    }
    return i.Resize(new_width, new_height)
}

func (i *Image) ResizeCrop(width, height int) (image *Image) {
    image = new(Image)
    defer func() {
        if r := recover(); r != nil {
            image.Error = errors.New(fmt.Sprintf("%v", r))
        }
    }()

//    image.Image, image.Error = i.Image.CropResize(width, height, mag.FQuadratic, mag.CSCenter)
    if image.Error != nil {
        return
    }
//    image.Format = image.Image.Format()
    return
}

func (i *Image) Resize(width, height int, params ...ResampleFilter) (image *Image) {
    image = new(Image)
    image.width = i.Width()
    image.height = i.Height()
    image.Quality = 100
    defer func() {
        if r := recover(); r != nil {
            image.Error = errors.New(fmt.Sprintf("%v", r))
        }
    }()

    filter := Lanczos
    if len(params) > 0 {
        filter = params[0]
    }

    if filter.Support <= 0.0 {
        image.Image, image.Error = i.resize(width, height)
    } else {
        if width != i.Width() {
            image.Image, image.Error = i.resizeW(width, filter)
            image.width = width
        }
        if image.Error != nil {
            return
        }
        if height != i.Height() {
            image.Image, image.Error = image.resizeH(height, filter)
            image.height = height
        }
    }
    image.Format = i.Format
    return
}

func toNRGBA(i image.Image) *image.NRGBA {
    srcBounds := i.Bounds()
    if srcBounds.Min.X == 0 && srcBounds.Min.Y == 0 {
        if src0, ok := i.(*image.NRGBA); ok {
            return src0
        }
    }
    return Clone(i)
}