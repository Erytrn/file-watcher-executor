# Sistem Mimarisi ve Çalışma Mantığı

Bu döküman, yazılımın arka planda nasıl çalıştığını açıklar.

## Akış Diyagramı (Mantıksal)

1. **Başlatma:** Program `main.go` üzerinden başlar.
2. **Watcher Kurulumu:** `fsnotify` kütüphanesi işletim sisteminin kernel seviyesindeki dosya olaylarına kanca (hook) atar.
3. **Event Loop (Döngü):** - Bir `Go Routine` (hafif iplikçik) sürekli olarak olay kanalını dinler.
   - Ana program bloklanmaz (Non-blocking).
4. **Filtreleme Katmanı:**
   - Gelen olay `.js` veya `.py` uzantılı mı?
   - EVET -> Devam et.
   - HAYIR -> Görmezden gel.
5. **Debounce Katmanı:**
   - Zamanlayıcı kontrol edilir. Gerekirse eski zamanlayıcı iptal edilir, yenisi başlatılır.
6. **Execution (Çalıştırma):**
   - Süre dolduğunda `exec` paketi ile terminal komutu çalıştırılır.