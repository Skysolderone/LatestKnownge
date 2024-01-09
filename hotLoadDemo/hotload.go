package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

//使用fsnotify  and os.exec

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()
	go watcherHandler(watcher)
	for {
		select {
		case event := <-watcher.Events:
			//process events
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("File modify :", event.Name)
				restartApp()
			}
		case err := <-watcher.Errors:
			log.Println("err:", err)

		}
	}
}
func watcherHandler(watcher *fsnotify.Watcher) {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".go" {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}
func restartApp() {
	log.Println("restarting log")
	cmd := exec.Command("go", "run", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Println("restarting err")
	}

	//Waitng close other operater
	time.Sleep(2 * time.Second)
	os.Exit(0)
}
