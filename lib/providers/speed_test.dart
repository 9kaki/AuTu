import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:outo/common/preferences.dart';
import 'package:outo/models/speed_test.dart';

class SpeedTestConfig {
  const SpeedTestConfig({
    this.downloadUrl = SpeedTestConfigNotifier.defaultDownloadUrl,
    this.uploadUrl = SpeedTestConfigNotifier.defaultUploadUrl,
    this.timeout = SpeedTestConfigNotifier.defaultTimeout,
  });

  final String downloadUrl;
  final String uploadUrl;
  final int timeout;

  SpeedTestConfig copyWith({
    String? downloadUrl,
    String? uploadUrl,
    int? timeout,
  }) {
    return SpeedTestConfig(
      downloadUrl: downloadUrl ?? this.downloadUrl,
      uploadUrl: uploadUrl ?? this.uploadUrl,
      timeout: timeout ?? this.timeout,
    );
  }
}

class SpeedTestConfigNotifier extends Notifier<SpeedTestConfig> {
  static const defaultDownloadUrl =
      'https://speed.cloudflare.com/__down?bytes=100000000';
  static const defaultUploadUrl = 'https://speed.cloudflare.com/__up';
  static const defaultTimeout = 10000;

  @override
  SpeedTestConfig build() {
    _load();
    return const SpeedTestConfig();
  }

  Future<void> _load() async {
    final p = Preferences();
    var d = await p.getSpeedTestDownloadUrl();
    if (d == 'https://speed.cloudflare.com/__down?bytes=25000000') {
      d = defaultDownloadUrl;
      await p.setSpeedTestDownloadUrl(defaultDownloadUrl);
    }
    final u = await p.getSpeedTestUploadUrl();
    final t = await p.getSpeedTestTimeout();
    state = SpeedTestConfig(
      downloadUrl: d != null && d.isNotEmpty ? d : defaultDownloadUrl,
      uploadUrl: u != null && u.isNotEmpty ? u : defaultUploadUrl,
      timeout: t != null && t > 0 ? t : defaultTimeout,
    );
  }

  Future<void> update({String? downloadUrl, String? uploadUrl, int? timeout}) async {
    final next = state.copyWith(
      downloadUrl: downloadUrl,
      uploadUrl: uploadUrl,
      timeout: timeout,
    );
    state = next;
    final p = Preferences();
    if (downloadUrl != null) await p.setSpeedTestDownloadUrl(downloadUrl);
    if (uploadUrl != null) await p.setSpeedTestUploadUrl(uploadUrl);
    if (timeout != null) await p.setSpeedTestTimeout(timeout);
  }

  Future<void> reset() async {
    state = const SpeedTestConfig();
    final p = Preferences();
    await p.setSpeedTestDownloadUrl(defaultDownloadUrl);
    await p.setSpeedTestUploadUrl(defaultUploadUrl);
    await p.setSpeedTestTimeout(defaultTimeout);
  }
}

final speedTestConfigProvider =
    NotifierProvider<SpeedTestConfigNotifier, SpeedTestConfig>(
  SpeedTestConfigNotifier.new,
);

class SpeedTestDataSourceNotifier
    extends Notifier<Map<String, SpeedTestResult>> {
  @override
  Map<String, SpeedTestResult> build() => const {};

  void setResult(SpeedTestResult result) {
    state = {...state, result.proxyName: result};
  }

  void setResults(Iterable<SpeedTestResult> results) {
    final copy = Map<String, SpeedTestResult>.from(state);
    for (final r in results) {
      copy[r.proxyName] = r;
    }
    state = copy;
  }

  void clear() {
    state = const {};
  }
}

final speedTestDataSourceProvider =
    NotifierProvider<SpeedTestDataSourceNotifier, Map<String, SpeedTestResult>>(
  SpeedTestDataSourceNotifier.new,
);

class SpeedTestingProxiesNotifier extends Notifier<Set<String>> {
  @override
  Set<String> build() => const {};

  bool start(String proxyName) {
    if (state.contains(proxyName)) return false;
    state = {...state, proxyName};
    return true;
  }

  void stop(String proxyName) {
    if (!state.contains(proxyName)) return;
    final next = Set<String>.from(state)..remove(proxyName);
    state = next;
  }

  void clear() {
    state = const {};
  }
}

final speedTestingProxiesProvider =
    NotifierProvider<SpeedTestingProxiesNotifier, Set<String>>(
  SpeedTestingProxiesNotifier.new,
);

class SpeedTestingGroupsNotifier extends Notifier<Set<String>> {
  @override
  Set<String> build() => const {};

  bool start(String groupName) {
    if (state.contains(groupName)) return false;
    state = {...state, groupName};
    return true;
  }

  void stop(String groupName) {
    if (!state.contains(groupName)) return;
    final next = Set<String>.from(state)..remove(groupName);
    state = next;
  }

  void clear() {
    state = const {};
  }
}

final speedTestingGroupsProvider =
    NotifierProvider<SpeedTestingGroupsNotifier, Set<String>>(
  SpeedTestingGroupsNotifier.new,
);

final speedTestResultForProxyProvider =
    Provider.family<SpeedTestResult?, String>((ref, proxyName) {
  return ref.watch(speedTestDataSourceProvider.select((m) => m[proxyName]));
});

final isSpeedTestingProxyProvider =
    Provider.family<bool, String>((ref, proxyName) {
  return ref
      .watch(speedTestingProxiesProvider.select((s) => s.contains(proxyName)));
});
