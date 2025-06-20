pragma ComponentBehavior: Bound

import QtQml.Models
import QtQuick
import QtQuick.Effects
import Quickshell
import Quickshell.Wayland
import Quickshell.Services.Notifications
import "root:/widgets/spaces"
import "root:/widgets/taskman"
import "root:/widgets/notifs"
import "root:/widgets"
import "root:/config"

Variants {
	id: help
	model: Quickshell.screens

	PanelWindow {
		id: root
		required property ShellScreen modelData
		screen: modelData
		WlrLayershell.exclusionMode: ExclusionMode.Ignore
		WlrLayershell.layer: WlrLayer.Bottom
		color: "transparent"
		anchors {
			top: true
			bottom: true
			right: true
			left: true
		}

		PanelWindow {
			id: dock
			anchors {
				top: true
				bottom: true
				right: true
			}

			implicitWidth: 96
			screen: root.modelData

			color: Sakura.layerBase

			Workspaces {
				id: workspaces
				screen: root.modelData.name
			}
			Line {
				id: workspace_separator
				anchors.topMargin: 8
				anchors.top: workspaces.bottom
			}
			Tasks {

				anchors.topMargin: 8
				anchors.top: workspace_separator.bottom
			}
			Clock {
				id: clockWidget
			}
		}

		Instantiator {
			id: notifPopupList
			model: Opts.notifCount
			PopupWindow {
				id: notifPopup
				required property int modelData
				property int notifPad: Opts.radius * 3
				// visible: true
				anchor.window: root
				anchor.rect.x: root.width - dock.width - this.width - 8 + Opts.radius / 2 + notifPad / 2
				anchor.rect.y: previousY - notifBox.height
				anchor.rect.width: notifBox.height
				anchor.rect.height: notifBox.width
				property var previousElement: notifPopupList.objectAt(modelData - 1)
				property int previousY: previousElement ? previousElement.anchor.rect.y : root.height + Opts.radius / 2 - notifPad / 2 - 8
				implicitWidth: notifBox.width + notifPad
				implicitHeight: notifBox.height + notifPad
				color: "transparent"
				visible: modelData < NotifSvr.floatingNotifs.length
				Canvas {
					id: notifBox
					width: 384
					height: realNotif.height
					anchors.verticalCenter: parent.verticalCenter
					anchors.horizontalCenter: parent.horizontalCenter
					anchors.verticalCenterOffset: notifPopup.modelData == 0 ? notifPopup.notifPad / 2 : 0
					MultiEffect {
						source: realNotif
						anchors.fill: parent
						shadowBlur: 1.0
						shadowEnabled: true
						shadowScale: 1.0
						shadowColor: Sakura.paintPine
						shadowVerticalOffset: 0
						shadowHorizontalOffset: 0
						shadowOpacity: 0.5
					}
					NotificationEntry {
						id: realNotif
						notif: NotifSvr.floatingNotifs[NotifSvr.floatingNotifs.length - Math.min(NotifSvr.floatingNotifs.length, Opts.notifCount) + notifPopup.modelData] || NotifSvr.notifs[0]
						inFloat: true
					}
				}
			}
		}

		Corner {
			id: lowerCorner
			px: Opts.radius
			color: Sakura.layerBase
			botRight: true
			anchors.right: parent.right
			anchors.rightMargin: dock.width
			anchors.bottom: parent.bottom
		}

		Corner {
			id: upperCorner
			px: Opts.radius
			color: lowerCorner.color
			topRight: true
			anchors.right: lowerCorner.anchors.right
			anchors.rightMargin: lowerCorner.anchors.rightMargin
			anchors.top: parent.top
		}

		PanelWindow {
			id: widgets
			anchors {
				top: true
				bottom: true
				right: true
			}

			implicitWidth: 512

			WlrLayershell.exclusionMode: ExclusionMode.Ignore
			color: "transparent"
			screen: root.modelData
			aboveWindows: this.shown
			readonly property bool shown: notifWidget.hovered
			margins.right: dock.implicitWidth

			Behavior on aboveWindows {
				NumberAnimation {
					duration: widgets.shown ? Opts.aniWidget.duration + 50 : 0
				}
			}

			onShownChanged: NotifSvr.floating = []

			Notifications {
				id: notifWidget
				trigger: clockWidget.area_.containsMouse
			}
		}
	}
}
