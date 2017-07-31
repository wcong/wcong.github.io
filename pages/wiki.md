---
layout: page
title: Wiki
description: Wiki page for me
keywords: Wiki
comments: false
menu: Wiki
permalink: /wiki/
---

<ul class="listing">
{% for wiki in site.wiki %}
{% if wiki.title != "Wiki Template" %}
<li class="listing-item"><a href="{{ wiki.url }}">{{ wiki.title }}</a></li>
{% endif %}
{% endfor %}
</ul>
