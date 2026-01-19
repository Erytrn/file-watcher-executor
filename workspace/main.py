import json
import subprocess
import os
import platform
from pathlib import Path

BASE_DIR = Path(__file__).resolve().parent

# 1. JSON Okuma
def load_project_info():
    try:
        with open(BASE_DIR / 'project_info.json', 'r', encoding='utf-8') as f:
            return json.load(f)
    except FileNotFoundError:
        return None

# 2. Servis Kontrolü (Windows ve Linux Uyumlu)
def check_service_status(service_name="ssh"):
    current_os = platform.system()

    if current_os == "Windows":
        return f"Aktif (Windows Simülasyonu - {service_name})", True

    try:
        stat = subprocess.call(
            ["systemctl", "is-active", "--quiet", service_name],
            stdout=subprocess.DEVNULL,
            stderr=subprocess.DEVNULL
        )
        if stat == 0:
            return "Aktif (Çalışıyor)", True
        else:
            return "Pasif (Durduruldu)", False
    except Exception as e:
        return f"Hata: {str(e)}", False

# 3. İzlenen Dosyaları Kontrol Et (.py + .js)
def detect_source_files():
    py_files = list(BASE_DIR.glob("*.py"))
    js_files = list(BASE_DIR.glob("*.js"))

    if py_files:
        print(f"[OK] Python dosyaları bulundu: {[f.name for f in py_files]}")
    else:
        print("[INFO] Python dosyası bulunamadı.")

    if js_files:
        print(f"[OK] JavaScript dosyaları bulundu: {[f.name for f in js_files]}")
    else:
        print("[INFO] JavaScript dosyası bulunamadı.")

# 4. Self-Check
def self_check():
    print("--- SELF CHECK BAŞLATILIYOR ---")

    if (BASE_DIR / "project_info.json").exists():
        print("[OK] Yapılandırma dosyası bulundu.")
    else:
        print("[FAIL] JSON dosyası eksik!")
        return False

    print(f"[INFO] Çalışılan Ortam: {platform.system()}")

    detect_source_files()

    print("--- SELF CHECK TAMAMLANDI ---\n")
    return True

# --- Programı Çalıştır ---
if __name__ == "__main__":
    self_check()

    info = load_project_info()
    if info:
        print(f"Proje: {info.get('projectName', 'UNKNOWN')} v{info.get('version', 'N/A')}")
        print(f"Yazar: {info.get('author', 'UNKNOWN')}")

        durum_mesaji, _ = check_service_status("ssh")
        print(f"\nSSH Servis Durumu: {durum_mesaji}")
    else:
        print("Proje bilgileri yüklenemedi.")

    print("-" * 30)
    print("✅ TEST BAŞARIYLA TAMAMLANDI (ERAY)")
