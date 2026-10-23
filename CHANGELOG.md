# Changelog

All notable changes to the AuTu project will be documented in this file.

## [v1.0.6] - 2026-10-23 (AuTu 1.0.6 预览版)

### 修复
- **修复节点卡片测速文本格式化错误**：彻底修复测速结果误展示为对象调试字符串 `TrafficShow(value: ..., unit: ...)` 的问题，正确格式化为清晰规范的 `↓ 2.6 MB/s`、`↑ 1.2 MB/s`。
- **防止小屏卡片文本截断与溢出**：测速行采用 `FittedBox` 动态自适应缩放，确保在任意屏幕尺寸和字号下上下行测速完整展现。

## [v1.0.5] - 2026-10-23 (AuTu 1.0.5 预览版)

### 节点卡片测速显示与布局优化
- **专属测速行呈现**：测速结果独立置于协议名称上方，测速中显示“测速中...”与加载动画；测速完成后以 `↓`（下行跑满最高速度）与 `↑`（上行跑满最高速度）清晰展示，彻底解决卡片空间狭小导致的重叠、残缺和乱码问题。
- **协议与延迟布局解耦**：协议名称（如 VLESS）居左，延迟（如 150 ms 或 ⚡）居右，各得其所，互不挤压。
- **长按菜单精简**：删除“导出剪贴板”选项，保持菜单聚焦于测速、延迟测试与节点详情。

### 测速引擎与响应机制改进
- **滑动窗口峰值采集**：引入 300ms 滑动窗口速率采集器，精确记录并展示测试期间跑满的最高速度（Peak Speed）。
- **3 节点并发工作池**：一键测所有节点速度采用并发调度，大幅缩短整体测速时间，各节点即测即刷新。
- **快速失败防卡死**：单个节点连接建立超时缩短至 2.5 秒，无效/虚拟节点快速跳过，彻底解决全组测速看似无响应的问题。

### 新增设置项
- **测速链接自定义**：在基础设置中新增“测下载速度、上传速度的测速链接”选项，支持自定义测速 URL 并持久化保存。

## [v1.0.4] - 2026-10-22 (AuTu 1.0.4 预览版)

### 节点长按操作菜单 (Proxy Card Context Menu)
- **卡片长按手势与触感反馈**：代理节点卡片支持长按手势（`onLongPress`），触发轻微触感振动反馈（`HapticFeedback.mediumImpact`）。
- **上下文菜单与功能集成**：
  - **测试当前节点速度**：支持快捷测试单个节点的速度，提供“测试上行+下行”、“测试上行”、“测试下行”三项子菜单。
  - **URLTEST (延迟测试)**：一键执行单个节点的网络延迟测试。
  - **查看节点详情**：弹窗展示节点名称、协议类型、延迟及上下行测速速率。
  - **导出剪贴板**：支持一键复制节点分享链接或节点名称到系统剪贴板。
- **卡片实时速率呈现**：测速完成后，节点卡片原有延迟区域自动动态展示测速结果（例如：`184 ms · ↓24.5 M/s ↑3.2 M/s`）。

### 策略组头部速度测试按钮 (Group Header Speed Test)
- **汽车油门仪表盘样式按钮**：在策略组操作栏最左侧新增汽车速度仪表盘图标按钮（`AppGlyphs.speed`）。
- **一键批量测速**：点击按钮弹出浮动菜单，支持 3 种批量测速模式：
  1. 测试上行+下行
  2. 测试上行
  3. 测试下行
- **测速状态动效**：测速过程中图标自动切换为旋转加载指示器，测速完成后恢复并即时同步各节点速率。

### 核心测速引擎 (Core Speed Test Engine)
- **专有出站拨号**：Go 核心层通过目标节点专属 `proxy.DialContext` 建立独立连接，实现无需切换当前全局/活动代理的定向测速。
- **流量节约保护**：测速时间窗口限制在 4 秒以内，下行采用流式丢弃计数，上行采用零拷贝流式载荷，杜绝内存堆积与流量浪费。
- **双层兼容降级**：客户端提供基于 `CoreMethod.speedTest` 与 `probe` 的双层兼容策略，确保各平台各版本环境下的高可用性。

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
