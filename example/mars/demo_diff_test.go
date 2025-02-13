package mars

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/mars"
)

const (
	testAk       = "your ak"
	testSk       = "your sk"
	appId  int64 = 12345
)

func defaultInstance() *mars.Diff {
	instance := mars.DefaultInstance

	// insert necessary params
	instance.Client.SetAccessKey(testAk)
	instance.Client.SetSecretKey(testSk)
	return instance
}

func printResult(status int, err error, response interface{}) {
	// print httpcode (and err)
	fmt.Println("status:", status, "error:", err)

	// print response
	b, _ := json.Marshal(response)
	fmt.Println("json string result:", string(b))
}

func TestCreateService(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &mars.CreateServiceReq{
		Name:    "your service name",
		Comment: "your service comment",
	}

	// do http query
	response, status, err := instance.CreateService(request, appId)

	// printResult
	printResult(status, err, response)
}

func TestDeletePackages(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &mars.DeletePackagesReq{
		ServiceId:   12345,
		OldVersions: []string{"your old package version1", "your old package version2"},
	}

	// do http query
	response, status, err := instance.DeletePackages(request, appId)

	// printResult
	printResult(status, err, response)
}

func TestDeleteService(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &mars.DeleteServiceReq{
		ServiceId: 12345,
	}

	// do http query
	response, status, err := instance.DeleteService(request, appId)

	// printResult
	printResult(status, err, response)
}

func TestQueryPatchByService(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &mars.QueryPatchByServiceReq{
		ServiceId:    12345,
		StartTime:    time.Now().UnixMilli(),
		EndTime:      time.Now().UnixMilli(),
		StartVersion: "your version",
		EndVersion:   "your version",
		NoPatches:    false,
	}

	// do http query
	response, status, err := instance.QueryPatchByService(request, appId)

	// printResult
	printResult(status, err, response)
}

func TestValidate(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &mars.ValidateReq{
		OldUrl:   "your old package url",
		PatchUrl: "your patch url",
	}

	// do http query
	response, status, err := instance.Validate(request)

	// printResult
	printResult(status, err, response)
}

func TestGenByPkg(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &mars.GenByPkgReq{
		ServiceId:    12345,
		Alg:          "wp",
		OldUrl:       "your old package url",
		OldVersion:   "1.0.0",
		OldExtraInfo: "your old package info",
		NewUrl:       "your new pacakge url",
		NewVersion:   "1.1.0",
		NewExtraInfo: "your new package info",
	}

	// do http query
	response, status, err := instance.GenByPkg(request, appId)

	// printResult
	printResult(status, err, response)
}

func TestGenByCount(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &mars.GenByCountReq{
		ServiceId:    12345,
		Alg:          "wp",
		NewUrl:       "your new pacakge url",
		NewVersion:   "1.1.0",
		NewExtraInfo: "your new package info",
		Count:        1,
	}

	// do http query
	response, status, err := instance.GenByCount(request, appId)

	// printResult
	printResult(status, err, response)
}

func TestGenByVersion(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &mars.GenByVersionReq{
		ServiceId:    12345,
		Alg:          "wp",
		NewUrl:       "your new pacakge url",
		NewVersion:   "1.1.0",
		NewExtraInfo: "your new package info",
		OldVersions:  []string{"1.0.0"},
	}

	// do http query
	response, status, err := instance.GenByVersion(request, appId)

	// printResult
	printResult(status, err, response)
}

func TestCheckResponse(t *testing.T) {
	instance := defaultInstance()

	// create request
	request := &mars.CheckResponseReq{
		TaskId: 12345,
	}

	// do http query
	response, result, status, err := instance.CheckResponse(request)

	// printResult
	printResult(status, err, response)

	if result == nil {
		return
	}

	// 打印实际的任务结果
	switch response.Data.Api {
	case 0:
		if actualResult, ok := result.(*mars.GenResult); ok {
			b, _ := json.Marshal(actualResult)
			fmt.Println("actual gen patch data:", string(b))
		}
	case 1:
		if actualResult, ok := result.(*mars.ValidateResult); ok {
			b, _ := json.Marshal(actualResult)
			fmt.Println("actual validate data:", string(b))
		}
	}
}
