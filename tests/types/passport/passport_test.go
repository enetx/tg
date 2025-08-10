package passport_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/types/passport"
)

func TestPassportElementType_String(t *testing.T) {
	tests := []struct {
		name        string
		elementType passport.PassportElementType
		expected    string
	}{
		{"PersonalDetails", passport.PersonalDetails, "personal_details"},
		{"Passport", passport.Passport, "passport"},
		{"DriverLicense", passport.DriverLicense, "driver_license"},
		{"IdentityCard", passport.IdentityCard, "identity_card"},
		{"InternalPassport", passport.InternalPassport, "internal_passport"},
		{"Address", passport.Address, "address"},
		{"UtilityBill", passport.UtilityBill, "utility_bill"},
		{"BankStatement", passport.BankStatement, "bank_statement"},
		{"RentalAgreement", passport.RentalAgreement, "rental_agreement"},
		{"PassportRegistration", passport.PassportRegistration, "passport_registration"},
		{"TemporaryRegistration", passport.TemporaryRegistration, "temporary_registration"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.elementType.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestNewDataFieldError(t *testing.T) {
	elementType := passport.PersonalDetails
	fieldName := g.String("first_name")
	dataHash := g.String("data_hash_123")
	message := g.String("Field contains invalid data")

	error := passport.NewDataFieldError(elementType, fieldName, dataHash, message)

	if error == nil {
		t.Error("Expected NewDataFieldError to return non-nil error")
	}

	built := error.Build()
	if built == nil {
		t.Error("Expected Build to return non-nil error")
	}

	if dataFieldError, ok := built.(*gotgbot.PassportElementErrorDataField); ok {
		if dataFieldError.Type != elementType.String() {
			t.Errorf("Expected Type %s, got %s", elementType.String(), dataFieldError.Type)
		}
		if dataFieldError.FieldName != fieldName.Std() {
			t.Errorf("Expected FieldName %s, got %s", fieldName.Std(), dataFieldError.FieldName)
		}
		if dataFieldError.DataHash != dataHash.Std() {
			t.Errorf("Expected DataHash %s, got %s", dataHash.Std(), dataFieldError.DataHash)
		}
		if dataFieldError.Message != message.Std() {
			t.Errorf("Expected Message %s, got %s", message.Std(), dataFieldError.Message)
		}
	} else {
		t.Error("Expected result to be PassportElementErrorDataField")
	}
}

func TestNewFrontSideError(t *testing.T) {
	elementType := passport.Passport
	fileHash := g.String("file_hash_123")
	message := g.String("Front side is blurry")

	error := passport.NewFrontSideError(elementType, fileHash, message)

	if error == nil {
		t.Error("Expected NewFrontSideError to return non-nil error")
	}

	built := error.Build()
	if frontSideError, ok := built.(*gotgbot.PassportElementErrorFrontSide); ok {
		if frontSideError.Type != elementType.String() {
			t.Errorf("Expected Type %s, got %s", elementType.String(), frontSideError.Type)
		}
		if frontSideError.FileHash != fileHash.Std() {
			t.Errorf("Expected FileHash %s, got %s", fileHash.Std(), frontSideError.FileHash)
		}
		if frontSideError.Message != message.Std() {
			t.Errorf("Expected Message %s, got %s", message.Std(), frontSideError.Message)
		}
	} else {
		t.Error("Expected result to be PassportElementErrorFrontSide")
	}
}

func TestNewReverseSideError(t *testing.T) {
	elementType := passport.DriverLicense
	fileHash := g.String("reverse_hash_456")
	message := g.String("Reverse side is damaged")

	error := passport.NewReverseSideError(elementType, fileHash, message)

	built := error.Build()
	if reverseSideError, ok := built.(*gotgbot.PassportElementErrorReverseSide); ok {
		if reverseSideError.Type != elementType.String() {
			t.Errorf("Expected Type %s, got %s", elementType.String(), reverseSideError.Type)
		}
	} else {
		t.Error("Expected result to be PassportElementErrorReverseSide")
	}
}

func TestNewSelfieError(t *testing.T) {
	elementType := passport.IdentityCard
	fileHash := g.String("selfie_hash_789")
	message := g.String("Selfie quality is too low")

	error := passport.NewSelfieError(elementType, fileHash, message)

	built := error.Build()
	if selfieError, ok := built.(*gotgbot.PassportElementErrorSelfie); ok {
		if selfieError.Type != elementType.String() {
			t.Errorf("Expected Type %s, got %s", elementType.String(), selfieError.Type)
		}
	} else {
		t.Error("Expected result to be PassportElementErrorSelfie")
	}
}

func TestNewFileError(t *testing.T) {
	elementType := passport.UtilityBill
	fileHash := g.String("file_hash_abc")
	message := g.String("File format not supported")

	error := passport.NewFileError(elementType, fileHash, message)

	built := error.Build()
	if fileError, ok := built.(*gotgbot.PassportElementErrorFile); ok {
		if fileError.Type != elementType.String() {
			t.Errorf("Expected Type %s, got %s", elementType.String(), fileError.Type)
		}
	} else {
		t.Error("Expected result to be PassportElementErrorFile")
	}
}

func TestNewFilesError(t *testing.T) {
	elementType := passport.BankStatement
	fileHashes := g.SliceOf(g.String("hash1"), g.String("hash2"))
	message := g.String("Multiple files have issues")

	error := passport.NewFilesError(elementType, fileHashes, message)

	built := error.Build()
	if filesError, ok := built.(*gotgbot.PassportElementErrorFiles); ok {
		if filesError.Type != elementType.String() {
			t.Errorf("Expected Type %s, got %s", elementType.String(), filesError.Type)
		}
		if len(filesError.FileHashes) != 2 {
			t.Errorf("Expected 2 file hashes, got %d", len(filesError.FileHashes))
		}
	} else {
		t.Error("Expected result to be PassportElementErrorFiles")
	}
}

func TestNewTranslationFileError(t *testing.T) {
	elementType := passport.RentalAgreement
	fileHash := g.String("translation_hash_def")
	message := g.String("Translation file is missing")

	error := passport.NewTranslationFileError(elementType, fileHash, message)

	built := error.Build()
	if translationError, ok := built.(*gotgbot.PassportElementErrorTranslationFile); ok {
		if translationError.Type != elementType.String() {
			t.Errorf("Expected Type %s, got %s", elementType.String(), translationError.Type)
		}
	} else {
		t.Error("Expected result to be PassportElementErrorTranslationFile")
	}
}

func TestNewTranslationFilesError(t *testing.T) {
	elementType := passport.PassportRegistration
	fileHashes := g.SliceOf(g.String("trans1"), g.String("trans2"))
	message := g.String("Translation files are incomplete")

	error := passport.NewTranslationFilesError(elementType, fileHashes, message)

	built := error.Build()
	if translationFilesError, ok := built.(*gotgbot.PassportElementErrorTranslationFiles); ok {
		if translationFilesError.Type != elementType.String() {
			t.Errorf("Expected Type %s, got %s", elementType.String(), translationFilesError.Type)
		}
	} else {
		t.Error("Expected result to be PassportElementErrorTranslationFiles")
	}
}

func TestNewUnspecifiedError(t *testing.T) {
	elementType := passport.TemporaryRegistration
	elementHash := g.String("element_hash_xyz")
	message := g.String("Unspecified error occurred")

	error := passport.NewUnspecifiedError(elementType, elementHash, message)

	built := error.Build()
	if unspecifiedError, ok := built.(*gotgbot.PassportElementErrorUnspecified); ok {
		if unspecifiedError.Type != elementType.String() {
			t.Errorf("Expected Type %s, got %s", elementType.String(), unspecifiedError.Type)
		}
	} else {
		t.Error("Expected result to be PassportElementErrorUnspecified")
	}
}

func TestErrors(t *testing.T) {
	error1 := passport.NewDataFieldError(passport.PersonalDetails, g.String("name"), g.String("hash1"), g.String("Error 1"))
	error2 := passport.NewFileError(passport.Passport, g.String("hash2"), g.String("Error 2"))

	errors := passport.Errors(error1, error2)

	if errors.Len() != 2 {
		t.Errorf("Expected 2 errors, got %d", errors.Len())
	}
}
