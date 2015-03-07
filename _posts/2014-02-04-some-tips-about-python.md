---
layout: post
title: Some Tips About Python
description: "some little but import thing about python"
modified: 2014-02-04
tags: [python]
---
### none print when nohup execute python 
sometimes there is no print when execute python,and I do not known the reason
but I known this could work

``` bash
nohup python -u ./cmd.py > cmd.log 2>&1 &
```

### package management about python

#### import all the file of a direction
some time I want import all the file in a direction(eg test)

1.  I could create a new file named "__init__.py" if not exist
2.  in the file I write the list("__all__"),and put all the file name as variables(eg the follow code)

``` python 
__all__ = ["lines","text","fill",...]
```

### encode convert of string

#### decode to unicode
    str.decode(‘utf8’)
#### encode unicode to some char encode
    str.encode(‘utf8’)

