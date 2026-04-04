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
	PhoneCode   string  `json:"phoneCode"`   // Tambahan kode telepon
	PhonePrefix string  `json:"phonePrefix"` // Tambahan prefix telepon
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

// Map untuk kode telepon berdasarkan negara
var phoneCodeMap = map[string]string{
	"Afghanistan":                      "93",
	"South Africa":                     "27",
	"Central African Republic":         "236",
	"Albania":                          "355",
	"Algeria":                          "213",
	"United States":                    "1",
	"Andorra":                          "376",
	"Angola":                           "244",
	"Antigua and Barbuda":              "1-268",
	"Saudi Arabia":                     "966",
	"Argentina":                        "54",
	"Armenia":                          "374",
	"Australia":                        "61",
	"Austria":                          "43",
	"Azerbaijan":                       "994",
	"Bahamas":                          "1-242",
	"Bahrain":                          "973",
	"Bangladesh":                       "880",
	"Barbados":                         "1-246",
	"Netherlands":                      "31",
	"Belarus":                          "375",
	"Belgium":                          "32",
	"Belize":                           "501",
	"Benin":                            "229",
	"Bhutan":                           "975",
	"Bolivia":                          "591",
	"Bosnia and Herzegovina":           "387",
	"Botswana":                         "267",
	"Brazil":                           "55",
	"United Kingdom":                   "44",
	"Brunei":                           "673",
	"Bulgaria":                         "359",
	"Burkina Faso":                     "226",
	"Burundi":                          "257",
	"Czech Republic":                   "420",
	"Chad":                             "235",
	"Chile":                            "56",
	"China":                            "86",
	"Denmark":                          "45",
	"Djibouti":                         "253",
	"Dominica":                         "1-767",
	"Ecuador":                          "593",
	"El Salvador":                      "503",
	"Eritrea":                          "291",
	"Estonia":                          "372",
	"Ethiopia":                         "251",
	"Fiji":                             "679",
	"Philippines":                      "63",
	"Finland":                          "358",
	"Gabon":                            "241",
	"Gambia":                           "220",
	"Georgia":                          "995",
	"Ghana":                            "233",
	"Grenada":                          "1-473",
	"Guatemala":                        "502",
	"Guinea":                           "224",
	"Guinea-Bissau":                    "245",
	"Equatorial Guinea":                "240",
	"Guyana":                           "592",
	"Haiti":                            "509",
	"Honduras":                         "504",
	"Hungary":                          "36",
	"Hong Kong":                        "852",
	"India":                            "91",
	"Indonesia":                        "62",
	"Iraq":                             "964",
	"Iran":                             "98",
	"Ireland":                          "353",
	"Iceland":                          "354",
	"Israel":                           "972",
	"Italy":                            "39",
	"Jamaica":                          "1-876",
	"Japan":                            "81",
	"Germany":                          "49",
	"Jordan":                           "962",
	"Cambodia":                         "855",
	"Cameroon":                         "237",
	"Canada":                           "1",
	"Kazakhstan":                       "7",
	"Kenya":                            "254",
	"Kyrgyzstan":                       "996",
	"Kiribati":                         "686",
	"Colombia":                         "57",
	"Comoros":                          "269",
	"Republic of the Congo":            "243",
	"South Korea":                      "82",
	"North Korea":                      "850",
	"Costa Rica":                       "506",
	"Croatia":                          "385",
	"Cuba":                             "53",
	"Kuwait":                           "965",
	"Laos":                             "856",
	"Latvia":                           "371",
	"Lebanon":                          "961",
	"Lesotho":                          "266",
	"Liberia":                          "231",
	"Libya":                            "218",
	"Liechtenstein":                    "423",
	"Lithuania":                        "370",
	"Luxembourg":                       "352",
	"Madagascar":                       "261",
	"Macau":                            "853",
	"Macedonia":                        "389",
	"Maldives":                         "960",
	"Malawi":                           "265",
	"Malaysia":                         "60",
	"Mali":                             "223",
	"Malta":                            "356",
	"Morocco":                          "212",
	"Marshall Islands":                 "692",
	"Mauritania":                       "222",
	"Mauritius":                        "230",
	"Mexico":                           "52",
	"Egypt":                            "20",
	"Micronesia":                       "691",
	"Moldova":                          "373",
	"Monaco":                           "377",
	"Mongolia":                         "976",
	"Montenegro":                       "382",
	"Mozambique":                       "258",
	"Myanmar":                          "95",
	"Namibia":                          "264",
	"Nauru":                            "674",
	"Nepal":                            "977",
	"Niger":                            "227",
	"Nigeria":                          "234",
	"Nicaragua":                        "505",
	"Norway":                           "47",
	"Oman":                             "968",
	"Pakistan":                         "92",
	"Palau":                            "680",
	"Panama":                           "507",
	"Ivory Coast":                      "225",
	"Papua New Guinea":                 "675",
	"Paraguay":                         "595",
	"France":                           "33",
	"Peru":                             "51",
	"Poland":                           "48",
	"Portugal":                         "351",
	"Qatar":                            "974",
	"Congo":                            "242",
	"Dominican Republic":               "1-809",
	"Romania":                          "40",
	"Russia":                           "7",
	"Rwanda":                           "250",
	"Saint Kitts and Nevis":            "1-869",
	"Saint Lucia":                      "1-758",
	"Saint Vincent and the Grenadines": "1-784",
	"Samoa":                            "685",
	"San Marino":                       "378",
	"Sao Tome and Principe":            "239",
	"New Zealand":                      "64",
	"Senegal":                          "221",
	"Serbia":                           "381",
	"Seychelles":                       "248",
	"Sierra Leone":                     "232",
	"Singapore":                        "65",
	"Cyprus":                           "357",
	"Slovenia":                         "386",
	"Slovakia":                         "421",
	"Solomon Islands":                  "677",
	"Somalia":                          "252",
	"Spain":                            "34",
	"Sri Lanka":                        "94",
	"Sudan":                            "249",
	"South Sudan":                      "211",
	"Syria":                            "963",
	"Suriname":                         "597",
	"Swaziland":                        "268",
	"Sweden":                           "46",
	"Switzerland":                      "41",
	"Tajikistan":                       "992",
	"Cape Verde":                       "238",
	"Tanzania":                         "255",
	"Taiwan":                           "886",
	"Thailand":                         "66",
	"Timor-Leste":                      "670",
	"Togo":                             "228",
	"Tonga":                            "676",
	"Trinidad and Tobago":              "1-868",
	"Tunisia":                          "216",
	"Turkey":                           "90",
	"Turkmenistan":                     "993",
	"Tuvalu":                           "688",
	"Uganda":                           "256",
	"Ukraine":                          "380",
	"United Arab Emirates":             "971",
	"Uruguay":                          "598",
	"Uzbekistan":                       "998",
	"Vanuatu":                          "678",
	"Venezuela":                        "58",
	"Vietnam":                          "84",
	"Yemen":                            "967",
	"Greece":                           "30",
	"Zambia":                           "260",
	"Zimbabwe":                         "263",
}

// Fungsi untuk mendapatkan kode telepon berdasarkan nama negara
func getPhoneCode(country string) string {
	if code, exists := phoneCodeMap[country]; exists {
		return code
	}
	return "" // Return empty string jika negara tidak ditemukan
}

// Fungsi untuk mendapatkan prefix telepon (format +XX)
func getPhonePrefix(country string) string {
	code := getPhoneCode(country)
	if code == "" {
		return ""
	}
	return "+" + code
}

func GetLocationData(c echo.Context) error {
	ip := c.RealIP() // Ambil IP client dari header/request

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

	// Tambahkan kode telepon berdasarkan negara
	location.PhoneCode = getPhoneCode(location.Country)
	location.PhonePrefix = getPhonePrefix(location.Country)

	return c.JSON(http.StatusOK, location)
}
