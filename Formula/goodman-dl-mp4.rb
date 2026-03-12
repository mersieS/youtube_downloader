class GoodmanDlMp4 < Formula
  desc "Free YouTube & Video Downloader CLI - MP4/MP3"
  homepage "https://github.com/mersieS/youtube_downloader"
  version "0.1.0"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/mersieS/youtube_downloader/releases/download/v#{version}/goodman-dl-mp4-darwin-arm64.tar.gz"
      sha256 "bc137725022b5639418a766d6e0350ff1f9905f16ae48176c795cb8d3102553c"

      def install
        bin.install "goodman-dl-mp4-darwin-arm64" => "goodman-dl-mp4"
      end
    end

    on_intel do
      url "https://github.com/mersieS/youtube_downloader/releases/download/v#{version}/goodman-dl-mp4-darwin-amd64.tar.gz"
      sha256 "481a45abfcfe5310f972b4d26b87fa772913a2fb44787045c064d97948b58d6e"

      def install
        bin.install "goodman-dl-mp4-darwin-amd64" => "goodman-dl-mp4"
      end
    end
  end

  depends_on "yt-dlp"
  depends_on "ffmpeg"

  test do
    assert_match "goodman-dl-mp4 v#{version}", shell_output("#{bin}/goodman-dl-mp4 --version")
  end
end
