# 📡 Copilot Araştırma İsteği: File Watcher & Executor Projesi

## 🎯 Amaç
Bu araştırma, "File Watcher & Executor" adlı sistem yazılımı projesinin mimari, güvenlik, modern C++ kod kalitesi ve teknik sunum yönlerinden en iyi uygulamalarla nasıl geliştirilebileceğini belirlemek amacıyla yapılmıştır.

## 🧩 Bağlam
Proje, dosya sistemindeki değişiklikleri izleyerek belirli regex kurallarına göre otomatik shell komutları çalıştıran bir araçtır. Kullanıcı, bu sistemin yüksek dosya trafiğinde stabil çalışmasını, güvenli olmasını ve teknik jüriye profesyonel bir sunumla aktarılmasını hedeflemektedir.

## 🧠 Sorulan Teknik Sorular
1. **Mimari & Performans:**  
   - High I/O altında concurrency nasıl yönetilmeli?  
   - Cross-platform uyumluluk ve debounce mekanizması nasıl “bulletproof” hale getirilir?

2. **Güvenlik (Security Audit):**  
   - Command injection, TOCTOU ve path traversal gibi zafiyetlere karşı nasıl korunulur?  
   - Config dosyaları nasıl sanitize edilir?

3. **Modern C++ ve Kod Kalitesi:**  
   - C++17/20 ile memory-safe ve okunabilir kod nasıl yazılır?  
   - Genişletilebilirlik için hangi design pattern’ler önerilir?

4. **Sunum ve Sunuş:**  
   - Projenin sistem yazılımı olduğunu kanıtlamak için hangi teknik zorluklar ve çözümler vurgulanmalı?

## 🛠️ Kullanıcının Paylaştığı Kod
Kullanıcı, Go dilinde fsnotify tabanlı bir prototip geliştirmiştir. Kod:
- `.py` ve `.js` dosyalarını izler,
- debounce ile tekrarlayan olayları engeller,
- `exec.Command` ile shell’siz komut çalıştırır,
- `workspace` dizininde çalışır.

## 🧪 Copilot’un Sağladığı Çıktılar
- Path-bazlı debounce önerisi (`map[string]*time.Timer`)
- Worker pool ile concurrency yönetimi
- Güvenli komut yürütme için shell-free exec stratejisi
- TOCTOU ve symlink saldırılarına karşı canonicalization
- C++ tarafında Strategy, Observer, Command pattern önerileri
- README.md için eksiksiz bir şablon
- project_info.json formatı
- Teknik jüriye sunumda vurgulanacak zorluklar ve çözümler

## 📁 Dosya Yapısı Önerisi
- `cmd/` → main.go  
- `internal/watcher/` → fsnotify wrapper  
- `internal/executor/` → safe exec  
- `configs/config.json` → uzantı ve komut tanımı  
- `README.md` → proje tanımı, mimari, güvenlik, test, lisans  
- `project_info.json` → meta veri

## 📌 Sonuç
Bu araştırma, sistem yazılımı seviyesinde bir dosya izleme ve otomasyon aracının mimari, güvenlik ve sunum yönlerinden nasıl profesyonelleştirileceğini kapsamlı şekilde ortaya koymuştur.
