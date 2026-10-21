package com.outo.app.service.models

import com.outo.app.common.GlobalState
import com.outo.app.core.Core
import com.google.gson.Gson

private val gson = Gson()

data class Traffic(
    val up: Long = 0L,
    val down: Long = 0L,
)

val Long.formatBytes: String
    get() {
        val units = arrayOf("B", "KB", "MB", "GB", "TB")
        var value = toDouble()
        var unit = 0
        while (value >= 1024 && unit < units.lastIndex) {
            value /= 1024
            unit++
        }
        return if (unit == 0) {
            "${value.toLong()}${units[unit]}"
        } else {
            "%.1f${units[unit]}".format(value)
        }
    }

val Traffic.speedText: String
    get() = "${up.formatBytes}/s↑  ${down.formatBytes}/s↓"

val Traffic.totalText: String
    get() = "${up.formatBytes}↑  ${down.formatBytes}↓"

data class NotificationTrafficInfo(
    val collapsedText: String,
    val bigText: String,
)

fun Core.getSpeedTrafficText(onlyStatisticsProxy: Boolean): String {
    return runCatching {
        gson.fromJson(getTraffic(onlyStatisticsProxy), Traffic::class.java).speedText
    }.onFailure { error ->
        GlobalState.log("Unable to read traffic: $error")
    }.getOrDefault("")
}

fun Core.getNotificationTrafficInfo(onlyStatisticsProxy: Boolean): NotificationTrafficInfo {
    val speed = runCatching {
        gson.fromJson(getTraffic(onlyStatisticsProxy), Traffic::class.java)
    }.onFailure { error ->
        GlobalState.log("Unable to read speed: $error")
    }.getOrNull()

    val total = runCatching {
        gson.fromJson(getTotalTraffic(onlyStatisticsProxy), Traffic::class.java)
    }.onFailure { error ->
        GlobalState.log("Unable to read total traffic: $error")
    }.getOrNull()

    val speedStr = speed?.let { "↑${it.up.formatBytes}/s  ↓${it.down.formatBytes}/s" } ?: ""
    val totalStr = total?.let { "↑${it.up.formatBytes}  ↓${it.down.formatBytes}" } ?: ""

    val collapsedText = speedStr.ifBlank { "已连接" }
    val bigText = buildString {
        if (speedStr.isNotBlank()) {
            append("实时速率：").append(speedStr)
        }
        if (totalStr.isNotBlank()) {
            if (isNotEmpty()) append("\n")
            append("已用流量：").append(totalStr)
        }
    }.ifBlank { collapsedText }

    return NotificationTrafficInfo(
        collapsedText = collapsedText,
        bigText = bigText,
    )
}
