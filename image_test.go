package images

import (
    "testing"
    "fmt"
)

func TestOpenAndResize(t *testing.T) {
    img, err := Open("./testdata/test.jpg")
    if img.Error != nil {
        t.Errorf("Error %q", err.Error())
    }
    if img.Format != "jpeg" {
        t.Errorf("Error Format, %q", img.Format)
    }
    if img.Width() != 457 || img.Height() != 343 {
        t.Errorf("Error size, %q %q", fmt.Sprintf("%v", img.Width()), fmt.Sprintf("%v", img.Height()))
    }
    if img.Quality != 100 {
        t.Errorf("Error default quality, %q", img.Quality)
    }

    img_small := img.Resize(100, 100)
    img_small.Format = "jpeg"

    if img_small.Error != nil {
        t.Errorf("Error resize %q", img_small.Error.Error())
    }
    if img_small.Format != "jpeg" {
        t.Errorf("Error resize Format, %q", img_small.Format)
    }
    if img_small.Width() != 100 || img_small.Height() != 100 {
        t.Errorf("Error resize size, %q %q", fmt.Sprintf("%v", img_small.Width()), fmt.Sprintf("%v", img_small.Height()))
    }
    if img_small.Quality != 100 {
        t.Errorf("Error resize default quality, %q", img_small.Quality)
    }

    err = img_small.Save("./testdata/test_save.jpg")
    if err != nil {
        t.Errorf("Error resize save, %q", err.Error())
    }
    img2, err := Open("./testdata/test_save.jpg")
    if err != nil {
        t.Errorf("Error resize open, %q", err.Error())
    }
    if img2.Format != "jpeg" {
        t.Errorf("Error Format, %q", img.Format)
    }
    if img2.Width() != 100 || img2.Height() != 100 {
        t.Errorf("Error resize open size, %q %q", fmt.Sprintf("%v", img2.Width()), fmt.Sprintf("%v", img2.Height()))
    }
}

func TestSave(t *testing.T) {
    img, _ := Open("./testdata/test.jpg")
    if img.Error != nil || img.Format != "jpeg" || img.Width() != 457 || img.Height() != 343 || img.Quality != 100 {
        t.Errorf("Error open image %q %q %q %q %q", img.Error, img.Format, img.Width(), img.Height(), img.Quality)
    }
    err := img.Save("./testdata/test_only_save.jpg")
    if err != nil {
        t.Errorf("Error save, %q", err.Error())
    }

    img2, err := Open("./testdata/test_only_save.jpg")
    if err != nil {
        t.Errorf("Error open save , %q", err.Error())
    }
    if img2.Format != "jpeg" {
        t.Errorf("Error save Format, %q", img.Format)
    }
    if img2.Width() != 457 || img2.Height() != 343 {
        t.Errorf("Error save size, %q %q", fmt.Sprintf("%v", img2.Width()), fmt.Sprintf("%v", img2.Height()))
    }
}
func TestResizeCrop(t *testing.T) {
    img, _ := Open("./testdata/test.jpg")
    if img.Error != nil || img.Format != "jpeg" || img.Width() != 457 || img.Height() != 343 || img.Quality != 100 {
        t.Errorf("Error open image %q %q %q %q %q", img.Error, img.Format, img.Width(), img.Height(), img.Quality)
    }
    img_crop := img.ResizeCrop(100, 100)
    if img_crop.Error != nil {
        t.Errorf("Error crop %q", img_crop.Error.Error())
    }
    if img_crop.Format != "jpeg" {
        t.Errorf("Error crop Format, %q", img_crop.Format)
    }
    if img_crop.Width() != 100 || img_crop.Height() != 100 {
        t.Errorf("Error crop size, %q %q", fmt.Sprintf("%v", img_crop.Width()), fmt.Sprintf("%v", img_crop.Height()))
    }
    if img_crop.Quality != 100 {
        t.Errorf("Error crop default quality, %q", img_crop.Quality)
    }

    err := img_crop.Save("./testdata/test_crop_save.jpg")
    if err != nil {
        t.Errorf("Error crop save, %q", err.Error())
    }

    img2, err := Open("./testdata/test_crop_save.jpg")
    if err != nil {
        t.Errorf("Error crop open, %q", err.Error())
    }
    if img2.Format != "jpeg" {
        t.Errorf("Error crop Format, %q", img.Format)
    }
    if img2.Width() != 100 || img2.Height() != 100 {
        t.Errorf("Error crop open size, %q %q", fmt.Sprintf("%v", img2.Width()), fmt.Sprintf("%v", img2.Height()))
    }
}