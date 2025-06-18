import QtQuick
import Quickshell
import Quickshell.Services.Notifications
import "root:/config"

Canvas {
	required property Notification notif

	width: 192
	height: 128

	Rectangle {
		width: parent.width
		height: parent.height
		color: Sakura.layerSurface
	}
}
