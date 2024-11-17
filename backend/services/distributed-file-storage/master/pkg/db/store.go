package db

// SaveFileMetadata stores file metadata and returns a new file ID
func SaveFileMetadata(filename string) int {
	fileIDLock.Lock()
	defer fileIDLock.Unlock()
	fileID := fileIDSeq
	fileIDSeq++
	files[fileID] = filename
	return fileID
}

// GetFileMetadata retrieves the filename for a given file ID
func GetFileMetadata(fileID int) string {
	return files[fileID]
}
