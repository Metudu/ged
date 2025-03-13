# ged
ged is a CLI tool helps you manage app launcher icons' visibility.

## Installation
Download the binary file from the releases and copy to a directory which is in PATH.

## How it works?
ged watches the files under `.local/share/applications` directory, fetches a specific file and adds or removes `NoDisplay = true` option depending on what user wants to do.

> [!WARNING]
> ged doesn't watch the `/usr/share/applications` directory, which may cause you to not find the application you want to change its visibility.

> [!NOTE]
> ged only supports changing visibility at the moment. New features will be added in the future releases.

## Usage
Run `ged` in terminal, select the application and select *YES* if you want to set it visible or invisibile.