// Package passport defines Telegram passport element error types and utilities.
package passport

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// ErrorType enumerates all supported Telegram passport element error types.
type ErrorType int

const (
	DataFieldError        ErrorType = iota // Error in data field
	FrontSideError                         // Error in front side of document
	ReverseSideError                       // Error in reverse side of document
	SelfieError                            // Error in selfie photo
	FileError                              // Error in single file
	FilesError                             // Error in multiple files
	TranslationFileError                   // Error in single translation file
	TranslationFilesError                  // Error in multiple translation files
	UnspecifiedError                       // Unspecified error
)

// PassportElementType enumerates all supported passport element types.
type PassportElementType string

const (
	PersonalDetails       PassportElementType = "personal_details"
	Passport              PassportElementType = "passport"
	DriverLicense         PassportElementType = "driver_license"
	IdentityCard          PassportElementType = "identity_card"
	InternalPassport      PassportElementType = "internal_passport"
	Address               PassportElementType = "address"
	UtilityBill           PassportElementType = "utility_bill"
	BankStatement         PassportElementType = "bank_statement"
	RentalAgreement       PassportElementType = "rental_agreement"
	PassportRegistration  PassportElementType = "passport_registration"
	TemporaryRegistration PassportElementType = "temporary_registration"
)

// String returns the string representation of the PassportElementType.
func (pet PassportElementType) String() string {
	return string(pet)
}

// PassportError represents a builder for creating passport element errors.
type PassportError struct {
	errorType   ErrorType
	elementType PassportElementType
	message     String
	fieldName   Option[String]
	dataHash    Option[String]
	fileHash    Option[String]
	fileHashes  Option[Slice[String]]
	elementHash Option[String]
}

// NewDataFieldError creates a new data field error.
func NewDataFieldError(elementType PassportElementType, fieldName, dataHash, message String) *PassportError {
	return &PassportError{
		errorType:   DataFieldError,
		elementType: elementType,
		message:     message,
		fieldName:   Some(fieldName),
		dataHash:    Some(dataHash),
	}
}

// NewFrontSideError creates a new front side error.
func NewFrontSideError(elementType PassportElementType, fileHash, message String) *PassportError {
	return &PassportError{
		errorType:   FrontSideError,
		elementType: elementType,
		message:     message,
		fileHash:    Some(fileHash),
	}
}

// NewReverseSideError creates a new reverse side error.
func NewReverseSideError(elementType PassportElementType, fileHash, message String) *PassportError {
	return &PassportError{
		errorType:   ReverseSideError,
		elementType: elementType,
		message:     message,
		fileHash:    Some(fileHash),
	}
}

// NewSelfieError creates a new selfie error.
func NewSelfieError(elementType PassportElementType, fileHash, message String) *PassportError {
	return &PassportError{
		errorType:   SelfieError,
		elementType: elementType,
		message:     message,
		fileHash:    Some(fileHash),
	}
}

// NewFileError creates a new file error.
func NewFileError(elementType PassportElementType, fileHash, message String) *PassportError {
	return &PassportError{
		errorType:   FileError,
		elementType: elementType,
		message:     message,
		fileHash:    Some(fileHash),
	}
}

// NewFilesError creates a new files error.
func NewFilesError(elementType PassportElementType, fileHashes Slice[String], message String) *PassportError {
	return &PassportError{
		errorType:   FilesError,
		elementType: elementType,
		message:     message,
		fileHashes:  Some(fileHashes),
	}
}

// NewTranslationFileError creates a new translation file error.
func NewTranslationFileError(elementType PassportElementType, fileHash, message String) *PassportError {
	return &PassportError{
		errorType:   TranslationFileError,
		elementType: elementType,
		message:     message,
		fileHash:    Some(fileHash),
	}
}

// NewTranslationFilesError creates a new translation files error.
func NewTranslationFilesError(
	elementType PassportElementType,
	fileHashes Slice[String],
	message String,
) *PassportError {
	return &PassportError{
		errorType:   TranslationFilesError,
		elementType: elementType,
		message:     message,
		fileHashes:  Some(fileHashes),
	}
}

// NewUnspecifiedError creates a new unspecified error.
func NewUnspecifiedError(elementType PassportElementType, elementHash, message String) *PassportError {
	return &PassportError{
		errorType:   UnspecifiedError,
		elementType: elementType,
		message:     message,
		elementHash: Some(elementHash),
	}
}

// Build converts the PassportError to the appropriate gotgbot.PassportElementError interface.
func (pe *PassportError) Build() gotgbot.PassportElementError {
	switch pe.errorType {
	case DataFieldError:
		return &gotgbot.PassportElementErrorDataField{
			Type:      pe.elementType.String(),
			FieldName: pe.fieldName.UnwrapOrDefault().Std(),
			DataHash:  pe.dataHash.UnwrapOrDefault().Std(),
			Message:   pe.message.Std(),
		}
	case FrontSideError:
		return &gotgbot.PassportElementErrorFrontSide{
			Type:     pe.elementType.String(),
			FileHash: pe.fileHash.UnwrapOrDefault().Std(),
			Message:  pe.message.Std(),
		}
	case ReverseSideError:
		return &gotgbot.PassportElementErrorReverseSide{
			Type:     pe.elementType.String(),
			FileHash: pe.fileHash.UnwrapOrDefault().Std(),
			Message:  pe.message.Std(),
		}
	case SelfieError:
		return &gotgbot.PassportElementErrorSelfie{
			Type:     pe.elementType.String(),
			FileHash: pe.fileHash.UnwrapOrDefault().Std(),
			Message:  pe.message.Std(),
		}
	case FileError:
		return &gotgbot.PassportElementErrorFile{
			Type:     pe.elementType.String(),
			FileHash: pe.fileHash.UnwrapOrDefault().Std(),
			Message:  pe.message.Std(),
		}
	case FilesError:
		return &gotgbot.PassportElementErrorFiles{
			Type:       pe.elementType.String(),
			FileHashes: pe.fileHashes.UnwrapOrDefault().ToStringSlice(),
			Message:    pe.message.Std(),
		}
	case TranslationFileError:
		return &gotgbot.PassportElementErrorTranslationFile{
			Type:     pe.elementType.String(),
			FileHash: pe.fileHash.UnwrapOrDefault().Std(),
			Message:  pe.message.Std(),
		}
	case TranslationFilesError:
		return &gotgbot.PassportElementErrorTranslationFiles{
			Type:       pe.elementType.String(),
			FileHashes: pe.fileHashes.UnwrapOrDefault().ToStringSlice(),
			Message:    pe.message.Std(),
		}
	case UnspecifiedError:
		return &gotgbot.PassportElementErrorUnspecified{
			Type:        pe.elementType.String(),
			ElementHash: pe.elementHash.UnwrapOrDefault().Std(),
			Message:     pe.message.Std(),
		}
	default:
		// Fallback to unspecified error
		return &gotgbot.PassportElementErrorUnspecified{
			Type:        pe.elementType.String(),
			ElementHash: pe.elementHash.UnwrapOrDefault().Std(),
			Message:     pe.message.Std(),
		}
	}
}

// Errors creates a slice of gotgbot.PassportElementError from multiple PassportError builders.
func Errors(errors ...*PassportError) Slice[gotgbot.PassportElementError] {
	result := NewSlice[gotgbot.PassportElementError](Int(len(errors)))
	for i, err := range errors {
		result[i] = err.Build()
	}

	return result
}
