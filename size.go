package images

func (i *Image) Width() int {
    return i.width
}

func (i *Image) Height() int {
    return i.height
}

func (i *Image) Width64() float64 {
    return float64(i.width)
}

func (i *Image) Height64() float64 {
    return float64(i.height)
}

func (i *Image) SetFormat(format string) {
    i.Format = format
}
