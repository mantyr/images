package image

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
    fmt.Println(fmt.Sprintf("%v", img.Quality))

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
