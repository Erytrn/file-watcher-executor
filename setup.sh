#!/bin/bash
# Renkli çıktılar
GREEN='\033[0;32m'
CYAN='\033[0;36m'
NC='\033[0m'

echo -e "${CYAN}🚀 File Watcher & Executor Kurulumu Başlıyor...${NC}"

echo -e "${GREEN}📦 Bağımlılıklar yükleniyor...${NC}"
go mod tidy

echo -e "${GREEN}🔨 Proje derleniyor...${NC}"
if go build -o file-watcher; then
    echo -e "${GREEN}✅ Build Başarılı!${NC}"
    echo -e "Çalıştırmak için: ./file-watcher"
else
    echo "❌ Build Hatası!"
    exit 1
fi
