package speedbandparseractivity

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-speedbandparseractivity")

// MyActivity is a stub for your Activity implementation
type XMLParserActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &XMLParserActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *XMLParserActivity) Metadata() *activity.Metadata {
	return a.metadata
}

//XSD

/*type TrafficInfo struct {
	XMLName            xml.Name         `xml:"TrafficInfo" json:"-"`
	SpeedBandInfoList  []SpeedBandInfo  `xml:"SpeedBandInfo" json:"SpeedBandInfo"`
	TDataList          []TData          `xml:"TData" json:"TData"`
}

type SpeedBandInfoList struct {
	XMLName        xml.Name `xml:"SpeedBandInfo" json:"-"`
	Band           string   `xml:"Band" json:"Band"`
	MinimumSpeed   string   `xml:"MinimumSpeed" json:"MinimumSpeed"`
	MaximumSpeed   string   `xml:"MaximumSpeed" json:"MaximumSpeed"`
}

type TDataList struct {
	XMLName     xml.Name `xml:"TData" json:"-"`
	LinkID      string   `xml:"LinkID" json:"LinkID"`
	SpeedBand   string   `xml:"SpeedBand" json:"SpeedBand"`
}
*/

type TrafficInfo struct {
		SpeedBandInfo []struct {
			Band         string `json:"Band"`
			MinimumSpeed string `json:"MinimumSpeed"`
			MaximumSpeed string `json:"MaximumSpeed"`
		} `json:"SpeedBandInfo"`
		TData []struct {
			LinkID    string `json:"LinkID"`
			SpeedBand string `json:"SpeedBand"`
		} `json:"TData"`
}

// end of XSD

// Eval implements activity.Activity.Eval
func (a *XMLParserActivity) Eval(ctx activity.Context) (done bool, err error) {

	XMLString := ctx.GetInput("xmlString").(string)

	activityLog.Debugf("XML String is : [%s]", XMLString)
	//fmt.Println("XML String is : ", XMLString)

	if len(XMLString) == 0 {
		activityLog.Debugf("value in the field is empty ")
		//fmt.Println("value in  the field is empty ")

	}
	//	XMLString = (string(XMLString))

	xml_data := TrafficInfo{}
	err = xml.Unmarshal([]byte(XMLString), &xml_data)

	jsondata, _ := json.Marshal(xml_data)
	if err != nil {
		activityLog.Debugf("Error ", err)
		fmt.Println("error: ", err)
		return
	}

	//fmt.Println(" JSON String ")
	//fmt.Println(string(jsondata))

	// Set the output as part of the context
	activityLog.Debugf("Activity has parsed SpeedBand XML Successfully")
	fmt.Println("Activity has parsed SpeedBand XML Successfully")

	ctx.SetOutput("output", string(jsondata))

	return true, nil
}
