// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LinksResponse links response
//
// swagger:model LinksResponse
type LinksResponse struct {

	// account settings Url
	AccountSettingsURL string `json:"accountSettingsUrl,omitempty"`

	// account top up Url
	AccountTopUpURL string `json:"accountTopUpUrl,omitempty"`

	// accounts Url
	AccountsURL string `json:"accountsUrl,omitempty"`

	// admin ivr settings
	AdminIvrSettings string `json:"adminIvrSettings,omitempty"`

	// admin product settings Url
	AdminProductSettingsURL string `json:"adminProductSettingsUrl,omitempty"`

	// admin public phones Url
	AdminPublicPhonesURL string `json:"adminPublicPhonesUrl,omitempty"`

	// affiliate Url
	AffiliateURL string `json:"affiliateUrl,omitempty"`

	// apps Url
	AppsURL string `json:"appsUrl,omitempty"`

	// basic Url
	BasicURL string `json:"basicUrl,omitempty"`

	// book connection Url
	BookConnectionURL string `json:"bookConnectionUrl,omitempty"`

	// book faxline Url
	BookFaxlineURL string `json:"bookFaxlineUrl,omitempty"`

	// book phoneline Url
	BookPhonelineURL string `json:"bookPhonelineUrl,omitempty"`

	// book phonenumber Url
	BookPhonenumberURL string `json:"bookPhonenumberUrl,omitempty"`

	// brand Url
	BrandURL string `json:"brandUrl,omitempty"`

	// chatbot Url
	ChatbotURL string `json:"chatbotUrl,omitempty"`

	// conference Url
	ConferenceURL string `json:"conferenceUrl,omitempty"`

	// console web Url
	ConsoleWebURL string `json:"consoleWebUrl,omitempty"`

	// contacts Url
	ContactsURL string `json:"contactsUrl,omitempty"`

	// contract upselling Url
	ContractUpsellingURL string `json:"contractUpsellingUrl,omitempty"`

	// fax app Url
	FaxAppURL string `json:"faxAppUrl,omitempty"`

	// fax rescue Url
	FaxRescueURL string `json:"faxRescueUrl,omitempty"`

	// feature store Url
	FeatureStoreURL string `json:"featureStoreUrl,omitempty"`

	// feedback form Url
	FeedbackFormURL string `json:"feedbackFormUrl,omitempty"`

	// forwardings Url
	ForwardingsURL string `json:"forwardingsUrl,omitempty"`

	// helpdesk phone configuration Url
	HelpdeskPhoneConfigurationURL string `json:"helpdeskPhoneConfigurationUrl,omitempty"`

	// helpdesk Url
	HelpdeskURL string `json:"helpdeskUrl,omitempty"`

	// legacy token authenticate Url
	LegacyTokenAuthenticateURL string `json:"legacyTokenAuthenticateUrl,omitempty"`

	// login Url
	LoginURL string `json:"loginUrl,omitempty"`

	// logo target Url
	LogoTargetURL string `json:"logoTargetUrl,omitempty"`

	// logout page Url
	LogoutPageURL string `json:"logoutPageUrl,omitempty"`

	// logout Url
	LogoutURL string `json:"logoutUrl,omitempty"`

	// mnp faq Url
	MnpFaqURL string `json:"mnpFaqUrl,omitempty"`

	// notifications Url
	NotificationsURL string `json:"notificationsUrl,omitempty"`

	// personal settings Url
	PersonalSettingsURL string `json:"personalSettingsUrl,omitempty"`

	// phonestats Url
	PhonestatsURL string `json:"phonestatsUrl,omitempty"`

	// presence container Url
	PresenceContainerURL string `json:"presenceContainerUrl,omitempty"`

	// presence Url
	PresenceURL string `json:"presenceUrl,omitempty"`

	// rescue Url
	RescueURL string `json:"rescueUrl,omitempty"`

	// send fax Url
	SendFaxURL string `json:"sendFaxUrl,omitempty"`

	// simquadrat Url
	SimquadratURL string `json:"simquadratUrl,omitempty"`

	// start up page Url
	StartUpPageURL string `json:"startUpPageUrl,omitempty"`

	// support Url
	SupportURL string `json:"supportUrl,omitempty"`

	// team Url
	TeamURL string `json:"teamUrl,omitempty"`

	// test programme Url
	TestProgrammeURL string `json:"testProgrammeUrl,omitempty"`

	// token authenticate Url
	TokenAuthenticateURL string `json:"tokenAuthenticateUrl,omitempty"`

	// user settings Url
	UserSettingsURL string `json:"userSettingsUrl,omitempty"`

	// voicemail settings Url
	VoicemailSettingsURL string `json:"voicemailSettingsUrl,omitempty"`

	// wbci opt in Url
	WbciOptInURL string `json:"wbciOptInUrl,omitempty"`

	// whats new Url
	WhatsNewURL string `json:"whatsNewUrl,omitempty"`
}

// Validate validates this links response
func (m *LinksResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinksResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinksResponse) UnmarshalBinary(b []byte) error {
	var res LinksResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
