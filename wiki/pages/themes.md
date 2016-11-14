Themes are what controls how your QuickWiki looks and what the final html files
will be structured like. You can download new themes and start using them by just
changing a configuration variable. They are also easy to customize and develop
on your own.

# How to configure theme?

In your `config.toml`, you can set the value `theme.name` to whatever theme
you want to use. Example to use the default theme:

```
[theme]
  name = "simpleblue"
```

# List of themes

## simpleblue

simpleblue is the default theme. It looks something like this:

![screenshot of simpleblue](../media/simpleblue-theme.png)

# How themes are loaded

When compilation and writing of posts happen, theme is loaded based on the name
variable in the configuration. It tries to find `./themes/:name` and load `style.css` and `template.html` from there. At that point magic variables are being replaced. Probably move these to golang templates
in the future.

# How to develop new themes

Take a look at how `simpleblue` was developed for now
