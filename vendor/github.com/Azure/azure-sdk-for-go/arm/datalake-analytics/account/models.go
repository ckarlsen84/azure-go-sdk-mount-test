package account

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// DataLakeAnalyticsAccountState enumerates the values for data lake analytics
// account state.
type DataLakeAnalyticsAccountState string

const (
	// Active specifies the active state for data lake analytics account state.
	Active DataLakeAnalyticsAccountState = "Active"
	// Suspended specifies the suspended state for data lake analytics account
	// state.
	Suspended DataLakeAnalyticsAccountState = "Suspended"
)

// DataLakeAnalyticsAccountStatus enumerates the values for data lake
// analytics account status.
type DataLakeAnalyticsAccountStatus string

const (
	// Creating specifies the creating state for data lake analytics account
	// status.
	Creating DataLakeAnalyticsAccountStatus = "Creating"
	// Deleted specifies the deleted state for data lake analytics account
	// status.
	Deleted DataLakeAnalyticsAccountStatus = "Deleted"
	// Deleting specifies the deleting state for data lake analytics account
	// status.
	Deleting DataLakeAnalyticsAccountStatus = "Deleting"
	// Failed specifies the failed state for data lake analytics account
	// status.
	Failed DataLakeAnalyticsAccountStatus = "Failed"
	// Patching specifies the patching state for data lake analytics account
	// status.
	Patching DataLakeAnalyticsAccountStatus = "Patching"
	// Resuming specifies the resuming state for data lake analytics account
	// status.
	Resuming DataLakeAnalyticsAccountStatus = "Resuming"
	// Running specifies the running state for data lake analytics account
	// status.
	Running DataLakeAnalyticsAccountStatus = "Running"
	// Succeeded specifies the succeeded state for data lake analytics account
	// status.
	Succeeded DataLakeAnalyticsAccountStatus = "Succeeded"
	// Suspending specifies the suspending state for data lake analytics
	// account status.
	Suspending DataLakeAnalyticsAccountStatus = "Suspending"
)

// AddDataLakeStoreParameters is additional Data Lake Store parameters.
type AddDataLakeStoreParameters struct {
	*DataLakeStoreAccountInfoProperties `json:"properties,omitempty"`
}

// AddStorageAccountParameters is storage account parameters for a storage
// account being added to a Data Lake Analytics account.
type AddStorageAccountParameters struct {
	*StorageAccountProperties `json:"properties,omitempty"`
}

// CreateStorageAccountInfo is azure Storage account information to add to the
// Data Lake analytics account being created.
type CreateStorageAccountInfo struct {
	Name                      *string `json:"name,omitempty"`
	*StorageAccountProperties `json:"properties,omitempty"`
}

// DataLakeAnalyticsAccount is a Data Lake Analytics account object,
// containing all information associated with the named Data Lake Analytics
// account.
type DataLakeAnalyticsAccount struct {
	autorest.Response                   `json:"-"`
	ID                                  *string             `json:"id,omitempty"`
	Name                                *string             `json:"name,omitempty"`
	Type                                *string             `json:"type,omitempty"`
	Location                            *string             `json:"location,omitempty"`
	Tags                                *map[string]*string `json:"tags,omitempty"`
	*DataLakeAnalyticsAccountProperties `json:"properties,omitempty"`
}

// DataLakeAnalyticsAccountListDataLakeStoreResult is data Lake Account list
// information.
type DataLakeAnalyticsAccountListDataLakeStoreResult struct {
	autorest.Response `json:"-"`
	Value             *[]DataLakeStoreAccountInfo `json:"value,omitempty"`
	NextLink          *string                     `json:"nextLink,omitempty"`
}

// DataLakeAnalyticsAccountListDataLakeStoreResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DataLakeAnalyticsAccountListDataLakeStoreResult) DataLakeAnalyticsAccountListDataLakeStoreResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// DataLakeAnalyticsAccountListResult is dataLakeAnalytics Account list
// information.
type DataLakeAnalyticsAccountListResult struct {
	autorest.Response `json:"-"`
	Value             *[]DataLakeAnalyticsAccount `json:"value,omitempty"`
	NextLink          *string                     `json:"nextLink,omitempty"`
}

// DataLakeAnalyticsAccountListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DataLakeAnalyticsAccountListResult) DataLakeAnalyticsAccountListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// DataLakeAnalyticsAccountListStorageAccountsResult is azure Storage Account
// list information.
type DataLakeAnalyticsAccountListStorageAccountsResult struct {
	autorest.Response `json:"-"`
	Value             *[]StorageAccountInfo `json:"value,omitempty"`
	NextLink          *string               `json:"nextLink,omitempty"`
}

// DataLakeAnalyticsAccountListStorageAccountsResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DataLakeAnalyticsAccountListStorageAccountsResult) DataLakeAnalyticsAccountListStorageAccountsResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// DataLakeAnalyticsAccountProperties is the account specific properties that
// are associated with an underlying Data Lake Analytics account.
type DataLakeAnalyticsAccountProperties struct {
	ProvisioningState            DataLakeAnalyticsAccountStatus `json:"provisioningState,omitempty"`
	State                        DataLakeAnalyticsAccountState  `json:"state,omitempty"`
	DefaultDataLakeStoreAccount  *string                        `json:"defaultDataLakeStoreAccount,omitempty"`
	MaxDegreeOfParallelism       *int32                         `json:"maxDegreeOfParallelism,omitempty"`
	QueryStoreRetention          *int32                         `json:"queryStoreRetention,omitempty"`
	MaxJobCount                  *int32                         `json:"maxJobCount,omitempty"`
	SystemMaxDegreeOfParallelism *int32                         `json:"systemMaxDegreeOfParallelism,omitempty"`
	SystemMaxJobCount            *int32                         `json:"systemMaxJobCount,omitempty"`
	DataLakeStoreAccounts        *[]DataLakeStoreAccountInfo    `json:"dataLakeStoreAccounts,omitempty"`
	StorageAccounts              *[]StorageAccountInfo          `json:"storageAccounts,omitempty"`
	CreationTime                 *date.Time                     `json:"creationTime,omitempty"`
	LastModifiedTime             *date.Time                     `json:"lastModifiedTime,omitempty"`
	Endpoint                     *string                        `json:"endpoint,omitempty"`
}

// DataLakeAnalyticsAccountUpdateParameters is the parameters that can be used
// to update an existing Data Lake Analytics account.
type DataLakeAnalyticsAccountUpdateParameters struct {
	Tags                                      *map[string]*string `json:"tags,omitempty"`
	*UpdateDataLakeAnalyticsAccountProperties `json:"properties,omitempty"`
}

// DataLakeStoreAccountInfo is data Lake Store account information.
type DataLakeStoreAccountInfo struct {
	autorest.Response                   `json:"-"`
	Name                                *string `json:"name,omitempty"`
	*DataLakeStoreAccountInfoProperties `json:"properties,omitempty"`
}

// DataLakeStoreAccountInfoProperties is data Lake Store account properties
// information.
type DataLakeStoreAccountInfoProperties struct {
	Suffix *string `json:"suffix,omitempty"`
}

// Error is generic resource error information.
type Error struct {
	Code       *string         `json:"code,omitempty"`
	Message    *string         `json:"message,omitempty"`
	Target     *string         `json:"target,omitempty"`
	Details    *[]ErrorDetails `json:"details,omitempty"`
	InnerError *InnerError     `json:"innerError,omitempty"`
}

// ErrorDetails is generic resource error details information.
type ErrorDetails struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Target  *string `json:"target,omitempty"`
}

// InnerError is generic resource inner error information.
type InnerError struct {
	Trace   *string `json:"trace,omitempty"`
	Context *string `json:"context,omitempty"`
}

// ListSasTokensResult is the SAS response that contains the storage account,
// container and associated SAS token for connection use.
type ListSasTokensResult struct {
	autorest.Response `json:"-"`
	Value             *[]SasTokenInfo `json:"value,omitempty"`
	NextLink          *string         `json:"nextLink,omitempty"`
}

// ListSasTokensResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ListSasTokensResult) ListSasTokensResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ListStorageContainersResult is the list of blob containers associated with
// the storage account attached to the Data Lake Analytics account.
type ListStorageContainersResult struct {
	autorest.Response `json:"-"`
	Value             *[]StorageContainer `json:"value,omitempty"`
	NextLink          *string             `json:"nextLink,omitempty"`
}

// ListStorageContainersResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ListStorageContainersResult) ListStorageContainersResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Resource is the Resource model definition.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// SasTokenInfo is sAS token information.
type SasTokenInfo struct {
	AccessToken *string `json:"accessToken,omitempty"`
}

// StorageAccountInfo is azure Storage account information.
type StorageAccountInfo struct {
	autorest.Response         `json:"-"`
	Name                      *string `json:"name,omitempty"`
	*StorageAccountProperties `json:"properties,omitempty"`
}

// StorageAccountProperties is azure Storage account properties information.
type StorageAccountProperties struct {
	AccessKey *string `json:"accessKey,omitempty"`
	Suffix    *string `json:"suffix,omitempty"`
}

// StorageContainer is azure Storage blob container information.
type StorageContainer struct {
	autorest.Response           `json:"-"`
	Name                        *string `json:"name,omitempty"`
	ID                          *string `json:"id,omitempty"`
	Type                        *string `json:"type,omitempty"`
	*StorageContainerProperties `json:"properties,omitempty"`
}

// StorageContainerProperties is azure Storage blob container properties
// information.
type StorageContainerProperties struct {
	LastModifiedTime *date.Time `json:"lastModifiedTime,omitempty"`
}

// UpdateDataLakeAnalyticsAccountProperties is the properties to update that
// are associated with an underlying Data Lake Analytics account to.
type UpdateDataLakeAnalyticsAccountProperties struct {
	MaxDegreeOfParallelism *int32 `json:"maxDegreeOfParallelism,omitempty"`
	QueryStoreRetention    *int32 `json:"queryStoreRetention,omitempty"`
	MaxJobCount            *int32 `json:"maxJobCount,omitempty"`
}

// UpdateStorageAccountParameters is storage account parameters for a storage
// account being updated in a Data Lake Analytics account.
type UpdateStorageAccountParameters struct {
	*UpdateStorageAccountProperties `json:"properties,omitempty"`
}

// UpdateStorageAccountProperties is azure Storage account properties
// information to update.
type UpdateStorageAccountProperties struct {
	AccessKey *string `json:"accessKey,omitempty"`
	Suffix    *string `json:"suffix,omitempty"`
}
