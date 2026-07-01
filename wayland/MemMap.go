package wayland

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

type MemMapInfo struct {
	Bytes          []byte
	Addr           unsafe.Pointer
	Size           uint64
	FileDescriptor int
	UnMapped       bool
}

func NewMemMapInfo(fd int, size uint64) (MemMapInfo, error) {
	data, err := unix.Mmap(fd, 0, int(size), unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return MemMapInfo{
			FileDescriptor: fd,
			Size:           size,
			UnMapped:       true,
		}, fmt.Errorf("failed to mmap fd %d: %w", fd, err)
	}

	info := MemMapInfo{
		Addr:           unsafe.Pointer(unsafe.SliceData(data)),
		Size:           uint64(len(data)),
		FileDescriptor: fd,
		UnMapped:       false,
		Bytes:          data,
	}
	return info, nil
}

func (m *MemMapInfo) Unmap() {
	if m.UnMapped {
		return
	}
	if m.Bytes != nil {
		unix.Munmap(m.Bytes)
	}
	m.UnMapped = true
	m.Bytes = nil
}
