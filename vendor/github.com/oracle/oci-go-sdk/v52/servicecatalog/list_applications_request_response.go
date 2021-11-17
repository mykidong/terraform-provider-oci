// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicecatalog

import (
	"github.com/oracle/oci-go-sdk/v52/common"
	"net/http"
)

// ListApplicationsRequest wrapper for the ListApplications operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ListApplications.go.html to see an example of how to use ListApplicationsRequest.
type ListApplicationsRequest struct {

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The unique identifier for the service catalog.
	ServiceCatalogId *string `mandatory:"false" contributesTo:"query" name:"serviceCatalogId"`

	// The type of the application in the service catalog.
	EntityType *string `mandatory:"false" contributesTo:"query" name:"entityType"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// How many records to return. Specify a value greater than zero and less than or equal to 1000. The default is 30.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Exact match name filter.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The unique identifier of the entity associated with service catalog.
	EntityId *string `mandatory:"false" contributesTo:"query" name:"entityId"`

	// Limit results to just this publisher.
	PublisherId []string `contributesTo:"query" name:"publisherId" collectionFormat:"multi"`

	// Name of the package type. If multiple package types are provided, then any resource with
	// one or more matching package types will be returned.
	PackageType []PackageTypeEnumEnum `contributesTo:"query" name:"packageType" omitEmpty:"true" collectionFormat:"multi"`

	// Name of the pricing type. If multiple pricing types are provided, then any resource with
	// one or more matching pricing models will be returned.
	Pricing []PricingTypeEnumEnum `contributesTo:"query" name:"pricing" omitEmpty:"true" collectionFormat:"multi"`

	// Indicates whether to show only featured resources. If this is set to `false` or is omitted, then all resources will be returned.
	IsFeatured *bool `mandatory:"false" contributesTo:"query" name:"isFeatured"`

	// The sort order to apply, either `ASC` or `DESC`. Default is `ASC`.
	SortOrder ListApplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApplicationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApplicationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListApplicationsResponse wrapper for the ListApplications operation
type ListApplicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ApplicationCollection instances
	ApplicationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListApplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApplicationsSortOrderEnum Enum with underlying type: string
type ListApplicationsSortOrderEnum string

// Set of constants representing the allowable values for ListApplicationsSortOrderEnum
const (
	ListApplicationsSortOrderAsc  ListApplicationsSortOrderEnum = "ASC"
	ListApplicationsSortOrderDesc ListApplicationsSortOrderEnum = "DESC"
)

var mappingListApplicationsSortOrder = map[string]ListApplicationsSortOrderEnum{
	"ASC":  ListApplicationsSortOrderAsc,
	"DESC": ListApplicationsSortOrderDesc,
}

// GetListApplicationsSortOrderEnumValues Enumerates the set of values for ListApplicationsSortOrderEnum
func GetListApplicationsSortOrderEnumValues() []ListApplicationsSortOrderEnum {
	values := make([]ListApplicationsSortOrderEnum, 0)
	for _, v := range mappingListApplicationsSortOrder {
		values = append(values, v)
	}
	return values
}
