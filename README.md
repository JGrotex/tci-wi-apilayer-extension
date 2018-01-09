# TCI WI APILayer.com Extension

[![Go Report Card](https://goreportcard.com/badge/github.com/JGrotex/tci-wi-apilayer-extension)](https://goreportcard.com/report/github.com/JGrotex/tci-wi-apilayer-extension) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

first Version with just International Phonenumber validation, more to come ...

As the Connector is available any other offering of APILayer.Com should be easy to create.

Attached ZIP contains the first release v.1.3 and can just uploaded under TIBCO Cloud Integration Extensions

Please create your own 'free' Access Key on http://www.APILayer.com to enter into the Connector Details.

## Activities
available Activities so far
### Phone Number Validation
Input
- Phone Number string

Output
- Valid               bool   `json:"valid"`
- Number              string `json:"number"`
- LocalFormat         string `json:"local_format"`
- InternationalFormat string `json:"international_format"`
- CountryPrefix       string `json:"country_prefix"`
- CountryCode         string `json:"country_code"`
- CountryName         string `json:"country_name"`
- Location            string `json:"location"`
- Carrier             string `json:"carrier"`
- LineType            string `json:"line_type"`
  
Sample Input Data
498003303000 or another invalid one 49800330300099

Sample Output Data

``json:
{"valid":true,"number":"498003303000","local_format":"08003303000","international_format":"+498003303000","country_prefix":"+49","country_code":"DE","country_name":"Germany (Federal Republic of)","location":"","carrier":"","line_type":"toll_free"}
``

<hr>
<sub><b>Note:</b> more TCI Extensions can be found here: https://tibcosoftware.github.io/tci-awesome/ </sub>
