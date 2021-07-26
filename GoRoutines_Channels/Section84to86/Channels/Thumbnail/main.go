package main

import (
	"TheGoProgLangBook/GoRoutines_Channels/Section84to/Channels/Thumbnail/thumbnail"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var filenames []string

func main() {
	filenames = []string{"image1.jpg", "image2.jpg", "image5.jpg"}
	t1 := time.Now()
	//create a channel to pass string
	ch_s := make(chan string, len(filenames))
	go func() {
		for _, v := range filenames {
			ch_s <- v
		}
		close(ch_s)
	}()
	sizes := makeThumbnails(ch_s)
	fmt.Println(sizes)
	fmt.Println(time.Since(t1))
}

func makeThumbnails(ch_s chan string) int64 {
	var wg sync.WaitGroup
	sizes := make(chan int64)
	for v := range ch_s {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(v)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(v)
	}
	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
