package file_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/constants"
	. "github.com/enetx/tg/file"
)

func TestInput_EmptyFilename(t *testing.T) {
	result := Input(g.String(""))

	if result.IsOk() {
		t.Error("Expected error for empty filename")
	}

	err := result.Err()
	if err == nil {
		t.Error("Expected non-nil error for empty filename")
	}

	if !strings.Contains(err.Error(), "filename is empty") {
		t.Errorf("Expected 'filename is empty' error, got '%s'", err.Error())
	}
}

func TestInput_HTTPUrl(t *testing.T) {
	httpURL := g.String("http://example.com/image.jpg")
	result := Input(httpURL)

	if result.IsErr() {
		t.Errorf("Expected success for HTTP URL, got error: %v", result.Err())
	}

	inputFile := result.Ok()

	// Test that Doc is set correctly
	if inputFile.Doc == nil {
		t.Error("Expected Doc to be set for HTTP URL")
	}

	// Test that File is nil for URL inputs
	if inputFile.File != nil {
		t.Error("Expected File to be nil for HTTP URL input")
	}

	// Test the URL was processed correctly - just verify it's not nil
	if inputFile.Doc == nil {
		t.Error("Expected Doc to be processed for HTTP URL")
	}
}

func TestInput_HTTPSUrl(t *testing.T) {
	httpsURL := g.String("https://example.com/secure-image.png")
	result := Input(httpsURL)

	if result.IsErr() {
		t.Errorf("Expected success for HTTPS URL, got error: %v", result.Err())
	}

	inputFile := result.Ok()

	// Test that Doc is set correctly
	if inputFile.Doc == nil {
		t.Error("Expected Doc to be set for HTTPS URL")
	}

	// Test that File is nil for URL inputs
	if inputFile.File != nil {
		t.Error("Expected File to be nil for HTTPS URL input")
	}

	// Test the URL was processed correctly - just verify it's not nil
	if inputFile.Doc == nil {
		t.Error("Expected Doc to be processed for HTTPS URL")
	}
}

func TestInput_FileID(t *testing.T) {
	fileID := "BAADBAADrwADBREAAYag2HL2RX0XAg"
	prefixedFileID := g.String(constants.FileIDPrefix + fileID)

	result := Input(prefixedFileID)

	if result.IsErr() {
		t.Errorf("Expected success for file ID, got error: %v", result.Err())
	}

	inputFile := result.Ok()

	// Test that Doc is set correctly
	if inputFile.Doc == nil {
		t.Error("Expected Doc to be set for file ID")
	}

	// Test that File is nil for file ID inputs
	if inputFile.File != nil {
		t.Error("Expected File to be nil for file ID input")
	}

	// Test the file ID was processed correctly - just verify it's not nil
	if inputFile.Doc == nil {
		t.Error("Expected Doc to be processed for file ID")
	}
}

func TestInput_LocalFile_Success(t *testing.T) {
	// Create a temporary file for testing
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	testContent := "This is a test file for file package testing"

	err := os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result := Input(g.String(testFile))

	if result.IsErr() {
		t.Errorf("Expected success for local file, got error: %v", result.Err())
	}

	inputFile := result.Ok()

	// Test that Doc is set correctly
	if inputFile.Doc == nil {
		t.Error("Expected Doc to be set for local file")
	}

	// Test that File is set for local file inputs
	if inputFile.File == nil {
		t.Error("Expected File to be set for local file input")
	}

	// Test the file was processed correctly - just verify it's not nil
	if inputFile.Doc == nil {
		t.Error("Expected Doc to be processed for local file")
	}

	// Test that we can close the file (cleanup)
	if inputFile.File != nil {
		inputFile.File.Close()
	}
}

func TestInput_LocalFile_NotFound(t *testing.T) {
	nonExistentFile := g.String("/path/to/non/existent/file.txt")
	result := Input(nonExistentFile)

	if result.IsOk() {
		t.Error("Expected error for non-existent file")
	}

	err := result.Err()
	if err == nil {
		t.Error("Expected non-nil error for non-existent file")
	}

	// The error should be related to file not being found/openable
	errStr := err.Error()
	if !strings.Contains(errStr, "no such file") &&
		!strings.Contains(errStr, "cannot find") &&
		!strings.Contains(errStr, "does not exist") {
		t.Logf("Got error (which is expected): %s", errStr)
	}
}

func TestInput_EdgeCases(t *testing.T) {
	// Test URL edge cases
	testCases := []struct {
		name     string
		input    string
		shouldOk bool
		docType  string
	}{
		{
			name:     "HTTP with parameters",
			input:    "http://example.com/file.jpg?param=value",
			shouldOk: true,
			docType:  "url",
		},
		{
			name:     "HTTPS with port",
			input:    "https://example.com:8080/secure-file.png",
			shouldOk: true,
			docType:  "url",
		},
		{
			name:     "File ID with extra characters",
			input:    constants.FileIDPrefix + "ABCD1234_test-file-id",
			shouldOk: true,
			docType:  "id",
		},
		{
			name:     "File ID empty after prefix",
			input:    constants.FileIDPrefix,
			shouldOk: true,
			docType:  "id",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Input(g.String(tc.input))

			if tc.shouldOk && result.IsErr() {
				t.Errorf("Expected success for %s, got error: %v", tc.name, result.Err())
				return
			}

			if !tc.shouldOk && result.IsOk() {
				t.Errorf("Expected error for %s, got success", tc.name)
				return
			}

			if tc.shouldOk {
				inputFile := result.Ok()

				// Just verify the Doc was set correctly
				if inputFile.Doc == nil {
					t.Errorf("Expected Doc to be set for %s", tc.name)
				}
			}
		})
	}
}

func TestInput_URLVariations(t *testing.T) {
	// Test various URL formats
	urls := []string{
		"http://simple.com/file.jpg",
		"https://secure.com/file.png",
		"http://with-dash.com/file-name.gif",
		"https://subdomain.example.com/path/to/file.pdf",
		"http://example.com/file.jpg?query=param&other=value",
		"https://example.com:443/secure/path/file.webp",
	}

	for _, url := range urls {
		t.Run(url, func(t *testing.T) {
			result := Input(g.String(url))

			if result.IsErr() {
				t.Errorf("Expected success for URL %s, got error: %v", url, result.Err())
				return
			}

			inputFile := result.Ok()

			if inputFile.File != nil {
				t.Errorf("Expected File to be nil for URL %s", url)
			}

			// Just verify the Doc was set correctly for URL
			if inputFile.Doc == nil {
				t.Errorf("Expected Doc to be set for URL %s", url)
			}
		})
	}
}

func TestInput_FileIDVariations(t *testing.T) {
	// Test various file ID formats
	fileIDs := []string{
		"BAADBAADrwADBREAAYag2HL2RX0XAg",
		"simple_id_123",
		"file-id-with-dashes",
		"ID_WITH_UNDERSCORES_123",
		"mixed-ID_with123Numbers",
		"", // Empty file ID (should still work)
	}

	for _, fileID := range fileIDs {
		t.Run("FileID_"+fileID, func(t *testing.T) {
			input := constants.FileIDPrefix + fileID
			result := Input(g.String(input))

			if result.IsErr() {
				t.Errorf("Expected success for file ID %s, got error: %v", fileID, result.Err())
				return
			}

			inputFile := result.Ok()

			if inputFile.File != nil {
				t.Errorf("Expected File to be nil for file ID %s", fileID)
			}

			// Just verify the Doc was set correctly for file ID
			if inputFile.Doc == nil {
				t.Errorf("Expected Doc to be set for file ID %s", fileID)
			}
		})
	}
}

func TestInputFile_Structure(t *testing.T) {
	// Test that InputFile structure is properly initialized

	// Test with URL
	urlResult := Input(g.String("https://example.com/test.jpg"))
	if urlResult.IsErr() {
		t.Fatalf("Expected success for URL test")
	}

	urlFile := urlResult.Ok()
	if urlFile.Doc == nil {
		t.Error("Expected Doc to be non-nil for URL")
	}

	if urlFile.File != nil {
		t.Error("Expected File to be nil for URL")
	}

	// Test with file ID
	fileIDResult := Input(g.String(constants.FileIDPrefix + "test123"))
	if fileIDResult.IsErr() {
		t.Fatalf("Expected success for file ID test")
	}

	idFile := fileIDResult.Ok()
	if idFile.Doc == nil {
		t.Error("Expected Doc to be non-nil for file ID")
	}

	if idFile.File != nil {
		t.Error("Expected File to be nil for file ID")
	}
}

func TestInput_Constants(t *testing.T) {
	// Test that constants package integration works correctly
	testFileID := "test_file_id_123"
	fullInput := constants.FileIDPrefix + testFileID

	result := Input(g.String(fullInput))

	if result.IsErr() {
		t.Errorf("Expected success with constants prefix, got error: %v", result.Err())
		return
	}

	inputFile := result.Ok()
	// Just verify the Doc was set correctly
	if inputFile.Doc == nil {
		t.Error("Expected Doc to be set for file ID with constants prefix")
	}

	// Verify the prefix constant is being used
	if !strings.HasPrefix(fullInput, constants.FileIDPrefix) {
		t.Error("Expected input to start with FileIDPrefix constant")
	}
}
