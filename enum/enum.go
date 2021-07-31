package enum

type authTypes struct {
	AuthDenied           int
	AuthValidationFailed int
	AuthAllowed          int
	AuthValidation       int
	AuthError            int
}

var AuthTypes = authTypes{0, 6, 1, 5, -1}