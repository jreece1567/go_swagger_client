package main

import (
	//"bytes"
	"encoding/json"
	//"errors"
	"fmt"
	"log"
	"os"
	"time"
	//"path/filepath"
	//"regexp"
	//"sort"
	//"strings"
	"github.com/go-openapi/runtime"
	"github.com/go-swagger/go-swagger/vendor/github.com/go-openapi/runtime/client"
	//
	"restclient/client_core"
	"restclient/models_core"
	/*
		"restclient/client_core/c_a_t_e_g_o_r_i_e_s"
	*/
	"restclient/client_core/c_e_n_t_r_e_e_v_e_n_t_s"
	"restclient/client_core/c_e_n_t_r_e_i_n_f_o_r_m_a_t_i_o_n"
	"restclient/client_core/c_e_n_t_r_e_n_o_t_i_c_e_s"
	/*
		"restclient/client_core/c_e_n_t_r_e_p_a_r_k_i_n_g_r_a_t_e_s"
	*/
	"restclient/client_core/c_e_n_t_r_e_s_e_r_v_i_c_e_s"
	/*
		"restclient/client_core/c_e_n_t_r_e_t_r_a_d_i_n_g_h_o_u_r_s"
		"restclient/client_core/c_e_n_t_r_e_z_o_n_e_s"
	*/
	"restclient/client_core/c_u_r_a_t_i_o_n_s"
	"restclient/client_core/d_e_a_l_s"
	/*
		"restclient/client_core/k_i_o_s_k"
		"restclient/client_core/m_o_v_i_e_i_n_f_o_r_m_a_t_i_o_n"
		"restclient/client_core/p_r_o_d_u_c_t_r_e_c_o_m_m_e_n_d_a_t_i_o_n_s"
		"restclient/client_core/p_r_o_d_u_c_t_s"
	*/
	"restclient/client_core/r_e_t_a_i_l_e_r_s"
	/*
		"restclient/client_core/s_e_a_r_c_h"
	*/
	"restclient/client_core/s_t_o_r_e_i_n_f_o_r_m_a_t_i_o_n"
	/*
		"restclient/client_core/s_t_o_r_e_t_r_a_d_i_n_g_h_o_u_r_s"

		"restclient/models_secure"
		"restclient/client_secure"
		"restclient/client_secure/i_n_t_e_g_r_a_t_i_o_n_s"
		"restclient/client_secure/p_a_r_k_i_n_g"
		"restclient/client_secure/p_a_y_m_e_n_t_s"
		"restclient/client_secure/p_e_o_p_l_e_a_u_t_h"
		"restclient/client_secure/p_e_o_p_l_e_c_o_n_s_u_m_e_r"
		"restclient/client_secure/p_e_o_p_l_e_m_a_n_a_g_e_m_e_n_t"
		"restclient/client_secure/s_t_a_f_f_a_u_t_h"
		"restclient/client_secure/s_t_a_f_f_m_a_n_a_g_e_m_e_n_t"
		"restclient/client_secure/w_i_s_h_l_i_s_t_s"
	*/)

func getCentreCurations(centreId string, service *client_core.WestfieldCoreApis, auth runtime.ClientAuthInfoWriter, c chan []*models_core.CurationListInstance) {
	defer timeTrack(time.Now(), "getCentreCurations()")

	// set up the call parameters for 'GET /centres/{centre_id}/services'
	GetCurationsParams := c_u_r_a_t_i_o_n_s.NewGetCurationsParams()
	GetCurationsParams.WithCentreID(&centreId)
	var page, per_page int64
	page = int64(1)
	per_page = int64(100)
	GetCurationsParams.WithPage(&page)
	GetCurationsParams.WithPerPage(&per_page)
	GetCurationsParams.WithFeatured(nil)
	GetCurationsParams.WithStatuses([]string{"live"})
	//GetCurationsParams.WithFields([]string{"long_title"})
	//log.Println("Parameters:")
	//log.Printf("%+v\n", *GetCurationsParams)

	// make the API call to 'GET /curations'
	log.Println("Calling API - curations:centre:index ...")
	GetCurationsOK, err := service.CURATIONS.GetCurations(GetCurationsParams, auth)
	if err == nil {

		pages := GetCurationsOK.Payload.Data

		for page < GetCurationsOK.Payload.Meta.PageCount {
			page = page + 1
			GetCurationsParams.WithPage(&page)
			log.Println("Calling API - curations:centre:index ...")
			GetCurationsOK, _ := service.CURATIONS.GetCurations(GetCurationsParams, auth)
			pages = append(pages, GetCurationsOK.Payload.Data...)
		}

		//return GetCurationsOK.Payload.Data
		c <- pages

	} else {

		// print the error
		log.Println(err)
		b2, _ := json.Marshal(GetCurationsOK)
		fmt.Println(string(b2))
		log.Println("")

		c <- nil
		//return nil
	}
}

func getCentreDeals(centreId string, service *client_core.WestfieldCoreApis, auth runtime.ClientAuthInfoWriter, c chan []*models_core.DealInstance) {
	defer timeTrack(time.Now(), "getCentreDeals()")

	// set up the call parameters for 'GET /deals'
	GetDealsParams := d_e_a_l_s.NewGetDealsParams()
	GetDealsParams.WithCentreID(&centreId)
	var page, per_page int64
	page = int64(1)
	per_page = int64(100)
	GetDealsParams.WithPage(&page)
	GetDealsParams.WithPerPage(&per_page)
	GetDealsParams.WithAll(nil)
	GetDealsParams.WithFeatured(nil)
	GetDealsParams.WithStatuses([]string{"live"})
	//GetDealsParams.WithFields([]string{"long_title"})
	//log.Println("Parameters:")
	//log.Printf("%+v\n", *GetDealsParams)

	// make the API call to 'GET /deals'
	log.Println("Calling API - deals:centre:index ...")
	GetDealsOK, err := service.DEALS.GetDeals(GetDealsParams, auth)
	if err == nil {

		pages := GetDealsOK.Payload.Data

		for page < GetDealsOK.Payload.Meta.PageCount {
			page = page + 1
			GetDealsParams.WithPage(&page)
			log.Println("Calling API - deals:centre:index ...")
			GetDealsOK, _ := service.DEALS.GetDeals(GetDealsParams, auth)
			pages = append(pages, GetDealsOK.Payload.Data...)
		}

		//return GetDealsOK.Payload.Data
		c <- pages

	} else {

		// print the error
		log.Println(err)
		b2, _ := json.Marshal(GetDealsOK)
		fmt.Println(string(b2))
		log.Println("")

		c <- nil
		//return nil
	}
}

func getCentreNotices(centreId string, service *client_core.WestfieldCoreApis, auth runtime.ClientAuthInfoWriter, c chan []*models_core.NoticeListInstance) {
	defer timeTrack(time.Now(), "getCentreNotices()")

	// set up the call parameters for 'GET /notices'
	GetNoticesParams := c_e_n_t_r_e_n_o_t_i_c_e_s.NewGetNoticesParams()
	GetNoticesParams.WithCentreID(&centreId)
	var page, per_page int64
	page = int64(1)
	per_page = int64(100)
	GetNoticesParams.WithPage(&page)
	GetNoticesParams.WithPerPage(&per_page)
	GetNoticesParams.WithStatuses([]string{"live"})
	//GetNoticesParams.WithFields([]string{"long_title"})
	//log.Println("Parameters:")
	//log.Printf("%+v\n", *GetNoticesParams)

	// make the API call to 'GET /notices'
	log.Println("Calling API - notices:centre:index ...")
	GetNoticesOK, err := service.CENTRENOTICES.GetNotices(GetNoticesParams, auth)
	if err == nil {

		pages := GetNoticesOK.Payload.Data

		for page < GetNoticesOK.Payload.Meta.PageCount {
			page = page + 1
			GetNoticesParams.WithPage(&page)
			log.Println("Calling API - notices:centre:index ...")
			GetNoticesOK, _ := service.CENTRENOTICES.GetNotices(GetNoticesParams, auth)
			pages = append(pages, GetNoticesOK.Payload.Data...)
		}

		//return GetNoticesOK.Payload.Data
		c <- pages

	} else {

		// print the error
		log.Println(err)
		b2, _ := json.Marshal(GetNoticesOK)
		fmt.Println(string(b2))
		log.Println("")

		c <- nil
		//return nil
	}
}

func getCentreEvents(centreId string, service *client_core.WestfieldCoreApis, auth runtime.ClientAuthInfoWriter, c chan []*models_core.EventInstance) {
	defer timeTrack(time.Now(), "getCentreEvents()")

	// set up the call parameters for 'GET /events'
	GetEventsParams := c_e_n_t_r_e_e_v_e_n_t_s.NewGetEventsParams()
	GetEventsParams.WithCentreID(&centreId)
	var page, per_page int64
	page = int64(1)
	per_page = int64(100)
	GetEventsParams.WithPage(&page)
	GetEventsParams.WithPerPage(&per_page)
	GetEventsParams.WithFeatured(nil)
	GetEventsParams.WithPublished(nil)
	GetEventsParams.WithStatuses([]string{"live"})
	//GetEventsParams.WithFields([]string{"long_title"})
	//log.Println("Parameters:")
	//log.Printf("%+v\n", *GetEventsParams)

	// make the API call to 'GET /events'
	log.Println("Calling API - events:centre:index ...")
	GetEventsOK, err := service.CENTREEVENTS.GetEvents(GetEventsParams, auth)
	if err == nil {

		pages := GetEventsOK.Payload.Data

		for page < GetEventsOK.Payload.Meta.PageCount {
			page = page + 1
			GetEventsParams.WithPage(&page)
			log.Println("Calling API - events:centre:index ...")
			GetEventsOK, _ := service.CENTREEVENTS.GetEvents(GetEventsParams, auth)
			pages = append(pages, GetEventsOK.Payload.Data...)
		}

		//return GetEventsOK.Payload.Data
		c <- pages

	} else {

		// print the error
		log.Println(err)
		b2, _ := json.Marshal(GetEventsOK)
		fmt.Println(string(b2))
		log.Println("")

		c <- nil
		//return nil
	}
}

func getCentreRetailers(centreId string, service *client_core.WestfieldCoreApis, auth runtime.ClientAuthInfoWriter, c chan []*models_core.RetailerInstance) {
	defer timeTrack(time.Now(), "getCentreRetailers()")

	// set up the call parameters for 'GET /retailers'
	GetRetailersParams := r_e_t_a_i_l_e_r_s.NewGetRetailersParams()
	GetRetailersParams.WithCentreID(&centreId)
	GetRetailersParams.WithCountry(nil)
	GetRetailersParams.WithStatuses([]string{"live"})
	var page, per_page int64
	page = int64(1)
	per_page = int64(100)
	GetRetailersParams.WithPage(&page)
	GetRetailersParams.WithPerPage(&per_page)
	//GetRetailersParams.WithFields([]string{"long_title"})
	//log.Println("Parameters:")
	//log.Printf("%+v\n", *GetRetailersParams)

	// make the API call to 'GET /retailers'
	log.Println("Calling API - retailers:centre:index ...")
	GetRetailersOK, err := service.RETAILERS.GetRetailers(GetRetailersParams, auth)
	if err == nil {

		pages := GetRetailersOK.Payload.Data

		for page < GetRetailersOK.Payload.Meta.PageCount {
			page = page + 1
			GetRetailersParams.WithPage(&page)
			log.Println("Calling API - retailers:centre:index ...")
			GetRetailersOK, _ := service.RETAILERS.GetRetailers(GetRetailersParams, auth)
			pages = append(pages, GetRetailersOK.Payload.Data...)
		}

		//return GetRetailersOK.Payload.Data
		c <- pages

	} else {

		// print the error
		log.Println(err)
		b2, _ := json.Marshal(GetRetailersOK)
		fmt.Println(string(b2))
		log.Println("")

		c <- nil
		//return nil
	}
}

func getCentreStores(centreId string, service *client_core.WestfieldCoreApis, auth runtime.ClientAuthInfoWriter, c chan []*models_core.StoreInstance) {
	defer timeTrack(time.Now(), "getCentreStores()")

	// set up the call parameters for 'GET /stores'
	GetStoresParams := s_t_o_r_e_i_n_f_o_r_m_a_t_i_o_n.NewGetStoresParams()
	GetStoresParams.WithCentreID(&centreId)
	GetStoresParams.WithCountry(nil)
	GetStoresParams.WithStatuses([]string{"live"})
	GetStoresParams.WithStoreIds([]int64{14857})
	var page, per_page int64
	page = int64(1)
	per_page = int64(100)
	GetStoresParams.WithPage(&page)
	GetStoresParams.WithPerPage(&per_page)
	//GetStoresParams.WithFields([]string{"long_title"})
	//log.Println("Parameters:")
	//log.Printf("%+v\n", *GetStoresParams)

	// make the API call to 'GET /stores'
	log.Println("Calling API - stores:centre:index ...")
	GetStoresOK, err := service.STOREINFORMATION.GetStores(GetStoresParams, auth)
	if err == nil {

		pages := GetStoresOK.Payload.Data

		for page < GetStoresOK.Payload.Meta.PageCount {
			page = page + 1
			GetStoresParams.WithPage(&page)
			log.Println("Calling API - stores:centre:index ...")
			GetStoresOK, _ := service.STOREINFORMATION.GetStores(GetStoresParams, auth)
			pages = append(pages, GetStoresOK.Payload.Data...)
		}

		//return GetRetailersOK.Payload.Data
		c <- pages

	} else {

		// print the error
		log.Println("getCentreStores")
		log.Println(err)
		b2, _ := json.Marshal(GetStoresOK)
		fmt.Println(string(b2))
		log.Println("")

		c <- nil
		//return nil
	}
}

func getStoreDetail(store_id int, service *client_core.WestfieldCoreApis, auth runtime.ClientAuthInfoWriter) *models_core.StoreInstance {
	defer timeTrack(time.Now(), "getStoreDetail()")

	// set up the call parameters for 'GET /stores/{store_id}'
	GetStoresStoreIDParams := s_t_o_r_e_i_n_f_o_r_m_a_t_i_o_n.NewGetStoresStoreIDParams()
	GetStoresStoreIDParams.WithStoreID(int64(store_id))
	//log.Println("Parameters:")
	//log.Printf("%+v\n", *GetStoresStoreIDParams)

	// make the API call to 'GET /stores/{store_id}'
	log.Println("Calling API - stores:store:show ...")
	GetStoresStoreIDOK, err := service.STOREINFORMATION.GetStoresStoreID(GetStoresStoreIDParams, auth)
	if err == nil {

		return GetStoresStoreIDOK.Payload.Data

	} else {

		// print the error
		log.Println("getStoreDetail")
		log.Println(err)
		b2, _ := json.Marshal(GetStoresStoreIDOK)
		fmt.Println(string(b2))
		log.Println("")

		return nil
	}
}

func getCentreServices(centreId string, serviceClass string, service *client_core.WestfieldCoreApis, auth runtime.ClientAuthInfoWriter, c chan []*models_core.Service) {
	defer timeTrack(time.Now(), "getCentreServices()")

	// set up the call parameters for 'GET /centres/{centre_id}/services'
	GetCentresCentreIDServicesParams := c_e_n_t_r_e_s_e_r_v_i_c_e_s.NewGetCentresCentreIDServicesParams()
	GetCentresCentreIDServicesParams.WithCentreID(centreId)
	GetCentresCentreIDServicesParams.WithServiceClass(&serviceClass)
	var page, per_page int64
	page = int64(1)
	per_page = int64(100)
	GetCentresCentreIDServicesParams.WithPage(&page)
	GetCentresCentreIDServicesParams.WithPerPage(&per_page)
	//GetCentresCentreIDServicesParams.WithFields([]string{"long_title"})
	//log.Println("Parameters:")
	//log.Printf("%+v\n", *GetCentresCentreIDServicesParams)

	// make the API call to 'GET /centres/{centre_id}/services'
	log.Println("Calling API - services:centre:index ...")
	GetCentresCentreIDServicesOK, err := service.CENTRESERVICES.GetCentresCentreIDServices(GetCentresCentreIDServicesParams, auth)
	if err == nil {

		pages := GetCentresCentreIDServicesOK.Payload.Data

		for page < GetCentresCentreIDServicesOK.Payload.Meta.PageCount {
			page = page + 1
			GetCentresCentreIDServicesParams.WithPage(&page)
			log.Println("Calling API - services:centre:index ...")
			GetCentresCentreIDServicesOK, _ := service.CENTRESERVICES.GetCentresCentreIDServices(GetCentresCentreIDServicesParams, auth)
			pages = append(pages, GetCentresCentreIDServicesOK.Payload.Data...)
		}

		//return GetCentresCentreIDServicesOK.Payload.Data
		c <- pages

	} else {

		// print the error
		log.Println(err)
		b2, _ := json.Marshal(GetCentresCentreIDServicesOK)
		fmt.Println(string(b2))
		log.Println("")

		c <- nil
		//return nil
	}
}

func getCentreDetail(centreId string, service *client_core.WestfieldCoreApis, auth runtime.ClientAuthInfoWriter) *models_core.CentreInstance {
	defer timeTrack(time.Now(), "getCentreDetail()")

	// set up the call parameters for 'GET /centres/{centre_id}'
	GetCentresCentreIDParams := c_e_n_t_r_e_i_n_f_o_r_m_a_t_i_o_n.NewGetCentresCentreIDParams()
	GetCentresCentreIDParams.WithCentreID(centreId)
	//GetCentresCentreIDParams.WithFields([]string{"short_name", "street_address", "phone_number", "enabled_at", "features"})
	//log.Println("Parameters:")
	//log.Printf("%+v\n", *GetCentresCentreIDParams)

	// make the API call to 'GET /centres/{centre_id}'
	log.Println("Calling API - centre:show ...")
	GetCentresCentreIDOK, err := service.CENTREINFORMATION.GetCentresCentreID(GetCentresCentreIDParams, auth)
	if err == nil {

		return GetCentresCentreIDOK.Payload.Data

	} else {

		// print the error
		log.Println(err)
		b1, _ := json.Marshal(GetCentresCentreIDOK)
		fmt.Println(string(b1))
		log.Println("")

		return nil
	}

}

func usage() {
	log.Println("Please provide an api_key and a centre_id on the command-line.")
	log.Println("Usage: ./restclient <api_key> <centre_id>")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer timeTrack(time.Now(), "main()")

	log.Println("Begin main()")
	log.Println("")

	// get the Apigee key and centre_id from the command-line
	args := os.Args
	if len(args) < 3 {
		usage()
		return
	}
	apiKey := string(args[1])
	centreId := string(args[2])

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// create the API interface-facade
	WestfieldCoreApis := client_core.NewHTTPClient(nil)

	// set up the api_key info for use in subsequent API calls
	AuthInfo := client.APIKeyAuth("api_key", "query", apiKey)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// get the centre details synchronously
	centre := getCentreDetail(centreId, WestfieldCoreApis, AuthInfo)
	if centre != nil {
		// no errors, print out a few fields from the response
		fmt.Printf("\t*** Centre Details:\n")
		fmt.Printf("\t%s\n", *centre.ShortName)
		fmt.Printf("\t%s\n", *centre.StreetAddress)
		fmt.Printf("\t%s, %s %s, %s\n", *centre.Suburb, *centre.State, *centre.Postcode, *centre.Country)
		fmt.Printf("\t%s\n", *centre.PhoneNumber)
		fmt.Printf("\tEnabled: %s\n", centre.EnabledAt)
		fmt.Printf("\tFeatures: (%d)\n", len(centre.Features))
		for i := 0; i < len(centre.Features); i++ {
			fmt.Printf("\t\t[%s]\n", centre.Features[i])
		}
	} else {
		log.Println("centre_id [" + centreId + "] and/or api_key [" + apiKey + "] is not valid, ending.")
		return
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// get the centre-services asynchronously
	serviceChannel := make(chan []*models_core.Service)
	go getCentreServices(centreId, "digital", WestfieldCoreApis, AuthInfo, serviceChannel)
	go getCentreServices(centreId, "physical", WestfieldCoreApis, AuthInfo, serviceChannel)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// get the centre-notices asynchronously
	noticesChannel := make(chan []*models_core.NoticeListInstance)
	go getCentreNotices(centreId, WestfieldCoreApis, AuthInfo, noticesChannel)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// get the centre-events asynchronously
	eventsChannel := make(chan []*models_core.EventInstance)
	go getCentreEvents(centreId, WestfieldCoreApis, AuthInfo, eventsChannel)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// get the centre-deals asynchronously
	dealsChannel := make(chan []*models_core.DealInstance)
	go getCentreDeals(centreId, WestfieldCoreApis, AuthInfo, dealsChannel)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// get the centre-curations asynchronously
	curationsChannel := make(chan []*models_core.CurationListInstance)
	go getCentreCurations(centreId, WestfieldCoreApis, AuthInfo, curationsChannel)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// get the centre-retailers asynchronously
	retailersChannel := make(chan []*models_core.RetailerInstance)
	go getCentreRetailers(centreId, WestfieldCoreApis, AuthInfo, retailersChannel)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// get the centre-stores asynchronously
	//storesChannel := make(chan []*models_core.StoreInstance)
	//go getCentreStores(centreId, WestfieldCoreApis, AuthInfo, storesChannel)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// now that all of the async API calls have been set up (above), fetch and display the results (below)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// lists of services of both service-classes
	digitalServices, physicalServices := <-serviceChannel, <-serviceChannel

	// build the list of *all* services
	allServices := append(digitalServices, physicalServices...)

	fmt.Printf("\t*** Centre Services: (%d)\n", len(allServices))
	for i := 0; i < len(allServices); i++ {
		fmt.Printf("\t%5d - %s %s\n", *allServices[i].ServiceID, *allServices[i].ServiceClass, *allServices[i].LongTitle)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// list of all notices
	centreNotices := <-noticesChannel

	fmt.Printf("\t*** Centre Notices: (%d)\n", len(centreNotices))
	for i := 0; i < len(centreNotices); i++ {
		fmt.Printf("\t%5d - Starts: %s Ends: %s - %s - %s\n", *centreNotices[i].NoticeID, *centreNotices[i].PublishedAt, *centreNotices[i].ExpiresAt, *centreNotices[i].Type, *centreNotices[i].Name)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// list of all events
	centreEvents := <-eventsChannel

	fmt.Printf("\t*** Centre Events: (%d)\n", len(centreEvents))
	for i := 0; i < len(centreEvents); i++ {
		fmt.Printf("\t%5d - Starts: %s Ends: %s - %s\n", *centreEvents[i].EventID, *centreEvents[i].EnabledAt, *centreEvents[i].DisabledAt, *centreEvents[i].Name)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// list of all deals
	centreDeals := <-dealsChannel

	fmt.Printf("\t*** Centre Deals: (%d)\n", len(centreDeals))
	for i := 0; i < len(centreDeals); i++ {
		fmt.Printf("\t%5d - Starts: %s Ends: %s - %s\n", *centreDeals[i].DealID, *centreDeals[i].StartsAt, *centreDeals[i].EndsAt, *centreDeals[i].Title)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// list of all curations
	centreCurations := <-curationsChannel

	fmt.Printf("\t*** Centre Curations: (%d)\n", len(centreCurations))
	for i := 0; i < len(centreCurations); i++ {
		fmt.Printf("\t%5d - Starts: %s Ends: %s - %s\n", *centreCurations[i].CurationID, *centreCurations[i].StartsAt, *centreCurations[i].EndsAt, *centreCurations[i].Name)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// list of all retailers
	centreRetailers := <-retailersChannel

	fmt.Printf("\t*** Centre Retailers: (%d)\n", len(centreRetailers))
	for i := 0; i < len(centreRetailers); i++ {
		fmt.Printf("\t%5d - %s\n", *centreRetailers[i].RetailerID, *centreRetailers[i].Name)
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	/*
		// list of all stores
		centreStores := <-storesChannel

		fmt.Printf("\t*** Centre Stores: (%d)\n", len(centreStores))
		for i := 0; i < len(centreStores); i++ {
			fmt.Printf("\t%d - %s\n", *centreStores[i].StoreID, *centreStores[i].Name)
		}
	*/
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	/*
		// get the store details synchronously
		store := getStoreDetail(28970, WestfieldCoreApis, AuthInfo)
		if store != nil {
			// no errors, print out a few fields from the response
			fmt.Printf("\t*** Store Details:\n")
			fmt.Printf("\t%d - %s\n", *store.StoreID, *store.Name)
		}
	*/
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	log.Println("")
	log.Println("End main()")
}
