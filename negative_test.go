package images

import (
    "testing"
    "fmt"
)

func TestNegative(t *testing.T) {
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

    img_negative := img.Negative()
    img_negative.Format = "jpeg"

    if img_negative.Error != nil {
        t.Errorf("Error negative %q", img_negative.Error.Error())
    }
    if img_negative.Format != "jpeg" {
        t.Errorf("Error negative Format, %q", img_negative.Format)
    }
    if img_negative.Width() != 457 || img_negative.Height() != 343 {
        t.Errorf("Error negative size, %q %q", fmt.Sprintf("%v", img_negative.Width()), fmt.Sprintf("%v", img_negative.Height()))
    }
    if img_negative.Quality != 100 {
        t.Errorf("Error negative default quality, %q", img_negative.Quality)
    }

    err = img_negative.Save("./testdata/test_negative.jpg")
    if err != nil {
        t.Errorf("Error negative save, %q", err.Error())
    }

    img2, err := Open("./testdata/test_negative.jpg")
    if err != nil {
        t.Errorf("Error negative open, %q", err.Error())
    }
    if img2.Format != "jpeg" {
        t.Errorf("Error negative Format, %q", img.Format)
    }
    if img2.Width() != 457 || img2.Height() != 343 {
        t.Errorf("Error negative open size, %q %q", fmt.Sprintf("%v", img2.Width()), fmt.Sprintf("%v", img2.Height()))
    }

}
