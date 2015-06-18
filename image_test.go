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
    if img.Format != "JPEG" {
        t.Errorf("Error Format, %q", img.Format)
    }
    if img.Width() != 457 || img.Height() != 343 {
        t.Errorf("Error size, %q %q", fmt.Sprintf("%v", img.Width()), fmt.Sprintf("%v", img.Height()))
    }

    img_small := img.Resize(100, 100)
    if img_small.Error != nil {
        t.Errorf("Error resize %q", img_small.Error.Error())
    }
    if img_small.Format != "JPEG" {
        t.Errorf("Error resize Format, %q", img_small.Format)
    }
    if img_small.Width() != 100 || img_small.Height() != 100 {
        t.Errorf("Error resize size, %q %q", fmt.Sprintf("%v", img_small.Width()), fmt.Sprintf("%v", img_small.Height()))
    }

    img_small = img.ResizeMin(100, 100)
    if img_small.Error != nil {
        t.Errorf("Error resize %q", img_small.Error.Error())
    }
    if img_small.Format != "JPEG" {
        t.Errorf("Error resize Format, %q", img_small.Format)
    }
    if img_small.Width() != 133 || img_small.Height() != 100 {
        t.Errorf("Error resize min size, %q %q", fmt.Sprintf("%v", img_small.Width()), fmt.Sprintf("%v", img_small.Height()))
    }


    img_small = img.ResizeMax(100, 100)
    if img_small.Error != nil {
        t.Errorf("Error resize %q", img_small.Error.Error())
    }
    if img_small.Format != "JPEG" {
        t.Errorf("Error resize Format, %q", img_small.Format)
    }
    if img_small.Width() != 100 || img_small.Height() != 75 {
        t.Errorf("Error resize max size, %q %q", fmt.Sprintf("%v", img_small.Width()), fmt.Sprintf("%v", img_small.Height()))
    }

    img_small = img.ResizeCrop(100, 100)
    if img_small.Error != nil {
        t.Errorf("Error resize %q", img_small.Error.Error())
    }
    if img_small.Format != "JPEG" {
        t.Errorf("Error resize Format, %q", img_small.Format)
    }
    if img_small.Width() != 100 || img_small.Height() != 100 {
        t.Errorf("Error resize crop size, %q %q", fmt.Sprintf("%v", img_small.Width()), fmt.Sprintf("%v", img_small.Height()))
    }
}

func TestSaveAndConvert(t *testing.T) {
//    img, err := Open("./testdata/test.jpg")
}