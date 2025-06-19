import QtQuick
import Quickshell
import Quickshell.Services.Notifications
import "root:/config"

Canvas {
	id: root
	required property Notification notif

	width: parent.width
	height: childrenRect.height

	property bool open: false

	Rectangle {
		id: header
		width: parent.width - 16
		anchors.horizontalCenter: parent.horizontalCenter
		anchors.horizontalCenterOffset: 4
		height: 24
		color: Sakura.layerOverlay
		topLeftRadius: 4
		topRightRadius: 4
		bottomLeftRadius: root.open ? 0 : 4
		bottomRightRadius: root.open ? 0 : 4
	}

	Rectangle {
		id: openRect
		width: 24
		height: 24
		anchors.left: header.left
		anchors.top: header.top
		topLeftRadius: 4
		bottomLeftRadius: root.open ? 0 : 4

		color: openBtn.renderColor

		BtnIcon {
			id: openBtn
			size: 16
			glyph: root.open ? "arrow-down" : "arrow-right"

			anchors.centerIn: parent
			pressColor: Sakura.hlHigh
			hoverColor: Sakura.hlMed
			defaultColor: Sakura.layerOverlay
			onClick: root.open = !root.open
		}
	}
	Text {
		text: root.notif.summary
		color: Sakura.textNormal
		font {
			pixelSize: 12
		}
		anchors {
			verticalCenter: header.verticalCenter
			left: openRect.right
			leftMargin: 4
		}
	}
	Rectangle {
		visible: root.open
		width: header.width
		anchors {
			left: header.left
			top: header.bottom
		}
		height: root.open ? childrenRect.height : 0
		color: Sakura.layerSurface

		Text {
			width: header.width
			height: 64
			text: root.notif.body
			color: Sakura.textNormal
			clip: true
			padding: 8
			lineHeight: 1.5
			wrapMode: Qt.TextWrapAnywhere
		}
	}
}
