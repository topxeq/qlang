package dysmsapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

	qlang "github.com/topxeq/qlang/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi",

	"EndpointMap":  dysmsapi.EndpointMap,
	"EndpointType": dysmsapi.EndpointType,

	"GetEndpointMap":          dysmsapi.GetEndpointMap,
	"GetEndpointType":         dysmsapi.GetEndpointType,
	"SetClientProperty":       dysmsapi.SetClientProperty,
	"SetEndpointDataToClient": dysmsapi.SetEndpointDataToClient,

	"AddSmsSignRequest":                qlang.StructOf((*dysmsapi.AddSmsSignRequest)(nil)),
	"CreateAddSmsSignRequest":          dysmsapi.CreateAddSmsSignRequest,
	"AddSmsSignResponse":               qlang.StructOf((*dysmsapi.AddSmsSignResponse)(nil)),
	"CreateAddSmsSignResponse":         dysmsapi.CreateAddSmsSignResponse,
	"AddSmsSignSignFileList":           qlang.StructOf((*dysmsapi.AddSmsSignSignFileList)(nil)),
	"AddSmsTemplateRequest":            qlang.StructOf((*dysmsapi.AddSmsTemplateRequest)(nil)),
	"CreateAddSmsTemplateRequest":      dysmsapi.CreateAddSmsTemplateRequest,
	"AddSmsTemplateResponse":           qlang.StructOf((*dysmsapi.AddSmsTemplateResponse)(nil)),
	"CreateAddSmsTemplateResponse":     dysmsapi.CreateAddSmsTemplateResponse,
	"Client":                           qlang.StructOf((*dysmsapi.Client)(nil)),
	"NewClient":                        dysmsapi.NewClient,
	"NewClientWithAccessKey":           dysmsapi.NewClientWithAccessKey,
	"NewClientWithEcsRamRole":          dysmsapi.NewClientWithEcsRamRole,
	"NewClientWithOptions":             dysmsapi.NewClientWithOptions,
	"NewClientWithProvider":            dysmsapi.NewClientWithProvider,
	"NewClientWithRamRoleArn":          dysmsapi.NewClientWithRamRoleArn,
	"NewClientWithRamRoleArnAndPolicy": dysmsapi.NewClientWithRamRoleArnAndPolicy,
	"NewClientWithRsaKeyPair":          dysmsapi.NewClientWithRsaKeyPair,
	"NewClientWithStsToken":            dysmsapi.NewClientWithStsToken,
	"DeleteSmsSignRequest":             qlang.StructOf((*dysmsapi.DeleteSmsSignRequest)(nil)),
	"CreateDeleteSmsSignRequest":       dysmsapi.CreateDeleteSmsSignRequest,
	"DeleteSmsTemplateRequest":         qlang.StructOf((*dysmsapi.DeleteSmsTemplateRequest)(nil)),
	"CreateDeleteSmsTemplateRequest":   dysmsapi.CreateDeleteSmsTemplateRequest,
	"ModifySmsSignRequest":             qlang.StructOf((*dysmsapi.ModifySmsSignRequest)(nil)),
	"CreateModifySmsSignRequest":       dysmsapi.CreateModifySmsSignRequest,
	"ModifySmsSignSignFileList":        qlang.StructOf((*dysmsapi.ModifySmsSignSignFileList)(nil)),
	"ModifySmsTemplateRequest":         qlang.StructOf((*dysmsapi.ModifySmsTemplateRequest)(nil)),
	"CreateModifySmsTemplateRequest":   dysmsapi.CreateModifySmsTemplateRequest,
	"QuerySendDetailsRequest":          qlang.StructOf((*dysmsapi.QuerySendDetailsRequest)(nil)),
	"CreateQuerySendDetailsRequest":    dysmsapi.CreateQuerySendDetailsRequest,
	"QuerySmsSignRequest":              qlang.StructOf((*dysmsapi.QuerySmsSignRequest)(nil)),
	"CreateQuerySmsSignRequest":        dysmsapi.CreateQuerySmsSignRequest,
	"QuerySmsTemplateRequest":          qlang.StructOf((*dysmsapi.QuerySmsTemplateRequest)(nil)),
	"CreateQuerySmsTemplateRequest":    dysmsapi.CreateQuerySmsTemplateRequest,
	"SendBatchSmsRequest":              qlang.StructOf((*dysmsapi.SendBatchSmsRequest)(nil)),
	"CreateSendBatchSmsRequest":        dysmsapi.CreateSendBatchSmsRequest,
	"SendSmsRequest":                   qlang.StructOf((*dysmsapi.SendSmsRequest)(nil)),
	"CreateSendSmsRequest":             dysmsapi.CreateSendSmsRequest,
	"SmsSendDetailDTO":                 qlang.StructOf((*dysmsapi.SmsSendDetailDTO)(nil)),
	"SmsSendDetailDTOs":                qlang.StructOf((*dysmsapi.SmsSendDetailDTOs)(nil)),

	// added by madarin
	"NewInteger": requests.NewInteger,
}
