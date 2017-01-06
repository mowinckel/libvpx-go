// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Fri, 06 Jan 2017 20:09:57 MSK.
// By https://git.io/cgogen. DO NOT EDIT.

package vpx

/*
#cgo pkg-config: vpx
#include <vpx/vpx_encoder.h>
#include <vpx/vpx_decoder.h>
#include <vpx/vp8.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// FixedBuf as declared in vpx-1.6.0/vpx_encoder.h:110
type FixedBuf struct {
	Buf            unsafe.Pointer
	Sz             uint
	refeac28dc0    *C.vpx_fixed_buf_t
	allocseac28dc0 interface{}
}

// CodecPts type as declared in vpx-1.6.0/vpx_encoder.h:118
type CodecPts int64

// CodecFrameFlags type as declared in vpx-1.6.0/vpx_encoder.h:128
type CodecFrameFlags uint32

// CodecErFlags type as declared in vpx-1.6.0/vpx_encoder.h:144
type CodecErFlags uint32

// CodecCxPkt as declared in vpx-1.6.0/vpx_encoder.h:223
type CodecCxPkt struct {
	Kind           CodecCxPktKind
	refa671fc83    *C.vpx_codec_cx_pkt_t
	allocsa671fc83 interface{}
}

// CodecEncOutputCxPktCbFn type as declared in vpx-1.6.0/vpx_encoder.h:233
type CodecEncOutputCxPktCbFn func(pkt *CodecCxPkt, userData unsafe.Pointer)

// CodecPrivOutputCxPktCbPair as declared in vpx-1.6.0/vpx_encoder.h:240
type CodecPrivOutputCxPktCbPair struct {
	OutputCxPkt    CodecEncOutputCxPktCbFn
	UserPriv       unsafe.Pointer
	ref5727a29d    *C.vpx_codec_priv_output_cx_pkt_cb_pair_t
	allocs5727a29d interface{}
}

// Rational as declared in vpx-1.6.0/vpx_encoder.h:249
type Rational struct {
	Num            int32
	Den            int32
	ref48ce5779    *C.vpx_rational_t
	allocs48ce5779 interface{}
}

// EncFrameFlags type as declared in vpx-1.6.0/vpx_encoder.h:291
type EncFrameFlags int

// CodecEncCfg as declared in vpx-1.6.0/vpx_encoder.h:751
type CodecEncCfg struct {
	GUsage                  uint32
	GThreads                uint32
	GProfile                uint32
	GW                      uint32
	GH                      uint32
	GBitDepth               BitDepth
	GInputBitDepth          uint32
	GTimebase               Rational
	GErrorResilient         CodecErFlags
	GPass                   EncPass
	GLagInFrames            uint32
	RcDropframeThresh       uint32
	RcResizeAllowed         uint32
	RcScaledWidth           uint32
	RcScaledHeight          uint32
	RcResizeUpThresh        uint32
	RcResizeDownThresh      uint32
	RcEndUsage              RcMode
	RcTwopassStatsIn        FixedBuf
	RcFirstpassMbStatsIn    FixedBuf
	RcTargetBitrate         uint32
	RcMinQuantizer          uint32
	RcMaxQuantizer          uint32
	RcUndershootPct         uint32
	RcOvershootPct          uint32
	RcBufSz                 uint32
	RcBufInitialSz          uint32
	RcBufOptimalSz          uint32
	Rc2passVbrBiasPct       uint32
	Rc2passVbrMinsectionPct uint32
	Rc2passVbrMaxsectionPct uint32
	KfMode                  KfMode
	KfMinDist               uint32
	KfMaxDist               uint32
	SsNumberLayers          uint32
	SsEnableAutoAltRef      [5]int32
	SsTargetBitrate         [5]uint32
	TsNumberLayers          uint32
	TsTargetBitrate         [5]uint32
	TsRateDecimator         [5]uint32
	TsPeriodicity           uint32
	TsLayerID               [16]uint32
	LayerTargetBitrate      [12]uint32
	TemporalLayeringMode    int32
	ref37e25db9             *C.vpx_codec_enc_cfg_t
	allocs37e25db9          interface{}
}

// SvcExtraCfg as declared in vpx-1.6.0/vpx_encoder.h:764
type SvcExtraCfg C.vpx_svc_extra_cfg_t

// CodecCaps type as declared in vpx-1.6.0/vpx_codec.h:153
type CodecCaps int

// CodecFlags type as declared in vpx-1.6.0/vpx_codec.h:165
type CodecFlags int

// CodecIface as declared in vpx-1.6.0/vpx_codec.h:173
type CodecIface C.vpx_codec_iface_t

// CodecPriv as declared in vpx-1.6.0/vpx_codec.h:181
type CodecPriv C.vpx_codec_priv_t

// CodecIter type as declared in vpx-1.6.0/vpx_codec.h:188
type CodecIter unsafe.Pointer

// CodecCtx as declared in vpx-1.6.0/vpx_codec.h:213
type CodecCtx C.vpx_codec_ctx_t

// Image as declared in vpx-1.6.0/vpx_image.h:133
type Image struct {
	Fmt            ImageFormat
	Cs             ColorSpace
	Range          ColorRange
	W              uint32
	H              uint32
	BitDepth       uint32
	DW             uint32
	DH             uint32
	RW             uint32
	RH             uint32
	XChromaShift   uint32
	YChromaShift   uint32
	Planes         [4]*byte
	Stride         [4]int32
	Bps            int32
	UserPriv       unsafe.Pointer
	ImgData        []byte
	ImgDataOwner   int32
	SelfAllocd     int32
	FbPriv         unsafe.Pointer
	refc09455e3    *C.vpx_image_t
	allocsc09455e3 interface{}
}

// ImageRect as declared in vpx-1.6.0/vpx_image.h:141
type ImageRect struct {
	X              uint32
	Y              uint32
	W              uint32
	H              uint32
	reff3ce051f    *C.vpx_image_rect_t
	allocsf3ce051f interface{}
}

// CodecStreamInfo as declared in vpx-1.6.0/vpx_decoder.h:93
type CodecStreamInfo struct {
	Sz             uint32
	W              uint32
	H              uint32
	IsKf           uint32
	ref342546e4    *C.vpx_codec_stream_info_t
	allocs342546e4 interface{}
}

// CodecDecCfg as declared in vpx-1.6.0/vpx_decoder.h:111
type CodecDecCfg struct {
	Threads        uint32
	W              uint32
	H              uint32
	ref7df355ac    *C.vpx_codec_dec_cfg_t
	allocs7df355ac interface{}
}

// CodecPutFrameCbFn type as declared in vpx-1.6.0/vpx_decoder.h:260
type CodecPutFrameCbFn func(userPriv unsafe.Pointer, img *Image)

// CodecPutSliceCbFn type as declared in vpx-1.6.0/vpx_decoder.h:300
type CodecPutSliceCbFn func(userPriv unsafe.Pointer, img *Image, valid *ImageRect, update *ImageRect)

// CodecFrameBuffer as declared in vpx-1.6.0/vpx_frame_buffer.h:43
type CodecFrameBuffer struct {
	Data           []byte
	Size           uint
	Priv           unsafe.Pointer
	refd319b8f1    *C.vpx_codec_frame_buffer_t
	allocsd319b8f1 interface{}
}

// GetFrameBufferCbFn type as declared in vpx-1.6.0/vpx_frame_buffer.h:63
type GetFrameBufferCbFn func(priv unsafe.Pointer, minSize uint, fb *CodecFrameBuffer) int32

// ReleaseFrameBufferCbFn type as declared in vpx-1.6.0/vpx_frame_buffer.h:76
type ReleaseFrameBufferCbFn func(priv unsafe.Pointer, fb *CodecFrameBuffer) int32
