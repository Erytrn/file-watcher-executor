# 🏗️ Sistem Mimarisi ve Çalışma Mantığı

Bu doküman, yazılımın arka planda nasıl çalıştığını, veri akışını ve bileşenler arası ilişkiyi açıklar.

---

## 1. Görsel Veri Akış Şeması (Data Flow Diagram)
Aşağıdaki diyagram, kullanıcının `CTRL+S` tuşuna basmasından terminalde sonucun çıkmasına kadar geçen süreci gösterir.

```mermaid
graph TD
    User[👤 Geliştirici] -->|Dosya Kaydeder| OS(📂 İşletim Sistemi Sinyali)
    OS -->|fsnotify| Watcher{👀 Go Watcher}
    
    Watcher -->|Sinyal Yakalandı| Filter{🔍 Filtreleme Katmanı}
    
    Filter -- .txt / .log --> Ignore[❌ Yoksay]
    Filter -- .py / .js --> Timer{⏳ Debounce (500ms)}
    
    Timer -- Yeni Sinyal Geldi --> Reset[🔄 Sayacı Sıfırla]
    Timer -- Süre Doldu --> Exec[🚀 Executor Modülü]
    
    Exec -->|Bağlam Değişimi| Workspace[📂 Workspace Dizini]
    Workspace -->|Python Çalıştır| Script[🐍 main.py]
    
    Script -->|Output| Terminal[🖥️ Terminal Çıktısı]
    ## 2. Mimari Katmanlar

Sistem 3 temel katman üzerine kurulmuştur:

### A. Dinleme Katmanı (Listener Layer)
* **Teknoloji:** `fsnotify` (Kernel hook)
* **Görevi:** İşletim sisteminin dosya sistemi olaylarını (Create, Write, Remove) anlık olarak dinler. Bu katman "non-blocking" (bloklamayan) yapıdadır, yani ana programı dondurmaz.

### B. Mantık Katmanı (Logic Layer)
Gelen ham sinyallerin işlendiği yerdir:
1.  **Filtreleme:** Gelen dosya uzantısı `.py` veya `.js` mi? (Değilse işlem iptal edilir).
2.  **Debounce:** Kullanıcı yazmaya devam ediyor mu? (Son olaydan sonra 500ms beklendi mi?).

### C. Yürütme Katmanı (Execution Layer)
* **Görevi:** Onaylanan işlemi gerçekleştirir.
* **İzolasyon:** Komut çalıştırılırken `cmd.Dir` parametresi ile çalışma dizini `/src` yerine `/workspace` olarak değiştirilir. Bu sayede Python scripti kendi yerel dosyalarına (JSON vb.) erişebilir.

---

## 3. Klasör İzolasyonu Prensibi
Sistemin "Sonsuz Döngüye" (Infinite Loop) girmemesi için **Gözlemci** ve **Hedef** birbirinden fiziksel olarak ayrılmıştır.

* **`src/` (Subject):** Gözlem yapan özne.
* **`workspace/` (Object):** Gözlemlenen nesne.

Bu yapı sayesinde, Go programı log oluşturduğunda veya kendi config dosyasını güncellediğinde, kendini yanlışlıkla tekrar tetiklemez.