# Changelog

All notable changes to the AuTu project will be documented in this file.

## [v1.0.3] - 2026-10-21 (AuTu 1.0.3 预览版)

### 超级岛 (Live Updates Capsule / Super Island)
- **Android 16 原生实况规范适配**：深度接入 Google Android 16 Rich Ongoing Notifications (Live Updates) 规范，通过 `requestPromotedOngoing` 与 `shortCriticalText` 在挖孔屏旁极简常驻显示。
- **主流 ROM 实况胶囊兼容**：全面兼容 OriginOS 原子岛、HyperOS 焦点通知、Flyme 灵动胶囊及 ColorOS 流体云协议。
- **智能地区提炼引擎 (`RegionExtractor`)**：内置 29+ 个国家与主要地区规则库，自动根据当前选中的代理节点提炼纯净的 2~4 字地区标签（如“香港”、“日本”、“新加坡”、“美国”）。
- **原生极简视觉规范**：移除彩色大圆盘底色，保持通透的原生 Logo 质感，清空 `subText` 避免系统强行穿插挤压排版。

### 通知中心 (Notification Center) 重构
- **节点名称智能联动**：通知标题动态同步显示选中的代理节点名称（`AuTu · [节点名称]`），节点变更实时生效。
- **双行精细化排版 (`BigTextStyle`)**：
  - 折叠态：极简单行实时上下行网速（`↑X.X KB/s  ↓X.X MB/s`）。
  - 展开态：双行精细化呈现，第一行实时速率，第二行累计已用流量（`Core.getTotalTraffic`）。
- **彻底杜绝幽灵残影**：优化生命周期销毁机制，停止 VPN 服务时显式取消前台通知与状态栏胶囊，彻底杜绝通知残影或白条。
- **息屏节能机制**：屏幕熄灭时自动挂起轮询刷新，屏幕点亮时每秒精准采集并进行差异比对，未发生变化不触发无谓重绘。

## [v1.0.2] - 2026-10-20 (AuTu 1.0.2 预览版)

### Icon Scaling Refinement
- **Balanced Logo Proportions**: Reduced center logo size to ~61% of visible circle diameter (~41% of adaptive canvas), providing generous, harmonious padding around the motif matching reference standard (FlClash) without being overly small.
- **Updated Full Asset Suite**: Synchronized the refined logo proportion across all density mipmaps (`mipmap-mdpi` through `mipmap-xxxhdpi`), night mipmaps, Android TV banner, notification drawables, and multi-size ICO/PNG desktop assets.

## [v1.0.1] - 2026-10-20 (AuTu 1.0.1 预览版)

### Icon & Adaptive Launcher Fixes
- **Fixed Launcher Icon Zoom/Crop Issue**: Re-engineered all Android adaptive icon foregrounds into density-specific mipmap buckets (`mipmap-mdpi` through `mipmap-xxxhdpi`), eliminating abnormal 400% zoom caused by unqualified drawable dimensions.
- **Enhanced Safe Zone & Aspect Ratio**: Motif is now precisely centered within the 72dp adaptive safe zone with clean 18dp padding, ensuring pristine rendering across all Android launchers (circle, squircle, round rectangle).
- **Seamless System Auto-Follow & Night Mode**:
  - Added dedicated `mipmap-night-anydpi-v26/` resource tree so launcher icons automatically switch between light and dark modes according to the phone's system settings.
  - Added full atomic component switching on Android 13+ via `setComponentEnabledSettings` to prevent launcher caching issues or icon flickering.
- **Refined Notification & TV Assets**: Fully synchronized notification drawables and Android TV banner with the new motif geometry.

## [v1.0.0] - 2026-10-18 (AuTu 1.0.0 预览版)

### Branding & Visuals
- **New Brand Identity**: Completely rebranded from OuTo to **AuTu**.
- **Dynamic Dual Icons**:
  - Introduced new brand icons featuring light mode (white background, black motif) and dark mode (black background, white motif).
  - **Desktop Launcher**: Switches dynamically between light and dark launcher icons, or automatically follows system night mode.
  - **Notification Center**: Notification badge and service icon seamlessly adjust between light and dark themes to match in-app and system modes.
  - **Auto Follow System**: Fully respects Android system-wide dark/light settings.

### Core & Stability
- **Fixed Android Startup Crash**: Resolved class stripping caused by stale R8 ProGuard rules; all Kotlin components, plugins, and state models are fully preserved.
- **Unified Release Signing**: Integrated unified, permanent release signing keys managed via GitHub Actions Secrets, guaranteeing zero signature mismatches across upgrades.
- **Mihomo (Clash.Meta) Core**: High-performance TUN proxy, multi-protocol splitting, and low latency routing.
- **Pure & Clean**: 100% open source, ad-free, and privacy-first.
