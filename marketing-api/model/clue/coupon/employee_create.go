package coupon

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// EmployeeCreateRequest 添加核销员 API Request
type EmployeeCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// List 核销员列表
	List []Employee `json:"list,omitempty"`
}

// Encode implement PostRequest interface
func (r EmployeeCreateRequest) Encode() []byte {
	b, _ := json.Marshal(r)
	return b
}

// EmployeeCreateResponse 添加核销员 API Response
type EmployeeCreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *EmployeeCreateResponseData `json:"data,omitempty"`
}

// EmployeeCreateResponseData json返回值
type EmployeeCreateResponseData struct {
	// ExistedList 本次添加已存在的核销员列表
	ExistedList []Employee `json:"existed_list,omitempty"`
	// SuccessList 本次添加成功的核销员列表
	SuccessList []Employee `json:"success_list,omitempty"`
}
