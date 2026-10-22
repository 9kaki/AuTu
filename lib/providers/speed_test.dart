import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:outo/models/speed_test.dart';

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
