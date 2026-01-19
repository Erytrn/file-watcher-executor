# 📊 Araştırma Özeti ve Karar Matrisi (Infographic)

Bu şema, **Go (fsnotify)** ve **Python (watchdog)** kütüphaneleri arasında yapılan kıyaslamayı ve nihai mimari kararını görselleştirir.

```mermaid
graph TD
    A[🔍 Araştırma Başlangıcı] --> B{Hangi Dil Seçilmeli?}
    
    subgraph Python_Watchdog
    C[Python Watchdog]
    C --> D[Kolay Yazım]
    C --> E[Yüksek CPU Kullanımı ⚠️]
    C --> F[Polling Gecikmesi]
    end
    
    subgraph Go_Fsnotify
    G[Go fsnotify]
    G --> H[Kernel Level Events]
    G --> I[Düşük Kaynak Tüketimi ✅]
    G --> J[Real-Time Tepki ⚡]
    end
    
    B -- Analiz --> Python_Watchdog
    B -- Analiz --> Go_Fsnotify
    
    Go_Fsnotify --> K[🏆 KAZANAN: Go]
    
    K --> L[MİMARİ TASARIM]
    L --> M[Watcher: Go]
    L --> N[Debounce: 500ms]
    L --> O[Executor: Python]
    
    style K fill:#bbf,stroke:#333,stroke-width:2px
    style E fill:#f9f,stroke:#333
    style I fill:#9f9,stroke:#333