# Changelog

All notable changes to the AuTu project will be documented in this file.

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
