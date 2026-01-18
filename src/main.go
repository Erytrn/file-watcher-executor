package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// --- AYARLAR ---
// Senin klasör adın "workspace" olduğu için burayı güncelledim:
const watchDir = "../workspace"
const debounceTime = 500 * time.Millisecond

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	var timer *time.Timer

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				// Sadece KAYDETME (Write) olayları
				if event.Op&fsnotify.Write == fsnotify.Write {

					// Filtreleme: .py veya .js
					if strings.HasSuffix(event.Name, ".py") || strings.HasSuffix(event.Name, ".js") {
						log.Printf("📝 Değişiklik: %s", filepath.Base(event.Name))

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

	// Klasörleri izlemeye başla
	err = filepath.Walk(watchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("❌ HATA: '%s' klasörü bulunamadı! Lütfen ismin doğru olduğundan emin ol.", watchDir)
	}

	fmt.Println("------------------------------------------------")
	fmt.Printf("👁️  GÖZCÜ DEVREDE! (Go Watcher)\n📁 İzlenen: %s\n", watchDir)
	fmt.Println("------------------------------------------------")

	<-make(chan struct{})
}

// --- DÜZELTİLEN FONKSİYON BURASI ---
func runCommand() {
	fmt.Println("\n🚀 OTOMASYON BAŞLATILIYOR...")

	// 1. Python'a sadece dosya adını veriyoruz (Yolunu değil)
	cmd := exec.Command("python", "main.py")

	// 2. İŞTE EKSİK OLAN SATIR BU:
	// Komut çalışmadan önce "workspace" klasörünün içine giriyor.
	cmd.Dir = watchDir 

	// Çıktıları terminale ver
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("❌ Çalıştırma Hatası: %v\n", err)
	} else {
		fmt.Println("✅ İşlem Başarıyla Tamamlandı.")
	}
	fmt.Println("------------------------------------------------")
}