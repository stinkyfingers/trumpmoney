package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// GOOS=js GOARCH=wasm go build -o main.wasm main.go && go run server/main.go

func main() {
	var err error
	done := make(chan bool)
	ev := make(chan bool)
	var cmd *exec.Cmd

	go func() {
		for range ev {
			err := quit(cmd)
			if err != nil {
				log.Fatal(err)
			}
			err = buildApp()
			if err != nil {
				log.Fatal(err)
			}
			cmd, err = run()
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	cmd, err = run()
	if err != nil {
		log.Fatal(err)
	}

	err = watch(ev, done)
	if err != nil {
		log.Fatal(err)
	}

}

func quit(cmd *exec.Cmd) error {
	if cmd != nil {
		err := cmd.Process.Kill()
		if err != nil {
			return err
		}
	}
	return nil
}

func run() (*exec.Cmd, error) {
	cmd := exec.Command("go", "run", "server/main.go")
	return cmd, cmd.Start()
}

func buildApp() error {
	os.Setenv("GOOS", "js")
	os.Setenv("GOARCH", "wasm")
	return exec.Command("go", "build", "-o", "main.wasm", "main.go").Run()
}

func watch(ev, done chan bool) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					ev <- true
				}

			case errors, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Fatal(errors)
			}
		}
	}()
	err = filepath.Walk(".", func(name string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		return watcher.Add(name)
	})
	if err != nil {
		return err
	}
	<-done
	return nil
}
