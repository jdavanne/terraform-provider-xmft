---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "xmft_st_sentinel Resource - xmft"
subcategory: ""
description: |-
  
---

# xmft_st_sentinel (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `enabled` (Boolean) Whether the reporting to Axway Sentinel is enabled.
- `event_states` (Map of String)
- `fips_enabled` (Boolean) default:false, Enables the communication between SecureTransport and Axway Sentinel in a mode that is fully compliant with FIPS 140-2. This setting overrides the AxwaySentinel.SecureConnection.Protocol and AxwaySentinel.SecureConnection.EnabledCipherSuites.
- `heartbeat_delay` (Number) default:10, Time between the heartbeat messages.
- `heartbeat_enabled` (Boolean) default:false, If true the SecureTransport server sends a message to the Axway Sentinel server to tell it that SecureTransport is alive and connected.
- `heartbeat_time_unit` (String) default:seconds, Time unit for the time between the heartbeat messages.
- `host` (String) The IP address of the remote Axway Sentinel host.
- `mapping_rules` (Map of String)
- `name` (String) default:sentinel
- `overflow_file_path` (String) File path to the file where SecureTransport writes events in case it cannot connect to Axway Sentinel.
- `overflow_file_policy` (String) default:ignore, Possible values (ignore - Stops Collecting New Events | stop - Pauses All File Transfers Sent From and Received by SecureTransport). The default value is ignore.
- `overflow_file_size` (Number) default:1, Maximum size of overflow file in MB.
- `overflow_file_threshold` (Number) default:94, Represents the percent of file size at which warnings are sent.
- `port` (Number) Port of the remote Axway Sentinel host.
- `should_persist_link_data` (Boolean) default:false, If true SecureTransport server will maintain link data when reporting is disabled.
- `should_verify_cert` (Boolean) default:false, Whether the SecureTransport server to check the validity, trust status and hostname of the certificate presented by Axway Sentinel server upon establishing connection. If it is “false”, connections could be established even if the certificate is expired or the certificate attributes are not correct.
- `use_secure_connection` (Boolean) default:false, Whether SecureTransport server should send messages to the Sentinel server over a secure channel.

### Read-Only

- `last_updated` (String)
