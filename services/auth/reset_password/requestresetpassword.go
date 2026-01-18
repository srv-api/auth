package auth

import (
	"strings"
	"time"

	dto "github.com/srv-cashpay/auth/dto/auth"
	util "github.com/srv-cashpay/util/s"
)

func (s *verifyService) RequestResetPassword(req dto.ResetPasswordRequest) (dto.ResetPasswordResponse, error) {
	req.Email = strings.ToLower(req.Email)

	// Encrypt the email
	encryptedEmail, err := util.Encrypt(req.Email)
	if err != nil {
		return dto.ResetPasswordResponse{}, err
	}

	// Set the encrypted email in the request
	req.Email = encryptedEmail
	user, err := s.Repo.GetUserByEmail(req.Email)
	if err != nil {
		return dto.ResetPasswordResponse{}, err
		// return errors.New("Email not found")
	}

	// Generate unique token
	token := util.GenerateRandomString()
	expirationDuration := time.Hour
	otp := generateRandomNumeric(4)

	// Save the token in the database with a link to the user
	data, err := s.Repo.SavePasswordResetToken(user.ID, token, otp, expirationDuration)
	if err != nil {
		return dto.ResetPasswordResponse{}, err
		// return errors.New("Error saving reset token")
	}

	// Build the reset link using the clientURL
	// resetLink := fmt.Sprintf("%s/resetpassword?token=%s", getClientURL(), token)
	// err = s.SendResetEmail(user.Email, resetLink)
	// if err != nil {
	// 	return errors.New("Error sending reset email")
	// }

	// Decrypt the email for the response
	decryptedEmail, err := util.Decrypt(user.Email)
	if err != nil {
		return dto.ResetPasswordResponse{}, err
	}

	if err := util.Mail(decryptedEmail, otp); err != nil {
		return dto.ResetPasswordResponse{}, err
		// return errors.New("Error sending reset email")
	}

	response := dto.ResetPasswordResponse{
		Email:     decryptedEmail,
		UserID:    data.UserID,
		Token:     data.Token,
		Otp:       data.Otp,
		ExpiredAt: data.ExpiredAt,
	}

	return response, nil
}

// func getClientURL() string {
// 	// Fetch the clientURL from environment variable or configuration
// 	clientURL := os.Getenv("CLIENT_URL")
// 	if clientURL == "" {
// 		// Provide a default value or handle the missing configuration accordingly
// 		clientURL = "http://localhost:8080"
// 	}
// 	return clientURL
// }

// func (u *authService) SendResetEmail(email, resetLink string) error {
// 	message := gomail.NewMessage()

// 	message.SetHeader("From", "aseprayana95@gmail.com")
// 	message.SetHeader("To", email)
// 	message.SetHeader("Subject", "Password Reset")
// 	message.SetBody("text/html", fmt.Sprintf("Click <a href='%s'>here</a> to reset your password.", resetLink))

// 	dialer := util.NewMailer()

// 	// Send the email
// 	if err := dialer.DialAndSend(message); err != nil {
// 		return err
// 	}

// 	return nil
// }
