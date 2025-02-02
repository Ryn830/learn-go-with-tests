package concurrency

type Websitechecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc Websitechecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
			results[url] = wc(u)
		}(url)
	}

	for range urls {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
