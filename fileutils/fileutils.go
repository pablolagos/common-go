package fileutils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

/**
 * Verify if a file or directory exists
 */
func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}

/**
 * FileSize returns file size in bytes
 */
func FileSize(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}

/**
 * Get the SHA-256 hash from a file
 */
func SHA256_File(file string) string {
	var result string
	result = ""
	f, err := os.Open(file)
	if err != nil {
		return result
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return result
	}

	result = hex.EncodeToString(h.Sum(nil))
	return result
}

/**
 * Get the SHA-256 hash from a file
 */
func MD5_File(file string) string {
	var result string
	result = ""
	f, err := os.Open(file)
	if err != nil {
		return result
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return result
	}

	result = hex.EncodeToString(h.Sum(nil))
	return result
}
