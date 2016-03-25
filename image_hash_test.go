package images

import (
    "testing"
    "os"
    "runtime"
    "strings"
    "regexp"
    "strconv"
)

func TestHashFile(t *testing.T) {
    file, err := os.Open("./testdata/test.jpg")
    if err != nil {
        t.Errorf("Error image open, %q", err)
    }
    defer file.Close()
    hashfile := GetHashFile(file)
    if hashfile != "e269a4995ad439664251b38951448022706e037b40d243475f1bb3ae74329212" {
        t.Errorf("Error image hash256 file, %q", hashfile)
    }
}

func TestGetHash(t *testing.T) {
    img, err := Open("./testdata/test.jpg")
    if err != nil {
        t.Errorf("Error image open, %q", err)
    }
    img.SetPng()
    hashfile := img.GetHash()
    if hashfile != "8151432a314a835448963a4c33a6822c16e8d0bcbe3d178541c373e0fdfdc99a" {
        t.Errorf("Error image hash256 file, %q", hashfile)
    }
    img.Save("./testdata/test_hash_png.png")

    img.SetJpeg()
    hashfile = img.GetHash()
    if hashfile != "d136596f089ee8e32bc6af040e108a6282a636392f102bef4eebfa1a7fa47dc7" {
        t.Errorf("Error image hash256 file, %q", hashfile)
    }
    img.Save("./testdata/test_hash_jpeg.jpg")

    /* recheck hash files */

    hashfile = GetHashFileA("./testdata/test_hash_png.png")
    if hashfile != "8151432a314a835448963a4c33a6822c16e8d0bcbe3d178541c373e0fdfdc99a" {
        t.Errorf("Error image hash256 file, %q", hashfile)
    }

    hashfile = GetHashFileA("./testdata/test_hash_jpeg.jpg")
    if hashfile != "d136596f089ee8e32bc6af040e108a6282a636392f102bef4eebfa1a7fa47dc7" {
        t.Errorf("Error image hash256 file, %q", hashfile)
    }
}

func TestHashFileNegative(t *testing.T) {
    file, err := os.Open("./testdata/test_negative.jpg")
    if err != nil {
        t.Errorf("Error negative open, %q", err)
    }
    defer file.Close()
    hashfile := GetHashFile(file)

    var go_version float64
    go_version_devel := strings.Contains(runtime.Version(), "devel")

    re := regexp.MustCompile("[0-9]+")
    go_version_chunk := re.FindAllString(runtime.Version(), 2)

    if len(go_version_chunk) > 0 {
        go_version, _ = strconv.ParseFloat(strings.Join(go_version_chunk, "."), 64)
    }

    if !go_version_devel && go_version < 1.5 {
        if hashfile != "b4d65104a11a52df7ece664680d7db58a8ec83992b64d8f4699e7b0c2b3e1cb8" {
            t.Errorf("Error negative hash256 file, %q, %q", runtime.Version(), hashfile)
        }
    } else {
        // see https://github.com/golang/go/commit/28388c4eb102f3218bbbdcca4699de6b078bdde6#diff-1e31509dba8d6eff03847d207acdb790R304
        if hashfile != "c0e19e49bde43035047619dca96bc906bdd7e3172f62cc34fc4f2be2683b0760" {
            t.Errorf("Error negative hash256 file, %q, %q", runtime.Version(), hashfile)
        }
    }
}

