package auth

import (
    "errors"
    "time"

    "github.com/dgrijalva/jwt-go"
)

// AuthService provides authentication functionality
type AuthService struct {
    // Define any necessary fields here
    users map[string]string // In-memory user store: username -> password
    secretKey string        // Secret key for JWT
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(secretKey string) *AuthService {
    // Initialize and return a new AuthService instance
    return &AuthService{
        users:     make(map[string]string),
        secretKey: secretKey,
    }
}

// Authenticate authenticates a user based on username and password
func (s *AuthService) Authenticate(username, password string) (string, error) {
    // Check if the user exists and the password is correct
    if storedPassword, ok := s.users[username]; ok {
        if storedPassword == password {
            // Generate JWT token
            token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
                "username": username,
                "exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
            })

            // Sign and get the complete encoded token as a string using the secret
            tokenString, err := token.SignedString([]byte(s.secretKey))
            if err != nil {
                return "", err
            }

            return tokenString, nil
        }
    }

    return "", errors.New("authentication failed")
}

// AddUser adds a new user to the in-memory store
func (s *AuthService) AddUser(username, password string) {
    s.users[username] = password
}
