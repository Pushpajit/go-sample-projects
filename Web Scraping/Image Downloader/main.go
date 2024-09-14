package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

type Image struct {
	Title string
	URL   string
}

// Number of workers
const numWorkers = 5

func main() {
	imageChannel := make(chan Image, 30)
	// var images []Image
	var promt string
	var URL string
	reader := bufio.NewReader(os.Stdin)

	// creating waitgroup
	wg := &sync.WaitGroup{}

	// take promt and search images based on the promt
	fmt.Printf("What kind of image do you want: ")
	promt, _ = reader.ReadString('\n')
	// promt = strings.Trim(strings.Join(strings.Split(promt, " "), "-"), "\r\n")
	fmt.Printf("You're searching for %+#v\n", promt)

	// Start the performance timer
	startTime := time.Now()

	c := colly.NewCollector(
		colly.AllowedDomains("www.desktophut.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Scrapping the URL: %v\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Fatal(err.Error())
	})

	counter := 0
	c.OnHTML("img[class='customimg img-fluid rounded  ']", func(h *colly.HTMLElement) {
		// Start sending the scrapped URL into the channel, then it will also start downloading.
		// images = append(images, Image{
		// 	Title: h.Attr("alt"),
		// 	URL:   h.Attr("src"),
		// })
		if counter <= 5 {
			imageChannel <- Image{
				Title: h.Attr("alt"),
				URL:   h.Attr("src"),
			}

			counter++
		}
	})

	// c.OnScraped(func(r *colly.Response) {
	// 	wg.Add(len(images)) // Add the go-routines

	// 	for _, item := range images {
	// 		go func() {
	// 			defer wg.Done()
	// 			if err := download(item.URL, fmt.Sprintf("./images/%v.jpg", item.Title)); err != nil {
	// 				fmt.Printf("Failed Downloading %v\n[ERROR]: %v", item.Title, err.Error())

	// 			}
	// 		}()
	// 	}
	// 	wg.Wait()
	// })

	// activate the go-routine for the channel, also channels need seperate thread thats why we make nested threads
	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go workers(imageChannel, wg)
	}

	// URL = fmt.Sprintf("https://unsplash.com/s/photos/%v?license=free&orientation=landscape", promt)
	URL = fmt.Sprintf("https://www.desktophut.com/search/%v", promt)
	fmt.Println(URL)

	c.Visit(URL) // Initiate the web sracpping

	close(imageChannel) // Close the channel
	wg.Wait()           // Wait for all the go ruutine to finish their work

	elapsed := time.Since(startTime)
	fmt.Printf("💿 All Download Completed ✅\n⏳ Time Taken %vs ", elapsed.Seconds())
}

// Download the image
func download(baseURl string, filepath string) error {
	// Create the file
	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Get the image data
	res, err := http.Get(baseURl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Check the status
	if res.StatusCode != http.StatusOK {
		log.Fatal("status is not 200")
	}

	// download the binary content into the file
	_, err = io.Copy(file, res.Body)

	if err == nil {
		fmt.Printf("%v download completed [✅]\n", filepath)
	}

	if err != nil {
		fmt.Printf("%v download failed [❌]\n", filepath)
	}

	return err

}

func workers(jobs <-chan Image, wg *sync.WaitGroup) {
	defer wg.Done()

	for item := range jobs {
		if err := download(item.URL, fmt.Sprintf("./images/%v.jpg", item.Title)); err != nil {
			fmt.Printf("Failed Downloading %v\n[ERROR]: %v", item.Title, err.Error())

		}

	}
}
