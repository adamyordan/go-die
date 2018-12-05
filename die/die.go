package die

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

const (
	DIE_SHOWERRORS         = 0x00000001
	DIE_SHOWOPTIONS        = 0x00000002
	DIE_SHOWVERSION        = 0x00000004
	DIE_SHOWENTROPY        = 0x00000008
	DIE_SINGLELINEOUTPUT   = 0x00000010
	DIE_SHOWFILEFORMATONCE = 0x00000020
)

func DIEScan(fileName string, flags int) (res string, err error) {
	diedll := windows.NewLazyDLL("diedll.dll")
	procDIEScanW := diedll.NewProc("_DIE_scanW@16")

	ptrFilename, err := syscall.UTF16PtrFromString(fileName)
	if err != nil {
		return "", err
	}

	bufSizeIncr := uint32(1024)
	bufSize := bufSizeIncr

	for {
		buf := make([]uint8, bufSize)

		// int __declspec(dllexport) __stdcall DIE_scanW(wchar_t *pwszFileName,char *pszOutBuffer,int nOutBufferSize,unsigned int nFlags);
		ret, _, _ := procDIEScanW.Call(
			uintptr(unsafe.Pointer(ptrFilename)),
			uintptr(unsafe.Pointer(&buf[0])),
			unsafe.Sizeof(buf[0]) * uintptr(len(buf)),
			uintptr(flags),
		)

		if ret == uintptr(bufSize) {
			bufSize += bufSizeIncr
			continue
		}

		res = uint8ToString(buf)
		break
	}
	return res, nil
}

func uint8ToString(arr []uint8) (res string) {
	for _, c := range arr {
		if c == 0 {
			break
		}
		res += string(c)
	}
	return res
}
