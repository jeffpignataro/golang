# GoLang Mono Repo

## Bazel Commands

_All these commands assume you are running them from the root of the project._

Build everything

```
bazel build //...
```

Build specific thing

```
bazel build //cmd/cli-demo:cli-demo
```

Run something

```
bazel run //cmd/cli-demo:cli-demo
```

Update `BUILD.bazel` files

```
bazel run //:gazelle
```

Update `WORKSPACE` file for new dependencies

```
bazel run //:gazelle -- update-repos -from_file go.mod
```

_Note if you `go mod tidy` it's always a good idea to `update-repos` at the same time._
