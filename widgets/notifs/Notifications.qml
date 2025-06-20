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
			height: content.contentBox.height
			width: content.contentBox.width - 16
			anchors.centerIn: parent
			Canvas {
				id: pad
				height: 8
				width: parent.width
				anchors.top: parent.top
				anchors.left: parent.left
			}
			ScrollView {
				id: appList
				anchors.left: parent.left
				anchors.top: pad.bottom
				height: parent.height
				width: 28
				ListView {
					model: NotifSvr.appNames
					spacing: 8
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
					cursor: Qt.ArrowCursor
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
				id: headerBg
				color: Sakura.layerBase
				width: notifList.width
				height: header.height + 8 + pad.height
				anchors.top: parent.top
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
				anchors.top: pad.bottom
				anchors.left: appList.right
				anchors.leftMargin: 16
				z: 5
				visible: NotifSvr.notifs.length > 0
			}
			BtnWithIcon {
				btnSize: 24
				glyph: "edit-clear-history"
				anchors.right: notifList.right
				anchors.rightMargin: (this.btnSize - this.iconSize) / 2
				anchors.verticalCenter: header.verticalCenter
				z: 5
				radius: 4
				onClick: NotifSvr.selectedNotifs.slice().forEach(e => e.dismiss())
				visible: NotifSvr.notifs.length > 0
			}
			ScrollView {
				id: notifList
				anchors.left: appList.right
				anchors.leftMargin: 4
				anchors.top: header.bottom
				anchors.topMargin: 8
				height: parent.height - headerBg.height - 8
				width: parent.width - this.anchors.leftMargin - appList.width
				z: 1
				ListView {
					id: listNotifs
					model: NotifSvr.selectedNotifs
					spacing: 8
					reuseItems: true
					delegate: NotificationEntry {
						required property Notification modelData
						notif: modelData
					}
				}
			}
		}
	}
}
