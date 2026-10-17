cask "outo" do
  version "VERSION"

  on_macos do
    arch arm: "arm64", intel: "amd64"

    sha256 arm:   "ARM_SHA256",
           intel: "AMD_SHA256"

    url "https://github.com/Own716/OuTo/releases/download/v#{version}/OuTo-#{version}-macos-#{arch}.dmg"
  end

  name "OuTo"
  desc "Multi-platform proxy client based on ClashMeta"
  homepage "https://github.com/Own716/OuTo"

  livecheck do
    url :url
    strategy :github_latest
  end

  depends_on :macos

  app "OuTo.app"

  postflight_steps do
    run "/usr/bin/xattr",
        args:           ["-rd", "com.apple.quarantine", "{{appdir}}/OuTo.app"],
        writable_paths: ["OuTo.app"],
        writable_base:  :appdir
  end

  uninstall quit: "com.follow.clash"

  zap trash: [
    "~/Library/Application Support/com.follow.clash",
    "~/Library/Caches/com.follow.clash",
    "~/Library/Preferences/com.follow.clash.plist",
    "~/Library/Saved Application State/com.follow.clash.savedState",
  ]
end
