'''
layout: page
title: search of elasticsearch
category: elasticsearch
tags: elasticsearch
'''
### _search  
#### path param 
*  [search request body](#request body json key)
*  search_type
*  query_cache
*  scroll
*  preference
*  template
*  _search_shards
*  exists 

##### request body json key
*  query
*  filtered
*  filter  
*  [aggregations](#aggregations)
*  [facets](#facets)
*  from
*  size
*  sort
*  return fields
    *  _source:"user*"
    *  fields:[]
    *  partial_fields:
    *  script_fields
    *  fielddata_fields
*  rescore
*  highlight
*  explain:true
*  version:true
*  indices_boost
*  min_score
*  suggest
*  valid:true

##### aggregations


### _msearch

### _count

### _explain
**must set to a single document**
