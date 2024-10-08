/*
Onfido API v3.6

The Onfido API (v3.6)

API version: v3.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// DocumentTypes the model 'DocumentTypes'
type DocumentTypes string

// List of document_types
const (
	PASSPORT DocumentTypes = "passport"
	DRIVING_LICENCE DocumentTypes = "driving_licence"
	NATIONAL_IDENTITY_CARD DocumentTypes = "national_identity_card"
	RESIDENCE_PERMIT DocumentTypes = "residence_permit"
	PASSPORT_CARD DocumentTypes = "passport_card"
	TAX_ID DocumentTypes = "tax_id"
	VISA DocumentTypes = "visa"
	VOTER_ID DocumentTypes = "voter_id"
	RESIDENCE_STATUS_DOCUMENT DocumentTypes = "residence_status_document"
	POSTAL_IDENTITY_CARD DocumentTypes = "postal_identity_card"
	SOCIAL_SECURITY_CARD DocumentTypes = "social_security_card"
	WORK_PERMIT DocumentTypes = "work_permit"
	ASYLUM_REGISTRATION_CARD DocumentTypes = "asylum_registration_card"
	NATIONAL_HEALTH_INSURANCE_CARD DocumentTypes = "national_health_insurance_card"
	MUNICIPALITY_IDENTITY_CARD DocumentTypes = "municipality_identity_card"
	PRIVATE_OPERATORS_CARD DocumentTypes = "private_operators_card"
	PROOF_OF_CITIZENSHIP DocumentTypes = "proof_of_citizenship"
	SERVICE_ID_CARD DocumentTypes = "service_id_card"
	IMMIGRATION_STATUS_DOCUMENT DocumentTypes = "immigration_status_document"
	INDIGENOUS_CARD DocumentTypes = "indigenous_card"
	VEHICLE_REGISTRATION_CARD DocumentTypes = "vehicle_registration_card"
	CERTIFICATE_OF_NATURALISATION DocumentTypes = "certificate_of_naturalisation"
	PROFESSIONAL_QUALIFICATION_CARD DocumentTypes = "professional_qualification_card"
	CONSULAR_ID DocumentTypes = "consular_id"
	INTERNATIONAL_DRIVING_LICENCE DocumentTypes = "international_driving_licence"
)

// All allowed values of DocumentTypes enum
var AllowedDocumentTypesEnumValues = []DocumentTypes{
	"passport",
	"driving_licence",
	"national_identity_card",
	"residence_permit",
	"passport_card",
	"tax_id",
	"visa",
	"voter_id",
	"residence_status_document",
	"postal_identity_card",
	"social_security_card",
	"work_permit",
	"asylum_registration_card",
	"national_health_insurance_card",
	"municipality_identity_card",
	"private_operators_card",
	"proof_of_citizenship",
	"service_id_card",
	"immigration_status_document",
	"indigenous_card",
	"vehicle_registration_card",
	"certificate_of_naturalisation",
	"professional_qualification_card",
	"consular_id",
	"international_driving_licence",
}

func (v *DocumentTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DocumentTypes(value)
	for _, existing := range AllowedDocumentTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DocumentTypes", value)
}

// NewDocumentTypesFromValue returns a pointer to a valid DocumentTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDocumentTypesFromValue(v string) (*DocumentTypes, error) {
	ev := DocumentTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DocumentTypes: valid values are %v", v, AllowedDocumentTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DocumentTypes) IsValid() bool {
	for _, existing := range AllowedDocumentTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to document_types value
func (v DocumentTypes) Ptr() *DocumentTypes {
	return &v
}

type NullableDocumentTypes struct {
	value *DocumentTypes
	isSet bool
}

func (v NullableDocumentTypes) Get() *DocumentTypes {
	return v.value
}

func (v *NullableDocumentTypes) Set(val *DocumentTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentTypes(val *DocumentTypes) *NullableDocumentTypes {
	return &NullableDocumentTypes{value: val, isSet: true}
}

func (v NullableDocumentTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

