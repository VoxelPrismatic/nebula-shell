import QtQuick
import QtQuick.Controls
import Quickshell
import Quickshell.Services.Notifications

import "root:/config"

Rectangle {
	id: root
	required property string appName
	readonly property DesktopEntry entry: DesktopEntries.byId(appName)
	readonly property int notifCount: NotifSvr.notifs.filter(e => e.desktopEntry == appName).length
	width: parent.width
	height: parent.width
	color: NotifSvr.selectedApp == this.appName ? Sakura.layerInverse : area.hovered ? Sakura.hlMed : Sakura.layerSurface
	radius: 4

	BtnIcon {
		id: ico
		glyph: parent.entry.icon
		anchors {
			top: parent.top
			topMargin: 2
			left: parent.left
			leftMargin: 2
		}
	}
	MouseArea {
		id: area
		anchors.fill: parent
		hoverEnabled: true
		onClicked: NotifSvr.selectedApp = root.appName
		cursorShape: Qt.PointingHandCursor
	}
	Rectangle {
		id: ping
		anchors {
			bottom: parent.bottom
			right: parent.right
		}
		width: 12
		height: 12
		radius: 2
		color: Sakura.paintRose
	}
	Rectangle {
		visible: area.containsMouse
		anchors.fill: parent
		opacity: 0.2
		color: Sakura.layerInverse
		radius: parent.radius
	}
	Text {
		anchors.centerIn: ping
		text: root.notifCount

		font {
			pixelSize: 8
			bold: true
		}
		color: Sakura.layerBase
	}
}
