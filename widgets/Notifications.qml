pragma ComponentBehavior: Bound

import QtQuick
import QtQuick.Effects
import QtQuick.Controls
import QtQuick.Layouts
import Quickshell
import Quickshell.Hyprland
import Quickshell.Services.Notifications

import "root:/config"
import "root:/modules"

Canvas {
	id: switcher

	width: parent.width
	height: 368

	readonly property bool hovered: content.hovered
	required property bool trigger

	anchors.bottom: parent.bottom
	anchors.right: parent.right
	Cute {
		id: content
		trigger: parent.trigger
		anchors.fill: parent
		botAttached: true
		topAttached: false
		botColor: Sakura.layerOverlay

		onOpened: {
			if (!NotifSvr.appNames.includes(NotifSvr.selectedApp)) {
				NotifSvr.selectedApp += "_";
			}
		}

		Rectangle {
			width: appList.width + 16
			height: parent.height
			anchors.left: parent.left
			anchors.top: parent.top
			color: Sakura.layerOverlay
			topLeftRadius: Opts.radius
		}
		Canvas {
			id: root
			clip: true
			height: content.contentBox.height - 16
			width: content.contentBox.width - 16
			anchors.centerIn: parent
			ScrollView {
				id: appList
				anchors.left: parent.left
				anchors.top: parent.top
				height: parent.height
				width: 28
				ListView {
					model: NotifSvr.appNames
					spacing: 4
					reuseItems: true
					delegate: NotificationGroup {
						required property string modelData
						appName: modelData
					}
				}
			}
			Canvas {
				anchors.centerIn: parent
				width: childrenRect.width
				height: childrenRect.height
				visible: NotifSvr.notifs.length == 0

				BtnIcon {
					id: emptyBell
					glyph: "notification-empty"
					size: 64
					anchors.horizontalCenter: parent.horizontalCenter
					fill: Sakura.textNormal
				}
				Text {
					id: emptyText
					text: "No Notifications"
					color: Sakura.textNormal
					font {
						pixelSize: 16
					}
					anchors {
						top: emptyBell.bottom
						topMargin: 8
						horizontalCenter: parent.horizontalCenter
					}
				}
				z: 5
			}
			Rectangle {
				color: Sakura.layerBase
				width: notifList.width
				height: header.height + 4
				anchors.top: header.top
				anchors.left: header.left
				z: 4
			}
			Text {
				id: header
				text: NotifSvr.selectedNotifs[0]?.appName || NotifSvr.entry?.name || ""
				color: Sakura.textNormal
				font {
					pixelSize: 16
				}
				anchors.top: parent.top
				anchors.left: appList.right
				anchors.leftMargin: 16
				z: 5
			}
			ScrollView {
				id: notifList
				anchors.left: appList.right
				anchors.leftMargin: 4
				anchors.top: header.bottom
				anchors.topMargin: 4
				height: parent.height - header.height - 4
				width: parent.width - this.anchors.leftMargin - appList.width
				ListView {
					model: NotifSvr.selectedNotifs
					spacing: 4
					reuseItems: true
					delegate: NotificationEntry {
						required property Notification modelData
						notif: modelData
					}
				}
				z: 1
			}
		}
	}
}
