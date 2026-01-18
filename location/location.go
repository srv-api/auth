package location

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IPGeoLocation struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
	Query       string  `json:"query"`
}

type Local struct {
	ID string `json:"id"`
}
type take struct {
	ID string `json:"id"`
}
type Stake struct {
	ID string `json:"id"`
}

func GetLocationData(c echo.Context) error {
	ip := c.RealIP() // Ini ambil IP client dari header/request

	// Panggil ip-api dengan IP tersebut
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)

	resp, err := http.Get(url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch IP data"})
	}
	defer resp.Body.Close()

	var location IPGeoLocation
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to decode IP data"})
	}

	return c.JSON(http.StatusOK, location)
}
