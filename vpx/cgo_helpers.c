// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sun, 25 Sep 2016 18:05:40 MSK.
// By https://git.io/cgogen. DO NOT EDIT.

#include "_cgo_export.h"
#include "cgo_helpers.h"

void vpx_codec_enc_output_cx_pkt_cb_fn_t_947769c1(vpx_codec_cx_pkt_t* pkt, void* user_data) {
	codecEncOutputCxPktCbFn947769C1(pkt, user_data);
}

void vpx_codec_put_frame_cb_fn_t_f425e4fc(void* user_priv, vpx_image_t* img) {
	codecPutFrameCbFnF425E4FC(user_priv, img);
}

void vpx_codec_put_slice_cb_fn_t_1c2e1e21(void* user_priv, vpx_image_t* img, vpx_image_rect_t* valid, vpx_image_rect_t* update) {
	codecPutSliceCbFn1C2E1E21(user_priv, img, valid, update);
}

int vpx_get_frame_buffer_cb_fn_t_f5fe44eb(void* priv, unsigned long int min_size, vpx_codec_frame_buffer_t* fb) {
	return getFrameBufferCbFnF5FE44EB(priv, min_size, fb);
}

int vpx_release_frame_buffer_cb_fn_t_ed49b503(void* priv, vpx_codec_frame_buffer_t* fb) {
	return releaseFrameBufferCbFnED49B503(priv, fb);
}

