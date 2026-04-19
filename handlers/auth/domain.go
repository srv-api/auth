package handlers

import (
	s "github.com/srv-api/auth/services/auth"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Signup(c echo.Context) error        //masuk
	Signin(c echo.Context) error        //masuk
	RefreshToken(c echo.Context) error  //refresh
	Signout(c echo.Context) error       //keluar
	Authenticator(c echo.Context) error //Authenticator
	Profile(c echo.Context) error       //Profile
	UpdateProfile(c echo.Context) error //UpdateProfile
	GoogleSignIn(c echo.Context) error
	GoogleSignInWeb(c echo.Context) error
	GetPicture(c echo.Context) error
	GetPictureGallery(c echo.Context) error
	UploadImage(c echo.Context) error
	Gallery(c echo.Context) error
	GetGallery(c echo.Context) error
	DeleteGallery(c echo.Context) error
}

type domainHandler struct {
	serviceAuth s.AuthService
}

func NewAuthHandler(service s.AuthService) DomainHandler {
	return &domainHandler{
		serviceAuth: service,
	}
}
