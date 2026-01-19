# Sistem Mimarisi ve Çalışma Mantığı

Bu doküman, yazılımın arka planda nasıl çalıştığını, veri akışını ve bileşenler arası ilişkiyi açıklar.

---

## 1. Görsel Veri Akış Şeması (Data Flow Diagram)

Aşağıdaki diyagram, kullanıcının bir dosyayı kaydetmesinden (CTRL+S) terminalde sonucun görüntülenmesine kadar geçen süreci göstermektedir.

```mermaid
graph TD
    User[Gelistirici] -->|Dosya Kaydeder| OS[Isletim Sistemi Olayi]
    OS -->|fsnotify| Watcher[Go Watcher]

    Watcher -->|Olay Yakalandi| Filter[Filtreleme Katmani]

    Filter -->|.txt / .log| Ignore[Yoksay]
    Filter -->|.py / .js| Timer[Debounce 500ms]

    Timer -->|Yeni Sinyal| Reset[Sayaci Sifirla]
    Timer -->|Sure Doldu| Exec[Executor Modulu]

    Exec -->|Calisma Dizini Degisir| Workspace[workspace dizini]
    Workspace -->|Script Calistir| Script[main.py veya test.js]

    Script -->|Cikti| Terminal[Terminal Ciktisi]
