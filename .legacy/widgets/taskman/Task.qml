import QtQuick
import QtQuick.Controls
import Quickshell
import Quickshell.Wayland
import Quickshell.Hyprland
import "root:/config"
import "root:/widgets/spaces"

Rectangle {
	id: root
	required property bool isMax
	required property Toplevel app

	property DesktopEntry entry: DesktopEntries.byId(app.appId)

	readonly property Tiles.Tile tile: Tiles.tiles.find(t => t.wmClass == root.app.appId || t.title == root.app.appId)
	visible: CurrentWs.hovered == -1 || Tiles.tiles.find(t => (t.wmClass == root.app.appId || t.title == root.app.appId) && t.workspace.id == CurrentWs.hovered) ? 1 : 0

	width: Opts.workspace.cellSize - 4
	height: this.width
	property real confidence: 0

	color: Sakura.layerBase
	radius: 4

	Rectangle {
		id: bg
		anchors.fill: parent
		color: Sakura.paintRose
		opacity: area.containsMouse ? 0 : 0.2
		radius: parent.radius
	}

	function jaccardSimilarity(str1: string, str2: string): double {
		var words1 = str1.toLowerCase().split(/\W+/g);
		var words2 = str2.toLowerCase().split(/\W+/g);
		var set1 = new Set(words1);
		var set2 = new Set(words2);
		var intersection = new Set([...set1].filter(x => set2.has(x)));
		var union = new Set([...set1, ...set2]);
		if (words1[0] != words2[0]) {
			return intersection.size / union.size / 2;
		}
		return intersection.size / union.size;
	}

	function locateIcon() {
		var bestEntry = null;
		var bestSimilarity = 0;
		const against = {
			"name": [tile?.initialClass, tile?.initialTitle, tile?.title]
		};
		for (var entry of DesktopEntries.applications.values) {
			for (var key in against) {
				if (!entry[key])
					continue;
				for (var val of against[key]) {
					var similarity = jaccardSimilarity(val, entry[key]);
					if (similarity > bestSimilarity) {
						bestSimilarity = similarity;
						bestEntry = entry;
					}
				}
			}
		}
		if (bestEntry) {
			var iconFromEntry = Quickshell.iconPath(bestEntry.icon, false);
			if (iconFromEntry !== "") {
				root.entry = bestEntry;
				root.confidence = bestSimilarity;
				return bestEntry.icon;
			}
		}
		return "application-x-executable";
	}

	Button {
		id: ico
		width: 24
		height: ico.width
		anchors.centerIn: parent
		icon {
			name: root.entry?.icon || root.locateIcon()
			width: ico.height
			height: ico.width
		}
		flat: true
		padding: 0
	}

	Rectangle {
		id: hover
		anchors.fill: parent
		color: Sakura.layerInverse
		opacity: area.containsMouse ? 0.2 : 0
		radius: parent.radius
	}

	ToolTip {}

	MouseArea {
		id: area
		anchors.fill: parent
		cursorShape: Qt.PointingHandCursor
		hoverEnabled: true
		property int idx: 0
		onHoveredChanged: function () {
			if (area.containsMouse) {
				let title = root.entry?.name || root.tile.initialClass || root.app.title;
				// if (root.confidence > 0) {
				// title = Math.round(root.confidence * 100) + "%: " + title;
				// }
				ToolTip.show(title);
			} else {
				ToolTip.hide();
			}
		}
		onClicked: {
			const siblings = Tiles.tiles.filter(t => t.wmClass == root.app.appId);
			this.idx++;
			if (this.idx >= siblings.length) {
				this.idx = 0;
			}
			const sibling = siblings[idx];
			ToolTip.show(`${idx + 1}/${siblings.length}`);
			Hyprland.dispatch(`focuswindow address:${sibling.address}`);
		}
		acceptedButtons: Qt.LeftButton | Qt.RightButton
	}
}
