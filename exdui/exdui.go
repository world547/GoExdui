package exdui

import (
	"syscall"
	"unsafe"
)

var (
	// Library
	libexdui    uintptr
	libkernel32 uintptr

	// Functions
	getModuleHandleA      uintptr
	exInit                uintptr
	exUnInit              uintptr
	exWndMsgLoop          uintptr
	exDUICreateFromLayout uintptr
	exDUIShowWindowEx     uintptr
)

func init() {
	// Library

	libexdui = MustLoadLibrary("libexdui.dll")
	libkernel32 = MustLoadLibrary("kernel32.dll")

	// Functions
	getModuleHandleA = MustGetProcAddress(libkernel32, "GetModuleHandleA")
	exInit = MustGetProcAddress(libexdui, "Ex_Init")
	exUnInit = MustGetProcAddress(libexdui, "Ex_UnInit")
	exWndMsgLoop = MustGetProcAddress(libexdui, "Ex_WndMsgLoop")
	exDUICreateFromLayout = MustGetProcAddress(libexdui, "Ex_DUICreateFromLayout")
	exDUIShowWindowEx = MustGetProcAddress(libexdui, "Ex_DUIShowWindowEx")

}

func GetModuleHandleA(lpModuleName uintptr) uintptr {
	ret, _, _ := syscall.Syscall(getModuleHandleA, 1,
		lpModuleName,
		0,
		0)
	return ret
}

// 初始化ExDirectUI引擎，初始化引擎内部的各项参数，加载默认主题、语言.
func ExInit(hInstance uintptr, dwGlobalFlags uintptr, hDefaultCursor uintptr, lpszDefaultClassName string, lpDefaultTheme []byte, dwDefaultThemeLen, lpDefaultI18N, dwDefaultI18NLen uintptr) bool {
	ret, _, _ := syscall.Syscall9(exInit, 8,
		hInstance,
		dwGlobalFlags,
		hDefaultCursor,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpszDefaultClassName))),
		uintptr(unsafe.Pointer(&lpDefaultTheme[0])),
		dwDefaultThemeLen,
		lpDefaultI18N,
		dwDefaultI18NLen,
		0)

	return ret != 0
}

// 反初始化ExDirectUI引擎，释放内部使用的各类资源.
func ExUnInit() bool {
	defer syscall.FreeLibrary(syscall.Handle(libexdui))
	defer syscall.FreeLibrary(syscall.Handle(libkernel32))

	ret, _, _ := syscall.Syscall(exUnInit, 0,
		0,
		0,
		0)
	return ret != 0
}

// 窗口消息循环.
func ExWndMsgLoop() uintptr {
	ret, _, _ := syscall.Syscall(exWndMsgLoop, 0,
		0,
		0,
		0)
	return ret
}

// 通过布局文件创建一个界面.
func ExDUICreateFromLayout(hParent uintptr, hTheme uintptr, lpData []byte, dwLen uintptr, hWnd, hExDui *int32) bool {
	ret, _, _ := syscall.Syscall6(exDUICreateFromLayout, 6,
		hParent,
		hTheme,
		uintptr(unsafe.Pointer(&lpData[0])),
		dwLen,
		uintptr(unsafe.Pointer(hWnd)),
		uintptr(unsafe.Pointer(hExDui)))
	return ret != 0
}

// 显示一个已被绑定的ExDUI界面（已被绑定的窗口请使用该函数显示）.
func ExDUIShowWindowEx(hExDui, nCmdShow, dwTimer, dwFrames, dwFlags, uEasing, wParam, lParam uintptr) bool {
	ret, _, _ := syscall.Syscall9(exDUIShowWindowEx, 8,
		hExDui,
		nCmdShow,
		dwTimer,
		dwFrames,
		dwFlags,
		uEasing,
		wParam,
		lParam,
		0)
	return ret != 0
}
