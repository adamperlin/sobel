package main

import (
  "fmt"
  "log"
  "os"
  "flag"
  "strings"
  "image/png"
  "image/jpeg"
  "image"
  "sobel"
)

var (
  inFile *string
  outFile *string
)

const (
  outputFileMode = 644
)

func init() {
  inFile = flag.String("f", "", "input file")
  outFile = flag.String("o", "sobel", "input file")
}

func main() {
  flag.Parse()
  if *inFile == "" {
    fmt.Printf("error: must provide input image")
    fmt.Println("usage: ")
    flag.PrintDefaults()
    os.Exit(0)
  }

  contents, err := os.Open(*inFile)
  if err != nil {
    log.Fatalf("error: %s\n", err.Error())
  }
  defer contents.Close()


  img, ftype, err := image.Decode(contents)
  if err != nil {
    log.Printf("error: %s\n", err.Error())
    log.Fatalf("filetype %s not supported\n", ftype)
  }

  var edged = sobel.Filter(img)

  var ext string

  if ftype != "png" && ftype != "jpeg" {
      log.Fatalf("can't encode file type")
      os.Exit(0)
  } else if ftype == "png" {
    if !strings.Contains(*outFile, ".png") {
      ext = ".png" 
    } else {
      ext = ""
    }
    out, err := os.Create(*outFile + ext)
    handleError(err)
    defer out.Close()
    err = png.Encode(out, edged)
    handleError(err)
  } else if ftype == "jpeg" {
    if !strings.Contains(*outFile, ".jpg") {
      ext = ".jpg"
    } else {
      ext = ""
    }
    out, err := os.Create(*outFile + ext)
    handleError(err)
    defer out.Close()
    err = jpeg.Encode(out, edged, nil)
    handleError(err)
  }
}

func handleError(err error) {
  if err != nil {
    log.Fatalf("error: %s\n", err.Error())
  }
}
