
---
layout: "intersight"
page_title: "Intersight: intersight_macpool_pool_member"
sidebar_current: "docs-intersight-data-source-macpool-pool-member"
description: |-
PoolMember represents a single MAC address that is part of a pool.
---

# Data Source: intersight_macpool._pool_member
PoolMember represents a single MAC address that is part of a pool.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `assigned`:(bool) Boolean to represent whether the ID is assigned or not. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `mac_address`:(string) MAC Address of this pool member. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
