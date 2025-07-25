package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/types/passport"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Simple command to show passport error types
	b.Command("start", func(ctx *ctx.Context) error {
		return ctx.Reply(`Passport Error Types:

• DataFieldError - Invalid data in fields
• FrontSideError - Front document image issues
• ReverseSideError - Back document image issues
• SelfieError - Selfie photo problems
• FileError - Single file issues
• FilesError - Multiple files problems
• TranslationFileError - Single translation issues
• TranslationFilesError - Multiple translation problems
• UnspecifiedError - General validation errors

Use /passport_demo to test all error types.`).Send().Err()
	})

	// Command to demonstrate all passport error types
	b.Command("passport_demo", func(ctx *ctx.Context) error {
		userID := ctx.EffectiveUser.Id

		// Demo: Create all types of passport errors
		result := ctx.SetPassportDataErrors(userID).
			Errors(
				// Data field error in personal details
				passport.NewDataFieldError(
					passport.PersonalDetails,
					"birth_date",
					"data_hash_001",
					"Invalid date format",
				),
				// Front side document error
				passport.NewFrontSideError(
					passport.Passport,
					"front_hash_002",
					"Photo is not clear enough",
				),
				// Reverse side error
				passport.NewReverseSideError(
					passport.DriverLicense,
					"reverse_hash_003",
					"Back side is damaged",
				),
				// Selfie error
				passport.NewSelfieError(
					passport.IdentityCard,
					"selfie_hash_004",
					"Face not visible in selfie",
				),
				// Single file error
				passport.NewFileError(
					passport.BankStatement,
					"file_hash_005",
					"Document format not supported",
				),
				// Multiple files error
				passport.NewFilesError(
					passport.UtilityBill,
					SliceOf[String]("file1_hash", "file2_hash", "file3_hash"),
					"All files are expired",
				),
				// Translation file error
				passport.NewTranslationFileError(
					passport.Passport,
					"translation_hash_006",
					"Translation is incomplete",
				),
				// Multiple translation files error
				passport.NewTranslationFilesError(
					passport.DriverLicense,
					SliceOf[String]("trans1_hash", "trans2_hash"),
					"Translation quality insufficient",
				),
				// Unspecified error
				passport.NewUnspecifiedError(
					passport.Address,
					"element_hash_007",
					"Address verification failed",
				),
			).Send()

		if result.IsOk() {
			return ctx.Reply("Demo passport errors set! All 9 types of passport validation errors have been created.").
				Send().
				Err()
		}

		return ctx.Reply("Failed to set passport errors").Send().Err()
	})

	// Handle messages with passport data (when user sends passport info)
	b.On.Message.Any(func(ctx *ctx.Context) error {
		// Check if message contains passport data
		if ctx.EffectiveMessage.PassportData != nil {
			passportData := ctx.EffectiveMessage.PassportData
			userID := ctx.EffectiveUser.Id

			// Simulate passport validation and create errors
			var errors Slice[*passport.PassportError]

			// Example validation scenarios
			for _, element := range passportData.Data {
				switch element.Type {
				case "personal_details":
					// Data field validation error
					errors.Push(
						passport.NewDataFieldError(
							passport.PersonalDetails,
							"first_name",
							"invalid_hash_123",
							"First name contains invalid characters",
						),
					)
				case "passport":
					// Front side image error
					errors.Push(
						passport.NewFrontSideError(
							passport.Passport,
							"front_side_hash_456",
							"Document image is too blurry",
						),
					)
				case "utility_bill":
					// Multiple files error
					fileHashes := Slice[String]{"file1_hash", "file2_hash"}
					errors.Push(
						passport.NewFilesError(
							passport.UtilityBill,
							fileHashes,
							"Documents are older than 90 days",
						),
					)
				}
			}

			// Send passport validation errors
			if errors.NotEmpty() {
				if result := ctx.SetPassportDataErrors(userID).Errors(errors...).Send(); result.IsOk() {
					return ctx.Reply("Passport validation failed. Please correct the errors.").Send().Err()
				}
			}

			return ctx.Reply("Passport data validated successfully!").Send().Err()
		}

		return nil
	})

	b.Polling().Start()
}
