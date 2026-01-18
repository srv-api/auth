package auth

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"strings"
	"time"

	dto "github.com/srv-cashpay/auth/dto/auth"
	"github.com/srv-cashpay/auth/entity"
	util "github.com/srv-cashpay/util/s"

	"gopkg.in/gomail.v2"
)

func (u *verifyService) ResendReset(req dto.ResendResetRequest) (*entity.PasswordResetToken, error) {

	resend := dto.ResendResetRequest{
		Otp:   generateRandomNumeric(4),
		Token: req.Token,
		Email: req.Email,
	}

	// Use your repository or service to fetch the user by token from the database
	user, err := u.Repo.ResendReset(resend)
	if err != nil {
		return nil, err
	}
	user.Otp = resend.Otp

	if err := util.Mail(resend.Email, user.Otp); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *verifyService) resendmail(to, otp string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "aseprayana95@gmail.com")
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Test Email")
	//if using otp kode
	mailer.SetBody("text/html", fmt.Sprintf("Your verification code is: <strong>%s</strong>", otp))
	//click button at email and verify
	// mailer.SetBody("text/html", fmt.Sprintf("Hello, this is a test email from "+
	// "Mailtrap: <a href='http://localhost:8080/verify/%s'>Verify Account</a>Your verification code is: <strong>%s</strong>",
	// verificationToken, otp))

	dialer := gomail.NewDialer("smtp.mailtrap.io", 587, "7de3a28724e886", "353081a2c62514")
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // Use this only for development, not secure for production

	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}

func generateRandomNumeric(length int) string {
	const chars = "0123456789"

	var result strings.Builder
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		result.WriteRune(rune(chars[rand.Intn(len(chars))]))
	}

	return result.String()
}
