
package phone

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Numverify struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}

const (
	ivConnection          = "apiConnection"
	ivPhone               = "phone"
	ovValid               = "valid"
	ovNumber              = "number"
	ovLocalFormat         = "local_format"
	ovInternationalFormat = "international_format"
	ovCountryPrefix       = "country_prefix"
	ovCountryCode         = "country_code"
	ovCountryName         = "country_name"
	ovLocation            = "location"
	ovCarrier             = "carrier"
	ovLineType            = "line_type"
)

var activityLog = logger.GetLogger("apilayer-activity-phone")

type phoneActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &phoneActivity{metadata: metadata}
}

func (a *phoneActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *phoneActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing Phone validation activity")
	//Read Inputs
	if context.GetInput(ivConnection) == nil {
		return false, activity.NewError("APILayer connection is not configured", "APILAyer-Phone-4001", nil)
	}
	if context.GetInput(ivPhone) == nil {
		// Phone string is not configured
		// return error to the engine
		return false, activity.NewError("APILayer Phone string is not configured", "APILayer-Phone-4002", nil)
	}
	Phonestr := context.GetInput(ivPhone).(string)
	activityLog.Info("Phone: " + Phonestr)

	//Read connection details
	connectionInfo := context.GetInput(ivConnection).(map[string]interface{})
	connectionSettings := connectionInfo["settings"].([]interface{})

	var accesskey string
	for _, v := range connectionSettings {
		setting := v.(map[string]interface{})
		if setting["name"] == "accessKeyId" {
			accesskey = setting["value"].(string)
		}
	}

	// execute validation - Start
	safePhone := url.QueryEscape(Phonestr)
	url := fmt.Sprintf("http://apilayer.net/api/validate?access_key="+accesskey+"&number=%s", safePhone)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Numverify

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	context.SetOutput(ovValid, record.Valid)
	context.SetOutput(ovNumber, record.Number)
	context.SetOutput(ovLocalFormat, record.LocalFormat)
	context.SetOutput(ovInternationalFormat, record.InternationalFormat)
	context.SetOutput(ovCountryPrefix, record.CountryPrefix)
	context.SetOutput(ovCountryName, record.CountryName)
	context.SetOutput(ovLocation, record.Location)
	context.SetOutput(ovCarrier, record.Carrier)
	context.SetOutput(ovLineType, record.LineType)
	return true, nil
}
