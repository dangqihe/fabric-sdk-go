package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

/************************************************
When the subjectAltName extension contains a URI, the name MUST be
stored in the uniformResourceIdentifier (an IA5String).
************************************************/

import (
	"unicode"

	"github.com/studyzy/fabric-sdk-go/third_party/github.com/zmap/zcrypto/x509"
	"github.com/studyzy/fabric-sdk-go/third_party/github.com/zmap/zlint/util"
)

type extSANURINotIA5 struct{}

func (l *extSANURINotIA5) Initialize() error {
	return nil
}

func (l *extSANURINotIA5) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *extSANURINotIA5) Execute(c *x509.Certificate) *LintResult {
	for _, uri := range c.URIs {
		for _, c := range uri {
			if c > unicode.MaxASCII {
				return &LintResult{Status: Error}
			}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_uri_not_ia5",
		Description:   "When subjectAlternateName contains a URI, the name MUST be an IA5 string",
		Citation:      "RFC5280: 4.2.1.6",
		Source:        RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          &extSANURINotIA5{},
	})
}