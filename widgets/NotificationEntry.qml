import QtQuick
import Quickshell
import Quickshell.Services.Notifications
import "root:/config"

Canvas {
	id: root
	required property Notification notif
	readonly property bool urgent: notif.urgency == NotificationUrgency.Critical

	width: parent.width
	height: childrenRect.height

	property bool open: NotifSvr.opened[notif.id] == true

	Rectangle {
		id: header
		width: parent.width - 16
		anchors.horizontalCenter: parent.horizontalCenter
		anchors.horizontalCenterOffset: 4
		height: 24
		color: root.urgent ? Sakura.paintLove : Sakura.layerOverlay
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
			fill: root.urgent ? Sakura.layerBase : Sakura.textNormal

			anchors.centerIn: parent
			pressColor: Sakura.hlHigh
			hoverColor: Sakura.hlMed
			defaultColor: Sakura.layerOverlay
			onClick: NotifSvr.opened[root.notif.id] = !root.open
		}
	}
	Rectangle {
		id: delRect
		width: 24
		height: 24
		anchors.right: header.right
		anchors.top: header.top
		topRightRadius: 4
		bottomRightRadius: root.open ? 0 : 4

		color: delBtn.renderColor

		BtnIcon {
			id: delBtn
			size: 16
			glyph: "dialog-close"
			fill: root.urgent ? Sakura.layerBase : Sakura.textNormal

			anchors.centerIn: parent
			pressColor: Sakura.hlHigh
			hoverColor: Sakura.hlMed
			defaultColor: Sakura.layerOverlay
			onClick: {
				root.notif.dismiss();
				if (!NotifSvr.appNames.includes(NotifSvr.selectedApp)) {
					NotifSvr.selectedApp += "_";
				}
			}
		}
	}
	Text {
		text: root.notif.summary
		color: root.urgent ? Sakura.layerBase : Sakura.textNormal
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
		bottomLeftRadius: 4
		bottomRightRadius: 4

		Text {
			width: header.width
			text: root.notif.body
			color: Sakura.textNormal
			clip: true
			padding: 8
			lineHeight: 1.5
			wrapMode: Text.Wrap
			horizontalAlignment: Text.AlignJustify
		}
	}
}
