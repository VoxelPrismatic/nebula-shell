pragma Singleton

import Quickshell
import Quickshell.Io
import Quickshell.Hyprland
import QtQuick

Singleton {
	id: root

	readonly property list<Client> tiles: []
	Process {
		id: getTiles

		command: ["hyprctl", "-j", "clients"]
		stdout: StdioCollector {
			onStreamFinished: {
				const newTiles = JSON.parse(text);
				const globalTiles = root.tiles;
				const destroyed = globalTiles.filter(globalTile => !newTiles.find(newTile => newTile.address === globalTile.address));
				for (const client of destroyed) {
					globalTiles.splice(globalTiles.indexOf(client), 1).forEach(c => c.destroy());
				}

				for (const tile of newTiles) {
					const match = globalTiles.find(globalTile => globalTile.address === tile.address);
					if (match) {
						match.lastIpcObject = tile;
					} else {
						globalTiles.push(refresh.createObject(root, {
							lastIpcObject: tile
						}));
					}
				}
			}
		}
	}
	component Client: QtObject {
		required property var lastIpcObject
		readonly property string address: lastIpcObject.address
		readonly property string wmClass: lastIpcObject.class
		readonly property string title: lastIpcObject.title
		readonly property string initialClass: lastIpcObject.initialClass
		readonly property string initialTitle: lastIpcObject.initialTitle
		readonly property int x: lastIpcObject.at[0]
		readonly property int y: lastIpcObject.at[1]
		readonly property int width: lastIpcObject.size[0]
		readonly property int height: lastIpcObject.size[1]
		readonly property HyprlandWorkspace workspace: Hyprland.workspaces.values.find(w => w.id === lastIpcObject.workspace.id) ?? null
		readonly property bool floating: lastIpcObject.floating
		readonly property bool fullscreen: lastIpcObject.fullscreen
		readonly property int pid: lastIpcObject.pid
		readonly property int focusHistoryId: lastIpcObject.focusHistoryID
	}
	function refresh() {
		Hyprland.refreshMonitors();
		Hyprland.refreshWorkspaces();
		getTiles.running = true;
	}

	Component {
		id: refresh

		Client {}
	}
	Component.onCompleted: root.refresh()
	Connections {
		target: Hyprland

		function onRawEvent(event: HyprlandEvent): void {
			if (!event.name.endsWith("v2"))
				root.refresh();
		}
	}
}
