# OuTo

<p align="center">
  <img src="assets/images/icon.png" width="128" height="128" alt="OuTo Logo">
  <br>
  <b>A modern, simple, and ad-free proxy client based on Mihomo (Clash.Meta) for Android</b>
  <br>
  <b>基于 Mihomo (Clash.Meta) 内核的 Android 现代代理客户端，简单易用、开源无广告</b>
</p>

<p align="center">
  <a href="https://github.com/Own716/OuTo/releases/latest"><img src="https://img.shields.io/badge/Release-v0.9.1-blue.svg?style=flat-square" alt="Version"></a>
  <a href="https://android-arsenal.com/api?level=21"><img src="https://img.shields.io/badge/Android-5.0%2B%20(API%2021%2B)-brightgreen.svg?style=flat-square" alt="API"></a>
  <a href="https://www.gnu.org/licenses/gpl-3.0"><img src="https://img.shields.io/badge/License-GPL--3.0-orange.svg?style=flat-square" alt="License"></a>
  <a href="https://t.me/OwnBoxs"><img src="https://img.shields.io/badge/Telegram-@OwnBoxs-2CA5E0.svg?logo=telegram&style=flat-square" alt="Telegram"></a>
</p>

---
<div align="center">

## Support this project

·If this project helps you, joining our TG channel is the best support you can give me.

---
</div>

## 📖 Introduction

**OuTo for Android** is a modern network proxy client for Android, deeply customized and built on top of the high-performance Mihomo (Clash.Meta) core. It brings together full-protocol-stack support, Material You dynamic color theming, lightning-fast latency testing, customizable rule routing, per-app proxy, WebDAV cloud backup and sync, a Quick Settings tile, and Android TV support — all combined with blazing-fast connections, low battery consumption, and an elegant, clean user interface.

**OuTo for Android** 是一款专为 Android 平台打造的现代代理客户端，深度适配主流 Android 手机、平板及 Android TV 设备。基于高性能 Mihomo (Clash.Meta) 内核构建，集成了全协议栈支持、Material You 动态取色、低延迟规则分流、分应用代理（黑白名单）、WebDAV 云端备份与同步、通知栏快捷磁贴等丰富特性，兼顾极速连接、低功耗与清爽直观的视觉交互。

---

## ✨ Features / 功能亮点

- ⚡ **高性能内核**：基于 Mihomo (Clash.Meta) 内核，支持 TUN 模式、低延迟分流及规则组代理。
- 🎨 **Material You 现代设计**：全面支持动态配色、深色/浅色自适应主题、连续曲率圆角与丝滑手势交互。
- 📱 **深度 Android 系统适配**：支持下拉通知栏快捷磁贴（Quick Settings Tile）、分应用代理（应用白名单/黑名单模式）及 Android TV 遥控器与大屏适配。
- 🔄 **灵活的配置管理**：支持一键订阅导入、内置在线配置文件编辑器、规则与覆写脚本、WebDAV 云端备份与同步。
- 📊 **可视化网络洞察**：实时连接监控、请求分析、DNS 查询日志记录与核心运行诊断。
- 🛡 **开源纯粹无广告**：完全透明、无任何第三方追踪或数据收集，纯净安全。

---

## 📥 Downloads

Please visit the official Releases page to download officially signed APK packages for all supported architectures:

👉 **[Go to the Official GitHub Releases Page](https://github.com/Own716/OuTo/releases)**

---

### 📦 Architecture Notes

* **ARM64 (v8a)**: 适用于绝大多数现代主流 64 位 Android 手机与平板设备（推荐）；
* **ARMeabi-v7a**: 适用于较老旧的 32 位 Android 设备或部分 TV 电视盒子；
* **x86_64 / x86**: 适用于 PC 端主流 Android 模拟器或 x86 架构 Android 设备。

---

## 🚀 Usage / 使用说明

### 链接导入配置
支持打开以下标准协议链接直接导入订阅配置（支持 `outo://` 与 `clash://` 协议）：

```text
outo://install-config?url=<经过 URL 编码的订阅链接>
clash://install-config?url=<经过 URL 编码的订阅链接>
```

### Android 自动化集成
支持通过 Tasker、MacroDroid 等自动化工具发送 Intent Action 控制代理启停：

```text
com.outo.app.action.START
com.outo.app.action.STOP
com.outo.app.action.TOGGLE
```

通过电脑 adb 调试控制：
```bash
adb shell am start -a com.outo.app.action.TOGGLE
```

---

## 🛠 Build from source / 从源码构建

编译 Android 版本需要安装 [Flutter](https://docs.flutter.dev/get-started/install) 3.47+、[Go](https://go.dev/dl/) 1.22+、JDK 17 以及 Android SDK / NDK。

```bash
# 克隆代码仓库
git clone https://github.com/Own716/OuTo.git
cd OuTo

# 获取依赖并构建 Android APK
flutter pub get
dart setup.dart android
```

构建生成的 APK 将输出至 `build/app/outputs/flutter-apk/` 目录。

---

## 💬 Issues & Feedback

If you encounter any issues, connection errors, or have feature requests, please feel free to submit an issue:

👉 **[Submit an Issue on GitHub](https://github.com/Own716/OuTo/issues)**

---

## ⚖️ License

OuTo is released under the [GNU General Public License v3.0 (GPL-3.0)](LICENSE).