// Code generated by protoc-gen-volcengine-sdk
// source: VodMediaService
// DO NOT EDIT!

package vod

import (
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/vod"
	"github.com/volcengine/volc-sdk-golang/service/vod/models/request"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Test_UpdateMediaInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateMediaInfoRequest{
		Vid:              "your Vid",
		PosterUri:        nil,
		Title:            nil,
		Description:      nil,
		Tags:             nil,
		ClassificationId: nil,
	}

	resp, status, err := instance.UpdateMediaInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateMediaPublishStatus(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateMediaPublishStatusRequest{
		Vid:    "your Vid",
		Status: "your Status",
	}

	resp, status, err := instance.UpdateMediaPublishStatus(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetMediaInfos(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetMediaInfosRequest{
		Vids: "your Vids",
	}

	resp, status, err := instance.GetMediaInfos(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetRecommendedPoster(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetRecommendedPosterRequest{
		Vids: "your Vids",
	}

	resp, status, err := instance.GetRecommendedPoster(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DeleteMedia(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDeleteMediaRequest{
		Vids:         "your Vids",
		CallbackArgs: "your CallbackArgs",
	}

	resp, status, err := instance.DeleteMedia(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DeleteTranscodes(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDeleteTranscodesRequest{
		Vid:          "your Vid",
		FileIds:      "your FileIds",
		CallbackArgs: "your CallbackArgs",
	}

	resp, status, err := instance.DeleteTranscodes(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetMediaList(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetMediaListRequest{
		SpaceName:         "your SpaceName",
		Vid:               "your Vid",
		Status:            "your Status",
		Order:             "your Order",
		Tags:              "your Tags",
		StartTime:         "your StartTime",
		EndTime:           "your EndTime",
		Offset:            "your Offset",
		PageSize:          "your PageSize",
		ClassificationIds: "your ClassificationIds",
	}

	resp, status, err := instance.GetMediaList(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetSubtitleInfoList(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetSubtitleInfoListRequest{
		Vid:         "your Vid",
		FileIds:     "your FileIds",
		Languages:   "your Languages",
		Formats:     "your Formats",
		LanguageIds: "your LanguageIds",
		SubtitleIds: "your SubtitleIds",
		Status:      "your Status",
		Title:       "your Title",
		Tag:         "your Tag",
		Offset:      "your Offset",
		PageSize:    "your PageSize",
		Ssl:         "your Ssl",
	}

	resp, status, err := instance.GetSubtitleInfoList(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateSubtitleStatus(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateSubtitleStatusRequest{
		Vid:       "your Vid",
		FileIds:   "your FileIds",
		Languages: "your Languages",
		Formats:   "your Formats",
		Status:    "your Status",
	}

	resp, status, err := instance.UpdateSubtitleStatus(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateSubtitleInfo(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateSubtitleInfoRequest{
		Vid:      "your Vid",
		FileId:   "your FileId",
		Language: "your Language",
		Format:   "your Format",
		Title:    nil,
		Tag:      nil,
	}

	resp, status, err := instance.UpdateSubtitleInfo(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetAuditFramesForAudit(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetAuditFramesForAuditRequest{
		Vid:               "your Vid",
		Strategy:          "your Strategy",
		MinNumberOfFrames: "your MinNumberOfFrames",
		MaxNumberOfFrames: "your MaxNumberOfFrames",
	}

	resp, status, err := instance.GetAuditFramesForAudit(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetMLFramesForAudit(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetMLFramesForAuditRequest{
		Vid:               "your Vid",
		Strategy:          "your Strategy",
		FrameOpt:          "your FrameOpt",
		FrameFps:          "your FrameFps",
		NumberOfFrames:    "your NumberOfFrames",
		CutTimeMills:      "your CutTimeMills",
		NeedFirstFrame:    "your NeedFirstFrame",
		NeedLastFrame:     "your NeedLastFrame",
		StartTimeMill:     "your StartTimeMill",
		EndTimeMill:       "your EndTimeMill",
		MinNumberOfFrames: "your MinNumberOfFrames",
		MaxNumberOfFrames: "your MaxNumberOfFrames",
	}

	resp, status, err := instance.GetMLFramesForAudit(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetBetterFramesForAudit(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetBetterFramesForAuditRequest{
		Vid:       "your Vid",
		Strategy:  "your Strategy",
		CoverRate: "your CoverRate",
	}

	resp, status, err := instance.GetBetterFramesForAudit(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetAudioInfoForAudit(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetAudioInfoForAuditRequest{
		Vid:      "your Vid",
		Strategy: "your Strategy",
	}

	resp, status, err := instance.GetAudioInfoForAudit(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetAutomaticSpeechRecognitionForAudit(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetAutomaticSpeechRecognitionForAuditRequest{
		Vid:      "your Vid",
		Strategy: "your Strategy",
	}

	resp, status, err := instance.GetAutomaticSpeechRecognitionForAudit(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_GetAudioEventDetectionForAudit(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodGetAudioEventDetectionForAuditRequest{
		Vid:      "your Vid",
		Strategy: "your Strategy",
	}

	resp, status, err := instance.GetAudioEventDetectionForAudit(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_CreateVideoClassification(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodCreateVideoClassificationRequest{
		SpaceName:      "your SpaceName",
		Level:          0,
		ParentId:       0,
		Classification: "your Classification",
	}

	resp, status, err := instance.CreateVideoClassification(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_UpdateVideoClassification(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodUpdateVideoClassificationRequest{
		SpaceName:        "your SpaceName",
		ClassificationId: 0,
		Classification:   "your Classification",
	}

	resp, status, err := instance.UpdateVideoClassification(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_DeleteVideoClassification(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodDeleteVideoClassificationRequest{
		SpaceName:        "your SpaceName",
		ClassificationId: 0,
	}

	resp, status, err := instance.DeleteVideoClassification(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListVideoClassifications(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodListVideoClassificationsRequest{
		SpaceName:        "your SpaceName",
		ClassificationId: 0,
	}

	resp, status, err := instance.ListVideoClassifications(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}

func Test_ListSnapshots(t *testing.T) {
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	query := &request.VodListSnapshotsRequest{
		Vid: "your Vid",
	}

	resp, status, err := instance.ListSnapshots(query)
	fmt.Println(status)
	fmt.Println(err)
	fmt.Println(resp.String())
}
