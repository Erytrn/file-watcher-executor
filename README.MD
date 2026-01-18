# 👁️ File Watcher & Executor (Otomasyon Aracı)

**Versiyon:** 1.0.0  
**Yazar:** Eray  
**Durum:** ✅ Tamamlandı

Bu proje, belirli bir dizindeki dosya değişikliklerini (Create/Write) gerçek zamanlı izleyen ve buna bağlı olarak tanımlanan test veya build komutlarını otomatik çalıştıran bir **CI/CD Simülasyon aracıdır.**

---

## 🚀 Özellikler

Bu araç, modern otomasyon sistemlerinde bulunan şu yeteneklere sahiptir:

* **⚡ Gerçek Zamanlı İzleme (Event Loop):** Dosya sistemi olaylarını (File System Events) anlık olarak dinler.
* **🔍 Akıllı Filtreleme:** Sadece hedeflenen uzantıdaki (`.py`, `.js`) dosya değişimlerine tepki verir; gereksiz dosyaları (`.txt`, `.log`) görmezden gelir.
* **⏳ Debounce (Geciktirme):** Hızlı ve peş peşe yapılan kaydetme işlemlerini gruplayarak sistemin gereksiz yere defalarca çalışmasını engeller (500ms).
* **🔄 Otomatik Çalıştırma (Auto-Execute):** Değişiklik algılandığında hedef Python projesini otomatik derler ve çalıştırır.
* **🛠️ Self-Check Mekanizması:** Sistem, çalışmadan önce kendi yapılandırma dosyalarını (`project_info.json`) ve çalışma ortamını (Windows/Linux) kontrol eder.

---

## 📂 Proje Yapısı

```text
FILE-WATCHER-EXECUTOR/
├── src/
│   └── main.go           # (GÖZCÜ) Dosyaları izleyen Go programı
│
├── workspace/            # (HEDEF) İzlenen proje klasörü
│   ├── main.py           # Otomatik çalıştırılan Python kodu
│   ├── project_info.json # Proje meta verileri
│   └── index.html        # Arayüz tasarımı
│
└── specs/                # Proje teknik gereksinimleri