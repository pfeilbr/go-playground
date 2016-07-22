package main

import "fmt"

func generateURLs(count int) []string {
	searchUrlBase := "https://www.google.com/#q="
	urls := make([]string, count)
	for i := 0; i < count; i++ {
		urls = append(urls, fmt.Sprintf("%s%d", searchUrlBase, i))
	}
	return urls
}

func main() {

	urlsToGenerate := 100
	urls := generateURLs(urlsToGenerate)

	for _, url := range urls {
		fmt.Println(url)
	}

	fmt.Println("worker-pool started")
}
