---
layout: page
title: Links
description: links
keywords: links
comments: true
menu: links
permalink: /links/
---

> God made relatives. Thank God we can choose our friends.

{% for link in site.data.links %}
* [{{ link.name }}]({{ link.url }})
{% endfor %}
