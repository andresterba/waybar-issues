# waybar-issues

Get your open issues into waybar.

## Api tokens

### GitLab

* Create a api token with scope `api`.

### GitHub

* Create a **Personal access token** with scope `repo` [here](https://github.com/settings/tokens).

## Waybar

Add a custom module to your waybar config eg.

```
    "custom/issues": {
        "interval": 60,
        "max-length": 50,
        "format": "{} ",
        "return-type": "json",
        "exec": "~/.config/waybar/modules/waybar-issues",
        "escape": true
    }
```

## Build

Create and save a configuration file in your `$HOME` (check `config.examle`).

* Build with `make`.
* Save the binary in your `waybar/modules` folder.
