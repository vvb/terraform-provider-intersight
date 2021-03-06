
---
layout: "intersight"
page_title: "Intersight: intersight_iaas_service_request"
sidebar_current: "docs-intersight-data-source-iaas-service-request"
description: |-
Gets last six months Service Requests from UCSD.
---

# Data Source: intersight_iaas._service_request
Gets last six months Service Requests from UCSD.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The fully-qualified name of the instantiated, concrete type.This property is used as a discriminator to identify the type of the payloadwhen marshaling and unmarshaling data.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `duration`:(string) Service request duration. 
* `initiating_user`:(string) Service Request initiating user. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. 
* `request_end_time`:(string) Service request end time. 
* `request_id`:(string) Service request id of an SR. 
* `request_start_time`:(string) Service request start time. 
* `request_type`:(string) Service request type of an SR. 
* `status`:(string) UCSD service request status. 
* `workflow_name`:(string) Executed workflow name for an SR. 
