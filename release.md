### 1、检查配置(不实际发布)

```
goreleaser check
```

### 2、测试配置(不实际发布)

```
goreleaser release --snapshot
```
### 3、设置环境变量
```
export GITHUB_TOKEN="your-github-token"
```

### 4、打Tag并发布
```
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
goreleaser release
```