package code

/*
	文件监听 Demo
*/

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func RunTestWatcher() {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatalf("fail to create new Watcher: %v\n", err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("config.json")
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}
