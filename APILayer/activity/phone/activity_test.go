package phone

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata
var connectionData = ``

func getActivityMetadata() *activity.Metadata {
	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}
		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}
	return activityMetadata
}

func TestActivityRegistration(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//Use your API Layer information
	dummyConnectionData := make(map[string]interface{})
	dummyConnectionSettings := make([]interface{}, 1)

	accesskeyId := make(map[string]interface{})
	accesskeyId["name"] = "accessKeyId"
	accesskeyId["value"] = "<YOUR ACCESS KEY ID>"
	dummyConnectionSettings[0] = accesskeyId

	dummyConnectionData["settings"] = dummyConnectionSettings
	tc.SetInput(ivConnection, dummyConnectionData)

	//setup attrs
	tc.SetInput("phone", "498003303000")

	_, err := act.Eval(tc)
	assert.Nil(t, err)

	result := tc.GetOutput("valid")
	assert.Equal(t, true, result)

	t.Log(result)
}

func TestInvalidEval(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//Use your API Layer information
	dummyConnectionData := make(map[string]interface{})
	dummyConnectionSettings := make([]interface{}, 1)

	accesskeyId := make(map[string]interface{})
	accesskeyId["name"] = "accessKeyId"
	accesskeyId["value"] = "<YOUR ACCESS KEY ID>"
	dummyConnectionSettings[0] = accesskeyId

	dummyConnectionData["settings"] = dummyConnectionSettings
	tc.SetInput(ivConnection, dummyConnectionData)

	//setup attrs
	tc.SetInput("phone", "49800330300099")

	_, err := act.Eval(tc)
	assert.Nil(t, err)

	result := tc.GetOutput("valid")
	assert.Equal(t, false, result)

	t.Log(result)
}
