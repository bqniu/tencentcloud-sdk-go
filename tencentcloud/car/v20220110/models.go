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

package v20220110

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type ApplyConcurrentRequestParams struct {
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户IP，用户客户端的公网IP，用于就近调度
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 应用版本ID
	ApplicationVersionId *string `json:"ApplicationVersionId,omitempty" name:"ApplicationVersionId"`

	// 应用ID。如果是独享项目，将忽略该参数，使用项目绑定的应用。如果是共享项目，使用该参数来指定应用。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

type ApplyConcurrentRequest struct {
	*tchttp.BaseRequest
	
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户IP，用户客户端的公网IP，用于就近调度
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 应用版本ID
	ApplicationVersionId *string `json:"ApplicationVersionId,omitempty" name:"ApplicationVersionId"`

	// 应用ID。如果是独享项目，将忽略该参数，使用项目绑定的应用。如果是共享项目，使用该参数来指定应用。
	ApplicationId *string `json:"ApplicationId,omitempty" name:"ApplicationId"`
}

func (r *ApplyConcurrentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConcurrentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserIp")
	delete(f, "ProjectId")
	delete(f, "ApplicationVersionId")
	delete(f, "ApplicationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ApplyConcurrentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ApplyConcurrentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ApplyConcurrentResponse struct {
	*tchttp.BaseResponse
	Response *ApplyConcurrentResponseParams `json:"Response"`
}

func (r *ApplyConcurrentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ApplyConcurrentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSessionRequestParams struct {
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户IP，用户客户端的公网IP，用于就近调度
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 客户端session信息，从SDK请求中获得。特殊的，当 RunMode 参数为 RunWithoutClient 时，该字段可以为空
	ClientSession *string `json:"ClientSession,omitempty" name:"ClientSession"`

	// 云端运行模式。
	// RunWithoutClient：允许无客户端连接的情况下仍保持云端 App 运行
	// 默认值（空）：要求必须有客户端连接才会保持云端 App 运行。
	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`

	// 【多人互动】房主用户ID，在多人互动模式下为必填字段。
	// 如果该用户是房主，HostUserId需要和UserId保持一致；
	// 如果该用户非房主，HostUserId需要填写房主的HostUserId。
	HostUserId *string `json:"HostUserId,omitempty" name:"HostUserId"`

	// 【多人互动】角色。
	// Player：玩家（可通过键鼠等操作应用）
	// Viewer：观察者（只能观看，无法操作）
	Role *string `json:"Role,omitempty" name:"Role"`
}

type CreateSessionRequest struct {
	*tchttp.BaseRequest
	
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 用户IP，用户客户端的公网IP，用于就近调度
	UserIp *string `json:"UserIp,omitempty" name:"UserIp"`

	// 客户端session信息，从SDK请求中获得。特殊的，当 RunMode 参数为 RunWithoutClient 时，该字段可以为空
	ClientSession *string `json:"ClientSession,omitempty" name:"ClientSession"`

	// 云端运行模式。
	// RunWithoutClient：允许无客户端连接的情况下仍保持云端 App 运行
	// 默认值（空）：要求必须有客户端连接才会保持云端 App 运行。
	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`

	// 【多人互动】房主用户ID，在多人互动模式下为必填字段。
	// 如果该用户是房主，HostUserId需要和UserId保持一致；
	// 如果该用户非房主，HostUserId需要填写房主的HostUserId。
	HostUserId *string `json:"HostUserId,omitempty" name:"HostUserId"`

	// 【多人互动】角色。
	// Player：玩家（可通过键鼠等操作应用）
	// Viewer：观察者（只能观看，无法操作）
	Role *string `json:"Role,omitempty" name:"Role"`
}

func (r *CreateSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserIp")
	delete(f, "ClientSession")
	delete(f, "RunMode")
	delete(f, "HostUserId")
	delete(f, "Role")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSessionResponseParams struct {
	// 服务端session信息，返回给SDK
	ServerSession *string `json:"ServerSession,omitempty" name:"ServerSession"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateSessionResponseParams `json:"Response"`
}

func (r *CreateSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroySessionRequestParams struct {
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type DestroySessionRequest struct {
	*tchttp.BaseRequest
	
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *DestroySessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroySessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroySessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroySessionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DestroySessionResponse struct {
	*tchttp.BaseResponse
	Response *DestroySessionResponseParams `json:"Response"`
}

func (r *DestroySessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroySessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamRequestParams struct {
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（UserId将作为StreamId进行推流，比如绑定推流域名为abc.livepush.myqcloud.com，那么推流地址为rtmp://abc.livepush.myqcloud.com/live/UserId?txSecret=xxx&txTime=xxx）
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type StartPublishStreamRequest struct {
	*tchttp.BaseRequest
	
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（UserId将作为StreamId进行推流，比如绑定推流域名为abc.livepush.myqcloud.com，那么推流地址为rtmp://abc.livepush.myqcloud.com/live/UserId?txSecret=xxx&txTime=xxx）
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *StartPublishStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartPublishStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartPublishStreamResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartPublishStreamResponse struct {
	*tchttp.BaseResponse
	Response *StartPublishStreamResponseParams `json:"Response"`
}

func (r *StartPublishStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartPublishStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishStreamRequestParams struct {
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

type StopPublishStreamRequest struct {
	*tchttp.BaseRequest
	
	// 唯一用户身份标识，由业务方自定义，平台不予理解。（可根据业务需要决定使用用户的唯一身份标识或是使用时间戳随机生成；在用户重连时应保持UserId不变）
	UserId *string `json:"UserId,omitempty" name:"UserId"`
}

func (r *StopPublishStreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishStreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopPublishStreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopPublishStreamResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopPublishStreamResponse struct {
	*tchttp.BaseResponse
	Response *StopPublishStreamResponseParams `json:"Response"`
}

func (r *StopPublishStreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopPublishStreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}