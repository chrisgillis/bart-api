package bart

type AdvisoryService struct{}

func (s *AdvisoryService) Bsa(r *BsaRequest) (resp BsaResponse, err error) {
	err = makeRequest("bsa.aspx", "bsa", r, &resp)
	return
}

func (s *AdvisoryService) TrainCount(r *CountRequest) (resp CountResponse, err error) {
	err = makeRequest("bsa.aspx", "count", r, &resp)
	return
}

func (s *AdvisoryService) ElevatorStatus(r *ElevRequest) (resp BsaResponse, err error) {
	err = makeRequest("bsa.aspx", "elev", r, &resp)
	return
}

func (s *AdvisoryService) Help(r *HelpRequest) (resp HelpResponse, err error) {
	err = makeRequest("bsa.aspx", "help", r, &resp)
	return
}

type BsaRequest struct {
	Orig string `url:"orig"`
}

type BsaResponse struct {
	BartResponseMeta
	Bsa Bsa `xml:"bsa"`
}

type Bsa struct {
	Station     string `xml:"station"`
	Description string `xml:"description"`
	SmsText     string `xml:"sms_text"`
}

type CountRequest struct{}

type CountResponse struct {
	BartResponseMeta
	TrainCount int `xml:"traincount"`
}

type ElevRequest struct{}

type HelpRequest struct{}

type Message struct {
	Help string `xml:"help"`
}

type HelpResponse struct {
	BartResponseMeta
	Message Message `xml:"message"`
}
