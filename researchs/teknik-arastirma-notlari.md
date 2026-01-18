# Teknik Araştırma ve Karar Süreci

## 1. Dil Seçimi: Neden Go (Golang)?
Bu proje için Rust ve Go dilleri arasında bir karşılaştırma yapılmıştır.

- **Rust:** Bellek güvenliği konusunda çok iyi olsa da, öğrenme eğrisi (Learning Curve) diktir ve prototip geliştirme süreci uzundur.
- **Go:** "Concurrency" (Eşzamanlılık) yapısı (Goroutines) sayesinde Event Loop mantığını kurmak çok daha basittir. Ayrıca `fsnotify` kütüphanesi Go ekosisteminde standartlaşmıştır.
-> **Karar:** Projenin süresi ve performans gereksinimleri gözetilerek **Go** dili seçilmiştir.

## 2. Debounce Algoritması Araştırması
Dosya kaydetme işlemi sırasında işletim sistemi (OS) bazen milisaniyeler içinde birden fazla "Write" sinyali gönderebilmektedir.
- **Sorun:** Her sinyalde komut çalıştırılırsa sistem yorulur.
- **Çözüm:** "Debounce" tekniği. İlk sinyalden sonra bir Timer (Zamanlayıcı) başlatılır. Belirlenen süre (örn. 500ms) içinde yeni sinyal gelirse sayaç sıfırlanır.