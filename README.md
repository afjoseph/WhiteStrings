# White Strings -- Go implementation of Simon White's String Similarity Algorithm

This is a Go port of the "Strike A Match" algorithm from Simon White of
Catalysoft (http://www.catalysoft.com/articles/StrikeAMatch.html)

## Usage

    go get github.com/afjoseph/whitestrings

## Example


    package main

    import (
            "fmt"

            "github.com/afjoseph/whitestrings"
    )

    func main() {
            // Result: 1.0
            fmt.Println(whitestrings.GetSimilarity("bunnyfoofoo", "bunnyfoofoo"))
            // Result: 0.5714285714285714
            fmt.Println(whitestrings.GetSimilarity("bunnyfoofoo", "bunny"))
            // Result: 0.9
            fmt.Println(whitestrings.GetSimilarity("bunnyfoofoo", "foofoobunny"))
            // Result: 0.9
            fmt.Println(whitestrings.GetSimilarity("bunnyfoofoo", "foobunnyfoo"))
    }
