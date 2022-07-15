/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionossdk

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// LocationApiService LocationApi service
type LocationApiService service

type ApiLocationsFindByRegionRequest struct {
	ctx             _context.Context
	ApiService      *LocationApiService
	regionId        string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiLocationsFindByRegionRequest) Pretty(pretty bool) ApiLocationsFindByRegionRequest {
	r.pretty = &pretty
	return r
}
func (r ApiLocationsFindByRegionRequest) Depth(depth int32) ApiLocationsFindByRegionRequest {
	r.depth = &depth
	return r
}
func (r ApiLocationsFindByRegionRequest) XContractNumber(xContractNumber int32) ApiLocationsFindByRegionRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiLocationsFindByRegionRequest) Execute() (Locations, *APIResponse, error) {
	return r.ApiService.LocationsFindByRegionExecute(r)
}

/*
 * LocationsFindByRegion List Locations within a region
 * Retrieve a list of Locations within a world's region
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param regionId
 * @return ApiLocationsFindByRegionRequest
 */
func (a *LocationApiService) LocationsFindByRegion(ctx _context.Context, regionId string) ApiLocationsFindByRegionRequest {
	return ApiLocationsFindByRegionRequest{
		ApiService: a,
		ctx:        ctx,
		regionId:   regionId,
	}
}

/*
 * Execute executes the request
 * @return Locations
 */
func (a *LocationApiService) LocationsFindByRegionExecute(r ApiLocationsFindByRegionRequest) (Locations, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Locations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationApiService.LocationsFindByRegion")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/locations/{regionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"regionId"+"}", _neturl.PathEscape(parameterToString(r.regionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:   localVarHTTPResponse,
		Method:     localVarHTTPMethod,
		RequestURL: localVarPath,
		Operation:  "LocationsFindByRegion",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiLocationsFindByRegionAndIdRequest struct {
	ctx             _context.Context
	ApiService      *LocationApiService
	regionId        string
	locationId      string
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiLocationsFindByRegionAndIdRequest) Pretty(pretty bool) ApiLocationsFindByRegionAndIdRequest {
	r.pretty = &pretty
	return r
}
func (r ApiLocationsFindByRegionAndIdRequest) Depth(depth int32) ApiLocationsFindByRegionAndIdRequest {
	r.depth = &depth
	return r
}
func (r ApiLocationsFindByRegionAndIdRequest) XContractNumber(xContractNumber int32) ApiLocationsFindByRegionAndIdRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiLocationsFindByRegionAndIdRequest) Execute() (Location, *APIResponse, error) {
	return r.ApiService.LocationsFindByRegionAndIdExecute(r)
}

/*
 * LocationsFindByRegionAndId Retrieve a Location
 * Retrieves the attributes of a given location
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param regionId
 * @param locationId
 * @return ApiLocationsFindByRegionAndIdRequest
 */
func (a *LocationApiService) LocationsFindByRegionAndId(ctx _context.Context, regionId string, locationId string) ApiLocationsFindByRegionAndIdRequest {
	return ApiLocationsFindByRegionAndIdRequest{
		ApiService: a,
		ctx:        ctx,
		regionId:   regionId,
		locationId: locationId,
	}
}

/*
 * Execute executes the request
 * @return Location
 */
func (a *LocationApiService) LocationsFindByRegionAndIdExecute(r ApiLocationsFindByRegionAndIdRequest) (Location, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Location
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationApiService.LocationsFindByRegionAndId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/locations/{regionId}/{locationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"regionId"+"}", _neturl.PathEscape(parameterToString(r.regionId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"locationId"+"}", _neturl.PathEscape(parameterToString(r.locationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:   localVarHTTPResponse,
		Method:     localVarHTTPMethod,
		RequestURL: localVarPath,
		Operation:  "LocationsFindByRegionAndId",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}

type ApiLocationsGetRequest struct {
	ctx             _context.Context
	ApiService      *LocationApiService
	pretty          *bool
	depth           *int32
	xContractNumber *int32
}

func (r ApiLocationsGetRequest) Pretty(pretty bool) ApiLocationsGetRequest {
	r.pretty = &pretty
	return r
}
func (r ApiLocationsGetRequest) Depth(depth int32) ApiLocationsGetRequest {
	r.depth = &depth
	return r
}
func (r ApiLocationsGetRequest) XContractNumber(xContractNumber int32) ApiLocationsGetRequest {
	r.xContractNumber = &xContractNumber
	return r
}

func (r ApiLocationsGetRequest) Execute() (Locations, *APIResponse, error) {
	return r.ApiService.LocationsGetExecute(r)
}

/*
 * LocationsGet List Locations
 * Retrieve a list of Locations. This list represents where you can provision your virtual data centers
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiLocationsGetRequest
 */
func (a *LocationApiService) LocationsGet(ctx _context.Context) ApiLocationsGetRequest {
	return ApiLocationsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return Locations
 */
func (a *LocationApiService) LocationsGetExecute(r ApiLocationsGetRequest) (Locations, *APIResponse, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Locations
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationApiService.LocationsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/locations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	if r.depth != nil {
		localVarQueryParams.Add("depth", parameterToString(*r.depth, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xContractNumber != nil {
		localVarHeaderParams["X-Contract-Number"] = parameterToString(*r.xContractNumber, "")
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Token Authentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)

	localVarAPIResponse := &APIResponse{
		Response:   localVarHTTPResponse,
		Method:     localVarHTTPMethod,
		RequestURL: localVarPath,
		Operation:  "LocationsGet",
	}

	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarAPIResponse.Payload = localVarBody
	if err != nil {
		return localVarReturnValue, localVarAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarAPIResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarAPIResponse, newErr
	}

	return localVarReturnValue, localVarAPIResponse, nil
}
