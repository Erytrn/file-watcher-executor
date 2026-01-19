# 👁️ File Watcher & Executor (Otomasyon Aracı)

**Versiyon:** 1.0.0  
**Yazar:** Eray  
**Durum:** ✅ Tamamlandı

Bu proje, belirli bir dizindeki dosya değişikliklerini (**Create / Modify**) gerçek zamanlı olarak izleyen ve bu değişikliklere bağlı olarak tanımlanan komutları otomatik tetikleyen, **CI/CD mantığını simüle eden** bir otomasyon aracıdır.

Amaç; modern yazılım geliştirme süreçlerinde kullanılan dosya izleme, otomatik test ve yeniden çalıştırma mekanizmalarının temelini performanslı bir yapı ile göstermektir.

---

## 🚀 Özellikler

Bu araç, modern otomasyon sistemlerinde bulunan aşağıdaki yeteneklere sahiptir:

* **⚡ Gerçek Zamanlı İzleme (Event Loop):**  
  Go dili ve `fsnotify` kütüphanesi kullanılarak dosya sistemi olayları anlık olarak dinlenir.

* **🔍 Akıllı Filtreleme:**  
  Sadece hedeflenen uzantılardaki (`.py`, `.js`) dosya değişikliklerine tepki verir.  
  Diğer dosya türleri bilinçli olarak yok sayılır.

* **⏳ Debounce (Geciktirme Mekanizması):**  
  Çok kısa sürede gerçekleşen ardışık dosya değişiklikleri gruplanır (500ms) ve gereksiz tekrar çalıştırmalar engellenir.

* **🔄 Otomatik Çalıştırma (Auto-Execute):**  
  Dosya değişikliği algılandığında hedef Python veya JavaScript dosyaları otomatik olarak çalıştırılır.

* **🛠️ Self-Check Mekanizması:**  
  Sistem, çalışmadan önce:
  - `project_info.json` yapılandırma dosyasının varlığını
  - Çalışılan işletim sistemini (Windows / Linux)
  - Python ve JavaScript dosyalarının erişilebilirliğini kontrol eder.

* **🧪 Servis Durumu Kontrolü (Simülasyon):**  
  Linux ortamında gerçek servis kontrolü (`systemctl`) yapılabilir.  
  Windows ortamında platform uyumlu olacak şekilde simülasyon uygulanır.

---

## 📂 Proje Yapısı

```text
FILE-WATCHER-EXECUTOR/
├── src/
│   ├── main.go            # (GÖZCÜ) Dosya değişikliklerini izleyen Go uygulaması
│   ├── go.mod
│   └── go.sum
│
├── workspace/             # (HEDEF) İzlenen proje dizini
│   ├── main.py            # Otomatik çalıştırılan Python uygulaması
│   ├── test.js            # Test amaçlı JavaScript dosyası
│   ├── project_info.json  # Proje meta verileri
│   └── index.html         # Sistem durumunu gösteren statik kontrol arayüzü
│
├── specs/                 # Proje teknik gereksinimleri ve notları
├── README.md
└── Sunum.pdf
