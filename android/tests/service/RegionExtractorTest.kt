package com.outo.app.service.utils

import org.junit.Assert.assertEquals
import org.junit.Test

class RegionExtractorTest {

    @Test
    fun `extracts known regions correctly`() {
        assertEquals("香港", RegionExtractor.extractRegion("🇭🇰 香港 01 [BGP]"))
        assertEquals("日本", RegionExtractor.extractRegion("🇯🇵 日本 Tokyo-01"))
        assertEquals("新加坡", RegionExtractor.extractRegion("SG 新加坡 02"))
        assertEquals("美国", RegionExtractor.extractRegion("US - 美国洛杉矶 05"))
        assertEquals("台湾", RegionExtractor.extractRegion("TW 台湾 03"))
        assertEquals("韩国", RegionExtractor.extractRegion("KR 韩国首尔"))
        assertEquals("英国", RegionExtractor.extractRegion("GB 英国伦敦 01"))
        assertEquals("德国", RegionExtractor.extractRegion("DE 德国法兰克福"))
    }

    @Test
    fun `extracts by country codes and abbreviations`() {
        assertEquals("日本", RegionExtractor.extractRegion("jp-direct-node"))
        assertEquals("香港", RegionExtractor.extractRegion("hk_premium_node"))
        assertEquals("新加坡", RegionExtractor.extractRegion("sg-01"))
        assertEquals("美国", RegionExtractor.extractRegion("us-west"))
    }

    @Test
    fun `handles fallback and empty cases`() {
        assertEquals("已连接", RegionExtractor.extractRegion(""))
        assertEquals("已连接", RegionExtractor.extractRegion(null))
        assertEquals("专线节点", RegionExtractor.extractRegion("专线节点-01"))
    }
}
