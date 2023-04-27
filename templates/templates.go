package templates

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
)

//go:embed index.html report.html
var EmbeddedFiles embed.FS

func CopyEmbeddedFile(srcFS fs.FS, srcPath, dstPath string) error {
	srcFile, err := srcFS.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}
	return nil
}
