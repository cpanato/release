/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cve

import (
	"errors"
	"fmt"
	"regexp"

	cvss "github.com/goark/go-cvss/v3/metric"
)

// CVE Information of a linked CVE vulnerability.
type CVE struct {
	ID            string  `json:"id"                 yaml:"id"`                 // CVE ID, eg CVE-2019-1010260
	Title         string  `json:"title"              yaml:"title"`              // Title of the vulnerability
	Description   string  `json:"description"        yaml:"description"`        // Description text of the vulnerability
	TrackingIssue string  `json:"issue"              yaml:"issue"`              // Link to the vulnerability tracking issue (url, optional)
	CVSSVector    string  `json:"vector"             yaml:"vector"`             // Full CVSS vector string, CVSS:3.1/AV:N/AC:H/PR:H/UI:R/S:U/C:H/I:H/A:H
	CVSSScore     float32 `json:"score"              yaml:"score"`              // Numeric CVSS score (eg 6.2)
	CVSSRating    string  `json:"rating"             yaml:"rating"`             // Severity bucket (eg Medium)
	CalcLink      string  `json:"calclink,omitempty" yaml:"calclink,omitempty"` // Link to the CVE calculator (automatic)
	LinkedPRs     []int   `json:"pullrequests"`                                 // List of linked PRs (to remove them from the release notes doc)
}

// ReadRawInterface populates the CVE data struct from the raw array
// as returned by the YAML parser.
func (cve *CVE) ReadRawInterface(cvedata interface{}) error {
	if val, ok := cvedata.(map[interface{}]interface{})["id"].(string); ok {
		cve.ID = val
	}

	if val, ok := cvedata.(map[interface{}]interface{})["title"].(string); ok {
		cve.Title = val
	}

	if val, ok := cvedata.(map[interface{}]interface{})["issue"].(string); ok {
		cve.TrackingIssue = val
	}

	if val, ok := cvedata.(map[interface{}]interface{})["vector"].(string); ok {
		cve.CVSSVector = val
	}

	if val, ok := cvedata.(map[interface{}]interface{})["score"].(float64); ok {
		cve.CVSSScore = float32(val)
	}

	if val, ok := cvedata.(map[interface{}]interface{})["rating"].(string); ok {
		cve.CVSSRating = val
	}

	if val, ok := cvedata.(map[interface{}]interface{})["description"].(string); ok {
		cve.Description = val
	}
	// Linked PRs is a list of the PR IDs
	if val, ok := cvedata.(map[interface{}]interface{})["linkedPRs"].([]interface{}); ok {
		cve.LinkedPRs = []int{}

		for _, prid := range val {
			if prid, ok := prid.(int); ok {
				cve.LinkedPRs = append(cve.LinkedPRs, prid)
			}
		}
	}

	return nil
}

// Validate checks the data defined in a CVE map is complete and valid.
func (cve *CVE) Validate() (err error) {
	// Verify that rating is defined and a known string
	if cve.CVSSRating == "" {
		return errors.New("missing CVSS rating from CVE data")
	}

	// Check rating is a valid string
	if _, ok := map[string]bool{
		"None": true, "Low": true, "Medium": true, "High": true, "Critical": true,
	}[cve.CVSSRating]; !ok {
		return errors.New("invalid CVSS rating")
	}

	// Check vector string is not empty
	if cve.CVSSVector == "" {
		return errors.New("string CVSS vector missing from CVE data")
	}

	var bm cvss.Metrics
	// Parse the vector string to make sure it is well formed
	if len(cve.CVSSVector) == 44 {
		bm, err = cvss.NewBase().Decode(cve.CVSSVector)
	} else {
		bm, err = cvss.NewTemporal().Decode(cve.CVSSVector)
	}

	if err != nil {
		return fmt.Errorf("parsing CVSS vector string: %w", err)
	}

	cve.CalcLink = fmt.Sprintf(
		"https://www.first.org/cvss/calculator/%s#%s", bm.BaseMetrics().Ver.String(), cve.CVSSVector,
	)

	if cve.CVSSScore == 0 {
		return errors.New("missing CVSS score from CVE data")
	}

	if cve.CVSSScore < 0 || cve.CVSSScore > 10 {
		return errors.New("out of range CVSS score, should be 0.0 - 10.0")
	}

	if err := ValidateID(cve.ID); err != nil {
		return fmt.Errorf("checking CVE ID: %w", err)
	}

	// Title and description must not be empty
	if cve.Title == "" {
		return errors.New("title missing from CVE data")
	}

	if cve.Description == "" {
		return errors.New("missing CVE description from CVE data")
	}

	return nil
}

// ValidateID checks if a CVE IS string is valid.
func ValidateID(cveID string) error {
	if cveID == "" {
		return errors.New("empty CVE ID string")
	}

	// Verify that the CVE ID is well formed
	if !regexp.MustCompile(CVEIDRegExp).MatchString(cveID) {
		return errors.New("not well formed CVS ID")
	}

	return nil
}
