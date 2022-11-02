// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BindParams bind params
// Example: {"address":"127.0.0.1","name":"http","port":80}
//
// swagger:model bind_params
type BindParams struct {

	// accept netscaler cip
	AcceptNetscalerCip int64 `json:"accept_netscaler_cip,omitempty"`

	// accept proxy
	AcceptProxy bool `json:"accept_proxy,omitempty"`

	// allow 0rtt
	Allow0rtt bool `json:"allow_0rtt,omitempty"`

	// alpn
	// Pattern: ^[^\s]+$
	Alpn string `json:"alpn,omitempty"`

	// backlog
	Backlog string `json:"backlog,omitempty"`

	// ca ignore err
	CaIgnoreErr string `json:"ca_ignore_err,omitempty"`

	// ca sign file
	CaSignFile string `json:"ca_sign_file,omitempty"`

	// ca sign pass
	CaSignPass string `json:"ca_sign_pass,omitempty"`

	// ca verify file
	CaVerifyFile string `json:"ca_verify_file,omitempty"`

	// ciphers
	Ciphers string `json:"ciphers,omitempty"`

	// ciphersuites
	Ciphersuites string `json:"ciphersuites,omitempty"`

	// crl file
	CrlFile string `json:"crl_file,omitempty"`

	// crt ignore err
	CrtIgnoreErr string `json:"crt_ignore_err,omitempty"`

	// crt list
	CrtList string `json:"crt_list,omitempty"`

	// curves
	Curves string `json:"curves,omitempty"`

	// defer accept
	DeferAccept bool `json:"defer_accept,omitempty"`

	// ecdhe
	Ecdhe string `json:"ecdhe,omitempty"`

	// expose fd listeners
	ExposeFdListeners bool `json:"expose_fd_listeners,omitempty"`

	// force sslv3
	ForceSslv3 bool `json:"force_sslv3,omitempty"`

	// force tlsv10
	ForceTlsv10 bool `json:"force_tlsv10,omitempty"`

	// force tlsv11
	ForceTlsv11 bool `json:"force_tlsv11,omitempty"`

	// force tlsv12
	ForceTlsv12 bool `json:"force_tlsv12,omitempty"`

	// force tlsv13
	ForceTlsv13 bool `json:"force_tlsv13,omitempty"`

	// generate certificates
	GenerateCertificates bool `json:"generate_certificates,omitempty"`

	// gid
	Gid int64 `json:"gid,omitempty"`

	// group
	Group string `json:"group,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// interface
	Interface string `json:"interface,omitempty"`

	// level
	// Enum: [user operator admin]
	Level string `json:"level,omitempty"`

	// maxconn
	Maxconn int64 `json:"maxconn,omitempty"`

	// mode
	Mode string `json:"mode,omitempty"`

	// mss
	Mss string `json:"mss,omitempty"`

	// name
	// Pattern: ^[^\s]+$
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// nice
	Nice int64 `json:"nice,omitempty"`

	// no ca names
	NoCaNames bool `json:"no_ca_names,omitempty"`

	// no sslv3
	NoSslv3 bool `json:"no_sslv3,omitempty"`

	// no tls tickets
	NoTLSTickets bool `json:"no_tls_tickets,omitempty"`

	// no tlsv10
	NoTlsv10 bool `json:"no_tlsv10,omitempty"`

	// no tlsv11
	NoTlsv11 bool `json:"no_tlsv11,omitempty"`

	// no tlsv12
	NoTlsv12 bool `json:"no_tlsv12,omitempty"`

	// no tlsv13
	NoTlsv13 bool `json:"no_tlsv13,omitempty"`

	// npn
	Npn string `json:"npn,omitempty"`

	// prefer client ciphers
	PreferClientCiphers bool `json:"prefer_client_ciphers,omitempty"`

	// process
	// Pattern: ^[^\s]+$
	Process string `json:"process,omitempty"`

	// proto
	Proto string `json:"proto,omitempty"`

	// severity output
	// Enum: [none number string]
	SeverityOutput string `json:"severity_output,omitempty"`

	// ssl
	Ssl bool `json:"ssl,omitempty"`

	// ssl cafile
	// Pattern: ^[^\s]+$
	SslCafile string `json:"ssl_cafile,omitempty"`

	// ssl certificate
	// Pattern: ^[^\s]+$
	SslCertificate string `json:"ssl_certificate,omitempty"`

	// ssl max ver
	// Enum: [SSLv3 TLSv1.0 TLSv1.1 TLSv1.2 TLSv1.3]
	SslMaxVer string `json:"ssl_max_ver,omitempty"`

	// ssl min ver
	// Enum: [SSLv3 TLSv1.0 TLSv1.1 TLSv1.2 TLSv1.3]
	SslMinVer string `json:"ssl_min_ver,omitempty"`

	// strict sni
	StrictSni bool `json:"strict_sni,omitempty"`

	// tcp user timeout
	TCPUserTimeout *int64 `json:"tcp_user_timeout,omitempty"`

	// tfo
	Tfo bool `json:"tfo,omitempty"`

	// tls ticket keys
	TLSTicketKeys string `json:"tls_ticket_keys,omitempty"`

	// transparent
	Transparent bool `json:"transparent,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`

	// user
	User string `json:"user,omitempty"`

	// v4v6
	V4v6 bool `json:"v4v6,omitempty"`

	// v6only
	V6only bool `json:"v6only,omitempty"`

	// verify
	// Enum: [none optional required]
	Verify string `json:"verify,omitempty"`
}

// Validate validates this bind params
func (m *BindParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlpn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverityOutput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslCafile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslMaxVer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslMinVer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerify(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BindParams) validateAlpn(formats strfmt.Registry) error {
	if swag.IsZero(m.Alpn) { // not required
		return nil
	}

	if err := validate.Pattern("alpn", "body", m.Alpn, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var bindParamsTypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","operator","admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bindParamsTypeLevelPropEnum = append(bindParamsTypeLevelPropEnum, v)
	}
}

const (

	// BindParamsLevelUser captures enum value "user"
	BindParamsLevelUser string = "user"

	// BindParamsLevelOperator captures enum value "operator"
	BindParamsLevelOperator string = "operator"

	// BindParamsLevelAdmin captures enum value "admin"
	BindParamsLevelAdmin string = "admin"
)

// prop value enum
func (m *BindParams) validateLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bindParamsTypeLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BindParams) validateLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.Level) { // not required
		return nil
	}

	// value enum
	if err := m.validateLevelEnum("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

func (m *BindParams) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", m.Name, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *BindParams) validateProcess(formats strfmt.Registry) error {
	if swag.IsZero(m.Process) { // not required
		return nil
	}

	if err := validate.Pattern("process", "body", m.Process, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var bindParamsTypeSeverityOutputPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","number","string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bindParamsTypeSeverityOutputPropEnum = append(bindParamsTypeSeverityOutputPropEnum, v)
	}
}

const (

	// BindParamsSeverityOutputNone captures enum value "none"
	BindParamsSeverityOutputNone string = "none"

	// BindParamsSeverityOutputNumber captures enum value "number"
	BindParamsSeverityOutputNumber string = "number"

	// BindParamsSeverityOutputString captures enum value "string"
	BindParamsSeverityOutputString string = "string"
)

// prop value enum
func (m *BindParams) validateSeverityOutputEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bindParamsTypeSeverityOutputPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BindParams) validateSeverityOutput(formats strfmt.Registry) error {
	if swag.IsZero(m.SeverityOutput) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityOutputEnum("severity_output", "body", m.SeverityOutput); err != nil {
		return err
	}

	return nil
}

func (m *BindParams) validateSslCafile(formats strfmt.Registry) error {
	if swag.IsZero(m.SslCafile) { // not required
		return nil
	}

	if err := validate.Pattern("ssl_cafile", "body", m.SslCafile, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *BindParams) validateSslCertificate(formats strfmt.Registry) error {
	if swag.IsZero(m.SslCertificate) { // not required
		return nil
	}

	if err := validate.Pattern("ssl_certificate", "body", m.SslCertificate, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var bindParamsTypeSslMaxVerPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SSLv3","TLSv1.0","TLSv1.1","TLSv1.2","TLSv1.3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bindParamsTypeSslMaxVerPropEnum = append(bindParamsTypeSslMaxVerPropEnum, v)
	}
}

const (

	// BindParamsSslMaxVerSSLv3 captures enum value "SSLv3"
	BindParamsSslMaxVerSSLv3 string = "SSLv3"

	// BindParamsSslMaxVerTLSv1Dot0 captures enum value "TLSv1.0"
	BindParamsSslMaxVerTLSv1Dot0 string = "TLSv1.0"

	// BindParamsSslMaxVerTLSv1Dot1 captures enum value "TLSv1.1"
	BindParamsSslMaxVerTLSv1Dot1 string = "TLSv1.1"

	// BindParamsSslMaxVerTLSv1Dot2 captures enum value "TLSv1.2"
	BindParamsSslMaxVerTLSv1Dot2 string = "TLSv1.2"

	// BindParamsSslMaxVerTLSv1Dot3 captures enum value "TLSv1.3"
	BindParamsSslMaxVerTLSv1Dot3 string = "TLSv1.3"
)

// prop value enum
func (m *BindParams) validateSslMaxVerEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bindParamsTypeSslMaxVerPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BindParams) validateSslMaxVer(formats strfmt.Registry) error {
	if swag.IsZero(m.SslMaxVer) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslMaxVerEnum("ssl_max_ver", "body", m.SslMaxVer); err != nil {
		return err
	}

	return nil
}

var bindParamsTypeSslMinVerPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SSLv3","TLSv1.0","TLSv1.1","TLSv1.2","TLSv1.3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bindParamsTypeSslMinVerPropEnum = append(bindParamsTypeSslMinVerPropEnum, v)
	}
}

const (

	// BindParamsSslMinVerSSLv3 captures enum value "SSLv3"
	BindParamsSslMinVerSSLv3 string = "SSLv3"

	// BindParamsSslMinVerTLSv1Dot0 captures enum value "TLSv1.0"
	BindParamsSslMinVerTLSv1Dot0 string = "TLSv1.0"

	// BindParamsSslMinVerTLSv1Dot1 captures enum value "TLSv1.1"
	BindParamsSslMinVerTLSv1Dot1 string = "TLSv1.1"

	// BindParamsSslMinVerTLSv1Dot2 captures enum value "TLSv1.2"
	BindParamsSslMinVerTLSv1Dot2 string = "TLSv1.2"

	// BindParamsSslMinVerTLSv1Dot3 captures enum value "TLSv1.3"
	BindParamsSslMinVerTLSv1Dot3 string = "TLSv1.3"
)

// prop value enum
func (m *BindParams) validateSslMinVerEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bindParamsTypeSslMinVerPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BindParams) validateSslMinVer(formats strfmt.Registry) error {
	if swag.IsZero(m.SslMinVer) { // not required
		return nil
	}

	// value enum
	if err := m.validateSslMinVerEnum("ssl_min_ver", "body", m.SslMinVer); err != nil {
		return err
	}

	return nil
}

var bindParamsTypeVerifyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","optional","required"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bindParamsTypeVerifyPropEnum = append(bindParamsTypeVerifyPropEnum, v)
	}
}

const (

	// BindParamsVerifyNone captures enum value "none"
	BindParamsVerifyNone string = "none"

	// BindParamsVerifyOptional captures enum value "optional"
	BindParamsVerifyOptional string = "optional"

	// BindParamsVerifyRequired captures enum value "required"
	BindParamsVerifyRequired string = "required"
)

// prop value enum
func (m *BindParams) validateVerifyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, bindParamsTypeVerifyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BindParams) validateVerify(formats strfmt.Registry) error {
	if swag.IsZero(m.Verify) { // not required
		return nil
	}

	// value enum
	if err := m.validateVerifyEnum("verify", "body", m.Verify); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this bind params based on context it is used
func (m *BindParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BindParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BindParams) UnmarshalBinary(b []byte) error {
	var res BindParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
