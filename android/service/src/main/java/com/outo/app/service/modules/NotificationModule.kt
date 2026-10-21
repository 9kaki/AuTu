package com.outo.app.service.modules

import android.app.Notification.FOREGROUND_SERVICE_IMMEDIATE
import android.app.Service
import android.app.Service.STOP_FOREGROUND_REMOVE
import android.content.Intent
import android.os.Build
import android.os.PowerManager
import androidx.core.app.NotificationCompat
import androidx.core.app.NotificationManagerCompat
import androidx.core.content.getSystemService
import com.outo.app.common.Components
import com.outo.app.common.GlobalState
import com.outo.app.common.QuickAction
import com.outo.app.common.ensureNotificationChannel
import com.outo.app.common.quickIntent
import com.outo.app.common.receiveBroadcastFlow
import com.outo.app.common.startForeground
import com.outo.app.common.toPendingIntent
import com.outo.app.core.Core
import com.outo.app.service.R
import com.outo.app.service.ServiceConfig
import com.outo.app.service.models.NotificationParams
import com.outo.app.service.models.NotificationTrafficInfo
import com.outo.app.service.models.getNotificationTrafficInfo
import com.outo.app.service.utils.RegionExtractor
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.delay
import kotlinx.coroutines.flow.collectLatest
import kotlinx.coroutines.flow.combine
import kotlinx.coroutines.flow.distinctUntilChanged
import kotlinx.coroutines.flow.flow
import kotlinx.coroutines.flow.map
import kotlinx.coroutines.flow.onStart
import kotlinx.coroutines.launch

private data class ExtendedNotificationParams(
    val title: String,
    val stopText: String,
    val showStopAction: Boolean,
    val contentText: String,
    val bigContentText: String,
)

private val NotificationParams.extended: ExtendedNotificationParams
    get() {
        val trafficInfo = Core.getNotificationTrafficInfo(onlyStatisticsProxy)
        return ExtendedNotificationParams(
            title = title,
            stopText = stopText,
            showStopAction = showStopAction,
            contentText = trafficInfo.collapsedText,
            bigContentText = trafficInfo.bigText,
        )
    }

internal class NotificationModule(
    private val service: Service,
    private val scope: CoroutineScope,
) : ServiceModule {
    private var isForegroundStarted = false

    override fun start() {
        service.ensureNotificationChannel()
        isForegroundStarted = false
        update(ServiceConfig.notificationParams.value.extended)
        scope.launch {
            // Some ROMs, vivo's among them, drop updates posted behind the
            // keyguard, so unlocking restarts the loop to post once more.
            // SystemUI, an app, sends USER_PRESENT, which a receiver that is
            // not exported never gets; all three actions are protected.
            service.receiveBroadcastFlow(exported = true) {
                addAction(Intent.ACTION_SCREEN_ON)
                addAction(Intent.ACTION_SCREEN_OFF)
                addAction(Intent.ACTION_USER_PRESENT)
            }.map { intent ->
                intent.action != Intent.ACTION_SCREEN_OFF
            }.onStart {
                emit(isScreenOn())
            }.collectLatest { screenOn ->
                if (!screenOn) return@collectLatest
                combine(
                    flow {
                        while (true) {
                            delay(1_000)
                            emit(Unit)
                        }
                    },
                    ServiceConfig.notificationParams,
                ) { _, params ->
                    params.extended
                }.distinctUntilChanged()
                    .collect(::update)
            }
        }
    }

    private fun isScreenOn() =
        service.getSystemService<PowerManager>()?.isInteractive ?: true

    private val notificationBuilder: NotificationCompat.Builder by lazy {
        val intent = Intent().setComponent(Components.mainActivity)

        NotificationCompat.Builder(
            service,
            GlobalState.NOTIFICATION_CHANNEL,
        ).apply {
            setWhen(0)
            setSmallIcon(R.drawable.ic_service)
            setContentTitle("AuTu")
            setContentIntent(intent.toPendingIntent)
            setPriority(NotificationCompat.PRIORITY_LOW)
            setCategory(NotificationCompat.CATEGORY_SERVICE)
            setOngoing(true)
            if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.S) {
                foregroundServiceBehavior = FOREGROUND_SERVICE_IMMEDIATE
            }
            setOnlyAlertOnce(true)
        }
    }

    private val stopIntent by lazy { QuickAction.STOP.quickIntent.toPendingIntent }

    private fun applyLiveUpdateCapsule(
        builder: NotificationCompat.Builder,
        title: String,
    ) {
        // 取消绿色/黄色强调色，移除展开卡片彩色大圆盘，与收起态统一为通透原生 Logo
        builder.color = NotificationCompat.COLOR_DEFAULT

        val displayChipText = RegionExtractor.extractRegion(title).ifBlank { "已连接" }

        // 谷歌原生 Android 16 实况岛 / Rich Ongoing Notifications (Live Updates) 规范
        // 收起状态：挖孔屏旁展示原生纯净 Logo 与地区汉字（如“日本”）
        builder.extras.apply {
            putBoolean("android.requestPromotedOngoing", true)
            putCharSequence("android.shortCriticalText", displayChipText)
            putString("capsule_text", displayChipText)
            remove("capsule_color")
            putString("live_activity_status", "active")
        }

        // 彻底清空 subText，避免系统在标题右侧强行并排追加网速文本
        builder.setSubText(null)
    }

    private fun update(params: ExtendedNotificationParams) {
        val notification = with(notificationBuilder) {
            setSmallIcon(R.drawable.ic_service)
            setContentTitle(params.title)
            setContentText(params.contentText)
            setStyle(NotificationCompat.BigTextStyle().bigText(params.bigContentText))
            applyLiveUpdateCapsule(this, params.title)
            clearActions()
            if (params.showStopAction) {
                val stopAction = NotificationCompat.Action.Builder(
                    0,
                    params.stopText,
                    stopIntent,
                ).setShowsUserInterface(false).build()
                addAction(stopAction)
            }
            build()
        }

        if (!isForegroundStarted) {
            service.startForeground(notification)
            isForegroundStarted = true
        } else {
            NotificationManagerCompat.from(service).notify(GlobalState.NOTIFICATION_ID, notification)
        }
    }

    @Suppress("DEPRECATION")
    override fun stop() {
        isForegroundStarted = false
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.N) {
            service.stopForeground(STOP_FOREGROUND_REMOVE)
        } else {
            service.stopForeground(true)
        }
        NotificationManagerCompat.from(service).cancel(GlobalState.NOTIFICATION_ID)
    }
}
