class Ytmp3 < Formula
  desc "Free YouTube & Video Downloader CLI - MP4/MP3"
  homepage "https://github.com/mersieS/youtube_downloader"
  version "0.1.0"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/mersieS/youtube_downloader/releases/download/v#{version}/ytmp3-darwin-arm64.tar.gz"
      sha256 "76d29da5ce00549421ce00955d9304975c819c332d9f94b02bc0bbe1ecbdbefa"

      def install
        bin.install "ytmp3-darwin-arm64" => "ytmp3"
      end
    end

    on_intel do
      url "https://github.com/mersieS/youtube_downloader/releases/download/v#{version}/ytmp3-darwin-amd64.tar.gz"
      sha256 "cdb2ea6a7b3cbb385f8967270a3225b07f48986309c4865316bd5e4837d4fcff"

      def install
        bin.install "ytmp3-darwin-amd64" => "ytmp3"
      end
    end
  end

  depends_on "yt-dlp"
  depends_on "ffmpeg"

  test do
    assert_match "ytmp3 v#{version}", shell_output("#{bin}/ytmp3 --version")
  end
end
