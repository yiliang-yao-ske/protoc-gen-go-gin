// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v0.0.1
// - protoc            v3.21.12
// source: api/article/article_ecode.proto

package article

import (
	errors "errors"
	ecode "github.com/sunmi-OS/gocore/v2/api/ecode"
)

const (
	UNKNOW_ERROR                     = 0
	ERR_PRODUCT_NOT_FOUND            = 16050030
	ERR_PRODUCT_MODULE_NOT_FOUND     = 16050033
	ERR_SAAS_DEVICE_BIND_FAIL        = 16050031
	ERR_PRODUCT_APP_FOUND_FAIL       = 16050032
	ERR_REGISTER_DEVICE_LIMIT        = 16050016
	ERR_REGISTER_DEVICE_SINGLE_FAILE = 16050017
	ERR_REGISTER_DEVICE_BATCH_FAILE  = 16050018
)

var (
	ErrMap = map[int]string{
		UNKNOW_ERROR:                     "unknow error",
		ERR_PRODUCT_NOT_FOUND:            "product not found",
		ERR_PRODUCT_MODULE_NOT_FOUND:     "product module not found",
		ERR_SAAS_DEVICE_BIND_FAIL:        "saas device bind fail",
		ERR_PRODUCT_APP_FOUND_FAIL:       "product app found fail",
		ERR_REGISTER_DEVICE_LIMIT:        "register device limit",
		ERR_REGISTER_DEVICE_SINGLE_FAILE: "register device single faile",
		ERR_REGISTER_DEVICE_BATCH_FAILE:  "register device batch faile",
	}
)

func makeNewErr(code int, msg ...string) *ecode.ErrorV2 {
	msgStr := ErrMap[code]
	if len(msg) > 0 {
		msgStr = msg[0]
	}
	return ecode.NewV2(code, msgStr)
}

func UnknowError(msg ...string) *ecode.ErrorV2 {
	return makeNewErr(UNKNOW_ERROR, msg...)
}

func ErrProductNotFound(msg ...string) *ecode.ErrorV2 {
	return makeNewErr(ERR_PRODUCT_NOT_FOUND, msg...)
}

func ErrProductModuleNotFound(msg ...string) *ecode.ErrorV2 {
	return makeNewErr(ERR_PRODUCT_MODULE_NOT_FOUND, msg...)
}

func ErrSaasDeviceBindFail(msg ...string) *ecode.ErrorV2 {
	return makeNewErr(ERR_SAAS_DEVICE_BIND_FAIL, msg...)
}

func ErrProductAppFoundFail(msg ...string) *ecode.ErrorV2 {
	return makeNewErr(ERR_PRODUCT_APP_FOUND_FAIL, msg...)
}

func ErrRegisterDeviceLimit(msg ...string) *ecode.ErrorV2 {
	return makeNewErr(ERR_REGISTER_DEVICE_LIMIT, msg...)
}

func ErrRegisterDeviceSingleFaile(msg ...string) *ecode.ErrorV2 {
	return makeNewErr(ERR_REGISTER_DEVICE_SINGLE_FAILE, msg...)
}

func ErrRegisterDeviceBatchFaile(msg ...string) *ecode.ErrorV2 {
	return makeNewErr(ERR_REGISTER_DEVICE_BATCH_FAILE, msg...)
}

func IsUnknowError(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == UNKNOW_ERROR
	}
	return false
}

func IsUnknowErrorDEEP(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == UNKNOW_ERROR && se.Message() == ErrMap[UNKNOW_ERROR]
	}
	return false
}

func IsErrProductNotFound(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_PRODUCT_NOT_FOUND
	}
	return false
}

func IsErrProductNotFoundDEEP(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_PRODUCT_NOT_FOUND && se.Message() == ErrMap[ERR_PRODUCT_NOT_FOUND]
	}
	return false
}

func IsErrProductModuleNotFound(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_PRODUCT_MODULE_NOT_FOUND
	}
	return false
}

func IsErrProductModuleNotFoundDEEP(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_PRODUCT_MODULE_NOT_FOUND && se.Message() == ErrMap[ERR_PRODUCT_MODULE_NOT_FOUND]
	}
	return false
}

func IsErrSaasDeviceBindFail(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_SAAS_DEVICE_BIND_FAIL
	}
	return false
}

func IsErrSaasDeviceBindFailDEEP(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_SAAS_DEVICE_BIND_FAIL && se.Message() == ErrMap[ERR_SAAS_DEVICE_BIND_FAIL]
	}
	return false
}

func IsErrProductAppFoundFail(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_PRODUCT_APP_FOUND_FAIL
	}
	return false
}

func IsErrProductAppFoundFailDEEP(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_PRODUCT_APP_FOUND_FAIL && se.Message() == ErrMap[ERR_PRODUCT_APP_FOUND_FAIL]
	}
	return false
}

func IsErrRegisterDeviceLimit(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_REGISTER_DEVICE_LIMIT
	}
	return false
}

func IsErrRegisterDeviceLimitDEEP(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_REGISTER_DEVICE_LIMIT && se.Message() == ErrMap[ERR_REGISTER_DEVICE_LIMIT]
	}
	return false
}

func IsErrRegisterDeviceSingleFaile(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_REGISTER_DEVICE_SINGLE_FAILE
	}
	return false
}

func IsErrRegisterDeviceSingleFaileDEEP(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_REGISTER_DEVICE_SINGLE_FAILE && se.Message() == ErrMap[ERR_REGISTER_DEVICE_SINGLE_FAILE]
	}
	return false
}

func IsErrRegisterDeviceBatchFaile(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_REGISTER_DEVICE_BATCH_FAILE
	}
	return false
}

func IsErrRegisterDeviceBatchFaileDEEP(err error) bool {
	if se := new(ecode.ErrorV2); errors.As(err, &se) {
		return se.Code() == ERR_REGISTER_DEVICE_BATCH_FAILE && se.Message() == ErrMap[ERR_REGISTER_DEVICE_BATCH_FAILE]
	}
	return false
}
