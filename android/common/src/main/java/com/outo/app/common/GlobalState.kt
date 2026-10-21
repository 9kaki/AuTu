package com.outo.app.common

import android.app.ActivityManager
import android.app.Application
import android.os.Build
import android.util.Log
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob

object GlobalState : CoroutineScope by CoroutineScope(SupervisorJob() + Dispatchers.Default) {
    const val NOTIFICATION_CHANNEL = "service-vpn-v2"
    const val NOTIFICATION_ID = 1
    private const val ANY_PID = 0
    private const val EVERY_EXIT_RECORD = 0

    val packageName: String
        get() = application.packageName

    val receiveBroadcastPermission: String
        get() = "$packageName.permission.RECEIVE_BROADCASTS"

    val application: Application
        get() = checkNotNull(appInstance) { "GlobalState is not initialized" }

    @Volatile
    private var appInstance: Application? = null

    fun init(application: Application) {
        appInstance = application
    }

    fun log(text: String) {
        Log.d("AuTu", text)
    }

    fun lastExitInfo(): Map<String, Any?>? {
        if (Build.VERSION.SDK_INT < Build.VERSION_CODES.R) return null
        val manager = application.getSystemService(ActivityManager::class.java) ?: return null
        val info = runCatching {
            manager.getHistoricalProcessExitReasons(
                packageName,
                ANY_PID,
                EVERY_EXIT_RECORD,
            ).firstOrNull()
        }.getOrNull() ?: return null
        return mapOf(
            "reason" to info.reason,
            "status" to info.status,
            "importance" to info.importance,
            "pss" to info.pss,
            "rss" to info.rss,
            "timestamp" to info.timestamp,
            "description" to info.description,
        )
    }
}
