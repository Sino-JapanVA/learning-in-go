package main

import (
	"fmt"
	"runtime"
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
	cpus := runtime.NumCPU()
	fmt.Println(cpus)
}

func array(p [4]int) {
	for index, val := range p {
		fmt.Printf("index:%d  val:%d\n", index, val)
	}

}
