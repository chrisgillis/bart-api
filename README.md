# bart-api
A Golang package for the BART API. Currently supports Advisories, Station Info, and Real Time Estimates. 

# Usage
```
	bartClient := bart.New()

	bsaResp, _ := bartClient.Advisories.Bsa(&bart.BsaRequest{Orig: "all"})
	fmt.Printf("Advisories.Bsa = %#v\n\n", bsaResp)

	countResp, _ := bartClient.Advisories.TrainCount(&bart.CountRequest{})
	fmt.Printf("Advisories.Count = %#v\n\n", countResp)

	elevResp, _ := bartClient.Advisories.ElevatorStatus(&bart.ElevRequest{})
	fmt.Printf("Advisories.Elev = %#v\n\n", elevResp)

	helpResp, _ := bartClient.Advisories.Help(&bart.HelpRequest{})
	fmt.Printf("Advisories.Help = %#v\n\n", helpResp)

	stationResp, _ := bartClient.StationInfo.Stations(&bart.StationListRequest{})
	fmt.Printf("StationInfo.Stations = %#v\n\n", stationResp)

	stationAccessResp, _ := bartClient.StationInfo.StationAccess(&bart.StationAccessRequest{Origin: "12th", LegendEnabled: 1})
	fmt.Printf("StationInfo.StationAccess = %#v\n\n", stationAccessResp)

	stationInfoResp, _ := bartClient.StationInfo.StationInfo(&bart.StationInfoRequest{Origin: "12th"})
	fmt.Printf("StationInfo.StationInfo = %#v\n\n", stationInfoResp)

	stationHelpResp, _ := bartClient.StationInfo.Help(&bart.HelpRequest{})
	fmt.Printf("StationInfo.Help = %#v\n\n", stationHelpResp)

	deptResp, _ := bartClient.RealTimeEstimates.Departures(&bart.EstimateDepartureRequest{Origin: "12th"})
	fmt.Printf("RealTimeEstimates.Departures = %#v\n\n", deptResp)

	estimateHelpResp, _ := bartClient.RealTimeEstimates.Help(&bart.HelpRequest{})
	fmt.Printf("RealTimeEstimates.Help = %#v\n\n", estimateHelpResp)
```
