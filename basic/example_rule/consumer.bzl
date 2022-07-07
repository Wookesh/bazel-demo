script_template = """\
#!/bin/bash
cat {file}
"""

def _consumer(ctx):
    file = ctx.file.src

    script = ctx.actions.declare_file("%s-consumer" % ctx.label.name)

    script_content = script_template.format(file = file.short_path)

    ctx.actions.write(script, script_content, is_executable = True)

    return [DefaultInfo(executable = script, runfiles = ctx.runfiles(files = [file]))]

consumer = rule(
    implementation = _consumer,
    attrs = {
        "src": attr.label(mandatory = True, allow_single_file = True),
    },
    executable = True,
)
