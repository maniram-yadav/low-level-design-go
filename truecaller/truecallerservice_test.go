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

func TestAddContact(t *testing.T) {
	service := NewTrueCallerServiceWithInMemoryDB()
	ctx := context.Background()

	tests := []struct {
		name    string
		userID  string
		contact *Contact
		wantErr bool
	}{
		{
			name:   "Valid new contact",
			userID: "user1",
			contact: &Contact{
				Name: "Eve",
				PhoneNumbers: []*PhoneNumber{
					{countryCode: "+1", number: "5551112222"},
				},
				Email: "eve@example.com",
			},
			wantErr: false,
		},
		{
			name:   "Duplicate contact",
			userID: "user1",
			contact: &Contact{
				Name: "Duplicate",
				PhoneNumbers: []*PhoneNumber{
					{countryCode: "+91", number: "5551234567"}, // Already exists
				},
			},
			wantErr: true,
		},
		{
			name:   "Invalid user",
			userID: "nonexistent",
			contact: &Contact{
				Name: "Test",
				PhoneNumbers: []*PhoneNumber{
					{countryCode: "+1", number: "5550000000"},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.AddContact(ctx, tt.userID, tt.contact)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddContact() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Verify contact was added
			if err == nil && !tt.wantErr {
				user, _ := service.db.GetUserByID(ctx, tt.userID)
				found := false
				for _, phone := range tt.contact.PhoneNumbers {
					if _, exists := user.Contacts[phone.String()]; exists {
						found = true
						break
					}
				}
				if !found {
					t.Error("AddContact() failed to add contact to user")
				}
			}
		})
	}
}

func TestReportSpam(t *testing.T) {
	service := NewTrueCallerServiceWithInMemoryDB()
	ctx := context.Background()

	spamNumber := &PhoneNumber{countryCode: "+1", number: "8887776666"}

	tests := []struct {
		name      string
		userID    string
		phone     *PhoneNumber
		wantErr   bool
		wantSpam  bool
		wantCount int
	}{
		{
			name:      "First spam report",
			userID:    "user1",
			phone:     spamNumber,
			wantErr:   false,
			wantSpam:  false, // Not over threshold yet
			wantCount: 1,
		},
		{
			name:      "Second spam report",
			userID:    "user2",
			phone:     spamNumber,
			wantErr:   false,
			wantSpam:  false,
			wantCount: 2,
		},
		{
			name:      "Third spam report (now marked as spam)",
			userID:    "user1",
			phone:     spamNumber,
			wantErr:   false,
			wantSpam:  true,
			wantCount: 3,
		},
		{
			name:      "Duplicate report from same user",
			userID:    "user1",
			phone:     spamNumber,
			wantErr:   true, // Should return ErrSpamReportExists
			wantSpam:  true,
			wantCount: 3, // Count shouldn't change
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.ReportSpam(ctx, tt.userID, tt.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReportSpam() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Verify spam count
			count, _ := service.db.GetSpamCount(ctx, tt.phone)
			if count != tt.wantCount {
				t.Errorf("ReportSpam() count = %d, want %d", count, tt.wantCount)
			}

			// Verify spam status
			info, _ := service.IdentifyCaller(ctx, "user1", tt.phone)
			if info.IsSpam != tt.wantSpam {
				t.Errorf("ReportSpam() isSpam = %v, want %v", info.IsSpam, tt.wantSpam)
			}
		})
	}
}

func TestBlockNumber(t *testing.T) {
	service := NewTrueCallerServiceWithInMemoryDB()
	ctx := context.Background()

	blockNumber := &PhoneNumber{countryCode: "+1", number: "5559876543"} // David's number

	tests := []struct {
		name        string
		userID      string
		phone       *PhoneNumber
		wantErr     bool
		wantBlocked bool
	}{
		{
			name:        "Block valid number",
			userID:      "user1",
			phone:       blockNumber,
			wantErr:     false,
			wantBlocked: true,
		},
		{
			name:        "Block already blocked number",
			userID:      "user1",
			phone:       blockNumber,
			wantErr:     false, // No error on re-blocking
			wantBlocked: true,
		},
		{
			name:        "Block with invalid user",
			userID:      "nonexistent",
			phone:       blockNumber,
			wantErr:     true,
			wantBlocked: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.BlockNumber(ctx, tt.userID, tt.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("BlockNumber() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				user, _ := service.db.GetUserByID(ctx, tt.userID)
				if blocked := user.Blocked[tt.phone.String()]; blocked != tt.wantBlocked {
					t.Errorf("BlockNumber() blocked = %v, want %v", blocked, tt.wantBlocked)
				}

				// Verify IdentifyCaller shows the number as blocked
				info, _ := service.IdentifyCaller(ctx, tt.userID, tt.phone)
				if info.IsBlocked != tt.wantBlocked {
					t.Errorf("IdentifyCaller() isBlocked = %v, want %v", info.IsBlocked, tt.wantBlocked)
				}
			}
		})
	}
}
