Yeah...

## Research

Where to put the plugin system? Need certain hooks...

### BUILD_WIKI

* BEFORE
* AFTER

### BUILD_FILE

* Before
* After
* BeforeWrite
* AfterWrite

## Examples

### Markdown

Parses .md files and saves .html files. Responsible for finding the files, transforming
the content and saving them again.

### AutoLink

Should automatically replace words with links if they have a corresponding page. 

If the word is multiple words, page will be matched like this:

```
Automagic Words === automagic-words.md
```

### SimpleLinks

Links by doing brackets around words.

`[A PagA Pagee]`

### DataTables

Set values for pages that can read elsewhere and inline

### Categories

Depends on DataTables

### Media


