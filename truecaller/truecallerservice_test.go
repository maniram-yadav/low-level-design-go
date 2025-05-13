package truecaller

import (
	"context"
	"testing"
)

func TestPhoneNumberValidation(t *testing.T) {
	tests := []struct {
		name        string
		countryCode string
		number      string
		expectError bool
	}{
		{"Valid US number", "+1", "6505551234", false},
		{"Valid UK number", "+44", "7911123456", false},
		{"Invalid country code", "1", "6505551234", true},
		{"Invalid number (too short)", "+1", "12345", true},
		{"Invalid number (non-digits)", "+1", "650ABC1234", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewPhoneNumber(tt.countryCode, tt.number)
			if (err != nil) != tt.expectError {
				t.Errorf("NewPhoneNumber() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestIdentifyCaller(t *testing.T) {
	service := NewTrueCallerServiceWithInMemoryDB()
	ctx := context.Background()

	tests := []struct {
		name         string
		userID       string
		callerNumber *PhoneNumber
		wantName     string
		wantSpam     bool
		wantErr      bool
	}{
		{
			name:         "Known contact",
			userID:       "user1",
			callerNumber: &PhoneNumber{countryCode: "+91", number: "5551234567"},
			wantName:     "Charlie",
			wantSpam:     false,
			wantErr:      false,
		},
		{
			name:         "Unknown number",
			userID:       "user1",
			callerNumber: &PhoneNumber{countryCode: "+1", number: "1231231234"},
			wantName:     "",
			wantSpam:     false,
			wantErr:      false,
		},
		{
			name:         "Spam number",
			userID:       "user1",
			callerNumber: &PhoneNumber{countryCode: "+1", number: "9998887777"},
			wantName:     "",
			wantSpam:     true,
			wantErr:      false,
		},
		{
			name:         "Invalid user",
			userID:       "nonexistent",
			callerNumber: &PhoneNumber{countryCode: "+1", number: "6505551234"},
			wantName:     "",
			wantSpam:     false,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.IdentifyCaller(ctx, tt.userID, tt.callerNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("IdentifyCaller() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if got.Name != tt.wantName {
					t.Errorf("IdentifyCaller() name = %v, want %v", got.Name, tt.wantName)
				}
				if got.IsSpam != tt.wantSpam {
					t.Errorf("IdentifyCaller() isSpam = %v, want %v", got.IsSpam, tt.wantSpam)
				}
			}
		})
	}
}
