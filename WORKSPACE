# Copyright 2019 Kindly Ops, LLC. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
workspace(name = "com_kindlyops_rules_manifest")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("//:defs.bzl", "manifest_deps")
manifest_deps()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
go_rules_dependencies()
go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    sha256 = "7fc87f4170011201b1690326e8c16c5d802836e3a0d617d8f75c3af2b23180c4",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz"],
)
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
gazelle_dependencies()

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "io_bazel_stardoc",
    remote = "https://github.com/bazelbuild/stardoc.git",
    commit = "4378e9b6bb2831de7143580594782f538f461180", # 0.4.0
    shallow_since = "1570829166 -0400",
)

load("@io_bazel_stardoc//:setup.bzl", "stardoc_repositories")
stardoc_repositories()