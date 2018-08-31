# Webview

A web based Desktop Application starter using the webview library.

## Requirements

You will need [Go](https://golang.org/) and [npm](https://www.npmjs.com/) installed to work with this repo.

## Developing (Windows)

From the `web` folder:

Install deps with `npm install` and then start the Vue dev server with `npm run serve`.

From the root folder run `go build -mod=vendor -tags=dev -ldflags="-s -w -H windowsgui" -o webview.exe`.

Now when you run `webview.exe` it will display the Vue development output.

(note: Building with `-mod=vendor` is required as I needed to change a file in the webview import.)