package images

import (
    "testing"
    "fmt"
    "os"
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

func TestHashFile(t *testing.T) {
    file, err := os.Open("./testdata/test.jpg")
    if err != nil {
        t.Errorf("Error image open, %q", err)
    }
    hashfile := GetHashFile(file)
    if hashfile != "e269a4995ad439664251b38951448022706e037b40d243475f1bb3ae74329212" {
        t.Errorf("Error image hash256 file, %q", hashfile)
    }
}

func TestHashFileNegative(t *testing.T) {
    file, err := os.Open("./testdata/test_negative.jpg")
    if err != nil {
        t.Errorf("Error negative open, %q", err)
    }
    hashfile := GetHashFile(file)
    if hashfile != "b4d65104a11a52df7ece664680d7db58a8ec83992b64d8f4699e7b0c2b3e1cb8" {
        t.Errorf("Error negative hash256 file, %q", hashfile)
    }
}

