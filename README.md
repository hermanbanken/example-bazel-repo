# Bazel for n00bs
Created by a Bazel n00b just learning Bazel for the first time.

Preparations:
```bash
brew install bazel
go install github.com/bazelbuild/buildtools/buildozer@latest
```

Generate something useful:
```bash
bazel run //projectB:tarball
```