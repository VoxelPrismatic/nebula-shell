import QtQuick
import QtQuick.Effects
import Quickshell
import Quickshell.Services.Notifications
import "root:/config"
import "root:/modules"

Canvas {
	id: root
	required property Notification notif
	readonly property bool urgent: notif.urgency == NotificationUrgency.Critical
	readonly property DesktopEntry entry: DesktopEntries.byId(notif.desktopEntry)
	visible: notif != null

	width: parent.width
	height: closing ? 0 : childrenRect.height

	readonly property bool floating: NotifSvr.floating[notif.id] == true
	readonly property bool open: inFloat || NotifSvr.opened[notif.id] == true
	property bool inFloat: false
	property bool closing: false

	Rectangle {
		id: header
		width: parent.width - (root.inFloat ? Opts.radius : 16)
		anchors.top: parent.top
		anchors.topMargin: root.inFloat ? Opts.radius / 2 : 0
		anchors.horizontalCenter: parent.horizontalCenter
		anchors.horizontalCenterOffset: root.inFloat ? 0 : 4
		height: 24
		color: root.urgent ? Sakura.paintLove : Sakura.layerOverlay
		topLeftRadius: 4
		topRightRadius: 4
		bottomLeftRadius: root.open ? 0 : 4
		bottomRightRadius: root.open ? 0 : 4
	}

	BtnWithIcon {
		id: openBtn
		btnSize: 24
		glyph: root.inFloat ? root.entry.icon : root.open ? "arrow-down" : "arrow-right"

		anchors.left: header.left
		anchors.top: header.top
		topLeftRadius: 4
		bottomLeftRadius: root.open ? 0 : 4
		onClick: NotifSvr.opened[root.notif.id] = !root.open
		fill: root.inFloat ? null : root.urgent ? Sakura.layerBase : Sakura.textNormal
	}
	BtnWithIcon {
		id: hideBtn
		btnSize: 24
		glyph: "collapse-all"
		anchors.right: delBtn.left
		anchors.rightMargin: -4
		anchors.top: header.top
		visible: root.inFloat && root.floating
		fill: root.urgent ? Sakura.layerBase : Sakura.textNormal
		onClick: NotifSvr.floating[root.notif.id] = false
	}
	BtnWithIcon {
		id: delBtn
		btnSize: 24
		glyph: "dialog-close"
		anchors.right: header.right
		anchors.top: header.top
		topRightRadius: 4
		bottomRightRadius: root.open ? 0 : 4
		fill: root.urgent ? Sakura.layerBase : Sakura.textNormal

		onClick: {
			root.notif.dismiss();
			if (!NotifSvr.appNames.includes(NotifSvr.selectedApp)) {
				NotifSvr.selectedApp += "_";
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
			left: openBtn.right
			leftMargin: 4
		}
	}
	Rectangle {
		id: content
		width: header.width
		opacity: root.open ? 1 : 0
		anchors {
			left: header.left
			top: header.bottom
		}
		height: root.open ? childrenRect.height : 0
		color: Sakura.layerSurface
		bottomLeftRadius: 4
		bottomRightRadius: 4

		Text {
			id: contentText
			width: header.width
			text: root.notif.body
			color: Sakura.textNormal
			clip: true
			padding: 8
			lineHeight: 1.5
			wrapMode: Text.Wrap
			horizontalAlignment: Text.AlignJustify
		}
		Behavior on height {
			NumberAnimation {
				duration: Opts.aniHover.duration / 2
				easing.type: Opts.aniHover.style
			}
		}
		Behavior on opacity {
			NumberAnimation {
				duration: Opts.aniHover.duration / 2
				easing.type: Opts.aniHover.style
			}
		}
	}
	Canvas {
		height: Opts.radius / 2
		width: header.width
		visible: root.inFloat
		anchors {
			left: header.left
			bottom: root.inFloat ? root.bottom : content.bottom
		}
	}
	Rectangle {
		id: timeoutBar
		width: 0
		height: 2
		anchors.top: header.bottom
		anchors.topMargin: -this.height / 2
		anchors.left: header.left
		color: Sakura.paintRose
		visible: root.inFloat
		onWidthChanged: NotifSvr.widths[root.notif.id] = width
	}
	Timer {
		id: floatTimeout
		running: timeoutBar.width > 0
		repeat: true
		interval: 10
		property real step: header.width * interval / Opts.notifTimeout
		onTriggered: {
			timeoutBar.width -= step;
			if (timeoutBar.width <= 0) {
				NotifSvr.floating[root.notif.id] = null;
			}
		}
	}
	onNotifChanged: {
		if (NotifSvr.widths[root.notif.id]) {
			timeoutBar.width = NotifSvr.widths[root.notif.id];
		} else {
			timeoutBar.width = header.width;
		}
	}
}
