package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}

	return true
}

func TestWebsiteChecker(t *testing.T) {
	t.Run("Website checker", func(t *testing.T) {
		websites := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}
		expected := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}

		actual := CheckWebsites(mockWebsiteChecker, websites)

		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("Actual %v not equal %v", actual, expected)
		}
	})
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := range urls {
		urls[i] = "URL"
	}

	for range b.N {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
