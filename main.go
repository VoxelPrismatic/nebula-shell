package main

import (
	"fmt"
	"nebula-shell/shell/dock"
	"nebula-shell/svc/layershell"
	"os"

	"github.com/mappu/miqt/qt6"
)

import "C"

func main() {
	layershell.UseLayerShell()
	qt6.NewQApplication(os.Args)
	defer qt6.QApplication_Exec()
	for _, screen := range qt6.QGuiApplication_Screens() {
		fmt.Println(screen.Name())
		dock.NewDock(screen)
	}
}
