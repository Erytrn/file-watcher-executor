# 🏗️ Sistem Mimarisi ve Çalışma Mantığı

Bu doküman, yazılımın arka planda nasıl çalıştığını, veri akışını ve bileşenler arası ilişkiyi açıklar.

---

## 1. Görsel Veri Akış Şeması (Data Flow Diagram)

Aşağıdaki diyagram, kullanıcının bir dosyayı kaydetmesinden (`CTRL+S`) terminalde sonucun görüntülenmesine kadar geçen süreci göstermektedir.

```mermaid
graph TD
    User[👤 Geliştirici] -->|Dosya Kaydeder| OS(📂 İşletim Sistemi Olayı)
    OS -->|fsnotify| Watcher{👀 Go Watcher}

    Watcher -->|Olay Yakalandı| Filter{🔍 Filtreleme Katmanı}

    Filter -- .txt / .log --> Ignore[❌ Yoksay]
    Filter -- .py / .js --> Timer{⏳ Debounce (500ms)}

    Timer -- Yeni Sinyal --> Reset[🔄 Sayaç Sıfırla]
    Timer -- Süre Doldu --> Exec[🚀 Executor Modülü]

    Exec -->|Çalışma Dizini Değişimi| Workspace[📂 workspace/]
    Workspace -->|Script Çalıştır| Script[🐍 main.py / 🟨 test.js]

    Script -->|Output| Terminal[🖥️ Terminal Çıktısı]
