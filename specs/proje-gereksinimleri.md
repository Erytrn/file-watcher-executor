# Proje Gereksinimleri ve Teknik Şartname

## Proje Amacı
Belirli bir dizindeki dosya değişikliklerini (oluşturma, silme, değiştirme) gerçek zamanlı izleyen ve buna bağlı olarak tanımlanan komutları çalıştıran bir otomasyon aracı geliştirmek.

## Beklenen Özellikler
1. **Event Loop:** Dosya sistemi olaylarını (File System Events) dinleyen performanslı bir yapı kurulmalı.
2. **Filtreleme:** Sadece belirli uzantıdaki (.js, .py) değişimler izlenmeli.
3. **Komut Çalıştırma:** Değişiklik algılandığında otomatik tetikleme yapılmalı.
4. **Debounce:** Çok sık gelen değişiklikleri gruplayarak gereksiz komut tekrarı önlenmeli.

## Teknik Kısıtlamalar
- **Dil:** Go (Golang) seçilmiştir.
- **Kütüphane:** fsnotify kullanılacaktır.
- **Platform:** VirtualBox kullanılmadan, Host OS üzerinde çalıştırılacaktır.