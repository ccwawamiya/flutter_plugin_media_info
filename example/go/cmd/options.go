package main

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/plugins/path_provider"
	file_picker "github.com/miguelpruivo/flutter_file_picker/go"
)

const VendorName = "media_info"
const ApplicationName = "media_info"

var options = []flutter.Option{
	flutter.WindowInitialDimensions(800, 1280),
	flutter.AddPlugin(&file_picker.FilePickerPlugin{}),
	flutter.AddPlugin(&path_provider.PathProviderPlugin{
		VendorName:      VendorName,
		ApplicationName: ApplicationName,
	}),
}
