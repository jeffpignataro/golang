# GoLang Mono Repo

## Bazel Commands

_All these commands assume you are running them from the root of the project._

Build everything

```bash
bazel build //...
```

Build specific thing

```bash
bazel build //cmd/cli-demo:cli-demo
```

Run something

```bash
bazel run //cmd/cli-demo:cli-demo
```

Update `BUILD.bazel` files

```bash
bazel run //:gazelle
```

Update `WORKSPACE` file for new dependencies

```bash
bazel run //:gazelle -- update-repos -from_file go.mod
```

_Note if you `go mod tidy` it's always a good idea to `update-repos` at the same time._

## Testing methods

Running CLI commands which were just built is a little awkward as they live in the `bazel-bin` folder and are buried under a large folder structure. To make this is a little easier you can create a simple `symbolic link` under the `/usr/local/bin` folder (which should already be in the `PATH`) making it easy to rerun the command every build.

```bash
sudo ln -s ~/Repos/golang/bazel-bin/cmd/cli-demo/cli-demo_/cli-demo /usr/local/bin/cli-demo

which cli-demo # should return /usr/local/bin/cli-demo
```

Then you can easily run `cli-demo` to execute the command after each `bazel build \\...`

## Avoid broke stuff...

```
bazel build --build_tag_filters=-excluded //...
```
