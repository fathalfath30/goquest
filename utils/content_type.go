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

package utils

type (
	ContentType struct {
		contentType string
	}
)

const (
	ContentTypeAppEDIX12       = "application/EDI-X12"
	ContentTypeAppEDIFAC       = "application/EDIFACT"
	ContentTypeAppJavascript   = "application/javascript"
	ContentTypeAppOctetStream  = "application/octet-stream"
	ContentTypeAppOgg          = "application/ogg"
	ContentTypeAppPdf          = "application/pdf"
	ContentTypeXhtmlXml        = "application/xhtml+xml"
	ContentTypeXShockwaveFlash = "application/x-shockwave-flash"
	ContentTypeAppJson         = "application/json"
	ContentTypeAppLdJson       = "application/ld+json"
	ContentTypeAppXml          = "application/xml"
	ContentTypeAppZip          = "application/zip"
	ContentTypeAppUrlEncoded   = "application/x-www-form-urlencoded"

	ContentTypeAudioMpeg         = "audio/mpeg"
	ContentTypeAudioWma          = "audio/x-ms-wma"
	ContentTypeAudioVndRealAudio = "audio/vnd.rn-realaudio"
	ContentTypeAudioWav          = "audio/x-wav"

	ContentTypeImageGif              = "image/gif"
	ContentTypeImageJpeg             = "image/jpeg"
	ContentTypeImagePng              = "image/png"
	ContentTypeImageTiff             = "image/tiff"
	ContentTypeImageVndMicrosoftIcon = "image/vnd.microsoft.icon"
	ContentTypeImageXIcon            = "image/x-icon"
	ContentTypeImageVndDjvu          = "image/vnd.djvu"
	ContentTypeImageXvgXml           = "image/svg+xml"

	ContentTypeMultipartMixed       = "multipart/mixed"
	ContentTypeMultipartAlternative = "multipart/alternative"
	ContentTypeMultipartRelated     = "multipart/related"
	ContentTypeMultipartFormData    = "multipart/form-data"

	ContentTypeTextCss        = "text/css"
	ContentTypeTextCsv        = "text/csv"
	ContentTypeTextHtml       = "text/html"
	ContentTypeTextJavascript = "text/javascript"
	ContentTypeTextPlain      = "text/plain"
	ContentTypeTextXml        = "text/xml"

	ContentTypeVideoMpeg      = "video/mpeg"
	ContentTypeVideoMp4       = "video/mp4"
	ContentTypeVideoQuickTime = "video/quicktime"
	ContentTypeVideoXMsWmv    = "video/x-ms-wmv"
	ContentTypeVideoXMsVideo  = "video/x-msvideo"
	ContentTypeVideoFlv       = "video/x-flv"
	ContentTypeVideoWebm      = "video/webm"

	ContentTypeVndOpenDocText           = "application/vnd.oasis.opendocument.text"
	ContentTypeVndOpenDocSpreadsheet    = "application/vnd.oasis.opendocument.spreadsheet"
	ContentTypeVndOpenDocPresentation   = "application/vnd.oasis.opendocument.presentation"
	ContentTypeVndOpenDocGraphic        = "application/vnd.oasis.opendocument.graphics"
	ContentTypeVndMsExcel               = "application/vnd.ms-excel"
	ContentTypeVndOpenXmlFormat         = "application/vnd.openxmlformats-"
	ContentTypeVndOfficeDocSpreadsheet  = "officedocument.spreadsheetml.sheet"
	ContentTypeVndMsPowerPoint          = "application/vnd.ms-powerpoint"
	ContentTypeVndOfficeDocPresentation = "officedocument.presentationml.presentation"
	ContentTypeVndMsWord                = "application/msword"
	ContentTypeVndWordProcessing        = "officedocument.wordprocessingml.document"
	ContentTypeVndMozillaXulXml         = "application/vnd.mozilla.xul+xml"
)

func NewContentType(value string) ContentType {
	return ContentType{
		contentType: value,
	}
}

func (ct ContentType) IsApplication() bool {
	switch ct.contentType {
	case ContentTypeAppEDIX12, ContentTypeAppEDIFAC, ContentTypeAppJavascript,
		ContentTypeAppOctetStream, ContentTypeAppOgg, ContentTypeAppPdf,
		ContentTypeXhtmlXml, ContentTypeXShockwaveFlash, ContentTypeAppJson,
		ContentTypeAppLdJson, ContentTypeAppXml, ContentTypeAppZip,
		ContentTypeAppUrlEncoded:
		return true
	default:
		return false
	}
}

func (ct ContentType) IsAudio() bool {
	switch ct.contentType {
	case
		ContentTypeAudioMpeg,
		ContentTypeAudioWma,
		ContentTypeAudioVndRealAudio,
		ContentTypeAudioWav:
		return true
	default:
		return false
	}
}

func (ct ContentType) IsImage() bool {
	switch ct.contentType {
	case
		ContentTypeImageGif,
		ContentTypeImageJpeg,
		ContentTypeImagePng,
		ContentTypeImageTiff,
		ContentTypeImageVndMicrosoftIcon,
		ContentTypeImageXIcon,
		ContentTypeImageVndDjvu,
		ContentTypeImageXvgXml:
		return true
	default:
		return false
	}
}

func (ct ContentType) IsMultipart() bool {
	switch ct.contentType {
	case
		ContentTypeMultipartMixed,
		ContentTypeMultipartAlternative,
		ContentTypeMultipartRelated,
		ContentTypeMultipartFormData:
		return true
	default:
		return false
	}
}

func (ct ContentType) IsText() bool {
	switch ct.contentType {
	case
		ContentTypeTextCss,
		ContentTypeTextCsv,
		ContentTypeTextHtml,
		ContentTypeTextJavascript,
		ContentTypeTextPlain,
		ContentTypeTextXml:
		return true
	default:
		return false
	}
}

func (ct ContentType) IsVideo() bool {
	switch ct.contentType {
	case
		ContentTypeVideoMpeg,
		ContentTypeVideoMp4,
		ContentTypeVideoQuickTime,
		ContentTypeVideoXMsWmv,
		ContentTypeVideoXMsVideo,
		ContentTypeVideoFlv,
		ContentTypeVideoWebm:
		return true
	default:
		return false
	}
}

func (ct ContentType) IsVnd() bool {
	switch ct.contentType {
	case
		ContentTypeVndOpenDocText,
		ContentTypeVndOpenDocSpreadsheet,
		ContentTypeVndOpenDocPresentation,
		ContentTypeVndOpenDocGraphic,
		ContentTypeVndMsExcel,
		ContentTypeVndOpenXmlFormat,
		ContentTypeVndOfficeDocSpreadsheet,
		ContentTypeVndMsPowerPoint,
		ContentTypeVndOfficeDocPresentation,
		ContentTypeVndMsWord,
		ContentTypeVndWordProcessing,
		ContentTypeVndMozillaXulXml:
		return true
	default:
		return false
	}
}
