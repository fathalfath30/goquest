/*
//
//  ______    _   _           _  __      _   _     ____   ___
// |  ____|  | | | |         | |/ _|    | | | |   |___ \ / _ \
// | |__ __ _| |_| |__   __ _| | |_ __ _| |_| |__   __) | | | |
// |  __/ _` | __| '_ \ / _` | |  _/ _` | __| '_ \ |__ <| | | |
// | | | (_| | |_| | | | (_| | | || (_| | |_| | | |___) | |_| |
// |_|  \__,_|\__|_| |_|\__,_|_|_| \__,_|\__|_| |_|____/ \___/
//
// Written by Fathalfath30.
// Email : fathalfath30@gmail.com
// Follow me on:
//  Github : https://github.com/fathalfath30
//  Gitlab : https://gitlab.com/Fathalfath30
//
*/

package utils_test

import "github.com/fathalfath30/goquest/utils"

func (ts *UtilsTestSuite) Test_ContentType_IsApplication() {
	ts.Run("should return true", func() {
		contentTypes := []string{
			utils.ContentTypeAppEDIX12, utils.ContentTypeAppEDIFAC, utils.ContentTypeAppJavascript,
			utils.ContentTypeAppOctetStream, utils.ContentTypeAppOgg, utils.ContentTypeAppPdf,
			utils.ContentTypeXhtmlXml, utils.ContentTypeXShockwaveFlash, utils.ContentTypeAppJson,
			utils.ContentTypeAppLdJson, utils.ContentTypeAppXml, utils.ContentTypeAppZip,
			utils.ContentTypeAppUrlEncoded,
		}

		for _, v := range contentTypes {
			ts.Require().True(utils.NewContentType(v).IsApplication())
		}
	})
	ts.Run("should return false", func() {
		contentTypes := []string{
			utils.ContentTypeAudioMpeg,
			utils.ContentTypeImageGif,
			utils.ContentTypeMultipartMixed,
			utils.ContentTypeTextCss,
			utils.ContentTypeVideoMpeg,
			utils.ContentTypeVndOpenDocText,
		}

		for _, v := range contentTypes {
			ts.Require().False(utils.NewContentType(v).IsApplication())
		}
	})
}
func (ts *UtilsTestSuite) Test_ContentType_IsAudio() {
	ts.Run("should return true", func() {
		contentTypes := []string{
			utils.ContentTypeAudioMpeg,
			utils.ContentTypeAudioWma,
			utils.ContentTypeAudioVndRealAudio,
			utils.ContentTypeAudioWav,
		}

		for _, v := range contentTypes {
			ts.Require().True(utils.NewContentType(v).IsAudio())
		}
	})
	ts.Run("should return false", func() {
		contentTypes := []string{
			utils.ContentTypeAppEDIX12,
			utils.ContentTypeImageGif,
			utils.ContentTypeMultipartMixed,
			utils.ContentTypeTextCss,
			utils.ContentTypeVideoMpeg,
			utils.ContentTypeVndOpenDocText,
		}

		for _, v := range contentTypes {
			ts.Require().False(utils.NewContentType(v).IsAudio())
		}
	})
}
func (ts *UtilsTestSuite) Test_ContentType_IsImage() {
	ts.Run("should return true", func() {
		contentTypes := []string{
			utils.ContentTypeImageGif,
			utils.ContentTypeImageJpeg,
			utils.ContentTypeImagePng,
			utils.ContentTypeImageTiff,
			utils.ContentTypeImageVndMicrosoftIcon,
			utils.ContentTypeImageXIcon,
			utils.ContentTypeImageVndDjvu,
			utils.ContentTypeImageXvgXml,
		}

		for _, v := range contentTypes {
			ts.Require().True(utils.NewContentType(v).IsImage())
		}
	})
	ts.Run("should return false", func() {
		contentTypes := []string{
			utils.ContentTypeAppEDIX12,
			utils.ContentTypeAudioMpeg,
			utils.ContentTypeMultipartMixed,
			utils.ContentTypeTextCss,
			utils.ContentTypeVideoMpeg,
			utils.ContentTypeVndOpenDocText,
		}

		for _, v := range contentTypes {
			ts.Require().False(utils.NewContentType(v).IsImage())
		}
	})
}
func (ts *UtilsTestSuite) Test_ContentType_IsMultipart() {
	ts.Run("should return true", func() {
		contentTypes := []string{
			utils.ContentTypeMultipartMixed,
			utils.ContentTypeMultipartAlternative,
			utils.ContentTypeMultipartRelated,
			utils.ContentTypeMultipartFormData,
		}

		for _, v := range contentTypes {
			ts.Require().True(utils.NewContentType(v).IsMultipart())
		}
	})
	ts.Run("should return false", func() {
		contentTypes := []string{
			utils.ContentTypeAppEDIX12,
			utils.ContentTypeAudioMpeg,
			utils.ContentTypeImageGif,
			utils.ContentTypeTextCss,
			utils.ContentTypeVideoMpeg,
			utils.ContentTypeVndOpenDocText,
		}

		for _, v := range contentTypes {
			ts.Require().False(utils.NewContentType(v).IsMultipart())
		}
	})
}
func (ts *UtilsTestSuite) Test_ContentType_IsText() {
	ts.Run("should return true", func() {
		contentTypes := []string{
			utils.ContentTypeTextCss,
			utils.ContentTypeTextCsv,
			utils.ContentTypeTextHtml,
			utils.ContentTypeTextJavascript,
			utils.ContentTypeTextPlain,
			utils.ContentTypeTextXml,
		}

		for _, v := range contentTypes {
			ts.Require().True(utils.NewContentType(v).IsText())
		}
	})
	ts.Run("should return false", func() {
		contentTypes := []string{
			utils.ContentTypeAppEDIX12,
			utils.ContentTypeAudioMpeg,
			utils.ContentTypeImageGif,
			utils.ContentTypeMultipartMixed,
			utils.ContentTypeVideoMpeg,
			utils.ContentTypeVndOpenDocText,
		}

		for _, v := range contentTypes {
			ts.Require().False(utils.NewContentType(v).IsText())
		}
	})
}
func (ts *UtilsTestSuite) Test_ContentType_IsVideo() {
	ts.Run("should return true", func() {
		contentTypes := []string{
			utils.ContentTypeVideoMpeg,
			utils.ContentTypeVideoMp4,
			utils.ContentTypeVideoQuickTime,
			utils.ContentTypeVideoXMsWmv,
			utils.ContentTypeVideoXMsVideo,
			utils.ContentTypeVideoFlv,
			utils.ContentTypeVideoWebm,
		}

		for _, v := range contentTypes {
			ts.Require().True(utils.NewContentType(v).IsVideo())
		}
	})
	ts.Run("should return false", func() {
		contentTypes := []string{
			utils.ContentTypeAppEDIX12,
			utils.ContentTypeAudioMpeg,
			utils.ContentTypeImageGif,
			utils.ContentTypeMultipartMixed,
			utils.ContentTypeTextCss,
			utils.ContentTypeVndOpenDocText,
		}

		for _, v := range contentTypes {
			ts.Require().False(utils.NewContentType(v).IsVideo())
		}
	})
}
func (ts *UtilsTestSuite) Test_ContentType_IsVnd() {
	ts.Run("should return true", func() {
		contentTypes := []string{
			utils.ContentTypeVndOpenDocText,
			utils.ContentTypeVndOpenDocSpreadsheet,
			utils.ContentTypeVndOpenDocPresentation,
			utils.ContentTypeVndOpenDocGraphic,
			utils.ContentTypeVndMsExcel,
			utils.ContentTypeVndOpenXmlFormat,
			utils.ContentTypeVndOfficeDocSpreadsheet,
			utils.ContentTypeVndMsPowerPoint,
			utils.ContentTypeVndOfficeDocPresentation,
			utils.ContentTypeVndMsWord,
			utils.ContentTypeVndWordProcessing,
			utils.ContentTypeVndMozillaXulXml,
		}

		for _, v := range contentTypes {
			ts.Require().True(utils.NewContentType(v).IsVnd())
		}
	})
	ts.Run("should return false", func() {
		contentTypes := []string{
			utils.ContentTypeAppEDIX12,
			utils.ContentTypeAudioMpeg,
			utils.ContentTypeImageGif,
			utils.ContentTypeMultipartMixed,
			utils.ContentTypeTextCss,
			utils.ContentTypeVideoMpeg,
		}

		for _, v := range contentTypes {
			ts.Require().False(utils.NewContentType(v).IsVnd())
		}
	})
}
