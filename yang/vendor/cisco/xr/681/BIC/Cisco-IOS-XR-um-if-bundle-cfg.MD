## Cisco-IOS-XR-um-if-bundle-cfg.yang

- [XPaths Obsoleted](#xpaths-obsoleted)
- [XPaths Deprecated](#xpaths-deprecated)
- [XPaths Added](#xpaths-added)
- [XPaths Removed](#xpaths-removed)
- [XPaths Modified](#xpaths-modified)

### Description

Modified type and primitive-type values for leaf destination under container ipv4. Removed container period-on-bundle due to sysdb path change.

## XPaths Obsoleted

N/A

## XPaths Deprecated

N/A

## XPaths Added

- (L75)	/{GROUP-BUNDLE}/bundle
- (L489)	/{GROUP-BUNDLE}/lacp

## XPaths Removed

- (L376)	/{GROUP-BUNDLE}/bfd/address-family/ipv4/ipv6
- (L650)	/{GROUP-BUNDLE}/lacp/period-on-bundle

## XPaths Modified

- (L344)	/{GROUP-BUNDLE}/bfd/address-family/ipv4/destination

