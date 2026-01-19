# 📚 Copilot Araştırma Kaynakları

## 🔹 Mimari & Performans
- **fsnotify (Go):** Dosya sistemi olaylarını izlemek için kullanılan kütüphane.  
- **inotify (Linux):** Kernel-level file system event API.  
- **ReadDirectoryChangesW (Windows):** Windows API ile dizin değişikliklerini izleme.  
- **FSEvents (macOS):** Apple’ın dosya sistemi event mekanizması.  
- **Concurrency Patterns:** Producer–Consumer, Worker Pool, Async I/O (epoll, IOCP).

---

## 🔹 Güvenlik
- **Defensive Programming:** Canonicalization (`filepath.Clean`, `filepath.Abs`), symlink kontrolü, whitelist komutlar.  
- **TOCTOU Mitigation:** Dosya varlığını `os.Stat()` ile tekrar kontrol etme.  
- **Safe Exec:** `exec.Command(program, args...)` → shell kullanılmadan komut çalıştırma.  
- **Sandboxing:** Least privilege, restricted environment, temiz environment variables.  
- **Regex Güvenliği:** RE2 veya timeout’lu regex engine kullanımı.

---

## 🔹 Modern Kod Kalitesi
- **Go:** Goroutines, channels, structured logging (`zap`, `logrus`).  
- **C++17/20:** `std::filesystem`, `std::optional`, `std::variant`, RAII, `std::jthread`.  
- **Design Patterns:** Observer, Strategy, Command, Builder, Policy.  
- **Testing:** Unit test (`go test`), fuzzing (libFuzzer, AFL), property-based testing.

---

## 🔹 Sunum & Dokümantasyon
- **README.md:** Proje tanımı, mimari diyagram, kurulum, config örneği, güvenlik önlemleri, test, lisans, demo.  
- **project_info.json:** Meta veri formatı (isim, kategori, versiyon, özellikler, gereksinimler).  
- **Demo Video/GIF:** Terminal çıktısı veya ekran kaydı ile otomasyonun gösterimi.  
- **CI/CD Pipeline:** GitHub Actions veya benzeri ile otomatik test ve build.

---

## 🔹 İlgili Araçlar & Kaynaklar
- **Burp Suite, OWASP ZAP, SQLMap, Nikto:** Web güvenlik test araçları (sunumda güvenlik bağlamı için referans).  
- **MobSF, Drozer:** Mobil güvenlik frameworkleri.  
- **Wireshark, Snort, Wazuh:** Ağ analizi ve IDS/IPS araçları.  
- **Kali Linux:** Pentest dağıtımı, demo ortamı için referans.  
- **Structured Logging Libraries:** `zap`, `logrus` (Go), `spdlog` (C++).  
- **Concurrency Libraries:** `moodycamel::ConcurrentQueue` (C++), Go channels.

---

## 📌 Özet
Bu kaynaklar, projenin:
- **Mimari sağlamlığını** (debounce, concurrency, cross-platform),
- **Güvenlik dayanıklılığını** (command injection, path traversal, TOCTOU),
- **Kod kalitesini** (modern standartlar, design patterns),
- **Sunum gücünü** (README, meta dosya, demo, CI/CD)

desteklemek için önerilen araçlar ve tekniklerdir.
