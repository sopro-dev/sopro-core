package fileformat

// SupportedAudioFormats is a list of supported audio formats
var SUPPORTED_AUDIO_FORMATS = []string{
	// Add here the supported audio formats
	AUDIO_3GP,
	AUDIO_AA,
	AUDIO_AAC,
	AUDIO_AAX,
	AUDIO_ACT,
	AUDIO_AIFF,
	AUDIO_ALAC,
	AUDIO_AMR,
	AUDIO_APE,
	AUDIO_AU,
	AUDIO_AWB,
	AUDIO_DSS,
	AUDIO_FLAC,
	AUDIO_GSM,
	AUDIO_M4A,
	AUDIO_M4B,
	AUDIO_M4P,
	AUDIO_MP3,
	AUDIO_MPC,
	AUDIO_OGG,
	AUDIO_OGA,
	AUDIO_MOGG,
	AUDIO_MULAW,
	AUDIO_OPUS,
	AUDIO_PCM,
	AUDIO_RA,
	AUDIO_RM,
	AUDIO_RAW,
	AUDIO_RF64,
	AUDIO_SLN,
	AUDIO_TTA,
	AUDIO_VOC,
	AUDIO_VOX,
	AUDIO_WAV,
	AUDIO_WMA,
	AUDIO_WEBM,
	AUDIO_8SVX,
	AUDIO_CAF,
	AUDIO_DVF,
	AUDIO_IKLAX,
	AUDIO_IVS,
	AUDIO_MMF,
	AUDIO_MSV,
	AUDIO_NMF,
	AUDIO_SPEEX,
	AUDIO_VORBIS,
}

// Audio formats
const (
	AUDIO_3GP   = "3gp"
	AUDIO_AA    = "aa"
	AUDIO_AAC   = "aac"
	AUDIO_AAX   = "aax"
	AUDIO_ACT   = "act"
	AUDIO_AIFF  = "aiff"
	AUDIO_ALAC  = "alac"
	AUDIO_AMR   = "amr"
	AUDIO_APE   = "ape"
	AUDIO_AU    = "au"
	AUDIO_AWB   = "awb"
	AUDIO_DSS   = "dss"
	AUDIO_FLAC  = "flac"
	AUDIO_GSM   = "gsm"
	AUDIO_M4A   = "m4a"
	AUDIO_M4B   = "m4b"
	AUDIO_M4P   = "m4p"
	AUDIO_MP3   = "mp3"
	AUDIO_MPC   = "mpc"
	AUDIO_OGG   = "ogg"
	AUDIO_OGA   = "oga"
	AUDIO_MOGG  = "mogg"
	AUDIO_MULAW = "mulaw"
	AUDIO_OPUS  = "opus"
	AUDIO_PCM   = "pcm"
	AUDIO_RA    = "ra"
	AUDIO_RM    = "rm"
	AUDIO_RAW   = "raw"
	AUDIO_RF64  = "rf64"
	AUDIO_SLN   = "sln"
	AUDIO_TTA   = "tta"
	AUDIO_VOC   = "voc"
	AUDIO_VOX   = "vox"
	AUDIO_WAV   = "wav"
	AUDIO_WMA   = "wma"
	AUDIO_WEBM  = "webm"
	AUDIO_8SVX  = "8svx"
	AUDIO_CAF   = "cda"
)

// DEPRECATED
const (
	AUDIO_DVF    = "dvf"
	AUDIO_IKLAX  = "iklax"
	AUDIO_IVS    = "ivs"
	AUDIO_MMF    = "mmf"
	AUDIO_MSV    = "msv"
	AUDIO_NMF    = "nmf"
	AUDIO_SPEEX  = "speex"
	AUDIO_VORBIS = "vorbis"
)

// PROPRIETARY CODECS (PAY ME AND I WILL ADD THEM)
// const AUDIO_G722 = "g722"
// const AUDIO_G726 = "g726"
// const AUDIO_G729 = "g729"
