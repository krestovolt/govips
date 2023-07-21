package vips

// #include "foreign.h"
import "C"

func maybeSetBoolParam(p BoolParameter, cp *C.Param) {
	if p.IsSet() {
		cp._type = C.PARAM_TYPE_BOOL
		if p.Get() {
			cp.b = C.TRUE
		} else {
			cp.b = C.FALSE
		}
		cp.is_set = C.TRUE
	}
}

func maybeSetIntParam(p IntParameter, cp *C.Param) {
	if p.IsSet() {
		cp._type = C.PARAM_TYPE_INT
		cp.i = C.int(p.Get())
		cp.is_set = C.TRUE
	}
}

func maybeSetVipAccess(p IntParameter, cp *C.struct_LoadParams) {
	if p.IsSet() {
		cp.access = C.VipsAccess(p.Get())
	}
}

func maybeSetDouble(p IntParameter, cp *C.Param) {
	if p.IsSet() {
		cp._type = C.PARAM_TYPE_DOUBLE
		cp.d = C.gdouble(p.Get())
		cp.is_set = C.TRUE
	}
}

func createImportParams(format ImageType, params ImportParams) C.LoadParams {
	defaultParam := C.struct_Param{}
	p := C.struct_LoadParams{
		access:        C.VipsAccess(VipsAccessRandom),
		inputFormat:   C.ImageType(format),
		inputBlob:     nil,
		outputImage:   nil,
		autorotate:    defaultParam,
		fail:          defaultParam,
		page:          defaultParam,
		n:             defaultParam,
		dpi:           defaultParam,
		jpegShrink:    defaultParam,
		heifThumbnail: defaultParam,
		svgUnlimited:  defaultParam,
	}

	maybeSetVipAccess(params.AccessMode, &p)
	maybeSetBoolParam(params.AutoRotate, &p.autorotate)
	maybeSetBoolParam(params.FailOnError, &p.fail)
	maybeSetIntParam(params.Page, &p.page)
	maybeSetIntParam(params.NumPages, &p.n)
	maybeSetIntParam(params.JpegShrinkFactor, &p.jpegShrink)
	maybeSetBoolParam(params.HeifThumbnail, &p.heifThumbnail)
	maybeSetBoolParam(params.SvgUnlimited, &p.svgUnlimited)
	maybeSetDouble(params.Density, &p.dpi)

	return p
}

// TODO: Change to same pattern as ImportParams
func createSaveParams(outputFormat C.ImageType) C.struct_SaveParams {
	var defaultSaveParams = C.struct_SaveParams{
		inputImage:             nil,
		outputBuffer:           nil,
		webpIccProfile:         nil,
		outputFormat:           outputFormat,
		outputLen:              0,
		interlace:              C.FALSE,
		quality:                0,
		stripMetadata:          C.FALSE,
		jpegOptimizeCoding:     C.FALSE,
		jpegSubsample:          C.VipsForeignJpegSubsample(VipsForeignSubsampleOn),
		jpegTrellisQuant:       C.FALSE,
		jpegOvershootDeringing: C.FALSE,
		jpegOptimizeScans:      C.FALSE,
		jpegQuantTable:         0,
		pngCompression:         6,
		pngPalette:             C.FALSE,
		pngBitdepth:            0,
		pngDither:              0,
		pngFilter:              C.VipsForeignPngFilter(PngFilterNone),
		gifDither:              0.0,
		gifEffort:              0,
		gifBitdepth:            0,
		webpLossless:           C.FALSE,
		webpNearLossless:       C.FALSE,
		webpReductionEffort:    4,
		heifBitdepth:           8,
		heifLossless:           C.FALSE,
		heifEffort:             5,
		tiffCompression:        C.VipsForeignTiffCompression(TiffCompressionLzw),
		tiffPredictor:          C.VipsForeignTiffPredictor(TiffPredictorHorizontal),
		tiffPyramid:            C.FALSE,
		tiffTile:               C.FALSE,
		tiffTileHeight:         256,
		tiffTileWidth:          256,
		jp2kLossless:           C.FALSE,
		jp2kTileHeight:         512,
		jp2kTileWidth:          512,
	}

	return defaultSaveParams
}
