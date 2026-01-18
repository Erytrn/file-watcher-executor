package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// İZLENECEK KLASÖR: "." koyarsan programın çalıştığı klasörü izler.
const watchDir = "."
const debounceTime = 500 * time.Millisecond

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	var timer *time.Timer

	// Olayları dinleyen ana döngü
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				
				// Sadece YAZMA veya OLUŞTURMA olayları
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					
					// FİLTRELEME: Sadece .js ve .py dosyalarındaki değişimleri yakala
					if strings.HasSuffix(event.Name, ".js") || strings.HasSuffix(event.Name, ".py") {
						log.Println("Değişiklik algılandı:", event.Name)

						if timer != nil {
							timer.Stop()
						}
						timer = time.AfterFunc(debounceTime, func() {
							runCommand()
						})
					}
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Hata:", err)
			}
		}
	}()

	// Klasörü izleyiciye ekle
	err = filepath.Walk(watchDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("İzleme başladı: %s (Çıkış için CTRL+C)", watchDir)
	<-make(chan struct{})
}

func runCommand() {
	log.Println("Komut tetiklendi! İşlem yapılıyor...")
	
	// Ekrana yazı yazan basit bir komut
	cmd := exec.Command("cmd", "/c", "echo DOSYA DEGISIKLIGI ALGILANDI!")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}