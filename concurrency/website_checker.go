package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// Questions:
// 1. What does make do?
// Map, Chan or slice empty

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		/// () at the end execute the anonymous function
		/// Call the anonymous func w/ url, which is every
		/// separate url. U is a copy so it cannot be changed
		/// go test -race can detect race conditions
		go func(u string) {
			// Sending results to a channel
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Receive expression
		r := <-resultChannel
		// Controlling writing to a map to avoid the race
		// condition of writing to a map
		results[r.string] = r.bool
	}

	return results
}
