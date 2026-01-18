# File Watcher & Executor (Dosya İzleyici ve Çalıştırıcı)

![Go Version](https://img.shields.io/badge/Go-1.25-blue?style=flat&logo=go)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey)
![Status](https://img.shields.io/badge/Status-Active-success)

Bu proje, belirlenen bir dizindeki dosya değişikliklerini (oluşturma, silme, değiştirme) gerçek zamanlı izleyen ve buna bağlı olarak otomatik komut çalıştıran bir otomasyon aracıdır. **VirtualBox gerektirmez**, doğrudan işletim sistemi çekirdeği (Kernel) üzerinden olayları dinler.

---

## 📂 Proje Yapısı

Proje, profesyonel geliştirme standartlarına uygun olarak modüler bir yapıda tasarlanmıştır:

```text
proje-root/
├── 📂 src/             # Kaynak kodlar (Source Code)
│   ├── main.go         # Ana uygulama dosyası
│   ├── go.mod          # Modül tanımları
│   └── go.sum          # Bağımlılık sağlama dosyası
├── 📂 docs/            # Teknik dökümantasyon ve mimari şema
├── 📂 specs/           # Proje gereksinimleri ve analizler
├── 📂 researchs/       # Teknik araştırma ve dil seçim notları
└── 📄 README.md        # Proje tanıtım dosyası