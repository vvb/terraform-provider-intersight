
---
layout: "intersight"
page_title: "Intersight: intersight_capability_io_card_descriptor"
sidebar_current: "docs-intersight-data-source-capability-io-card-descriptor"
description: |-
Descriptor that uniquely identifies an IO card module.
---

# Data Source: intersight_capability._io_card_descriptor
Descriptor that uniquely identifies an IO card module.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `description`:(string) Detailed information about the endpoint. 
* `model`:(string) The model of the endpoint, for which this capability information is applicable. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `num_hif_ports`:(int) Number of hif ports per blade for the iocard module. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `revision`:(string) Revision for the iocard module. 
* `vendor`:(string) The vendor of the endpoint, for which this capability information is applicable. 
* `version`:(string) The firmware or software version of the endpoint, for which this capability information is applicable. 
