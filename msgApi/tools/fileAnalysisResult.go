package tools

import "path/filepath"

/**
 * @author Threema GmbH
 * @copyright Copyright (c) 2015-2016 Threema GmbH
 */

type FileAnalysisResult struct {
	/**
	 * @var string
	 */
	mimeType string

	/**
	 * @var int
	 */
	size int

	/**
	 * @var string
	 */
	path string
}

/**
 * @param string $mimeType
 * @param int $size
 * @param string $path
 */
func NewFileAnalysisResult(mimeType string, size int, path string) *FileAnalysisResult {
	path = filepath.Abs(path)
	return &FileAnalysisResult{mimeType, size, path}
}

/**
 * @return string
 */
func (self *FileAnalysisResult) GetMimeType() string {
	return self.mimeType
}

/**
 * @return int
 */
func (self *FileAnalysisResult) GetSize() int {
	return self.size
}

/**
 * @return string
 */
func (self *FileAnalysisResult) GetPath() string {
	return self.path
}

/**
 * @return string
 */
func (self *FileAnalysisResult) GetFileName() string {
	return filepath.Base(self.path)
}
