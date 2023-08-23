# goquest
Simple http request for GO, inspired by Guzzle

## Creating a Client

```go
// import the package
import "github.com/fathalfath30/goquest"

func main()  {
	// create client with default configuration
	gq := goquest.New(nil)
	
	// create client with custom config 
	gq := goquest.New(&goquest.Config{ 
		Header:    &http.Header{},
		Transport: &http.Transport{},
	})
}
```