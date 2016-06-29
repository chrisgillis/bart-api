package bart

type StationInfoService struct{}

func (s *StationInfoService) Stations(r *StationInfoRequest) (resp StationInfoResponse, err error) {
	err = makeRequest("stn.aspx", "stns", r, &resp)
	return
}

func (s *StationInfoService) StationAccess(r *StationAccessRequest) (resp StationAccessResponse, err error) {
	err = makeRequest("stn.aspx", "stnaccess", r, &resp)
	return
}

type StationInfoRequest struct{}

type StationInfoResponse struct {
	BartResponseMeta
	Stations []Station `xml:"stations>station"`
}

type Station struct {
	Name      string  `xml:"name"`
	Abbr      string  `xml:"abbr"`
	Latitude  float64 `xml:"gtfs_latitude"`
	Longitude float64 `xml:"gtfs_longitude"`
	Address   string  `xml:"address"`
	City      string  `xml:"city"`
	County    string  `xml:"county"`
	State     string  `xml:"state"`
	ZipCode   string  `xml:"zipcode"`
}

type StationAccessRequest struct {
	Origin        string `url:"orig"`
	LegendEnabled int    `url:"l"`
}

type StationAccessResponse struct {
	BartResponseMeta
	Station StationAccess `xml:"stations>station"`
	Legend  string        `xml:"message>legend"`
}

type StationAccess struct {
	Name            string `xml:"name"`
	Abbr            string `xml:"abbr"`
	Link            string `xml:"link"`
	ParkingFlag     int    `xml:"parking_flag,attr"`
	BikeFlag        int    `xml:"bike_flag,attr"`
	BikeStationFlag int    `xml:"bike_station_flag,attr"`
	LockerFlag      int    `xml:"locker_flag,attr"`
	Entering        string `xml:"entering"`
	Exiting         string `xml:"exiting"`
	Parking         string `xml:"parking"`
	FillTime        string `xml:"fill_time"`
	CarShare        string `xml:"car_share"`
	Lockers         string `xml:"lockers"`
	BikeStationText string `xml:"bike_station_text"`
	Destinations    string `xml:"destinations"`
	TransitInfo     string `xml:"transit_info"`
}
