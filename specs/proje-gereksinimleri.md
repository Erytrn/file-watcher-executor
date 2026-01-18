# 📑 Teknik Şartname ve Sistem Tasarım Belgesi

**Proje Adı:** File Watcher & Executor  
**Geliştirici:** Eray  
**Versiyon:** 1.0.0  
**Tarih:** 18.01.2026

---

## 1. Proje Özeti
Bu proje, yazılım geliştirme süreçlerini hızlandırmak amacıyla tasarlanmış bir **CI/CD Simülasyon Aracıdır**. Belirli bir çalışma alanındaki (`workspace`) dosya değişikliklerini gerçek zamanlı izler ve tanımlanan kurallara göre test/derleme süreçlerini otomatik tetikler.

## 2. Sistem Mimarisi
Proje, **Gözlemci (Watcher)** ve **Yürütücü (Executor)** olmak üzere iki ana modülden oluşur:

### 2.1. Klasör Yapısı (Folder Structure)
* **`/src` (Core):** Sistemin beyni olan Go tabanlı izleme mekanizması burada çalışır.
* **`/workspace` (Target):** İzlenen hedef proje dosyaları (Python, HTML, JSON) burada bulunur.
* **`/specs` (Documentation):** Proje teknik gereksinimleri ve tasarım notları.

---

## 3. Fonksiyonel Gereksinimler

### 3.1. Olay Dinleme (Event Loop)
* Sistem, işletim sistemi seviyesindeki dosya olaylarını (`File System Events`) dinlemelidir.
* **Kullanılan Teknoloji:** Go `fsnotify` kütüphanesi.
* **Hedef:** CPU kullanımını minimumda tutarak anlık tepki vermek.

### 3.2. Akıllı Filtreleme (Smart Filtering)
Sistem her dosyaya tepki vermemelidir. Sadece geliştirme ile ilgili dosyalara odaklanmalıdır:
* **Kabul Edilenler:** `.py` (Python), `.js` (JavaScript)
* **Reddedilenler:** `.txt`, `.log`, `.tmp` ve klasör değişimleri.

### 3.3. Debounce Mekanizması (Zaman Yönetimi)
Kullanıcının hızlı yazma veya peş peşe kaydetme (`CTRL+S`) işlemlerinde sistemin çökmemesi için:
* **Süre:** 500ms (milisaniye).
* **Mantık:** Son olaydan sonra 500ms boyunca yeni bir olay gelmezse komut çalıştırılır.

### 3.4. Çapraz Platform Uyumluluğu (Cross-Platform)
Sistem hem **Windows** hem de **Linux** ortamlarında hatasız çalışmalıdır:
* **Windows:** `python` komutu kullanılır, Linux komutları (`systemctl`) simüle edilir.
* **Linux:** `python3` ve yerel `systemctl` komutları kullanılır.

---

## 4. Veri Yapısı ve Konfigürasyon
Proje meta verileri standart bir JSON formatında tutulmalıdır.

**Dosya:** `project_info.json`
**Format:**
```json
{
  "projectName": "String",
  "author": "String",
  "version": "String",
  "features": ["Array"],
  "requirements": { "os": "String" }
}
## 5. Kullanılan Teknolojiler ve Araçlar

| Teknoloji      | Amaç                                                 |
| :------------- | :--------------------------------------------------- |
| **Go (Golang)**| Yüksek performanslı dosya izleme ve arka plan servisi.|
| **Python** | Hedef proje dili ve otomasyon senaryoları.           |
| **fsnotify** | Dosya sistemi olaylarını yakalayan kütüphane.        |
| **JSON** | Veri taşıma ve konfigürasyon standardı.              |