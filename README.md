# waybar-issues

Get your open issues and merge/pull requests into waybar.

## Api tokens

### GitLab

* Create a api token with scope `api`.

### GitHub

* Create a **Personal access token** with scope `repo` [here](https://github.com/settings/tokens).

## Waybar

Add a custom module to your waybar config eg.

```json
    "custom/issues": {
        "interval": 60,
        "max-length": 50,
        "format": "{} ",
        "return-type": "json",
        "exec": "waybar-issues",
        "escape": true
    }
```

## Installation

Use the [AUR](https://aur.archlinux.org/packages/waybar-issues) package or follow the build instructions.

## Build

Create and save the main configuration file in your `$HOME` as `.waybar-issues` (check the provided`config.example`).

* Build with `make`.
* Save the binary in your `waybar/modules` folder.
