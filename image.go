package image

import (
    mag "github.com/mantyr/golang-magick"
    "fmt"
    "errors"
    "strings"
    "os"
)

type Image struct {
    Image *mag.Image
    Address string
    Format string
    Quality int
    Error error
}

const (
    FormatTIFF string = "TIFF"
    FormatJPEG string = "JPEG"
    FormatGIF  string = "GIF"
    FormatPNG  string = "PNG"
)

var Formats []string

func init() {
    Formats, _ = mag.SupportedFormats()
}

func Open(address string) (image *Image, err error){
    image = new(Image)
    image.Address = address
    image.Quality = 100

    image.Image, image.Error = mag.DecodeFile(address)
    err = image.Error
    if err != nil {
        return
    }

    image.Format = image.Image.Format()
    return
}

func (i *Image) Width() int {
    return i.Image.Width()
}

func (i *Image) Height() int {
    return i.Image.Height()
}

func (i *Image) Width64() float64 {
    return float64(i.Image.Width())
}

func (i *Image) Height64() float64 {
    return float64(i.Image.Height())
}

func (i *Image) SetFormat(format string) {
    i.Format = format
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

    if err != nil {
        return
    }

    info := mag.NewInfo()
    info.SetFormat(i.Format)
    info.SetQuality(uint(i.Quality))

    err = i.Image.Encode(file, info)
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

    image.Image, image.Error = i.Image.CropResize(width, height, mag.FQuadratic, mag.CSCenter)
    if image.Error != nil {
        return
    }
    image.Format = image.Image.Format()
    return
}

func (i *Image) Resize(width, height int) (image *Image) {
    image = new(Image)
    defer func() {
        if r := recover(); r != nil {
            image.Error = errors.New(fmt.Sprintf("%v", r))
        }
    }()

    image.Image, image.Error = i.Image.Resize(width, height, mag.FQuadratic)
    if image.Error != nil {
        return
    }
    image.Format = image.Image.Format()
    return
}

