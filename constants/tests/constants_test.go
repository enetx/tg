package constants_test

import (
	"strings"
	"testing"

	. "github.com/enetx/tg/constants"
)

func TestFileIDPrefix(t *testing.T) {
	// Test that the constant is defined and has expected value
	if FileIDPrefix == "" {
		t.Error("Expected FileIDPrefix to be non-empty")
	}

	// Test the expected value
	expectedPrefix := "__TG_FILEID__:"
	if FileIDPrefix != expectedPrefix {
		t.Errorf("Expected FileIDPrefix to be '%s', got '%s'", expectedPrefix, FileIDPrefix)
	}

	// Test that it's a reasonable prefix format
	if !strings.HasSuffix(FileIDPrefix, ":") {
		t.Error("Expected FileIDPrefix to end with ':'")
	}

	if !strings.HasPrefix(FileIDPrefix, "__") {
		t.Error("Expected FileIDPrefix to start with '__'")
	}

	if !strings.Contains(FileIDPrefix, "TG") {
		t.Error("Expected FileIDPrefix to contain 'TG'")
	}

	if !strings.Contains(FileIDPrefix, "FILEID") {
		t.Error("Expected FileIDPrefix to contain 'FILEID'")
	}
}

func TestFileIDPrefix_Usage(t *testing.T) {
	// Test typical usage patterns with the constant

	// Test concatenation
	testFileID := "BAADBAADrwADBREAAYag2HL2RX0XAg"
	fullFileRef := FileIDPrefix + testFileID

	expectedFull := "__TG_FILEID__:BAADBAADrwADBREAAYag2HL2RX0XAg"
	if fullFileRef != expectedFull {
		t.Errorf("Expected concatenated result '%s', got '%s'", expectedFull, fullFileRef)
	}

	// Test prefix detection
	if !strings.HasPrefix(fullFileRef, FileIDPrefix) {
		t.Error("Expected full file reference to start with FileIDPrefix")
	}

	// Test prefix stripping (simulated)
	if strings.HasPrefix(fullFileRef, FileIDPrefix) {
		stripped := strings.TrimPrefix(fullFileRef, FileIDPrefix)
		if stripped != testFileID {
			t.Errorf("Expected stripped file ID '%s', got '%s'", testFileID, stripped)
		}
	}
}

func TestFileIDPrefix_Properties(t *testing.T) {
	// Test various properties of the constant

	// Test length is reasonable (not too short, not too long)
	if len(FileIDPrefix) < 5 {
		t.Error("Expected FileIDPrefix to have reasonable length (at least 5 characters)")
	}

	if len(FileIDPrefix) > 50 {
		t.Error("Expected FileIDPrefix to not be excessively long (max 50 characters)")
	}

	// Test it doesn't contain spaces
	if strings.Contains(FileIDPrefix, " ") {
		t.Error("Expected FileIDPrefix to not contain spaces")
	}

	// Test it's uppercase (design choice verification)
	if strings.ToUpper(FileIDPrefix) != FileIDPrefix {
		t.Error("Expected FileIDPrefix to be in uppercase")
	}

	// Test uniqueness indicators
	if !strings.Contains(FileIDPrefix, "__") {
		t.Error("Expected FileIDPrefix to contain double underscores for uniqueness")
	}
}

func TestFileIDPrefix_EdgeCases(t *testing.T) {
	// Test edge cases and boundary conditions

	// Test with empty string
	emptyResult := FileIDPrefix + ""
	if emptyResult != FileIDPrefix {
		t.Error("Expected concatenation with empty string to equal the prefix")
	}

	// Test with special characters
	specialFileID := "file_with-special.chars"
	specialResult := FileIDPrefix + specialFileID
	if !strings.HasPrefix(specialResult, FileIDPrefix) {
		t.Error("Expected result with special characters to have correct prefix")
	}

	if !strings.HasSuffix(specialResult, specialFileID) {
		t.Error("Expected result to end with the special file ID")
	}

	// Test prefix uniqueness (shouldn't appear in normal file IDs)
	normalFileIDs := []string{
		"BAADBAADrwADBREAAYag2HL2RX0XAg",
		"simple_file_id",
		"file-with-dashes",
		"FILE_WITH_CAPS",
		"123456789",
	}

	for _, fileID := range normalFileIDs {
		if strings.Contains(fileID, FileIDPrefix) {
			t.Errorf("Normal file ID '%s' unexpectedly contains the prefix", fileID)
		}
	}
}

func TestFileIDPrefix_Consistency(t *testing.T) {
	// Test that the constant is consistent across multiple access

	first := FileIDPrefix
	second := FileIDPrefix

	if first != second {
		t.Error("Expected FileIDPrefix to be consistent across multiple access")
	}

	// Test immutability (constants should not change)
	original := FileIDPrefix

	// Simulate some operations that shouldn't affect the constant
	_ = strings.ToLower(FileIDPrefix)
	_ = strings.ToUpper(FileIDPrefix)
	_ = FileIDPrefix + "test"

	if FileIDPrefix != original {
		t.Error("Expected FileIDPrefix to remain unchanged after operations")
	}
}

func TestFileIDPrefix_Documentation(t *testing.T) {
	// Test that the prefix format is documented-friendly

	// Should be human-readable enough to understand its purpose
	prefix := FileIDPrefix

	// Should clearly indicate it's related to Telegram (TG)
	if !strings.Contains(prefix, "TG") {
		t.Error("Expected prefix to clearly indicate Telegram relationship")
	}

	// Should clearly indicate it's for file IDs
	if !strings.Contains(prefix, "FILEID") {
		t.Error("Expected prefix to clearly indicate file ID purpose")
	}

	// Should use a clear separator
	if !strings.Contains(prefix, ":") {
		t.Error("Expected prefix to use a clear separator (:)")
	}

	// Should be distinctive enough to avoid conflicts
	distinctiveChars := []string{"_", ":", "__"}
	hasDistinctive := false
	for _, char := range distinctiveChars {
		if strings.Contains(prefix, char) {
			hasDistinctive = true
			break
		}
	}

	if !hasDistinctive {
		t.Error("Expected prefix to contain distinctive characters to avoid conflicts")
	}
}
