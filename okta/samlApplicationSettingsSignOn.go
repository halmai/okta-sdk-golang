/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// Code generated by okta openapi generator. DO NOT EDIT.

package okta

type SamlApplicationSettingsSignOn struct {
	AcsEndpoints              []*AcsEndpoint            `json:"acsEndpoints,omitempty"`
	AllowMultipleAcsEndpoints *bool                     `json:"allowMultipleAcsEndpoints,omitempty"`
	AssertionSigned           *bool                     `json:"assertionSigned,omitempty"`
	AttributeStatements       []*SamlAttributeStatement `json:"attributeStatements"`
	Audience                  string                    `json:"audience,omitempty"`
	AudienceOverride          string                    `json:"audienceOverride"`
	AuthnContextClassRef      string                    `json:"authnContextClassRef,omitempty"`
	DefaultRelayState         string                    `json:"defaultRelayState"`
	Destination               string                    `json:"destination,omitempty"`
	DestinationOverride       string                    `json:"destinationOverride"`
	DigestAlgorithm           string                    `json:"digestAlgorithm,omitempty"`
	HonorForceAuthn           *bool                     `json:"honorForceAuthn,omitempty"`
	IdpIssuer                 string                    `json:"idpIssuer,omitempty"`
	Recipient                 string                    `json:"recipient,omitempty"`
	RecipientOverride         string                    `json:"recipientOverride"`
	RequestCompressed         *bool                     `json:"requestCompressed,omitempty"`
	ResponseSigned            *bool                     `json:"responseSigned,omitempty"`
	SignatureAlgorithm        string                    `json:"signatureAlgorithm,omitempty"`
	Slo                       *SingleLogout             `json:"slo,omitempty"`
	SpCertificate             *SpCertificate            `json:"spCertificate,omitempty"`
	SpIssuer                  string                    `json:"spIssuer,omitempty"`
	SsoAcsUrl                 string                    `json:"ssoAcsUrl,omitempty"`
	SsoAcsUrlOverride         string                    `json:"ssoAcsUrlOverride"`
	SubjectNameIdFormat       string                    `json:"subjectNameIdFormat,omitempty"`
	SubjectNameIdTemplate     string                    `json:"subjectNameIdTemplate,omitempty"`
}

func NewSamlApplicationSettingsSignOn() *SamlApplicationSettingsSignOn {
	return &SamlApplicationSettingsSignOn{}
}

func (a *SamlApplicationSettingsSignOn) IsApplicationInstance() bool {
	return true
}
