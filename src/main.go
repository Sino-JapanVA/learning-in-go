package main

import (
	"fmt"
)

type (
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func main() {
	x := [...]int{1, 2, 3, 4}
	// y := []int{5, 6, 7, 8}
	array(x)
	// array(y)
}

func array(p [4]int) {
	for index, val := range p {
		fmt.Printf("index:%d  val:%d\n", index, val)
	}
}
