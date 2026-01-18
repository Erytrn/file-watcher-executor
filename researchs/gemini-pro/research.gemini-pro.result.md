# Geliştirme Süreci ve Teknik Kararlar

**1. Teknoloji Seçimi:**
Sistem kaynaklarını yormayan ve VirtualBox gerektirmeyen bir yapı için **Go (Golang)** dili seçildi. Python (Watchdog) yerine Go'nun seçilme sebebi, "Goroutines" yapısı sayesinde asenkron dosya izleme (Event Loop) işlemlerinde daha yüksek performans sunmasıdır.

**2. Mimari Kararlar:**
* **Kütüphane:** Kernel seviyesindeki olayları yakalamak için `fsnotify/fsnotify` kullanıldı.
* **Debounce Algoritması:** İşletim sisteminin dosya kaydetme sırasında oluşturduğu mükerrer sinyalleri engellemek için 500ms'lik bir gecikme mekanizması kuruldu.

**3. Proje Yapısı:**
Proje, QLineSec standartlarına uygun olarak Template yapısına taşınmış; `src`, `docs` ve `specs` klasörleri ile modüler hale getirilmiştir.