import 'package:outo/common/num.dart';

enum SpeedTestMode {
  both,
  upload,
  download,
}

enum SpeedTestPhase {
  queued,
  running,
  finished,
  failed,
}

class SpeedTestParams {
  const SpeedTestParams({
    required this.proxyName,
    this.mode = SpeedTestMode.both,
    this.downloadUrl = '',
    this.uploadUrl = '',
    this.timeout = 10000,
  });

  final String proxyName;
  final SpeedTestMode mode;
  final String downloadUrl;
  final String uploadUrl;
  final int timeout;

  Map<String, dynamic> toJson() => {
        'proxy-name': proxyName,
        'mode': mode.name,
        'download-url': downloadUrl,
        'upload-url': uploadUrl,
        'timeout': timeout,
      };
}

class SpeedTestResult {
  const SpeedTestResult({
    required this.proxyName,
    this.downloadSpeed,
    this.uploadSpeed,
    this.delay,
    this.phase = SpeedTestPhase.finished,
    this.error,
  });

  final String proxyName;
  final int? downloadSpeed; // bytes per second
  final int? uploadSpeed; // bytes per second
  final int? delay; // ms
  final SpeedTestPhase phase;
  final String? error;

  String? get formattedDownloadSpeed {
    if (downloadSpeed == null || downloadSpeed! <= 0) return null;
    return '${downloadSpeed!.traffic}/s';
  }

  String? get formattedUploadSpeed {
    if (uploadSpeed == null || uploadSpeed! <= 0) return null;
    return '${uploadSpeed!.traffic}/s';
  }

  factory SpeedTestResult.fromJson(Map<String, dynamic> json) {
    return SpeedTestResult(
      proxyName: json['proxy-name'] as String? ?? '',
      downloadSpeed: (json['download-speed'] as num?)?.toInt(),
      uploadSpeed: (json['upload-speed'] as num?)?.toInt(),
      delay: (json['delay'] as num?)?.toInt(),
      error: json['error'] as String?,
    );
  }

  Map<String, dynamic> toJson() => {
        'proxy-name': proxyName,
        'download-speed': downloadSpeed,
        'upload-speed': uploadSpeed,
        'delay': delay,
        'error': error,
      };
}
