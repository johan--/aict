//go:build windows

package df

func getFsInfo(device, mount, fstype string) (FsEntry, error) {
	return FsEntry{
		Device: device,
		Mount:  mount,
		Type:   fstype,
	}, nil
}
