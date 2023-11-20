// 监控目录文件变化
package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	err = watcher.Add("./test")
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			log.Println(event)
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println(err)
		}
	}
}
