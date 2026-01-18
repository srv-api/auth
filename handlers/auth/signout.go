package handlers

import (
	"fmt"
	"net/http"
	"time"

	res "github.com/srv-cashpay/util/s/response"

	"github.com/labstack/echo/v4"
)

func (u *domainHandler) Signout(c echo.Context) error {
	// Hapus token cookie
	cookies := http.Cookie{
		Name:     "token",                             // Nama cookie yang digunakan untuk token
		Value:    "",                                  // Setel nilai cookie menjadi kosong
		Path:     "/dashboard",                        // Path di mana cookie ini berlaku
		Expires:  time.Now().Add(-7 * 24 * time.Hour), // Setel waktu kadaluarsa untuk menghapus cookie
		MaxAge:   -1,                                  // MaxAge negatif untuk menghapus cookie segera
		HttpOnly: true,                                // Hanya dapat diakses melalui HTTP (bukan JavaScript)
	}
	http.SetCookie(c.Response().Writer, &cookies)

	// Hapus refresh_token cookie jika diperlukan
	cookiesRefresh := http.Cookie{
		Name:     "refresh_token", // Nama cookie untuk refresh_token
		Value:    "",
		Path:     "/dashboard", // Path yang sama
		Expires:  time.Now().Add(-7 * 24 * time.Hour),
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(c.Response().Writer, &cookiesRefresh)

	// Kirimkan response sukses (tidak ada data tambahan)
	err := res.SuccessResponse(nil).Send(c)
	if err != nil {
		// Log error jika gagal mengirimkan response
		fmt.Println("Error sending response:", err)
	}

	return err
}
