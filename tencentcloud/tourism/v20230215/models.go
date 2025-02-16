// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20230215

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeDrawResourceListRequestParams struct {
	// PageNumber
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// PageSize
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeDrawResourceListRequest struct {
	*tchttp.BaseRequest
	
	// PageNumber
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// PageSize
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeDrawResourceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrawResourceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDrawResourceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDrawResourceListResponseParams struct {
	// 返回数据条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 返回数据数组
	ResourceDrawList []*ResourceDrawListType `json:"ResourceDrawList,omitempty" name:"ResourceDrawList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDrawResourceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDrawResourceListResponseParams `json:"Response"`
}

func (r *DescribeDrawResourceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrawResourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceDrawListType struct {
	// 记录id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 资源记录id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`

	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 本订单资源序列号
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexId *string `json:"IndexId,omitempty" name:"IndexId"`

	// 客户的uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitempty" name:"Uin"`

	// 大订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	BigDealId *string `json:"BigDealId,omitempty" name:"BigDealId"`

	// 小订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmallOrderId *string `json:"SmallOrderId,omitempty" name:"SmallOrderId"`

	// 资源创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceNewStartTime *string `json:"ResourceNewStartTime,omitempty" name:"ResourceNewStartTime"`

	// 资源到期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceNewEndTime *string `json:"ResourceNewEndTime,omitempty" name:"ResourceNewEndTime"`

	// 资源状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceStatus *int64 `json:"ResourceStatus,omitempty" name:"ResourceStatus"`

	// 本记录状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 项目类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *uint64 `json:"ResourceType,omitempty" name:"ResourceType"`
}