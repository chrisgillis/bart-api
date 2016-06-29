package bart

type StationInfoService struct{}

func (s *StationInfoService) Stations(r *StationListRequest) (resp StationListResponse, err error) {
	err = makeRequest("stn.aspx", "stns", r, &resp)
	return
}

func (s *StationInfoService) StationAccess(r *StationAccessRequest) (resp StationAccessResponse, err error) {
	err = makeRequest("stn.aspx", "stnaccess", r, &resp)
	return
}

func (s *StationInfoService) StationInfo(r *StationInfoRequest) (resp StationInfoResponse, err error) {
	err = makeRequest("stn.aspx", "stninfo", r, &resp)
	return
}

func (s *StationInfoService) Help(r *HelpRequest) (resp HelpResponse, err error) {
	err = makeRequest("stn.aspx", "help", r, &resp)
	return
}

type StationListRequest struct{}

type StationListResponse struct {
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

type StationInfoRequest struct {
	Origin string `url:"orig"`
}

type StationInfoResponse struct {
	BartResponseMeta
	Station StationDetail `xml:"stations>station"`
}

type StationDetail struct {
	Station
	NorthRoutes    []string `xml:"north_routes>route"`
	SouthRoutes    []string `xml:"south_routes>route"`
	NorthPlatforms []string `xml:"north_platforms>platform"`
	SouthPlatforms []string `xml:"south_platforms>platform"`
	PlatformInfo   string   `xml:"platform_info"`
	Intro          string   `xml:"intro"`
	CrossStreet    string   `xml:"cross_street"`
	Food           string   `xml:"food"`
	Shopping       string   `xml:"shopping"`
	Attraction     string   `xml:"attraction"`
}
