package uaparser

import (
	"sync"
	"testing"
)

var uagent []string = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0 Firefox/72.0",
	"Mozilla/5.0 (X11; Linux armv7l) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.81 Safari/537.36 CrKey/1.42.172094",
	"libhttp libhttp/7.02 (PlayStation 4)",
	"ReactNativeVideo/7.2.0 (Linux;Android 5.1.1) AmznExoPlayerLib/2.9.6",
	"discoveryPlayer/8.3.0 (Linux;Android 9) ExoPlayerLib/2.10.8",
	"AppleCoreMedia/1.0.0.17K449 (Apple TV; U; CPU OS 13_3 like Mac OS X; en_us)",
	"Mozilla/5.0 (Windows NT 6.0) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 TV Safari/537.36",
	"Mozilla/5.0 (SMART-TV; LINUX; Tizen 3.0) AppleWebKit/538.1 (KHTML, like Gecko) Version/3.0 TV Safari/538.1",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
	"Mozilla/5.0 (X11; Linux armv7l) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/ AppleWebKit/537.36 (KHTML, like Gecko) Version/5.0 TV Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One; WebView/3.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18363",
	"AppleCoreMedia/1.0.0.17C54 (iPhone; U; CPU OS 13_3 like Mac OS X; en_us)",
	"AppleCoreMedia/1.0.0.17C54 (iPad; U; CPU OS 13_3 like Mac OS X; en_gb)",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:72.0) Gecko/20100101 Firefox/72.0",
	"Mozilla/5.0 (Linux; Android 4.4.4; SM-T560) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Safari/537.36",
	"Mozilla/5.0 (Linux; Android 8.1.0; SM-T580) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Safari/537.36",
	"Mozilla/5.0 (Linux; Android 9; SM-G950F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 9; SM-A505FN) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
	"AppleCoreMedia/1.0.0.17J586 (Apple TV; U; CPU OS 13_0 like Mac OS X; nl_be)",
	"Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 10; Pixel 3a) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 9; SM-A202F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 4.4.2; YOGA Tablet 2-1050F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Safari/537.36",
	"Mozilla/5.0 (Linux; Android 9.0; tablet) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Safari/537.36",
	"Mozilla/5.0 (Android 5.1.1; Tablet; rv:68.0) Gecko/68.0 Firefox/68.0",
	"Mozilla/5.0 (Linux; Android 5.0.2; SM-T530) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.116 Safari/537.36",
	"Mozilla/5.0 (SMART-TV; Linux; Tizen 2.3) AppleWebkit/538.1 (KHTML, like Gecko) SamsungBrowser/1.0 TV Safari/538.1",
	"Mozilla/5.0 (Linux; Android 9; SHIELD Android TV Build/PPR1.180610.011; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/79.0.3945.136 Mobile Safari/537.36"}

var globalClient *Client

func BenchmarkUaParser(b *testing.B) {
	b.ReportAllocs()
	parser := NewFromSaved()
	var client *Client
	for n := 0; n < b.N; n++ {
		for _, userAgent := range uagent {
			client = parser.Parse(userAgent)
		}
	}
	globalClient = client
}

var globalClients []*Client

func uaParserWorker(userAgent string, i int, wg *sync.WaitGroup, parser *Parser, clients []*Client) {
	client := parser.Parse(userAgent)
	clients[i] = client
	wg.Done()
}

func BenchmarkUaParserParallel(b *testing.B) {
	b.ReportAllocs()
	parser := NewFromSaved()
	for n := 0; n < b.N; n++ {
		waitGroup := sync.WaitGroup{}
		clients := make([]*Client, len(uagent))
		for i, userAgent := range uagent {
			waitGroup.Add(1)
			go uaParserWorker(userAgent, i, &waitGroup, parser, clients)
		}
		waitGroup.Wait()
		globalClients = clients
	}
}
