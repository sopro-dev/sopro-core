package transcoding

import (
	"bytes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/cpuarch"
	"github.com/pablodz/sopro/pkg/encoding"
	"github.com/pablodz/sopro/pkg/fileformat"
	"github.com/pablodz/sopro/pkg/method"
	"github.com/pablodz/sopro/pkg/sopro"
	"github.com/pablodz/sopro/pkg/transcoder"
)

func TranscodeFile(c *fiber.Ctx) error {

	// get the uploaded file from the request
	file, err := c.FormFile("audio")
	if err != nil {
		return err
	}

	// open the uploaded file
	in, err := file.Open()
	if err != nil {
		return err
	}
	defer in.Close()

	// read in as buffer and pass to data
	buf := make([]byte, file.Size)
	if _, err := in.Read(buf); err != nil {
		return err
	}

	// convert to bytes buffer
	data := bytes.NewBuffer(buf)

	// Create the output file
	out, err := os.Create("./internal/samples/output.wav")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// create a transcoder
	t := &transcoder.Transcoder{
		MethodT: method.BIT_LOOKUP_TABLE,
		InConfigs: sopro.AudioConfig{
			Endianness: cpuarch.LITTLE_ENDIAN,
		},
		OutConfigs: sopro.AudioConfig{
			Endianness: cpuarch.LITTLE_ENDIAN,
		},
		SizeBuffer: 1024,
		Verbose:    true,
	}

	// Transcode the file
	err = t.Mulaw2Wav(
		&sopro.In{
			Data: data,
			AudioFileGeneral: sopro.AudioFileGeneral{
				Format: fileformat.AUDIO_MULAW,
				Config: audioconfig.MulawConfig{
					BitDepth:   8,
					Channels:   1,
					Encoding:   encoding.SPACE_LOGARITHMIC, // ulaw is logarithmic
					SampleRate: 8000,
				},
			},
		},
		&sopro.Out{
			Data: out,
			AudioFileGeneral: sopro.AudioFileGeneral{
				Format: fileformat.AUDIO_WAV,
				Config: audioconfig.WavConfig{
					BitDepth:   8,
					Channels:   1,
					Encoding:   encoding.SPACE_LOGARITHMIC,
					SampleRate: 8000,
				},
			},
		},
	)

	if err != nil {
		panic(err)
	}

	// set the response header to indicate that a file is being returned
	c.Set(fiber.HeaderContentType, "audio/mpeg")
	c.Set(fiber.HeaderContentDisposition, "attachment; filename=processed-audio.wav")

	// send the processed audio file to the client
	if err := c.SendFile(out.Name()); err != nil {
		return err
	}

	// return nil to indicate success
	return c.SendStatus(fiber.StatusOK)
}
