
# TF integration

- default (optional) vs unknown value (optional/computed)
- sub struct management with unknown pointer vs object => onsequence on documentation
- dependencies on fields / bool enabled

# ST API issues

- /api/v2.0/version : schema/example error string vs bool
- /api/v2.0/sites 
  - example: ssh maxConcurrent : unsupported parameter  -> maxConcurrentConnection
  - example: uploadPermissions int (644) -> string (644)
  - example: uploadPermissions int (0644) /octal
- /api/2.0/routes
  - We should have routeTemplate, routeComposite, simpleRoute
    - routeTemplate, routeComposite only accept ExecuteRoute ?
    - simpleRoute doesn't accept ExecuteRoute?
  - routeComposite vs routePackage !!!!
  - simpleRoute lifecycle is inconsistent : create independly but removed by composite/template !!!

# CFT API issues
- cftpart/cfttcp logic


# ST gaps

- version (datasource)
- account : user
- route_template : ???
- route_composite
- application : advanced_routing
- site : ssh
- subscriptions
- route_simple (+ steps)

- site : ftp
- site : http
- site : pesit
- transferProfiles
- site : as/2
- businessUnits
- certificates
- zones
- administrators
- configurations / options
- configurations / fileArchiving
- configurations / loginSettings
- configurations / sentinel
- icapServers
- ldapDomains
- loginRestrictionPolicies
- mailTemplates
- siteTemplates
- userClasses

# CFT gaps

- cftpart/cfttcp
- cftsend
- cftrecv

- cftsendi
- cftssl
- cftssh
- pki...
- 
- cftprot
- cftauth
- cft

