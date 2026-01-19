# SONUÇLAR: File Watcher Çözümleri

## PERFORMANS:
- Buffered channel (1000 kapasite) + worker pool
- Goroutine başına CPU sayısı × 2
- Event batching (100ms window)

## GÜVENLİK:
- Whitelist allowed commands
- exec.Command'da shell=false
- Path traversal için filepath.Clean() + ".." kontrolü
- syscall.Credential ile privilege drop

## DEBOUNCE:
- Per-file: Her dosya için timer (precise ama memory yüksek)
- Per-dir: Dizin bazlı (balanced)
- Global: Tek timer (hızlı ama yanlış pozitif)

## CROSS-PLATFORM:
- runtime.GOOS kontrolü
- Windows: `\\?\` long path prefix
- Linux: inotify tuning

## CONFIG:
- atomic.Value ile hot-reload
- fsnotify ile config dosyası izleme
- 1sn debounce ile reload

## METRİKLER:
- Prometheus counter, histogram
- Event sayısı, komut süresi, debounce gecikmesi