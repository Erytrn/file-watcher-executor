# 🔬 Teknik Araştırma ve Geliştirme Notları

**Konu:** Dosya Sistemi İzleme (File Watching) ve Otomasyon  
**Tarih:** 18.01.2026  
**Araştıran:** Eray

---

## 1. Mevcut Çözümlerin İncelenmesi

Proje öncesinde piyasadaki hazır dosya izleme ve otomasyon araçları incelenmiştir:

* **Watchexec (Rust):**  
  Yüksek performanslıdır ancak yapılandırma ve özelleştirme süreci karmaşıktır.

* **Nodemon (JavaScript):**  
  Node.js bağımlılığı gerektirir ve uzun süreli çalışmalarda sistem kaynaklarını (RAM) nispeten fazla tüketir.

**Karar:**  
Bu nedenlerle, daha hafif (lightweight), bağımsız ve performans odaklı bir çözüm geliştirmek amacıyla **Go (Golang)** dili tercih edilmiştir.

---

## 2. Dil ve Kütüphane Seçimi: Go vs Rust

Proje için iki düşük seviyeli sistem dili değerlendirilmiştir:

| Özellik | Go (Golang) | Rust |
|------|-------------|------|
| **Dosya İzleme Kütüphanesi** | `fsnotify/fsnotify` | `notify-rs` |
| **Öğrenme Eğrisi** | Düşük | Yüksek (Ownership kuralları) |
| **Eşzamanlılık** | Goroutines (Basit) | Async/Await (Daha karmaşık) |
| **Geliştirme Hızı** | Yüksek | Orta |
| **Sonuç** | ✅ **SEÇİLDİ** | ❌ Elendi |

**Sonuç:**  
Projenin süresi ve kapsamı göz önünde bulundurularak **Go**, daha hızlı prototipleme ve sade mimari avantajı nedeniyle tercih edilmiştir.

---

## 3. Karşılaşılan Teknik Zorluklar ve Çözümler

### 🔴 Sorun 1: Sonsuz Döngü (Infinite Loop)

**Durum:**  
Watcher (izleyici) ve izlenen dosyalar aynı klasörde bulunduğunda, watcher kendi oluşturduğu log veya çıktı dosyalarını tekrar tetikliyordu.

**Çözüm:**  
- İzleyici kodları `src/`
- İzlenen hedef dosyalar `workspace/`

olacak şekilde klasörler fiziksel olarak ayrıldı.

---

### 🔴 Sorun 2: Hayalet Tetiklenmeler (Debounce Problemi)

**Durum:**  
Modern kod editörleri (ör. VS Code), tek bir kaydetme işleminde birden fazla `Write` olayı üretebilmektedir.  
Bu durum, komutların art arda ve gereksiz şekilde çalışmasına neden oldu.

**Çözüm:**  
- 500ms gecikmeli bir **Timer (Debounce)** mekanizması uygulandı.
- Yeni bir olay geldiğinde sayaç sıfırlanarak sadece son olay işlendi.

---

### 🔴 Sorun 3: Windows ve Linux Ortam Farklılıkları

**Durum:**  
Linux ortamında çalışan `python3` ve `systemctl` komutları, Windows ortamında (`exit status 9009`) hata üretmiştir.

**Çözüm:**
* **Go tarafında:**  
  - `exec.Command("python")` kullanıldı  
  - `cmd.Dir` ile doğru çalışma dizini ayarlandı

* **Python tarafında:**  
  - `platform.system()` ile işletim sistemi tespit edildi  
  - Windows ortamında servis kontrolleri için simülasyon modu uygulandı

---

## 4. Güvenlik ve Kararlılık Değerlendirmesi

* Komut çalıştırma işlemleri sabit ve kontrollüdür.
* Kullanıcıdan doğrudan shell girdisi alınmamaktadır.
* JSON dosyaları okunurken hata kontrolü yapılmaktadır.
* Platform bağımlı komutlar koşullu olarak çalıştırılmaktadır.

Bu yaklaşım, **command injection** ve **yetkisiz komut çalıştırma** risklerini minimize etmektedir.

---

## 5. Geliştirme Fikirleri (Future Work)

Bu projenin ilerleyen sürümlerinde aşağıdaki geliştirmeler yapılabilir:

* Alt klasörlerin (recursive) otomatik izlenmesi
* Regex veya glob tabanlı gelişmiş dosya filtreleme
* Harici yapılandırma dosyası (JSON / YAML) ile dinamik komut tanımlama
* Aynı anda birden fazla komut çalıştırabilme
* Linux servisleri için gerçek zamanlı durum izleme

Bu geliştirmeler, mevcut mimari bozulmadan eklenebilir yapıdadır.

---

## 6. Kaynakça ve Referanslar

1. **fsnotify GitHub:** https://github.com/fsnotify/fsnotify  
2. **Go `os/exec` Dokümantasyonu:** https://pkg.go.dev/os/exec  
3. **Python `subprocess` Modülü:** https://docs.python.org/3/library/subprocess.html  
