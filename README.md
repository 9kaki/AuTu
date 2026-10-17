<div align="center">

<img src="assets/images/icon.png" alt="OuTo" width="88">

# OuTo

An Android proxy client based on ClashMeta. Simple to use, open source and ad-free.

**English** · [简体中文](README_zh_CN.md)

[![Release](https://img.shields.io/github/v/release/Own716/OuTo?style=flat-square&label=release)](https://github.com/Own716/OuTo/releases/latest)
[![Downloads](https://img.shields.io/github/downloads/Own716/OuTo/total?style=flat-square&logo=github)](https://github.com/Own716/OuTo/releases)
[![License](https://img.shields.io/github/license/Own716/OuTo?style=flat-square)](LICENSE)

[Website](https://Own716.github.io/OuTo) · [Download](#download) · [Changelog](CHANGELOG.md) · [Build from source](#build-from-source)

</div>

<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="snapshots/preview-dark.png">
    <img alt="OuTo dashboard" src="snapshots/preview.png" width="92%">
  </picture>
</p>

## Features

- **Built Specifically for Android**: Tailored for modern Android phones, tablets, and Android TV devices.
- **mihomo (Clash.Meta) core** with rule routing, proxy groups, latency tests, system proxy and TUN mode.
- **Profiles** from a subscription link or a file, with a built-in editor, override scripts, and custom rules,
  proxies and proxy groups.
- **Live views** of connections, requests, DNS queries and logs.
- **Material You design**: Dynamic color, light and dark themes, continuous corners, and sheets that snap between heights.
- **Backup and restore** through WebDAV or local storage.
- **Android platform integration**: Quick Settings tile, per-app proxy (whitelist/blacklist mode), and Android TV support.
- **Open source and ad-free**, licensed under GPL-3.0.

## Download

Get the latest build from [GitHub Releases](https://github.com/Own716/OuTo/releases/latest):

| Architecture | Package | Description |
| --- | --- | --- |
| ARM64 | `OuTo-*-android-arm64-v8a.apk` | Recommended for most modern Android devices |
| ARMv7 | `OuTo-*-android-armeabi-v7a.apk` | Older 32-bit Android phones and legacy TV boxes |
| x86_64 | `OuTo-*-android-x86_64.apk` | Android emulators or x86-based Android devices |

<p>
  <a href="https://github.com/Own716/OuTo/releases/latest"><img alt="Get it on GitHub" src="snapshots/get-it-on-github.svg" height="56"></a>
</p>

## Usage

**Import a profile from a link.** Opening a link in this form imports the subscription into OuTo. The `clashmeta://`
and `outo://` schemes work the same way:

```text
clash://install-config?url=<URL-encoded subscription link>
```

**Automate on Android.** Tasker, MacroDroid and similar apps can start, stop or toggle the proxy by starting an
activity with one of these actions:

```text
com.outo.app.action.START
com.outo.app.action.STOP
com.outo.app.action.TOGGLE
```

From a computer, the same works over adb:

```bash
adb shell am start -a com.outo.app.action.TOGGLE
```

## Build from source

Building Android requires [Flutter](https://docs.flutter.dev/get-started/install) 3.47+,
[Go](https://go.dev/dl/) 1.22+, JDK 17, and the Android SDK / NDK.

```bash
git clone --recursive https://github.com/Own716/OuTo.git
cd OuTo
flutter pub get
dart setup.dart android
```

Packages land in `dist/`. The Go core (via gomobile) is compiled as part of the Android build process.

## License

OuTo is released under the [GPL-3.0 license](LICENSE).

---

## Provenance & License Notice

- **Upstream Origin**: This project is an independent fork and complete rebrand of the open-source project [FlClash](https://github.com/chen08209/FlClash) at release [v0.9.1](https://github.com/chen08209/FlClash/releases/tag/v0.9.1).
- **Attribution**: Original code and architecture are copyright (c) 2023-2026 by chen08209 and the respective upstream contributors.
- **License**: Released under the terms of the [GNU General Public License v3.0 (GPL-3.0)](LICENSE).
- **Modifications**: Rebranded to OuTo, purely focused on the Android platform, package namespace migrated to `com.outo.app`, configured with a fresh release keystore, and non-Android platform code pruned.