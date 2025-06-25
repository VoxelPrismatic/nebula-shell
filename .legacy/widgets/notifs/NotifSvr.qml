pragma Singleton
pragma ComponentBehavior: Bound

import QtQuick
import Quickshell
import Quickshell.Services.Notifications

Singleton {
	id: root

	readonly property list<Notification> notifs: server.trackedNotifications.values
	readonly property list<string> appNames: notifs.map(e => e.desktopEntry).filter((e, i, s) => s.lastIndexOf(e) == i).reverse()
	readonly property list<Notification> selectedNotifs: notifs.filter(e => e.desktopEntry == selectedApp).reverse()
	readonly property list<Notification> floatingNotifs: notifs.filter(e => floating[e.id] == true)
	readonly property DesktopEntry entry: DesktopEntries.byId(selectedApp)
	property list<bool> opened: []
	property list<date> timestamps: []
	property list<bool> floating: []
	property list<real> widths: []
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
			root.timestamps[n.id] = new Date();
			if (!n.lastGeneration) {
				root.floating[n.id] = true;
			}
			root.selectAvailApp();
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
