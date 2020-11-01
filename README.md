# The Original

This is a trivial Go implementation of the `sponge` command from the [moreutils](https://joeyh.name/code/moreutils/) collection.

# Description

`sponge` is a command that

1. Soaks up the entirety of the data on its standard input (until EOF).
2. Wrings that data out into a destination file.

It is primarily useful along with a data transformation command that does not itself support intelligent "in-place" editing of files. For example, [jq](https://stedolan.github.io/jq/).

Here is an example use of `sponge` along with `jq` to do in-place formatting of JSON files.

    $ FILE=inventory.json

    $ jq . "$FILE" | sponge "$FILE"

# Why is a special command needed to do this?

Upon first glance, it seems like normal shell output redirection, i.e. `jq . "$FILE" >"$FILE"` should be sufficient to do this. But that will truncate the file to 0 bytes immediately and there will be nothing for the formatter command to read.

A more capable normal shell approach, `(rm -f "$FILE" ; jq . >"$FILE") <"$FILE"`, is [mentioned here](https://github.com/kkinnear/zprint/issues/159#issuecomment-720005601). However, that is still not robust if the formatter command itself fails, such as with a parsing error.

# Installation

With the [Go compiler](https://golang.org/dl/) installed, run:

`go get github.com/frou/sponge`

The `sponge` command will now be compiled and located in `$GOPATH/bin`
