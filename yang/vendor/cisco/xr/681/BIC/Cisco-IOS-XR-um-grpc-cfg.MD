## Cisco-IOS-XR-um-grpc-cfg.yang

- [XPaths Obsoleted](#xpaths-obsoleted)
- [XPaths Deprecated](#xpaths-deprecated)
- [XPaths Added](#xpaths-added)
- [XPaths Removed](#xpaths-removed)
- [XPaths Modified](#xpaths-modified)

### Description

Added must statements under tls-cipher and address-family containers for mutually exclusive cases (these are missing constraints in the existing nodes thus practically does not break backward-compatibility).

## XPaths Obsoleted

N/A

## XPaths Deprecated

N/A

## XPaths Added

- (L161)	/grpc/address-family/ipv4
- (L167)	/grpc/address-family/ipv6
- (L173)	/grpc/address-family/dual
- (L340)	/grpc/tls-cipher/default/enable
- (L346)	/grpc/tls-cipher/default/disable

## XPaths Removed

N/A

## XPaths Modified

N/A

