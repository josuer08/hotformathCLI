package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rocketlaunchr/google-search"
)

func main() {
    ctx := context.Background()
    opts := googlesearch.SearchOptions{
    Limit: 3,
    CountryCode: "us",
    LanguageCode: "en",
    }



    common:=" -inurl:(jsp|pl|php|html|aspx|htm|cf|shtml) -inurl:(index_of|listen77|mp3raid|mp3toss|mp3drug|index_of|wallywashis) intitle:\"index.of./\" "

    VIDEO:=" (avi|mkv|mov|mp4|mpg|wmv)"
    AUDIO:=" (ac3|flac|m4a|mp3|ogg|wav|wma)"
    EBOOK:=" (CBZ|CBR|CHM|DOC|DOCX|EPUB|MOBI|ODT|PDF|RTF|txt)"
    PICTURES:=" (bmp|gif|jpg|png|psd|tif|tiff)"
    SOFTWARE:=" (apk|exe|iso|rar|tar|zip)"
    COMPRESSED:=" (7z|bz2|gz|iso|rar|zip)"

    // implement getModifier()

    modifier:=""
    switch os.Args[1] {
        case "-v":
            modifier = VIDEO
        case "-a":
            modifier = AUDIO
        case "-e":
            modifier = EBOOK
        case "-p":
            modifier = PICTURES
        case "-s":
            modifier = SOFTWARE
        case "-c":
            modifier = COMPRESSED
        default:
            fmt.Println("Usage: scrapper -[flag] \"[your search]\"")
            fmt.Println("flags = [a]udio [c]ompressed [e]book [p]ictures [s]oftware [v]ideo")
            os.Exit(0)
    }

    query := "intext:\"" + os.Args[2] + "\"" + common + modifier


    returnLinks, err := googlesearch.Search(ctx, query, opts)
    if err != nil {
        fmt.Println("Something went wrong: ", err)
    return
    }
    if len(returnLinks) == 0 {
        fmt.Println("no results returned: ", returnLinks)
    }

    for _, result := range returnLinks {
        fmt.Println("*", result.URL)

    }
}
