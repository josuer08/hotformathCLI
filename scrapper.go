package main

import (
    "context"
    "fmt"
    "github.com/rocketlaunchr/google-search"
)

func main() {
    ctx := context.Background()
    opts := googlesearch.SearchOptions{
    Limit: 3,
    CountryCode: "us",
    LanguageCode: "en",
    }
    q:="Mr Robot"
    common:=" -inurl:(jsp|pl|php|html|aspx|htm|cf|shtml) -inurl:(index_of|listen77|mp3raid|mp3toss|mp3drug|index_of|wallywashis) intitle:\"index.of./\" "
    returnLinks, err := googlesearch.Search(ctx, q, opts)
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
    //fmt.Println(googlesearch.Search(ctx, "cars for sale in Toronto, Canada"))
}
