
# TF integration

- default (optional) vs unknown value (optional/computed)
   - input can be null,unknown,empty/def 
   - output can be null,unknown,empty
   - "" becomes null -> yes : use default "" and emptyIsNull, or use optional
   - null keeps null -> yes : use default else use default and emptyIsNull
     - emptyIsNull doesn't send any value
     - emptyIsNull convert null to ""   
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

- /api/v2.0/sites (ssh) - "clientCertificate": "" : Error "message" : "Id cannot be null or empty",
- /api/v2.0/sites (pesit)
  - "Id cannot be null or empty" on missing certificate
  - pesit.Protocols : PreconnectionPartnerId / PreconnectionServerId : not used but require something and sendback empty !! *workaround* noread
  - 

- inconsistence between null and "" ? 
  - account.businessUnit null -> return "", "" is not allowed 
- choose between default, computed/optional vs default emptyIsNull

- businessUnit : no id ???? (*inconsistency*)
- sentinel.host : bad host resolution : cannot configure -> no check ?

- /transferProfiles : sudden error: "Incorrect Accept header. Allowed values are application/json and multipart/mixed"
- inconsistency of enums : UPPERCASE vs Lowercase vs camelCase vs pascal case
  - Proposal : add options with choosen case, and deprecate others (???) but what to return !!!!

- references by name vs id : rename may break or not the links
  - rename account : 
    -  site
    -  profile
    -  route
    -  subscription
    -  application
   -  rename business unit
  
- inconsistent glob "case"

- step character-replace/strip : weird \u character....

- missing api for available plugins /version : authz/authn/site/step...

- file-maintenance-application : scheduler startDate ??? what for

- /api/v2.0/sites type=ftp bad field leads to "Incorrect JSON format"

- swagger schema : default as "string" !!!!
- 
# CFT API issues
- cftpart/cfttcp logic


