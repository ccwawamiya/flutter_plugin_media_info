# media_info

This Go package implements the host-side of the Flutter [media_info](https://github.com/ccwawamiya/flutter_plugin_media_info) plugin.

## Usage

Import as:

```go
import media_info "github.com/ccwawamiya/flutter_plugin_media_info/go"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&media_info.MediaInfoPlugin{}),
```
