def _lambda_manifest_impl(ctx):
    tree = ctx.actions.declare_directory(ctx.attr.name + ".artifacts")
    args = [str(ctx.outputs.manifest.path)] + [f.path for f in ctx.files.srcs]

    ctx.actions.run(
        inputs = ctx.files.srcs,
        arguments = args,
        outputs = [ctx.outputs.manifest, tree],
        progress_message = "Generating %s manifest" % str(ctx.outputs.manifest.path),
        executable = ctx.executable._manifester,
    )
    return [DefaultInfo(files = depset([tree]))]

lambda_manifest = rule(
    implementation = _lambda_manifest_impl,
    doc = """
Builds a content-addressable artifact repo with manifest file.

This is useful for collecting together one or more lambda zip artifacts and
publishing them to S3 as content-addressed artifacts.
""",
    attrs = {
        "srcs": attr.label_list(),
        "_manifester": attr.label(
            default = Label("//lambdamanifester:manifester"),
            allow_single_file = True,
            executable = True,
            cfg = "host",
        ),
    },
    outputs = {
        "manifest": "%{name}.artifacts/%{name}.json",
    },
)

def _docker_manifest_impl(ctx):
    args = [str(ctx.outputs.manifest.path)] + [f.path for f in ctx.files.srcs]

    ctx.actions.run(
        inputs = ctx.files.srcs,
        arguments = args,
        outputs = [ctx.outputs.manifest],
        progress_message = "Generating %s manifest" % str(ctx.outputs.manifest.path),
        executable = ctx.executable._manifester,
    )
    return [DefaultInfo(files = depset([ctx.outputs.manifest]))]

docker_manifest = rule(
    implementation = _docker_manifest_impl,
    attrs = {
        "srcs": attr.label_list(),
         "_manifester": attr.label(
            default = Label("//containermanifester:manifester"),
            allow_single_file = True,
            executable = True,
            cfg = "host",
        ),
    },
    outputs = {
        "manifest": "%{name}.json",
    },
)
