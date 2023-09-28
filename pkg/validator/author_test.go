package validator

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestAuthorValidator(t *testing.T) {
	// Initialize a new validator
	v := validator.New()

	// Register the AuthorValidator function
	v.RegisterValidation("author", AuthorValidator)

	// Define a struct with the field to be validated
	type TestStruct struct {
		Author string `validate:"required,author"`
	}

	// Define an array of test cases
	testCases := []struct {
		Name     string
		Author   string
		Expected bool
	}{
		{
			Name:     "ValidAuthor",
			Author:   "John Doe",
			Expected: true,
		},
		{
			Name:     "InvalidAuthor",
			Author:   "John123", // Contains digits, should be invalid
			Expected: false,
		},
	}

	// Iterate through the test cases
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Create a test case with the current author name
			testCase := TestStruct{
				Author: tc.Author,
			}

			// Validate the struct
			err := v.Struct(testCase)

			// Check the validation result based on the expected value
			if tc.Expected {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)

				// You can also check the specific validation error if needed
				validationErrors, ok := err.(validator.ValidationErrors)
				assert.True(t, ok)
				assert.Equal(t, 1, len(validationErrors))

				// You can further inspect the specific validation error if needed
				firstError := validationErrors[0]
				assert.Equal(t, "Author", firstError.Field())
				assert.Equal(t, "author", firstError.Tag())
			}
		})
	}
}
