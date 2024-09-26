# Release

```
make testcover
make generate
// to get diff with swagger 
// may need to enable doUpdate... 
go test -timeout 20s -run ^TestResources$ terraform-provider-xmft/internal/provider -v -coverpkg=all
make install
git commit -m $TAG
git tag $TAG
git push origin $v0.0.13  git push origin --tags
```