
---
layout: "intersight"
page_title: "Intersight: intersight_processor_unit"
sidebar_current: "docs-intersight-data-source-processor-unit"
description: |-
The CPU present on a server.
---

# Data Source: intersight_processor._unit
The CPU present on a server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `architecture`:(string) The architecture of the installed processor. 
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_cores`:(int) The number of cores present in a given processor. 
* `num_cores_enabled`:(string) The number of enabled cores in the installed processor. 
* `num_threads`:(string) The maximum number of threads available in the installed processor. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. 
* `oper_power_state`:(string) The power state of the processor. 
* `oper_state`:(string) The health indicator of the processor, 'OK' indicates the processor is operatinal. 
* `operability`:(string) Operability state of the central processing unit. 
* `presence`:(string) The valid values are 'equipped' and 'absent'. 
* `processor_id`:(int) The ID number of a given processor. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `socket_designation`:(string) The socket ID of the installed processor. 
* `speed`:(float) The maximum speed of the installed processor in GHz. 
* `stepping`:(string) The CPU stepping of the installed processor. 
* `thermal`:(string) The temperature of the processor in centigrade. 
* `vendor`:(string) This field identifies the vendor of the given component. 
