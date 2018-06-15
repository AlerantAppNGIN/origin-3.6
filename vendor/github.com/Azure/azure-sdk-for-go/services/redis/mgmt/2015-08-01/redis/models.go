package redis

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// KeyType enumerates the values for key type.
type KeyType string

const (
	// Primary ...
	Primary KeyType = "Primary"
	// Secondary ...
	Secondary KeyType = "Secondary"
)

// RebootType enumerates the values for reboot type.
type RebootType string

const (
	// AllNodes ...
	AllNodes RebootType = "AllNodes"
	// PrimaryNode ...
	PrimaryNode RebootType = "PrimaryNode"
	// SecondaryNode ...
	SecondaryNode RebootType = "SecondaryNode"
)

// SkuFamily enumerates the values for sku family.
type SkuFamily string

const (
	// C ...
	C SkuFamily = "C"
	// P ...
	P SkuFamily = "P"
)

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic ...
	Basic SkuName = "Basic"
	// Premium ...
	Premium SkuName = "Premium"
	// Standard ...
	Standard SkuName = "Standard"
)

// AccessKeys redis cache access keys.
type AccessKeys struct {
	// PrimaryKey - The current primary key that clients can use to authenticate with Redis cache.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - The current secondary key that clients can use to authenticate with Redis cache.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// CreateOrUpdateParameters parameters supplied to the CreateOrUpdate Redis operation.
type CreateOrUpdateParameters struct {
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags *map[string]*string `json:"tags,omitempty"`
	// Properties - Redis cache properties.
	*Properties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for CreateOrUpdateParameters struct.
func (coup *CreateOrUpdateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["properties"]
	if v != nil {
		var properties Properties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		coup.Properties = &properties
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		coup.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		coup.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		coup.Type = &typeVar
	}

	v = m["location"]
	if v != nil {
		var location string
		err = json.Unmarshal(*m["location"], &location)
		if err != nil {
			return err
		}
		coup.Location = &location
	}

	v = m["tags"]
	if v != nil {
		var tags map[string]*string
		err = json.Unmarshal(*m["tags"], &tags)
		if err != nil {
			return err
		}
		coup.Tags = &tags
	}

	return nil
}

// ListKeysResult the response of Redis list keys operation.
type ListKeysResult struct {
	autorest.Response `json:"-"`
	// PrimaryKey - The current primary key that clients can use to authenticate with Redis cache.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - The current secondary key that clients can use to authenticate with Redis cache.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// ListResult the response of list Redis operation.
type ListResult struct {
	autorest.Response `json:"-"`
	// Value - Results of the list operation.
	Value *[]ResourceType `json:"value,omitempty"`
	// NextLink - Link for next set of locations.
	NextLink *string `json:"nextLink,omitempty"`
}

// ListResultIterator provides access to a complete listing of ResourceType values.
type ListResultIterator struct {
	i    int
	page ListResultPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListResultIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListResultIterator) Response() ListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListResultIterator) Value() ResourceType {
	if !iter.page.NotDone() {
		return ResourceType{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (lr ListResult) IsEmpty() bool {
	return lr.Value == nil || len(*lr.Value) == 0
}

// listResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lr ListResult) listResultPreparer() (*http.Request, error) {
	if lr.NextLink == nil || len(to.String(lr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lr.NextLink)))
}

// ListResultPage contains a page of ResourceType values.
type ListResultPage struct {
	fn func(ListResult) (ListResult, error)
	lr ListResult
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListResultPage) Next() error {
	next, err := page.fn(page.lr)
	if err != nil {
		return err
	}
	page.lr = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListResultPage) NotDone() bool {
	return !page.lr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListResultPage) Response() ListResult {
	return page.lr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListResultPage) Values() []ResourceType {
	if page.lr.IsEmpty() {
		return nil
	}
	return *page.lr.Value
}

// Properties parameters supplied to CreateOrUpdate Redis operation.
type Properties struct {
	// RedisVersion - RedisVersion parameter has been deprecated. As such, it is no longer necessary to provide this parameter and any value specified is ignored.
	RedisVersion *string `json:"redisVersion,omitempty"`
	// Sku - What SKU of Redis cache to deploy.
	Sku *Sku `json:"sku,omitempty"`
	// RedisConfiguration - All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
	RedisConfiguration *map[string]*string `json:"redisConfiguration,omitempty"`
	// EnableNonSslPort - If the value is true, then the non-SLL Redis server port (6379) will be enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`
	// TenantSettings - tenantSettings
	TenantSettings *map[string]*string `json:"tenantSettings,omitempty"`
	// ShardCount - The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int32 `json:"shardCount,omitempty"`
	// VirtualNetwork - The exact ARM resource ID of the virtual network to deploy the Redis cache in. Example format: /subscriptions/{subid}/resourceGroups/{resourceGroupName}/Microsoft.ClassicNetwork/VirtualNetworks/vnet1
	VirtualNetwork *string `json:"virtualNetwork,omitempty"`
	// Subnet - Required when deploying a Redis cache inside an existing Azure Virtual Network.
	Subnet *string `json:"subnet,omitempty"`
	// StaticIP - Required when deploying a Redis cache inside an existing Azure Virtual Network.
	StaticIP *string `json:"staticIP,omitempty"`
}

// ReadableProperties parameters describing a Redis instance
type ReadableProperties struct {
	// RedisVersion - RedisVersion parameter has been deprecated. As such, it is no longer necessary to provide this parameter and any value specified is ignored.
	RedisVersion *string `json:"redisVersion,omitempty"`
	// Sku - What SKU of Redis cache to deploy.
	Sku *Sku `json:"sku,omitempty"`
	// RedisConfiguration - All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
	RedisConfiguration *map[string]*string `json:"redisConfiguration,omitempty"`
	// EnableNonSslPort - If the value is true, then the non-SLL Redis server port (6379) will be enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`
	// TenantSettings - tenantSettings
	TenantSettings *map[string]*string `json:"tenantSettings,omitempty"`
	// ShardCount - The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int32 `json:"shardCount,omitempty"`
	// VirtualNetwork - The exact ARM resource ID of the virtual network to deploy the Redis cache in. Example format: /subscriptions/{subid}/resourceGroups/{resourceGroupName}/Microsoft.ClassicNetwork/VirtualNetworks/vnet1
	VirtualNetwork *string `json:"virtualNetwork,omitempty"`
	// Subnet - Required when deploying a Redis cache inside an existing Azure Virtual Network.
	Subnet *string `json:"subnet,omitempty"`
	// StaticIP - Required when deploying a Redis cache inside an existing Azure Virtual Network.
	StaticIP *string `json:"staticIP,omitempty"`
	// ProvisioningState - Redis instance provisioning status.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// HostName - Redis host name.
	HostName *string `json:"hostName,omitempty"`
	// Port - Redis non-SSL port.
	Port *int32 `json:"port,omitempty"`
	// SslPort - Redis SSL port.
	SslPort *int32 `json:"sslPort,omitempty"`
}

// ReadablePropertiesWithAccessKey properties generated only in response to CreateOrUpdate Redis operation.
type ReadablePropertiesWithAccessKey struct {
	// RedisVersion - RedisVersion parameter has been deprecated. As such, it is no longer necessary to provide this parameter and any value specified is ignored.
	RedisVersion *string `json:"redisVersion,omitempty"`
	// Sku - What SKU of Redis cache to deploy.
	Sku *Sku `json:"sku,omitempty"`
	// RedisConfiguration - All Redis Settings. Few possible keys: rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value etc.
	RedisConfiguration *map[string]*string `json:"redisConfiguration,omitempty"`
	// EnableNonSslPort - If the value is true, then the non-SLL Redis server port (6379) will be enabled.
	EnableNonSslPort *bool `json:"enableNonSslPort,omitempty"`
	// TenantSettings - tenantSettings
	TenantSettings *map[string]*string `json:"tenantSettings,omitempty"`
	// ShardCount - The number of shards to be created on a Premium Cluster Cache.
	ShardCount *int32 `json:"shardCount,omitempty"`
	// VirtualNetwork - The exact ARM resource ID of the virtual network to deploy the Redis cache in. Example format: /subscriptions/{subid}/resourceGroups/{resourceGroupName}/Microsoft.ClassicNetwork/VirtualNetworks/vnet1
	VirtualNetwork *string `json:"virtualNetwork,omitempty"`
	// Subnet - Required when deploying a Redis cache inside an existing Azure Virtual Network.
	Subnet *string `json:"subnet,omitempty"`
	// StaticIP - Required when deploying a Redis cache inside an existing Azure Virtual Network.
	StaticIP *string `json:"staticIP,omitempty"`
	// ProvisioningState - Redis instance provisioning status.
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// HostName - Redis host name.
	HostName *string `json:"hostName,omitempty"`
	// Port - Redis non-SSL port.
	Port *int32 `json:"port,omitempty"`
	// SslPort - Redis SSL port.
	SslPort *int32 `json:"sslPort,omitempty"`
	// AccessKeys - Redis cache access keys.
	AccessKeys *AccessKeys `json:"accessKeys,omitempty"`
}

// RebootParameters specifies which Redis node(s) to reboot.
type RebootParameters struct {
	// RebootType - Which Redis node(s) to reboot. Depending on this value data loss is possible. Possible values include: 'PrimaryNode', 'SecondaryNode', 'AllNodes'
	RebootType RebootType `json:"rebootType,omitempty"`
	// ShardID - If clustering is enabled, the ID of the shared be rebooted.
	ShardID *int32 `json:"shardId,omitempty"`
}

// RegenerateKeyParameters specifies which Redis access keys to reset.
type RegenerateKeyParameters struct {
	// KeyType - Which Redis access key to reset. Possible values include: 'Primary', 'Secondary'
	KeyType KeyType `json:"keyType,omitempty"`
}

// Resource ...
type Resource struct {
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags *map[string]*string `json:"tags,omitempty"`
}

// ResourceType a single Redis item in List or Get Operation.
type ResourceType struct {
	autorest.Response `json:"-"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags *map[string]*string `json:"tags,omitempty"`
	// ReadableProperties - Redis cache properties.
	*ReadableProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for ResourceType struct.
func (rt *ResourceType) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["properties"]
	if v != nil {
		var properties ReadableProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		rt.ReadableProperties = &properties
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		rt.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		rt.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		rt.Type = &typeVar
	}

	v = m["location"]
	if v != nil {
		var location string
		err = json.Unmarshal(*m["location"], &location)
		if err != nil {
			return err
		}
		rt.Location = &location
	}

	v = m["tags"]
	if v != nil {
		var tags map[string]*string
		err = json.Unmarshal(*m["tags"], &tags)
		if err != nil {
			return err
		}
		rt.Tags = &tags
	}

	return nil
}

// ResourceWithAccessKey a Redis item in CreateOrUpdate Operation response.
type ResourceWithAccessKey struct {
	autorest.Response `json:"-"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags *map[string]*string `json:"tags,omitempty"`
	// ReadablePropertiesWithAccessKey - Redis cache properties.
	*ReadablePropertiesWithAccessKey `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for ResourceWithAccessKey struct.
func (rwak *ResourceWithAccessKey) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["properties"]
	if v != nil {
		var properties ReadablePropertiesWithAccessKey
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		rwak.ReadablePropertiesWithAccessKey = &properties
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		rwak.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		rwak.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		rwak.Type = &typeVar
	}

	v = m["location"]
	if v != nil {
		var location string
		err = json.Unmarshal(*m["location"], &location)
		if err != nil {
			return err
		}
		rwak.Location = &location
	}

	v = m["tags"]
	if v != nil {
		var tags map[string]*string
		err = json.Unmarshal(*m["tags"], &tags)
		if err != nil {
			return err
		}
		rwak.Tags = &tags
	}

	return nil
}

// Sku SKU parameters supplied to the create Redis operation.
type Sku struct {
	// Name - What type of Redis cache to deploy. Valid values: (Basic, Standard, Premium). Possible values include: 'Basic', 'Standard', 'Premium'
	Name SkuName `json:"name,omitempty"`
	// Family - Which family to use. Valid values: (C, P). Possible values include: 'C', 'P'
	Family SkuFamily `json:"family,omitempty"`
	// Capacity - What size of Redis cache to deploy. Valid values: for C family (0, 1, 2, 3, 4, 5, 6), for P family (1, 2, 3, 4).
	Capacity *int32 `json:"capacity,omitempty"`
}
