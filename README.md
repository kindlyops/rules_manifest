# rules_manifest
Bazel rules for building CI/CD manifests for serverless functions and docker images

## Setup

To use rules from this repo, add the following to your `WORKSPACE` file:

```python
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "com_kindlyops_rules_manifest",
    urls = ["https://github.com/kindlyops/rules_manifest/archive/v0.2.1.tar.gz"],
    strip_prefix = "rules_manifest-0.2.1",
    sha256 = "476f374a5b125032ffdeca8541302fc87fb37207bba4792c4f4baa1e19ee5222",
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
