# rules_manifest
Bazel rules for building CI/CD manifests for serverless functions and docker images

## Setup

To use rules from this repo, add the following to your `WORKSPACE` file:

```python
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "com_kindlyops_rules_manifest",
    urls = ["https://github.com/kindlyops/rules_manifest/archive/v0.1.0.tar.gz"],
    strip_prefix = "rules_manifest-0.1.0",
    sha256 = "99873d31226aa32dc025d651ca80628fbe2faa4e28d283436e0b82199200b7af",
)

load("@com_kindlyops_rules_manifest//:defs.bzl", "manifest_deps")
manifest_deps()
```

## Use

To generate a lambda manifest, invoke the rules like this:

```python
load("@com_kindlyops_rules_manifest//:manifest.bzl", "lambda_manifest")

lambda_manifest(
    name = "manifest",
    srcs = [
        "//lambdas/demo:lambda_deploy",
        "//lambdas/demo2:lambda_deploy",
    ],
)
```
