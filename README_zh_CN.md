<div align="center">

<img src="assets/images/icon.png" alt="OuTo" width="88">

# OuTo

基于 ClashMeta 的 Android 代理客户端，简单易用，开源无广告。

[English](README.md) · **简体中文**

[![Release](https://img.shields.io/github/v/release/Own716/OuTo?style=flat-square&label=release)](https://github.com/Own716/OuTo/releases/latest)
[![Downloads](https://img.shields.io/github/downloads/Own716/OuTo/total?style=flat-square&logo=github)](https://github.com/Own716/OuTo/releases)
[![License](https://img.shields.io/github/license/Own716/OuTo?style=flat-square)](LICENSE)

[官网](https://Own716.github.io/OuTo/zh) · [下载](#下载) · [更新日志](CHANGELOG.md) · [从源码构建](#从源码构建)

</div>

<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="snapshots/preview-dark.png">
    <img alt="OuTo 仪表盘" src="snapshots/preview.png" width="92%">
  </picture>
</p>

## 功能

- **专为 Android 设备深度优化**：支持主流 Android 手机、平板以及 Android TV，纯粹专注 Android 端体验。
- **mihomo（Clash.Meta）内核**：规则分流、代理组、延迟测试、系统代理与 TUN 模式。
- **配置管理**：通过订阅链接或文件导入配置，内置编辑器，支持覆写脚本以及自定义规则、代理和代理组。
- **实时查看**：连接详情、请求、DNS 查询与实时核心日志。
- **Material You 设计**：支持动态取色、浅色与深色主题，连续曲率圆角，弹出面板多档高度吸附。
- **备份与恢复**：支持通过 WebDAV 或本地存储导出与导入配置。
- **Android 系统特性**：支持快捷设置磁贴、分应用代理（应用白名单/黑名单模式）以及 Android TV 遥控器与大屏交互适配。
- **开源无广告**，以 GPL-3.0 协议发布。

## 下载

从 [GitHub Releases](https://github.com/Own716/OuTo/releases/latest) 获取最新版本：

| 架构 | 安装包 | 适用设备 |
| --- | --- | --- |
| ARM64 | `OuTo-*-android-arm64-v8a.apk` | 绝大多数现代 Android 手机推荐下载 |
| ARMv7 | `OuTo-*-android-armeabi-v7a.apk` | 较老旧的 32 位 Android 设备或部分 TV 盒子 |
| x86_64 | `OuTo-*-android-x86_64.apk` | 模拟器或 x86 架构 Android 设备 |

<p>
  <a href="https://github.com/Own716/OuTo/releases/latest"><img alt="Get it on GitHub" src="snapshots/get-it-on-github.svg" height="56"></a>
</p>

## 使用

**通过链接导入配置。** 打开下面格式的链接，即可把订阅导入 OuTo。`clashmeta://` 和 `outo://` 开头的链接同样可用：

```text
clash://install-config?url=<经过 URL 编码的订阅链接>
```

**Android 自动化。** Tasker、MacroDroid 等应用可以用下面的 Action 启动 Activity，从而启动、停止或切换代理：

```text
com.outo.app.action.START
com.outo.app.action.STOP
com.outo.app.action.TOGGLE
```

在电脑上也可以通过 adb 触发：

```bash
adb shell am start -a com.outo.app.action.TOGGLE
```

## 从源码构建

编译 Android 版本需要 [Flutter](https://docs.flutter.dev/get-started/install) 3.47+、[Go](https://go.dev/dl/) 1.22+、JDK 17 以及 Android SDK / NDK。

```bash
git clone --recursive https://github.com/Own716/OuTo.git
cd OuTo
flutter pub get
dart setup.dart android
```

APK 安装包将直接输出至 `dist/` 目录。Go 内核（通过 gomobile）和相关库会在构建过程中一并编译完成。

## 许可证

OuTo 以 [GPL-3.0 协议](LICENSE)发布。

---

## 开源溯源与许可证声明

- **上游来源**：本项目基于开源项目 [FlClash](https://github.com/chen08209/FlClash) 的 [v0.9.1](https://github.com/chen08209/FlClash/releases/tag/v0.9.1) 稳定版本进行独立复刻与全面品牌重命名（Rebranded to **OuTo**）。
- **致谢与版权**：原始核心架构与代码版权归原作者 chen08209 及相关贡献者所有。
- **开源许可证**：遵循 [GNU General Public License v3.0 (GPL-3.0)](LICENSE) 开源协议。
- **架构调整**：专精聚焦于 Android 平台，剥离非 Android（Windows / macOS / Linux）桌面原生代码与外部特权服务，包名定义为 `com.outo.app`，配置独立发布签名与精简 CI/CD 流水线。