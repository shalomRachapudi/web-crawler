package main 

type WebCrawler struct {
	url string
	wc int
}

// channels
var pages = make(chan WebPageInfo)
var unseenUrls = make(chan string)
var visitedUrls = make(map[string]bool)


func NewWebCrawler(url string, workers int) *WebCrawler {
	return &WebCrawler{url, workers}
}

func (this WebCrawler) Crawl() error {

	for id := 0; id < this.wc; id++ { //dispatch wc number of worker threads
		worker := NewWorkerThread()
		go worker.Run(id, unseenUrls, pages);
	}

	go func() { //initialize with the given URL
		pages <- WebPageInfo{url: this.url, linksWithin: map[string]bool {this.url:true}}
	}()

	for pageCount := 1; pageCount > 0; pageCount-- {
		webPageInfo := <- pages
		pageCount += this.processWebPageInfo(webPageInfo); // adding links to pageCount because every unseen URL is a newly found page
		print(webPageInfo)
	}

	return nil // everything's okay
}

func (this WebCrawler) processWebPageInfo(wpInfo WebPageInfo) (links int) {
	links = 0;
	for key, _ := range wpInfo.linksWithin {
		if visitedUrls[key]  {
			continue
		}

		links++;
		visitedUrls[key] = true;
		unseenUrls <- key
	}
	return
}

type WorkerThread struct {
	downloader Downloader //required to avoid race condition
	webPage WebPage
} 

func NewWorkerThread() *WorkerThread {
	return &WorkerThread{downloader: Downloader{}, webPage: WebPage{}}
}

func (this WorkerThread) Run(id int, unseenUrls chan string, pages chan WebPageInfo) {
	for link := range unseenUrls {
		info := this.extractPageInfo(link)
		//pages <- info
		go func() {pages <- info}()
	}
}

func (this WorkerThread) extractPageInfo(url string) (info WebPageInfo) {
	htmlScript, err := this.downloader.Download(url) //download/fetch html page at the give url
	if err != nil {
		return
	}

	webPage := NewWebPage(htmlScript, url);
	info = webPage.ParseWebPage();
	return
}