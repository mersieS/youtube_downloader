# ytmp3

Free YouTube & Video Downloader CLI / Ücretsiz YouTube & Video İndirici CLI

YouTube ve 1000+ siteden MP4/MP3 indirmenizi sağlayan interaktif terminal uygulaması.

An interactive terminal app to download MP4/MP3 from YouTube and 1000+ sites.

---

## Özellikler / Features

- YouTube + diğer siteler (yt-dlp destekli 1000+ site) / YouTube + other sites (1000+ yt-dlp supported sites)
- MP4 video indirme (4K, 2K, 1080p, 720p, 480p, 360p) / MP4 video download with quality selection
- MP3 ses çıkarma (en yüksek kalite) / MP3 audio extraction (highest quality)
- Ok tuşlarıyla interaktif menü / Interactive arrow-key navigation
- Özel klasör seçimi / Custom output directory

---

## Gereksinimler / Requirements

Uygulamayı kullanabilmek için aşağıdaki araçlar yüklü olmalıdır:

The following tools must be installed before using ytmp3:

| Araç / Tool | Kurulum / Install |
|---|---|
| [yt-dlp](https://github.com/yt-dlp/yt-dlp) | `brew install yt-dlp` / `winget install yt-dlp` |
| [ffmpeg](https://ffmpeg.org) | `brew install ffmpeg` / `winget install ffmpeg` |

---

## Kurulum / Installation

### Homebrew (macOS)

```bash
brew tap sbuker/ytmp3
brew install ytmp3
```

### Doğrudan İndirme / Direct Download

#### macOS (Apple Silicon - M1/M2/M3/M4)

```bash
curl -L https://github.com/mersieS/youtube_downloader/releases/latest/download/ytmp3-darwin-arm64.tar.gz | tar xz
sudo mv ytmp3-darwin-arm64 /usr/local/bin/ytmp3
```

#### macOS (Intel)

```bash
curl -L https://github.com/mersieS/youtube_downloader/releases/latest/download/ytmp3-darwin-amd64.tar.gz | tar xz
sudo mv ytmp3-darwin-amd64 /usr/local/bin/ytmp3
```

#### Windows

1. [Releases](https://github.com/mersieS/youtube_downloader/releases/latest) sayfasından `ytmp3-windows-amd64.zip` dosyasını indirin / Download `ytmp3-windows-amd64.zip` from Releases
2. ZIP dosyasını çıkartın / Extract the ZIP file
3. `ytmp3-windows-amd64.exe` dosyasını PATH'e ekleyin veya istediğiniz klasöre taşıyın / Add to PATH or move to desired folder

### Kaynaktan Derleme / Build from Source

```bash
git clone https://github.com/mersieS/youtube_downloader.git
cd youtube_downloader
make install
```

---

## Kullanım / Usage

Terminali açıp şu komutu çalıştırın / Open terminal and run:

```bash
ytmp3
```

### Adım Adım / Step by Step

**1. Platform Seçimi / Platform Selection**

Ok tuşlarıyla (↑↓) platform seçin, Enter ile onaylayın:

Use arrow keys (↑↓) to select, Enter to confirm:

- **YouTube** — Sadece YouTube linkleri / YouTube links only
- **Diğer (Tüm siteler)** — Tüm desteklenen siteler / All supported sites

**2. URL Girişi / URL Input**

Video linkini yapıştırın / Paste the video URL:

```
YouTube URL girin (geri 'b' · çıkış 'q'): https://www.youtube.com/watch?v=...
```

**3. Format Seçimi / Format Selection**

```
❯ MP4 (Video)
  MP3 (Ses)
```

**4. Klasör Seçimi / Output Directory**

Enter'a basarak varsayılan klasörü kullanın veya özel bir yol girin:

Press Enter for default directory or type a custom path:

```
Kayıt klasörü [Enter = ~/Music/YouTube/Video]:
Klasör:
```

**5. Kalite Seçimi (MP4) / Quality Selection (MP4)**

```
❯ 2160p (4K)
  1440p (2K)
  1080p (Full HD)
  720p  (HD)
  480p
  360p
  En iyi kalite (otomatik)
```

### Kısayollar / Shortcuts

| Tuş / Key | İşlev / Action |
|---|---|
| ↑ / ↓ | Seçenekler arası gezinme / Navigate options |
| k / j | Yukarı / Aşağı (vim tarzı) / Up / Down (vim style) |
| Enter | Onayla / Confirm |
| b | Geri dön / Go back |
| q | Çıkış / Quit |

### Versiyon Kontrolü / Version Check

```bash
ytmp3 --version
```

---

## Varsayılan Kayıt Klasörleri / Default Download Directories

| Mod / Mode | MP4 | MP3 |
|---|---|---|
| YouTube | `~/Music/YouTube/Video` | `~/Music/YouTube/Muzik` |
| Diğer / Other | `~/Music/YouTube/Other/Video` | `~/Music/YouTube/Other/Muzik` |

---

## Lisans / License

MIT License - Salo Goodman
