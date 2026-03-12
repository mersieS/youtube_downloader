class GoodmanDl < Formula
  desc "Free YouTube & Video Downloader CLI - MP4/MP3"
  homepage "https://github.com/mersieS/youtube_downloader"
  version "0.1.0"
  license "MIT"

  on_macos do
    on_arm do
      url "https://github.com/mersieS/youtube_downloader/releases/download/v#{version}/goodman-dl-mp4-darwin-arm64.tar.gz"
      sha256 "b7cb0948c5ea93c14beee75e8b46cc211462677ff83107ec63b53da3c92ba0c7"

      def install
        bin.install "goodman-dl-mp4-darwin-arm64" => "goodman-dl-mp4"
      end
    end

    on_intel do
      url "https://github.com/mersieS/youtube_downloader/releases/download/v#{version}/goodman-dl-mp4-darwin-amd64.tar.gz"
      sha256 "e6a11a3bddbd26cdace4ce05857dcc308718ed6a62f49ae1323d7c4347c021f9"

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
