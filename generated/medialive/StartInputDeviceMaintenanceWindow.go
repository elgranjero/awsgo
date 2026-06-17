package medialive

// StartInputDeviceMaintenanceWindow is generated as a reference stub.
// Executable command wiring lives under cmd/medialive.go.
//
// Start a maintenance window for the specified input device. Starting a
// maintenance window will give the device up to two hours to install software. If
// the device was streaming prior to the maintenance, it will resume streaming when
// the software is fully installed. Devices automatically install updates while
// they are powered on and their MediaLive channels are stopped. A maintenance
// window allows you to update a device without having to stop MediaLive channels
// that use the device. The device must remain powered on and connected to the
// internet for the duration of the maintenance.
