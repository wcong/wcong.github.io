---
layout: post
title: scrapy main compoment communicate
tags: scrapt compoment communicate
---
### main compoment
*   spider  
    *	    spider middleware
*   item pipelines
*   scrapy engine
    *	   download middleware
*   scheduler
*   downloader

### ExecutionEngine
ExecutionEngine combine all of them

### when engine shut down
_next_request in ExecutionEngine
``` python
if self.spider_is_idle(spider) and slot.close_if_idle:
   self._spider_idle(spider)
```

### spider middleware
Scraper crawl deep and  deal item