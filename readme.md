# QuickWiki

> Generate simple static wikis

QuickWiki is a simple wiki, driven by a simple binary that you run yourself
and generates a static website of it. Together with markdown, themes and plugins,
QuickWiki offers a fast, easy and extensible way of running your own, personal wiki.


## Table of Contents

- [Background](#background)
- [Installation](#installation)
- [Usage](#usage)
- [API](#api)
- [Contribute](#contribute)
- [License](#license)

## Background

Developed for my own need to have a portable wiki that I can access offline
and doesn't require a server. Previously I've used Mediawiki and Dokuwiki which
worked ok but it's time to have something better.

## Installation

You can binaries over here: https://github.com/VictorBjelkholm/quickwiki/releases

Or, install it with go if you have that installed:

```
go get -u github.com/victorbjelkholm/quickwiki
```

## Usage

QuickWiki is composed of a few commands.

### `build`

`quickwiki build` builds your wiki and puts it by default in a `./public/` directory.

You can change which directory to use as input as the first argument

`quickwiki build /path/to/mywiki`

### `server`

`quickwiki server` is a quickhand command meant for while you're writing and you
want to see a preview of your wiki. It will do the following:

* Start a server on :8085
* Build an initial version of your wiki
* Listen for changes to `pages/`, `media/` and your loaded theme
* On any change, build the entire wiki again
* Reload any page loaded in the browser

### `new`

`quickwiki new <name>` initializes a new wiki in a directory specified as the first
argument with the default settings and one demo page.

### `version`

`quickwiki version` simply prints the current version of your quickwiki binary,
so you can verify from source you're actually running the right binary. Good stuff!

## API

Currently, the API is limited except for what the CLI offers. Might change in the
future. If you're interested in extending QuickWiki, check out the pages about
{{themes}} and/or {{themes}}.

## Contribute

Contributions are very welcome in any area! Things that notably needs improvement:

* Error handling
* More default themes to chose from
* Plugins that can be activated
* Performance testing and improvements
* General code style (I'm not really used to Go yet...)
* Better documentation! Quickwiki documentation is driven by QuickWiki itself,
see the source for our wiki here: {{link_to_wiki_source}}
* Finding bugs! There is bugs, we just have to find and fix them!

Remember to check out our contribute.md file for instructions on how to get a
development environment setup for QuickWiki {{link_to_contribute_md}}

If you're looking for features to implement, we have todo.md file that contains
planned features that are wanted. Open an issue if you have any questions about them.

## License

The MIT License (MIT)

Copyright (c) 2016 Victor Bjelkholm

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
