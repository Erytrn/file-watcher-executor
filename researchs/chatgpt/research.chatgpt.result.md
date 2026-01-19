# 🤖 ChatGPT Kullanım Sonuçları ve Katkı Değerlendirmesi

Bu doküman, "File Watcher & Executor" projesi süresince ChatGPT kullanımının proje çıktıları üzerindeki etkisini ve elde edilen sonuçları açıklamak amacıyla hazırlanmıştır.

---

## 🎯 Genel Değerlendirme

ChatGPT ile yapılan etkileşimler sonucunda proje:

- Teknik olarak daha **net ve tutarlı** hale getirilmiştir
- Mimari yapı **daha anlaşılır** biçimde belgelenmiştir
- Dokümantasyon kalitesi **akademik seviyeye** çıkarılmıştır

ChatGPT, proje geliştirme sürecinde **rehber ve doğrulayıcı** rol üstlenmiştir.

---

## 🧩 Elde Edilen Somut Kazanımlar

### 1️⃣ Kod ve Sistem Doğrulaması
- Mevcut **Go (main.go)** ve **Python (main.py)** kodlarının;
  - Dosya izleme
  - Filtreleme (`.py`, `.js`)
  - Debounce mekanizması
  - Otomatik çalıştırma

  isterlerini karşıladığı doğrulanmıştır.

- Kod üzerinde **zorunlu bir değişiklik ihtiyacı olmadığı** teyit edilmiştir.

---

### 2️⃣ Çoklu Dosya Desteği Bilinci
- Sistem tasarımının yalnızca `.py` değil, `.js` gibi farklı dosya türlerine de uygun olduğu netleştirilmiştir.
- Genişletilebilirlik açısından mimarinin doğru kurulduğu sonucuna varılmıştır.

---

### 3️⃣ Test ve Çalışma Senaryoları
- Projenin çalıştığının nasıl test edileceği netleştirilmiştir:
  - `workspace/` klasöründe dosya değişikliği
  - Kaydetme (`CTRL + S`)
  - Terminal çıktısının gözlemlenmesi

- Bu senaryolar README ve teknik dokümanlara eklenmiştir.

---

### 4️⃣ Dokümantasyon Kalitesinin Artırılması
ChatGPT katkısı ile:
- `README.md`
- `teknik-arastirma-notlari.md`
- `mimari-yapi.md`
- `research.chatgpt.prompt.md`

dosyaları **daha sistematik**, **daha anlaşılır** ve **teslim edilebilir** hale getirilmiştir.

---

### 5️⃣ Mimari Anlatım ve Diyagram Netliği
- Sistem mimarisi:
  - Katmanlı yapı (Listener / Logic / Execution)
  - Klasör izolasyonu
  - Veri akışı

  açık ve teknik bir dille ifade edilmiştir.

- Mermaid diyagramlarında GitHub uyumluluğu sağlanmıştır.

---

## ⚖️ Akademik ve Etik Sonuç

Bu proje kapsamında:
- ChatGPT **kod yazan bir kaynak** olarak kullanılmamıştır
- Öğrenci tarafından yazılan kodlar **korunmuştur**
- Yapay zekâ, yalnızca:
  - Rehberlik
  - Analiz
  - Dokümantasyon desteği

  amacıyla kullanılmıştır.

Bu kullanım, **akademik etik ilkelere uygundur**.

---

## 📌 Genel Sonuç

ChatGPT kullanımı sayesinde:
- Proje daha profesyonel bir yapıya kavuşmuştur
- Teknik anlatım güçlenmiştir
- Teslim kalitesi artmıştır

Ancak proje sahipliği ve teknik uygulama **tamamen öğrenciye aittir**.

---

## 👤 Hazırlayan
**Eray**  
Bilişim Güvenliği – 2. Sınıf  
İstinye Üniversitesi  

**Tarih:** Ocak 2026
