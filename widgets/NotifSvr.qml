pragma Singleton
pragma ComponentBehavior: Bound

import QtQuick
import Quickshell
import Quickshell.Services.Notifications

Singleton {
	id: root

	readonly property list<Notification> notifs: server.trackedNotifications.values
	readonly property list<string> appNames: notifs.sort((a, b) => b.id - a.id).map(e => e.desktopEntry).filter((e, i, s) => s.indexOf(e) == i)
	readonly property list<Notification> selectedNotifs: notifs.filter(e => e.desktopEntry == selectedApp)
	readonly property DesktopEntry entry: DesktopEntries.byId(selectedApp)
	property string selectedApp: ""
	NotificationServer {
		id: server
		actionIconsSupported: true
		actionsSupported: true
		bodyHyperlinksSupported: true
		bodyImagesSupported: true
		bodyMarkupSupported: true
		bodySupported: true
		imageSupported: true
		keepOnReload: true
		persistenceSupported: true
		onNotification: function (n) {
			n.tracked = true;
		}
	}
	function selectAvailApp() {
		if (!appNames.includes(selectedApp) && appNames.length > 0) {
			selectedApp = appNames[0];
		}
	}
	onSelectedAppChanged: {
		root.selectAvailApp();
	}
}
