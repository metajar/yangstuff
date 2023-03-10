## Cisco-IOS-XR-um-router-isis-cfg.yang

- [XPaths Obsoleted](#xpaths-obsoleted)
- [XPaths Deprecated](#xpaths-deprecated)
- [XPaths Added](#xpaths-added)
- [XPaths Removed](#xpaths-removed)
- [XPaths Modified](#xpaths-modified)

### Description

Updated lsp-password and summary-prefix trees with proper constaraint statements. Bandwidth node type moved from leaf to container and few backward in-compatible nodes were been removed. Also must constraints were been alingned in many nodes(Adding missing must constraints does not break backward-compatibility). Also, levels under lsp-password was moved out of lsp-password and renamed as lsp-password-levels to fix a model bug. Moved container levels from container max-metric, this will break backward compatiblity, Moved presence containers as leafs under snmp-server traps. Re-ordered nodes under lsp-gen-interval container. Added cisco-support task yang nodes, logically it will not create backward compatiblity as customer wont have cisco-support cli configured. Re-ordered per-link, per-prefix nodes under interface fast-reroute container which will create non-backward compatibility.

## XPaths Obsoleted

N/A

## XPaths Deprecated

N/A

## XPaths Added

- (L203)	/snmp-server/traps/isis/all
- (L220)	/snmp-server/traps/isis/database-overload
- (L237)	/snmp-server/traps/isis/manual-address-drops
- (L254)	/snmp-server/traps/isis/corrupted-lsp-detected
- (L271)	/snmp-server/traps/isis/attempt-to-exceed-max-sequence
- (L288)	/snmp-server/traps/isis/id-len-mismatch
- (L305)	/snmp-server/traps/isis/max-area-addresses-mismatch
- (L322)	/snmp-server/traps/isis/own-lsp-purge
- (L339)	/snmp-server/traps/isis/sequence-number-skip
- (L356)	/snmp-server/traps/isis/authentication-type-failure
- (L373)	/snmp-server/traps/isis/authentication-failure
- (L390)	/snmp-server/traps/isis/version-skew
- (L407)	/snmp-server/traps/isis/area-mismatch
- (L424)	/snmp-server/traps/isis/rejected-adjacency
- (L441)	/snmp-server/traps/isis/lsp-too-large-to-propagate
- (L458)	/snmp-server/traps/isis/orig-lsp-buff-size-mismatch
- (L475)	/snmp-server/traps/isis/protocols-supported-mismatch
- (L492)	/snmp-server/traps/isis/adjacency-change
- (L509)	/snmp-server/traps/isis/lsp-error-detected
- (L186)	/snmp-server/traps/isis
- (L611)	/router/isis/processes/process[process-id]/lsp-password/text
- (L623)	/router/isis/processes/process[process-id]/lsp-password/hmac-md5
- (L635)	/router/isis/processes/process[process-id]/lsp-password/keychain
- (L859)	/router/isis/processes/process[process-id]/lsp-gen-interval
- (L989)	/router/isis/processes/process[process-id]/min-lsp-arrivaltime
- (L1170)	/router/isis/processes/process[process-id]/max-metric/on-startup/advertise-maximum-metric
- (L1187)	/router/isis/processes/process[process-id]/max-metric/on-startup/wait-for-bgp
- (L1779)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/segment-routing/mpls/connected-prefix-sid-map/prefixes/prefix[address-prefix]/index
- (L1796)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/segment-routing/mpls/connected-prefix-sid-map/prefixes/prefix[address-prefix]/absolute
- (L1844)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/segment-routing/mpls/connected-prefix-sid-map/flex-algo/prefix[address-prefix][flex-algo]/index
- (L1861)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/segment-routing/mpls/connected-prefix-sid-map/flex-algo/prefix[address-prefix][flex-algo]/absolute
- (L1899)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/segment-routing/mpls/connected-prefix-sid-map/strict-spf/prefix[address-prefix]/index
- (L1916)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/segment-routing/mpls/connected-prefix-sid-map/strict-spf/prefix[address-prefix]/absolute
- (L2172)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/redistribute/ospf[process-id]/match/internal
- (L2178)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/redistribute/ospf[process-id]/match/external
- (L2184)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/redistribute/ospf[process-id]/match/nssa-external
- (L2269)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/redistribute/ospfv3[process-id]/match/internal
- (L2275)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/redistribute/ospfv3[process-id]/match/external
- (L2281)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/redistribute/ospfv3[process-id]/match/nssa-external
- (L2851)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/redistribute/eigrp[as-number]/match/internal
- (L2857)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/redistribute/eigrp[as-number]/match/external
- (L3062)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/microloop/avoidance/protected
- (L3068)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/microloop/avoidance/segment-routing
- (L3089)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/apply-weight/ecmp-only
- (L3100)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/apply-weight/ucmp-only
- (L3207)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/metric-style/narrow
- (L3213)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/metric-style/wide
- (L3219)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/metric-style/transition
- (L3242)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/metric-style/levels/level[level-id]/narrow
- (L3248)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/metric-style/levels/level[level-id]/wide
- (L3254)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/metric-style/levels/level[level-id]/transition
- (L3479)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/spf/prefix-priority/prefix-priority[priority]/tag
- (L3490)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/spf/prefix-priority/prefix-priority[priority]/access-list-name
- (L3447)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/spf/prefix-priority/prefix-priority[priority]
- (L3541)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/spf/prefix-priority/levels/level[level-id][priority]/tag
- (L3552)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/spf/prefix-priority/levels/level[level-id][priority]/access-list-name
- (L3499)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/spf/prefix-priority/levels/level[level-id][priority]
- (L3666)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/attached-bit/send/always-set
- (L3672)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/attached-bit/send/never-set
- (L3709)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/fast-reroute/per-link/priority-limit/critical
- (L3715)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/fast-reroute/per-link/priority-limit/high
- (L3721)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/fast-reroute/per-link/priority-limit/medium
- (L3745)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/fast-reroute/per-link/priority-limit/levels/level[level-id]/critical
- (L3751)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/fast-reroute/per-link/priority-limit/levels/level[level-id]/high
- (L3757)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/fast-reroute/per-link/priority-limit/levels/level[level-id]/medium
- (L3808)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/fast-reroute/per-prefix/priority-limit/critical
- (L3814)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/fast-reroute/per-prefix/priority-limit/high
- (L3820)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/fast-reroute/per-prefix/priority-limit/medium
- (L3844)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/fast-reroute/per-prefix/priority-limit/levels/level[level-id]/critical
- (L3850)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/fast-reroute/per-prefix/priority-limit/levels/level[level-id]/high
- (L3856)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/fast-reroute/per-prefix/priority-limit/levels/level[level-id]/medium
- (L4678)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/redistribute/ospf[process-id]/match/internal
- (L4684)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/redistribute/ospf[process-id]/match/external
- (L4690)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/redistribute/ospf[process-id]/match/nssa-external
- (L4775)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/redistribute/ospfv3[process-id]/match/internal
- (L4781)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/redistribute/ospfv3[process-id]/match/external
- (L4787)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/redistribute/ospfv3[process-id]/match/nssa-external
- (L5357)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/redistribute/eigrp[as-number]/match/internal
- (L5363)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/redistribute/eigrp[as-number]/match/external
- (L5554)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/microloop/avoidance/protected
- (L5560)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/microloop/avoidance/segment-routing
- (L5581)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/apply-weight/ecmp-only
- (L5592)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/apply-weight/ucmp-only
- (L5698)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/metric-style/narrow
- (L5704)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/metric-style/wide
- (L5710)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/metric-style/transition
- (L5733)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/metric-style/levels/level[level-id]/narrow
- (L5739)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/metric-style/levels/level[level-id]/wide
- (L5745)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/metric-style/levels/level[level-id]/transition
- (L6010)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/spf/prefix-priority/prefix-priority[priority]/tag
- (L6021)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/spf/prefix-priority/prefix-priority[priority]/access-list-name
- (L5978)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/spf/prefix-priority/prefix-priority[priority]
- (L6072)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/spf/prefix-priority/levels/level[level-id][priority]/tag
- (L6083)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/spf/prefix-priority/levels/level[level-id][priority]/access-list-name
- (L6030)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/spf/prefix-priority/levels/level[level-id][priority]
- (L6197)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/attached-bit/send/always-set
- (L6203)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/attached-bit/send/never-set
- (L6346)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/hello-padding/disable
- (L6352)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/hello-padding/sometimes
- (L6375)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/hello-padding/levels/level[level-id]/disable
- (L6381)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/hello-padding/levels/level[level-id]/sometimes
- (L6527)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/hello-password/text
- (L6539)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/hello-password/hmac-md5
- (L6551)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/hello-password/keychain
- (L6592)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/hello-password-levels/level[level-id]/text
- (L6604)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/hello-password-levels/level[level-id]/hmac-md5
- (L6616)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/hello-password-levels/level[level-id]/keychain
- (L6955)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/override/metrics/maximum
- (L6961)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/override/metrics/high
- (L6967)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/override/metrics/disabled
- (L7042)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/passive
- (L7048)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/suppressed
- (L7054)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/shutdown
- (L7788)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/address-families/address-family[af-name][saf-name]/prefix-sid/sid/index
- (L7805)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/address-families/address-family[af-name][saf-name]/prefix-sid/sid/absolute
- (L7838)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/address-families/address-family[af-name][saf-name]/prefix-sid/strict-spf/index
- (L7855)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/address-families/address-family[af-name][saf-name]/prefix-sid/strict-spf/absolute
- (L7900)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/address-families/address-family[af-name][saf-name]/prefix-sid/algorithms/algorithm[algorithm-number]/index
- (L7917)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/address-families/address-family[af-name][saf-name]/prefix-sid/algorithms/algorithm[algorithm-number]/absolute
- (L8992)	/router/isis/processes/process[process-id]/adjacency/stagger/initial-number-of-neighbors
- (L9003)	/router/isis/processes/process[process-id]/adjacency/stagger/maximum-simultaneous-neighbors
- (L9108)	/router/isis/processes/process[process-id]/flex-algos/flex-algo[algorithm-number]/metric-type/delay
- (L9178)	/router/isis/processes/process[process-id]/affinity-maps/affinity-map[affinity-map-name]/bit-position

## XPaths Removed

- (L355)	/router/isis/processes/process[process-id]/lsp-password/levels
- (L435)	/router/isis/processes/process[process-id]/lsp-password/accept
- (L543)	/router/isis/processes/process[process-id]/lsp-gen-interval/maximum-wait/maximum-wait-time
- (L553)	/router/isis/processes/process[process-id]/lsp-gen-interval/maximum-wait/levels
- (L587)	/router/isis/processes/process[process-id]/lsp-gen-interval/initial-wait/initial-wait-time
- (L597)	/router/isis/processes/process[process-id]/lsp-gen-interval/initial-wait/levels
- (L631)	/router/isis/processes/process[process-id]/lsp-gen-interval/secondary-wait/secondary-wait-time
- (L641)	/router/isis/processes/process[process-id]/lsp-gen-interval/secondary-wait/levels
- (L723)	/router/isis/processes/process[process-id]/min-lsp-arrivaltime/maximum-wait/maximum-wait-time
- (L733)	/router/isis/processes/process[process-id]/min-lsp-arrivaltime/maximum-wait/levels
- (L767)	/router/isis/processes/process[process-id]/min-lsp-arrivaltime/initial-wait/initial-wait-time
- (L777)	/router/isis/processes/process[process-id]/min-lsp-arrivaltime/initial-wait/levels
- (L811)	/router/isis/processes/process[process-id]/min-lsp-arrivaltime/secondary-wait/secondary-wait-time
- (L821)	/router/isis/processes/process[process-id]/min-lsp-arrivaltime/secondary-wait/levels
- (L988)	/router/isis/processes/process[process-id]/max-metric/on-startup/levels
- (L1708)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/segment-routing/srv6/capability-subsubtlv
- (L2815)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/microloop/avoidance/rib-update-delay
- (L2998)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/spf-interval/maximum-wait/levels
- (L3042)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/spf-interval/initial-wait/levels
- (L3086)	/router/isis/processes/process[process-id]/address-families/address-family[af-name][saf-name]/spf-interval/secondary-wait/levels
- (L5234)	/router/isis/processes/process[process-id]/address-families/address-family-topology[af-name][saf-name][topology-name]/microloop/avoidance/rib-update-delay
- (L7012)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/address-families/address-family[af-name][saf-name]/fast-reroute/per-link/per-link-computation
- (L7017)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/address-families/address-family[af-name][saf-name]/fast-reroute/per-link/levels
- (L7099)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/address-families/address-family[af-name][saf-name]/fast-reroute/per-prefix/per-prefix-computation
- (L7104)	/router/isis/processes/process[process-id]/interfaces/interface[interface-name]/address-families/address-family[af-name][saf-name]/fast-reroute/per-prefix/levels
- (L7932)	/router/isis/processes/process[process-id]/flex-algos/flex-algo[algorithm-number]/affinity/exclude-any

## XPaths Modified

- (L190)	/snmp-server/traps/isis/all
- (L207)	/snmp-server/traps/isis/database-overload
- (L224)	/snmp-server/traps/isis/manual-address-drops
- (L241)	/snmp-server/traps/isis/corrupted-lsp-detected
- (L258)	/snmp-server/traps/isis/attempt-to-exceed-max-sequence
- (L275)	/snmp-server/traps/isis/id-len-mismatch
- (L292)	/snmp-server/traps/isis/max-area-addresses-mismatch
- (L309)	/snmp-server/traps/isis/own-lsp-purge
- (L326)	/snmp-server/traps/isis/sequence-number-skip
- (L343)	/snmp-server/traps/isis/authentication-type-failure
- (L360)	/snmp-server/traps/isis/authentication-failure
- (L377)	/snmp-server/traps/isis/version-skew
- (L394)	/snmp-server/traps/isis/area-mismatch
- (L411)	/snmp-server/traps/isis/rejected-adjacency
- (L428)	/snmp-server/traps/isis/lsp-too-large-to-propagate
- (L445)	/snmp-server/traps/isis/orig-lsp-buff-size-mismatch
- (L462)	/snmp-server/traps/isis/protocols-supported-mismatch
- (L479)	/snmp-server/traps/isis/adjacency-change
- (L496)	/snmp-server/traps/isis/lsp-error-detected
- (L606)	/router/isis/processes/process[process-id]/lsp-password
- (L863)	/router/isis/processes/process[process-id]/lsp-gen-interval/maximum-wait
- (L873)	/router/isis/processes/process[process-id]/lsp-gen-interval/initial-wait
- (L883)	/router/isis/processes/process[process-id]/lsp-gen-interval/secondary-wait
- (L993)	/router/isis/processes/process[process-id]/min-lsp-arrivaltime/maximum-wait
- (L1003)	/router/isis/processes/process[process-id]/min-lsp-arrivaltime/initial-wait
- (L1013)	/router/isis/processes/process[process-id]/min-lsp-arrivaltime/secondary-wait

