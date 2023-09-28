package validator

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestPasswordValidator(t *testing.T) {
	// Initialize a new validator
	v := validator.New()

	// Register the PasswordValidator function
	v.RegisterValidation("password", PasswordValidator)

	// Define a struct with the field to be validated
	type TestStruct struct {
		Password string `validate:"required,password"`
	}

	// Define an array of test cases
	testCases := []struct {
		Name     string
		Password string
		Expected bool
	}{
		{
			Name:     "ValidPassword",
			Password: "Abc12345!", // Meets all criteria, should be valid
			Expected: true,
		},
		{
			Name:     "ShortPassword",
			Password: "Abc12!", // Too short, should be invalid
			Expected: false,
		},
		{
			Name:     "NoUppercase",
			Password: "abc12345!", // Missing uppercase letter, should be invalid
			Expected: false,
		},
		{
			Name:     "NoLowercase",
			Password: "ABC12345!", // Missing lowercase letter, should be invalid
			Expected: false,
		},
		{
			Name:     "NoSpecialChar",
			Password: "Abcdefgh", // Missing special character, should be invalid
			Expected: false,
		},
	}

	// Iterate through the test cases
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Create a test case with the current password
			testCase := TestStruct{
				Password: tc.Password,
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
				assert.Equal(t, "Password", firstError.Field())
				assert.Equal(t, "password", firstError.Tag())
			}
		})
	}
}
