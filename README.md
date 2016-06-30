# bart-api
A Golang package for the BART API. Currently supports Advisories, Station Info, and Real Time Estimates. 

# Usage
```
bartClient := bart.New()
 12 
 13     bsaResp, _ := bartClient.Advisories.Bsa(&bart.BsaRequest{Orig: "all"})
 14     fmt.Printf("Advisories.Bsa = %#v\n\n", bsaResp)
 15 
 16     countResp, _ := bartClient.Advisories.TrainCount(&bart.CountRequest{})
 17     fmt.Printf("Advisories.Count = %#v\n\n", countResp)
 18 
 19     elevResp, _ := bartClient.Advisories.ElevatorStatus(&bart.ElevRequest{})
 20     fmt.Printf("Advisories.Elev = %#v\n\n", elevResp)
 21 
 22     helpResp, _ := bartClient.Advisories.Help(&bart.HelpRequest{})
 23     fmt.Printf("Advisories.Help = %#v\n\n", helpResp)
 24 
 25     stationResp, _ := bartClient.StationInfo.Stations(&bart.StationListRequest{})
 26     fmt.Printf("StationInfo.Stations = %#v\n\n", stationResp)
 27 
 28     stationAccessResp, _ := bartClient.StationInfo.StationAccess(&bart.StationAccess
 29     fmt.Printf("StationInfo.StationAccess = %#v\n\n", stationAccessResp)
 30 
 31     stationInfoResp, _ := bartClient.StationInfo.StationInfo(&bart.StationInfoReques
 32     fmt.Printf("StationInfo.StationInfo = %#v\n\n", stationInfoResp)
 33 
 34     stationHelpResp, _ := bartClient.StationInfo.Help(&bart.HelpRequest{})
 35     fmt.Printf("StationInfo.Help = %#v\n\n", stationHelpResp)
 36 
 37     deptResp, _ := bartClient.RealTimeEstimates.Departures(&bart.EstimateDepartureRe
 38     fmt.Printf("RealTimeEstimates.Departures = %#v\n\n", deptResp)
 39 
 40     estimateHelpResp, _ := bartClient.RealTimeEstimates.Help(&bart.HelpRequest{})
 41     fmt.Printf("RealTimeEstimates.Help = %#v\n\n", estimateHelpResp)
```
