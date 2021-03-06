package client_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/client_secure/i_n_t_e_g_r_a_t_i_o_n_s"
	"restclient/client_secure/p_a_r_k_i_n_g"
	"restclient/client_secure/p_a_y_m_e_n_t_s"
	"restclient/client_secure/p_e_o_p_l_e_a_u_t_h"
	"restclient/client_secure/p_e_o_p_l_e_c_o_n_s_u_m_e_r"
	"restclient/client_secure/p_e_o_p_l_e_m_a_n_a_g_e_m_e_n_t"
	"restclient/client_secure/s_t_a_f_f_a_u_t_h"
	"restclient/client_secure/s_t_a_f_f_m_a_n_a_g_e_m_e_n_t"
	"restclient/client_secure/w_i_s_h_l_i_s_t_s"
)

// Default westfield apis HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new westfield secure-apis HTTP client.
func NewHTTPClient(formats strfmt.Registry) *WestfieldSecureApis {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("secure.westfield.io", "/v1", []string{"https"})
	return New(transport, formats)
}

// New creates a new westfield apis client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *WestfieldSecureApis {
	cli := new(WestfieldSecureApis)
	cli.Transport = transport

	cli.INTEGRATIONS = i_n_t_e_g_r_a_t_i_o_n_s.New(transport, formats)

	cli.PARKING = p_a_r_k_i_n_g.New(transport, formats)

	cli.PAYMENTS = p_a_y_m_e_n_t_s.New(transport, formats)

	cli.PEOPLEAUTH = p_e_o_p_l_e_a_u_t_h.New(transport, formats)

	cli.PEOPLECONSUMER = p_e_o_p_l_e_c_o_n_s_u_m_e_r.New(transport, formats)

	cli.PEOPLEMANAGEMENT = p_e_o_p_l_e_m_a_n_a_g_e_m_e_n_t.New(transport, formats)

	cli.STAFFAUTH = s_t_a_f_f_a_u_t_h.New(transport, formats)

	cli.STAFFMANAGEMENT = s_t_a_f_f_m_a_n_a_g_e_m_e_n_t.New(transport, formats)

	cli.WISHLISTS = w_i_s_h_l_i_s_t_s.New(transport, formats)

	return cli
}

// WestfieldSecureApis is a client for westfield secure-apis
type WestfieldSecureApis struct {
	INTEGRATIONS *i_n_t_e_g_r_a_t_i_o_n_s.Client

	PARKING *p_a_r_k_i_n_g.Client

	PAYMENTS *p_a_y_m_e_n_t_s.Client

	PEOPLEAUTH *p_e_o_p_l_e_a_u_t_h.Client

	PEOPLECONSUMER *p_e_o_p_l_e_c_o_n_s_u_m_e_r.Client

	PEOPLEMANAGEMENT *p_e_o_p_l_e_m_a_n_a_g_e_m_e_n_t.Client

	STAFFAUTH *s_t_a_f_f_a_u_t_h.Client

	STAFFMANAGEMENT *s_t_a_f_f_m_a_n_a_g_e_m_e_n_t.Client

	WISHLISTS *w_i_s_h_l_i_s_t_s.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *WestfieldSecureApis) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.INTEGRATIONS.SetTransport(transport)

	c.PARKING.SetTransport(transport)

	c.PAYMENTS.SetTransport(transport)

	c.PEOPLEAUTH.SetTransport(transport)

	c.PEOPLECONSUMER.SetTransport(transport)

	c.PEOPLEMANAGEMENT.SetTransport(transport)

	c.STAFFAUTH.SetTransport(transport)

	c.STAFFMANAGEMENT.SetTransport(transport)

	c.WISHLISTS.SetTransport(transport)

}
