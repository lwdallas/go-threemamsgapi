package tools

import "os"

/**
 * @author Threema GmbH
 * @copyright Copyright (c) 2015-2016 Threema GmbH
 */

type FileAnalysisTool struct {
}

/**
 * @param string $file
 * @return FileAnalysisResult
 */
func (self *FileAnalysisTool) Analyse(file string) FileAnalysisResult {
	size := 0
	mimeType := ""

	f, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return nil
	}
	mode := fi.Mode()
	if mode.IsDir() {
		return nil
	} else if mode.IsRegular() {
		//get file size
		size = fi.Size()

		mimeType = ""
		//TODO: mime type getter

		//default mime type
		if len(mimeType) == 0 {
			//default mime type
			mimeType = "application/octet-stream"
		}

	}
	return FileAnalysisResult{mimeType, size, file}
}
