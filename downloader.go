package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/* 'Downloader' is necessary. Each worker thread will have a Downloader "instance"
	If implemented as a function, it results in race condition as multiple worker threads
	will access it.
*/
type Downloader struct {
}

func (this Downloader) Download(url string) (string, error) {
	
	//  fetch the data at `url`
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Fatal: Download failed %s: %s", url, err)
	}
	defer response.Body.Close()

	// read response.Body
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response.Body %s: %s", url, err)
	}

	return string(responseData), nil
}
