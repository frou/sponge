# The Original

This is a trivial Go implementation of the `sponge` command from the [moreutils](https://joeyh.name/code/moreutils/) collection.

# Description

`sponge` is a command that

1. Soaks up the entirety of the data on its standard input (until EOF).
2. Wrings that data out into a destination file.

It is primarily useful along with a data transformation command that does not itself support intelligent "in-place" editing of files. For example, [jq](https://stedolan.github.io/jq/).

Here is an example use of `sponge` along with `jq` to make an automated formatter for JSON files:

    $ FILE=inventory.json
    $ jq . $FILE | sponge $FILE

Upon first glance, it seems like normal shell output redirection (`jq . $FILE >$FILE`) would be sufficient for this, but that would truncate the file to 0 bytes immediately and there would be nothing for `jq` to read.

# Installation

With the [Go compiler](https://golang.org/dl/) installed, run:

`go get github.com/frou/sponge`

The `sponge` command will now be compiled and located in `$GOPATH/bin`
