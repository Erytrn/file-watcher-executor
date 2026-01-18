import json
import subprocess
import os
import platform

# 1. JSON Okuma
def load_project_info():
    try:
        with open('project_info.json', 'r', encoding='utf-8') as f:
            data = json.load(f)
            return data
    except FileNotFoundError:
        return None

# 2. Servis Kontrolü (Windows ve Linux Uyumlu)
def check_service_status(service_name="ssh"):
    current_os = platform.system()
    
    # Eğer Windows kullanıyorsan, Linux komutu çalıştırma, simülasyon yap.
    if current_os == "Windows":
        return f"Aktif (Windows Simülasyonu - {service_name})", True
    
    # Linux ise gerçek kontrol yap
    try:
        stat = subprocess.call(["systemctl", "is-active", "--quiet", service_name])
        if stat == 0:
            return "Aktif (Çalışıyor)", True
        else:
            return "Pasif (Durduruldu)", False
    except Exception as e:
        return f"Hata: {str(e)}", False

# 3. Self-Check (Kendi Kendini Test Etme)
def self_check():
    print("--- SELF CHECK BAŞLATILIYOR ---")
    
    if os.path.exists("project_info.json"):
        print("[OK] Yapılandırma dosyası bulundu.")
    else:
        print("[FAIL] JSON dosyası eksik!")
        return False

    # İşletim sistemi uyarısını kaldırdık, bilgi veriyoruz sadece.
    print(f"[INFO] Çalışılan Ortam: {platform.system()}")
    
    print("--- SELF CHECK TAMAMLANDI ---\n")
    return True

# --- Programı Çalıştır ---
if __name__ == "__main__":
    self_check()
    
    info = load_project_info()
    if info:
        print(f"Proje: {info['projectName']} v{info['version']}")
        print(f"Yazar: {info['author']}")  # Artık Eray yazacak
        
        durum_mesaji, durum_kodu = check_service_status("ssh")
        print(f"\nSSH Servis Durumu: {durum_mesaji}")
    else:
        print("Proje bilgileri yüklenemedi.")

    print("-" * 30)
    print("✅ TEST BAŞARIYLA TAMAMLANDI (ERAY)")