package bart

type RealTimeEstimateService struct{}

func (s *RealTimeEstimateService) Departures(r *EstimateDepartureRequest) (resp EstimateDepartureResponse, err error) {
	err = makeRequest("etd.aspx", "etd", r, &resp)
	return
}

type EstimateDepartureRequest struct {
	Origin    string `url:"orig"`
	Platform  string `url:"plat,omitempty"`
	Direction string `url:"dir,omitempty"`
}

type EstimateDepartureResponse struct {
	BartResponseMeta
	Station DepartureStation `xml:"station"`
}

type DepartureStation struct {
	Name       string      `xml:"name"`
	Abbr       string      `xml:"abbr"`
	Departures []Departure `xml:"etd"`
}

type Departure struct {
	Destination  string     `xml:"destination"`
	Abbreviation string     `xml:"abbreviation"`
	Limited      int        `xml:"limited"`
	Estimates    []Estimate `xml:"estimate"`
}

type Estimate struct {
	Minutes   int    `xml:"minutes"`
	Platform  int    `xml:"platform"`
	Direction string `xml:"direction"`
	Length    int    `xml:"length"`
	Color     string `xml:"color"`
	HexColor  string `xml:"hexcolor"`
	BikeFlag  int    `xml:"bikeflag"`
}
