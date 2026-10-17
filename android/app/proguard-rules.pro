# Keep all OuTo application, service, and core classes
-keep class com.outo.app.** { *; }
-keep interface com.outo.app.** { *; }
-keep class com.outo.app.models.** { *; }
-keep class com.outo.app.service.models.** { *; }
-keep class com.outo.app.core.mobile.** { *; }

# Keep Gomobile & JNI
-keep class go.** { *; }

# Keep WiFi SSID plugin
-keep class com.follow.clash.wifi_ssid.** { *; }

# Keep Gson serialization targets
-keep class com.google.gson.** { *; }
-keepattributes *Annotation*,Signature,InnerClasses,EnclosingMethod

-dontwarn **
