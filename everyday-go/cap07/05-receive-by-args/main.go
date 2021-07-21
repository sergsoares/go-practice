package main

//https://github.com/bensooter/URLchecker/blob/master/top-1000-websites.txt
import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

type ScrapeRun struct {
	Results map[string]ScrapeResult
	Lock    *sync.Mutex
}

type ScrapeResult struct {
	HTTPCode   int
	HTTPLength int64
	Error      error
}

type ScrapeOutput struct {
	Url        string `json:"Url"`
	HTTPCode   int    `json:"HTTPCode"`
	HTTPLength int64  `json:"HTTPLength"`
	Error      error  `json:"Error"`
}

func scrape(url string, sr *ScrapeRun) {
	res, err := http.Get(url)
	if err != nil {
		sr.Lock.Lock()
		defer sr.Lock.Unlock()

		sr.Results[url] = ScrapeResult{
			Error:    err,
			HTTPCode: http.StatusBadGateway,
		}
		return
	}

	defer res.Body.Close()

	sr.Lock.Lock()
	defer sr.Lock.Unlock()

	length := res.ContentLength
	if length == -1 {
		body, _ := ioutil.ReadAll(res.Body)
		length = int64(len(body))
	}

	sr.Results[url] = ScrapeResult{
		HTTPCode:   res.StatusCode,
		HTTPLength: length,
	}
}

func main() {
	var path string
	var outputPath string
	flag.StringVar(&path, "file", "list.txt", "")
	flag.StringVar(&outputPath, "output", "output.json", "")
	flag.Parse()
	log.Println("Parsed flags - Path: ", path, "output: ", outputPath)

	log.Println("Reading file from path")
	content, _ := ioutil.ReadFile(path)
	test := string(content)
	v := strings.Split(test, "\n")
	log.Println("Total of urls: ", len(v))

	sr := ScrapeRun{
		Lock:    &sync.Mutex{},
		Results: make(map[string]ScrapeResult),
	}

	wg := sync.WaitGroup{}
	wg.Add(len(v))
	for _, e := range v {
		log.Println("Initializing scrape for: ", e)
		newValue, _ := url.Parse(e)
		if newValue.Scheme == "" {
			newValue.Scheme = "https"
		}
		go func(u string) {
			defer wg.Done()
			scrape(u, &sr)
		}(newValue.String())
	}
	wg.Wait()

	log.Print("Init output path", outputPath)
	file, err := os.Create(outputPath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	for k, v := range sr.Results {
		log.Print("Parsing results of: ", k)
		output := ScrapeOutput{
			Url:        k,
			HTTPCode:   v.HTTPCode,
			HTTPLength: v.HTTPLength,
			Error:      v.Error,
		}
		parsed, _ := json.Marshal(output)
		if _, err := file.WriteString(string(parsed) + "\n"); err != nil {
			log.Fatal(err)
		}
		// ioutil.WriteFile("output.json", parsed, 0644)
		// fmt.Println(string(parsed))
	}

	log.Printf("Finished")
}
