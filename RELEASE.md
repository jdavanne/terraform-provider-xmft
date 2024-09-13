# Release

```
make testcover
make generate
// to get diff with swagger 
// may need to enable doUpdate... 
go test -timeout 20s -run ^TestResources$ terraform-provider-xmft/internal/provider -v -coverpkg=all
make install
```