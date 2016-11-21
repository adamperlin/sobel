# Sobel Operator Edge Detection

This is a command line utility which implements edge detection using the [sobel filter](https://en.wikipedia.org/wiki/Sobel_operator). 

# Installation 
Given that this is a go project, installation is simple. 
The only tool you'll need to run this implementation is [go](https://golang.org). 

If go is installed, run: 
```
 $ git clone https://github.com/adamperlin/sobel && cd sobel
 $ go build 
``` 
# Usage 
Usage is relatively simple, there are just a few options.
This tool supports both png and jpeg image input formats, and file type is determined automatically. 

To run detection on a **specific file**, use: 

`$ ./edgedetect -f <file.[png, jpg]>`

You can optionally specify an **output file**:

`$ ./edgedetect -f <file.[png, jpg]> -o <output.[png, jpg]>`

Default output is `sobel.jpg` or `sobel.png`

*Note: this is not a conversion program. 
The output file will be of the same format as the inut file, so name files accordingly.* 

# Usage as a go library 
The internal package `sobel` can be used in any standard go program. It exposes a single function.
Here is an example:
```go 
package main 

import (
  "image"
  "os"
  "github.com/adamperlin/sobel" //package which implements the filter
 )
 
 var edge image.Image
 
 func main() {
    f, err := os.Open("example.jpg")
    if err != nil { panic(err) }
    defer f.Close()
    
    img, _, err := image.Decode(f)
    if err != nil { panic(err) }
    
    edge = sobel.Filter(img) //converts "img" to grayscale and runs edge detect. Returns an image.Image with changes.
    //do something with detected image...
 }
```
