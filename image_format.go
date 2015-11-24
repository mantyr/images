package images

func (i *Image) SetJpeg() *Image {
    i.Format = "jpeg"
    return i
}
func (i *Image) SetPng() *Image {
    i.Format = "png"
    return i
}
func (i *Image) SetTiff() *Image {
    i.Format = "tiff"
    return i
}
func (i *Image) SetBmp() *Image {
    i.Format = "bmp"
    return i
}
func (i *Image) SetGif() *Image {
    i.Format = "gif"
    return i
}
