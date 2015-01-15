---
layout: post
title:	basicConfig of python logging
tags: python logging basicConfig
---
### python logging
#### why we should init python config before any log movement
##### how basicConfig work

``` python
        if len(root.handlers) == 0:
            filename = kwargs.get("filename")
	    ......
```
so you can see,it only init once, if root.handlers is not empty ,this method doing nothing
if you use first

``` python
logging.info('this is a test')
```

then init it

``` python
logging.basicConfig(format="[%(asctime)s %(name)s %(module)s %(lineno)d]%(levelname)s:%(message)s",
                    datefmt='%Y-%m-%d %H:%M:%S',
                    level=logging.INFO)
```

if will not work,the level of logging is still warning
