import QtQuick
import Quickshell
import Quickshell.Services.Notifications
import QtQuick.Effects
import "root:/config"

PopupWindow {
	id: notifPopup
	implicitWidth: listFloating.width
	implicitHeight: listFloating.height
	ListView {
		id: listFloating
		model: NotifSvr.notifs.slice(-5)
		spacing: 8
		reuseItems: true
		delegate: Canvas {
			id: notifBox
			required property Notification modelData
			width: 384
			height: realNotif.height
			MultiEffect {
				source: realNotif
				anchors.fill: parent
				shadowBlur: 1.0
				shadowEnabled: true
				shadowScale: 1.0
				shadowColor: Sakura.paintPine
				shadowVerticalOffset: 0
				shadowHorizontalOffset: 0
				opacity: 0.5
			}
			NotificationEntry {
				id: realNotif
				notif: notifBox.modelData
				inFloat: true
			}
		}
	}
}
