package images

import (
    "crypto/sha256"
    "encoding/hex"
    "os"
    "io"
)

func GetHashFile(file *os.File) string {
    hash := sha256.New()

    io.Copy(hash, file)

    return hex.EncodeToString(hash.Sum(nil))
}

func GetHashFileA(address string) string {
    file, err := os.Open(address)
    if err != nil {
        return ""
    }
    defer file.Close()

    hash := sha256.New()

    io.Copy(hash, file)

    return hex.EncodeToString(hash.Sum(nil))
}

