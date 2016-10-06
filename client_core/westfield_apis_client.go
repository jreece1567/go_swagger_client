package client_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/client_core/c_a_t_e_g_o_r_i_e_s"
	"restclient/client_core/c_e_n_t_r_e_e_v_e_n_t_s"
	"restclient/client_core/c_e_n_t_r_e_i_n_f_o_r_m_a_t_i_o_n"
	"restclient/client_core/c_e_n_t_r_e_n_o_t_i_c_e_s"
	"restclient/client_core/c_e_n_t_r_e_p_a_r_k_i_n_g_r_a_t_e_s"
	"restclient/client_core/c_e_n_t_r_e_s_e_r_v_i_c_e_s"
	"restclient/client_core/c_e_n_t_r_e_t_r_a_d_i_n_g_h_o_u_r_s"
	"restclient/client_core/c_e_n_t_r_e_z_o_n_e_s"
	"restclient/client_core/c_u_r_a_t_i_o_n_s"
	"restclient/client_core/d_e_a_l_s"
	"restclient/client_core/k_i_o_s_k"
	"restclient/client_core/m_o_v_i_e_i_n_f_o_r_m_a_t_i_o_n"
	"restclient/client_core/p_r_o_d_u_c_t_r_e_c_o_m_m_e_n_d_a_t_i_o_n_s"
	"restclient/client_core/p_r_o_d_u_c_t_s"
	"restclient/client_core/r_e_t_a_i_l_e_r_s"
	"restclient/client_core/s_e_a_r_c_h"
	"restclient/client_core/s_t_o_r_e_i_n_f_o_r_m_a_t_i_o_n"
	"restclient/client_core/s_t_o_r_e_t_r_a_d_i_n_g_h_o_u_r_s"
)

// Default westfield apis HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new westfield core-apis HTTP client.
func NewHTTPClient(formats strfmt.Registry) *WestfieldCoreApis {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("api.westfield.io", "/v1", []string{"https"})
	return New(transport, formats)
}

// New creates a new westfield apis client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *WestfieldCoreApis {
	cli := new(WestfieldCoreApis)
	cli.Transport = transport

	cli.CATEGORIES = c_a_t_e_g_o_r_i_e_s.New(transport, formats)

	cli.CENTREEVENTS = c_e_n_t_r_e_e_v_e_n_t_s.New(transport, formats)

	cli.CENTREINFORMATION = c_e_n_t_r_e_i_n_f_o_r_m_a_t_i_o_n.New(transport, formats)

	cli.CENTRENOTICES = c_e_n_t_r_e_n_o_t_i_c_e_s.New(transport, formats)

	cli.CENTREPARKINGRATES = c_e_n_t_r_e_p_a_r_k_i_n_g_r_a_t_e_s.New(transport, formats)

	cli.CENTRESERVICES = c_e_n_t_r_e_s_e_r_v_i_c_e_s.New(transport, formats)

	cli.CENTRETRADINGHOURS = c_e_n_t_r_e_t_r_a_d_i_n_g_h_o_u_r_s.New(transport, formats)

	cli.CENTREZONES = c_e_n_t_r_e_z_o_n_e_s.New(transport, formats)

	cli.CURATIONS = c_u_r_a_t_i_o_n_s.New(transport, formats)

	cli.DEALS = d_e_a_l_s.New(transport, formats)

	cli.KIOSK = k_i_o_s_k.New(transport, formats)

	cli.MOVIEINFORMATION = m_o_v_i_e_i_n_f_o_r_m_a_t_i_o_n.New(transport, formats)

	cli.PRODUCTRECOMMENDATIONS = p_r_o_d_u_c_t_r_e_c_o_m_m_e_n_d_a_t_i_o_n_s.New(transport, formats)

	cli.PRODUCTS = p_r_o_d_u_c_t_s.New(transport, formats)

	cli.RETAILERS = r_e_t_a_i_l_e_r_s.New(transport, formats)

	cli.SEARCH = s_e_a_r_c_h.New(transport, formats)

	cli.STOREINFORMATION = s_t_o_r_e_i_n_f_o_r_m_a_t_i_o_n.New(transport, formats)

	cli.STORETRADINGHOURS = s_t_o_r_e_t_r_a_d_i_n_g_h_o_u_r_s.New(transport, formats)

	return cli
}

// WestfieldCoreApis is a client for westfield core-apis
type WestfieldCoreApis struct {
	CATEGORIES *c_a_t_e_g_o_r_i_e_s.Client

	CENTREEVENTS *c_e_n_t_r_e_e_v_e_n_t_s.Client

	CENTREINFORMATION *c_e_n_t_r_e_i_n_f_o_r_m_a_t_i_o_n.Client

	CENTRENOTICES *c_e_n_t_r_e_n_o_t_i_c_e_s.Client

	CENTREPARKINGRATES *c_e_n_t_r_e_p_a_r_k_i_n_g_r_a_t_e_s.Client

	CENTRESERVICES *c_e_n_t_r_e_s_e_r_v_i_c_e_s.Client

	CENTRETRADINGHOURS *c_e_n_t_r_e_t_r_a_d_i_n_g_h_o_u_r_s.Client

	CENTREZONES *c_e_n_t_r_e_z_o_n_e_s.Client

	CURATIONS *c_u_r_a_t_i_o_n_s.Client

	DEALS *d_e_a_l_s.Client

	KIOSK *k_i_o_s_k.Client

	MOVIEINFORMATION *m_o_v_i_e_i_n_f_o_r_m_a_t_i_o_n.Client

	PRODUCTRECOMMENDATIONS *p_r_o_d_u_c_t_r_e_c_o_m_m_e_n_d_a_t_i_o_n_s.Client

	PRODUCTS *p_r_o_d_u_c_t_s.Client

	RETAILERS *r_e_t_a_i_l_e_r_s.Client

	SEARCH *s_e_a_r_c_h.Client

	STOREINFORMATION *s_t_o_r_e_i_n_f_o_r_m_a_t_i_o_n.Client

	STORETRADINGHOURS *s_t_o_r_e_t_r_a_d_i_n_g_h_o_u_r_s.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *WestfieldCoreApis) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.CATEGORIES.SetTransport(transport)

	c.CENTREEVENTS.SetTransport(transport)

	c.CENTREINFORMATION.SetTransport(transport)

	c.CENTRENOTICES.SetTransport(transport)

	c.CENTREPARKINGRATES.SetTransport(transport)

	c.CENTRESERVICES.SetTransport(transport)

	c.CENTRETRADINGHOURS.SetTransport(transport)

	c.CENTREZONES.SetTransport(transport)

	c.CURATIONS.SetTransport(transport)

	c.DEALS.SetTransport(transport)

	c.KIOSK.SetTransport(transport)

	c.MOVIEINFORMATION.SetTransport(transport)

	c.PRODUCTRECOMMENDATIONS.SetTransport(transport)

	c.PRODUCTS.SetTransport(transport)

	c.RETAILERS.SetTransport(transport)

	c.SEARCH.SetTransport(transport)

	c.STOREINFORMATION.SetTransport(transport)

	c.STORETRADINGHOURS.SetTransport(transport)

}
