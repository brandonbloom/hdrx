# HDRX: Headers Refined

This repository contains the specification and a Go-based reference
implementation of a modern take on the headers formats of traditional network
text protocols, such as HTTP, NNTP, SMTP, and MIME.

```hdrx
title: Example Document
created: 2024-01-26
synposis: Example document showcasing the hdrx format.
author {
  name: Brandon Bloom
  github: https://github.com/brandonbloom/
}

Headers documents are reminiscent of an HTTP response,
but with some convenient features such as multi-line
values and support for nesting in values.
```

## Format Specification

### Structure

Headers appear at the top of the file, followed by optional content.  When
content is present, a single blank line separates it from the headers section.

```ebnf
file  ::=  headers (LF content)?
```

Note that both headers and content are optional, and so the format permits
headers without content, content without headers, and completely emtpy files.

In the case of headerless content, a `# No Headers.` comment is conventional.

Headers are encoded in UTF-8, but the content section is an arbitrary octet
stream.

### Keys and Headers

Each header in the headers section begins with a key and is terminated by a
newline.  Headers are composed of a key, separator, and value.

Keys begin with an ascii letter and may contain decimal digits, ascii letters,
and the hyphen (`-`) character. Keys may not end with a hyphen. All other
characters are reserved for future use; as the format stablizes, the set of
allowed characters will be expanded.

```ebnf
headers  ::=  header*
header   ::=  key separator value LF

key  ::=  [a-z][a-zA-Z0-9-]*  (* May not end with "-" *)
```

There are two separators for varying value syntax behavior:

- The colon and space (`: `) separator preceeds a value.
- A space and open curly (` {` separator begins a block, whose content is the
  value.

Headers containing these separators are colloquially known as "line headers"
and "block headers" respectively.  However, this pertains to their common
usage, as both styles support single-line and multi-line usages.

```ebnf
separator  ::=  ": " / " {"
```

### Values

Values are octet streams containing more-or-less arbitrary content. However,
they adhere to three special rules pertaining to:

1. line terminators
2. nested brackets (`{` and `}`)
3. and leading whitespace

In the absence of brackets, a value is terminated by a newline.  In the
confines of brackets, newlines are treated literally. Nested brackets must be
balanced.

```hdrx
single-line-json: {"works": "fine"}
multi-line-json: {
  "works": "just as well"
}
javascript {
  function double(x) {
    return x * 2;
  }
}
```

Brackets are literal within values, but note that the outermost brackets
signaling block header syntax is not part of the value.

For both line and block headers, values are stripped of leading and trailing
whitespace from the first and last lines respectively.

Within the outermost block's brackets, leading whitespace of each line is
trimmed with respect to the whitespace on the first non-blank line. In the
example above, `function` is preceeded by a two-space indent, which is also
trimmed of each of the subsequent lines. If the prefix does not match exactly,
it is preserved literally; implementations may emit a warning.

Block headers end with one or more closing brackets on a single,
otherwise-blank line.  This line may be the same line that begun the header,
and so an inline, empty block is valid. Any other content on the line should
abort parsing with an error.

Blank lines are allowed within braces and do not signal the separation between
headers and content.

#### Escaping

The tilde (`~`) character is used to escape curly braces, allowing for
unbalanced braces. `~{` and `~}` produce literal `{` and `}` respectively.
A tilde preceeding any other character is not treated specially and does
not need to be escaped.

To produce a literal `~{` or `~}`, the tilde character can be repeated. This
escaping rule works for any number of tildes, escaping one fewer tildes and
the subsequent bracket.

### Whitespace

"Whitespace" refers to Unicode's "White Space" property; in the Latin-1 space
this is: tab, line feed, vertical tab, form feed, carriage return, space,
next line, next line, and non-breaking space.

Line terminators are expected to be Unix-style `LF` (`\n`), though
Windows-style `CR LN` are harmless because of trimming rules.

Trimming removes leading and trailing whitespace.

A blank line is one that is empty after trimming whitespace.

## Conventions

### File Extension

The "refined headers" format should use the file extension `.hdrx` and may be
preceeded by a file extension for the content format. For example
`README.md.hdrx` for a markdown file with headers on top.

### Case Insensitivity

Implementations should treat keys as case-preserving, but compare them
case-insensitively. Case may be normalized via lowercasing.

Casing within values should be preserved, delegating case sensitivity to
whatever format the value happens to be encoded with.

### Header Ordering

General hdrx processing tools should preserve the presence and ordering of
headers. Duplicate header heys are syntacticaly valid. The semantic meaning
of header order or duplicate keys is application-specific.

### Value Handling

Values are not expected to be any particular format, although nested hdrx
values are well supported. The only constraints on the values is that they
follow conventional whitespace flexibility.

### Chaining

Multiple headers documents can be chained together into a "linked list", where
each document's content is the next document. The enables the format to be
used for collections of headers.

```hdrx
name: Brandon
species: Human

name: Hudson
species: Dog
```

## Tool

TODO: Describe the `hdrx` tool.

## Go API

TODO: Describe the Go library API.
