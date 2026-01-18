# 🔬 Teknik Araştırma ve Geliştirme Notları

**Konu:** Dosya Sistemi İzleme (File Watching) ve Otomasyon  
**Tarih:** 18.01.2026  
**Araştıran:** Eray

---

## 1. Mevcut Çözümlerin İncelenmesi
Proje öncesinde piyasadaki hazır araçlar incelendi:
* **Watchexec (Rust):** Çok hızlı ama özelleştirmesi zor.
* **Nodemon (JS):** Node.js bağımlılığı var, sistem kaynaklarını (RAM) fazla tüketiyor.
* **Karar:** Kendi hafif (lightweight) aracımızı **Go (Golang)** ile yazmaya karar verdik.

---

## 2. Dil ve Kütüphane Seçimi: Go vs Rust
Proje için iki düşük seviyeli dil değerlendirildi:

| Özellik | Go (Golang) | Rust |
| :--- | :--- | :--- |
| **Kütüphane** | `fsnotify/fsnotify` | `notify-rs` |
| **Öğrenme Eğrisi** | Düşük (Daha hızlı kodlandı) | Yüksek (Ownership kuralları zor) |
| **Eşzamanlılık** | Goroutines (Çok basit) | Async/Await (Karmaşık) |
| **SONUÇ** | ✅ **SEÇİLDİ** | ❌ EELENDİ |

---

## 3. Karşılaşılan Teknik Zorluklar ve Çözümler

### 🔴 Sorun 1: Sonsuz Döngü (Infinite Loop)
**Durum:** İzleyici (Watcher) ve Hedef Dosya aynı klasörde olduğunda, watcher log dosyasına yazdığında kendini tekrar tetikliyordu.
**Çözüm:** "Gözlemci" (`src`) ve "Kobay" (`workspace`) klasörleri birbirinden fiziksel olarak ayrıldı.

### 🔴 Sorun 2: "Hayalet" Tetiklenmeler (Debounce)
**Durum:** Bir dosya kaydedildiğinde editörler (VS Code) bazen birden fazla "Write" sinyali gönderiyor. Bu da kodun 2-3 kere çalışmasına neden oluyordu.
**Çözüm:** 500ms'lik bir `Timer` (Debounce) mekanizması eklendi. Sinyal gelince sayaç sıfırlanıyor, sadece son sinyal işleniyor.

### 🔴 Sorun 3: Windows vs Linux Farkı
**Durum:** Linux'ta `python3` ve `systemctl` komutları çalışırken, Windows geliştirme ortamında bu komutlar hata veriyordu (`exit status 9009`).
**Çözüm:**
* **Go Tarafında:** `exec.Command("python")` kullanıldı ve `cmd.Dir` ile çalışma dizini düzeltildi.
* **Python Tarafında:** `platform.system()` kontrolü ile Windows'ta olduğumuz algılanıp "Simülasyon Modu" devreye alındı.

---

## 4. Kaynakça ve Referanslar
1.  **fsnotify GitHub:** https://github.com/fsnotify/fsnotify
2.  **Go `os/exec` Dokümantasyonu:** https://pkg.go.dev/os/exec
3.  **Python `subprocess` Modülü:** https://docs.python.org/3/library/subprocess.html