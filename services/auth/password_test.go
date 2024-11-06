package auth

import "testing"

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "Valid password",
			password: "validPassword123",
			wantErr:  false,
		},
		{
			name:     "Empty password",
			password: "",
			wantErr:  true,
		},
		{
			name:     "Long password",
			password: string(make([]byte, 73)),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := HashPassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if hash == "" {
					t.Error("HashPassword() returned empty hash")
				}

				if hash == tt.password {
					t.Error("HashPassword() returned unhashed password")
				}
			}
		})
	}
}

func TestComparePasswords(t *testing.T) {
	password := "testPassword123"
	hash, _ := HashPassword(password)

	tests := []struct {
		name       string
		hashedPass string
		plainPass  []byte
		wantMatch  bool
	}{
		{
			name:       "Matching passwords",
			hashedPass: hash,
			plainPass:  []byte(password),
			wantMatch:  true,
		},
		{
			name:       "Non-matching passwords",
			hashedPass: hash,
			plainPass:  []byte("wrongPassword"),
			wantMatch:  false,
		},
		{
			name:       "Empty plain password",
			hashedPass: hash,
			plainPass:  []byte(""),
			wantMatch:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComparePasswords(tt.hashedPass, tt.plainPass); got != tt.wantMatch {
				t.Errorf("ComparePasswords() = %v, want %v", got, tt.wantMatch)
			}
		})
	}
}
