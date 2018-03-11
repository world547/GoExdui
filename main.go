package main

import (
	"GoExdui/exdui"
	"io/ioutil"
)

var (
	hWnd   int32
	hExDui int32
)

func main() {
	hInstance := exdui.GetModuleHandleA(0)
	themeFile, _ := ioutil.ReadFile("./Default.ext")
	themelen := len(themeFile)
	exdui.ExInit(hInstance, exdui.EXGF_DPI_ENABLE|exdui.EXGF_RENDER_METHOD_D2D, 0, "Godemo", themeFile, uintptr(themelen), 0, 0)

	layoutFile, _ := ioutil.ReadFile("./main.xml")
	layoutlen := len(layoutFile)
	exdui.ExDUICreateFromLayout(0, 0, layoutFile, uintptr(layoutlen), &hWnd, &hExDui)

	exdui.ExDUIShowWindowEx(uintptr(hExDui), 1, 0, 0, 0, 0, 0, 0)

	exdui.ExWndMsgLoop()

	exdui.ExUnInit()

}
