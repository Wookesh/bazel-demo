def list_generator(name, count, prefix = None, outname = None, visibility = None):
    if not outname:
        outname = name + ".txt"

    cmd = "$(location generator) --count %s" % count
    if prefix:
        cmd += " --prefix %s" % prefix

    native.genrule(
        name = name,
        outs = [outname],
        cmd = cmd + " > $@",
        tools = [":generator"],
        visibility = visibility,
    )
