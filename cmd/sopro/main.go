package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pablodz/sopro/pkg/audioconfig"
	"github.com/pablodz/sopro/pkg/cpuarch"
	"github.com/pablodz/sopro/pkg/encoding"
	"github.com/pablodz/sopro/pkg/fileformat"
	"github.com/pablodz/sopro/pkg/sopro"
	"github.com/pablodz/sopro/pkg/transcoder"
	"github.com/urfave/cli/v2"
)

const VERSION = "v0.1.3"

func main() {
	app := &cli.App{
		Name:    "sopro",
		Usage:   "High performance audio processing tool",
		Version: VERSION,
		Commands: []*cli.Command{
			{
				Name:  "transcoder",
				Usage: "Transcode audio files",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "method",
						Usage:    "transcode method",
						Category: "method",
						Required: true,
						Value:    1,
					},

					&cli.IntFlag{
						Name:     "buffer",
						Usage:    "buffer size",
						Category: "method",
						Required: false,
						Value:    1024,
					},

					&cli.StringFlag{
						Name:     "in",
						Usage:    "input filename, can be any",
						Category: "input",
						Required: true,
					},

					&cli.StringFlag{
						Name:     "out",
						Usage:    "output filename, can be any",
						Required: true,
					},

					&cli.StringFlag{
						Name:     "in-format",
						Usage:    "input file format",
						Category: "input",
						// Required: true,
					},

					&cli.StringFlag{
						Name:     "out-format",
						Usage:    "output file format",
						Category: "output",
						// Required: true,
					},

					&cli.IntFlag{
						Name:     "in-sr",
						Usage:    "input sample rate",
						Required: true,
					},

					&cli.IntFlag{
						Name:     "out-sr",
						Usage:    "output sample rate",
						Required: true,
					},

					&cli.IntFlag{
						Name:     "in-c",
						Usage:    "input channels",
						Required: true,
					},

					&cli.IntFlag{
						Name:     "out-c",
						Usage:    "output channels",
						Required: true,
					},

					&cli.IntFlag{
						Name:     "in-b",
						Usage:    "input bits",
						Required: true,
					},

					&cli.IntFlag{
						Name:     "out-b",
						Usage:    "output bits",
						Required: true,
					},

					&cli.BoolFlag{
						Name:     "v",
						Usage:    "verbose",
						Value:    false,
						Required: false,
					},
				},

				Action: func(ctx *cli.Context) error {
					fmt.Println("[Version]    ", ctx.App.Version)

					methodT := ctx.Int("method")
					buffer := ctx.Int("buffer")
					inputFilename := ctx.String("in")
					outputFilename := ctx.String("out")
					// inputFormat := ctx.String("in-format")
					// outputFormat := ctx.String("out-format")
					inputSampleRate := ctx.Int("in-sr")
					outputSampleRate := ctx.Int("out-sr")
					inputChannels := ctx.Int("in-c")
					outputChannels := ctx.Int("out-c")
					inputBits := ctx.Int("in-b")
					outputBits := ctx.Int("out-b")
					verbose := ctx.Bool("v")

					// check required flags
					if methodT == -1 {
						return cli.Exit("method is required", 1)
					}

					if buffer <= 1 {
						return cli.Exit("buffer is not valid", 1)
					}

					if inputFilename == "" {
						return cli.Exit("input file is required", 1)
					}

					if outputFilename == "" {
						return cli.Exit("output file is required", 1)
					}

					// if inputFormat == "" {
					// 	return cli.Exit("input format is required", 1)
					// }

					// if outputFormat == "" {
					// 	return cli.Exit("output format is required", 1)
					// }

					if inputSampleRate == 0 {
						return cli.Exit("input sample rate is required", 1)
					}

					if outputSampleRate == 0 {
						return cli.Exit("output sample rate is required", 1)
					}

					if inputChannels == 0 {
						return cli.Exit("input channels is required", 1)
					}

					if outputChannels == 0 {
						return cli.Exit("output channels is required", 1)
					}

					if inputBits == 0 {
						return cli.Exit("input bits is required", 1)
					}

					if outputBits == 0 {
						return cli.Exit("output bits is required", 1)
					}

					// print the flags
					fmt.Println("[Method]     ", methodT)
					fmt.Println("[Buffer]     ", buffer)
					fmt.Println("[Input]      ", inputFilename)
					fmt.Println("[Output]     ", outputFilename)
					// fmt.Println("[In Format]  ", inputFormat)
					// fmt.Println("[Out Format] ", outputFormat)
					fmt.Println("[In SR]      ", inputSampleRate)
					fmt.Println("[Out SR]     ", outputSampleRate)
					fmt.Println("[In C]       ", inputChannels)
					fmt.Println("[Out C]      ", outputChannels)
					fmt.Println("[In B]       ", inputBits)
					fmt.Println("[Out B]      ", outputBits)
					fmt.Println("[Verbose]    ", verbose)

					// Open the input file
					in, err := os.Open(inputFilename)
					if err != nil {
						panic(err)
					}
					defer in.Close()

					// Create the output file
					out, err := os.Create(outputFilename)
					if err != nil {
						panic(err)
					}
					defer out.Close()

					// create a transcoder
					t := &transcoder.Transcoder{
						MethodT: methodT, // method.BIT_LOOKUP_TABLE=1
						InConfigs: sopro.AudioConfig{
							Endianness: cpuarch.LITTLE_ENDIAN,
						},
						OutConfigs: sopro.AudioConfig{
							Endianness: cpuarch.LITTLE_ENDIAN,
						},
						SizeBuffer: buffer,
						Verbose:    verbose,
					}

					// Transcode the file
					err = t.Mulaw2Wav(
						&sopro.In{
							Data: in,
							AudioFileGeneral: sopro.AudioFileGeneral{
								Format: fileformat.AUDIO_MULAW,
								Config: audioconfig.MulawConfig{
									BitDepth:   inputBits,
									Channels:   inputChannels,
									Encoding:   encoding.SPACE_LOGARITHMIC, // ulaw is logarithmic
									SampleRate: inputSampleRate,
								},
							},
						},
						&sopro.Out{
							Data: out,
							AudioFileGeneral: sopro.AudioFileGeneral{
								Format: fileformat.AUDIO_WAV,
								Config: audioconfig.WavConfig{
									BitDepth:   outputBits,
									Channels:   outputChannels,
									Encoding:   encoding.SPACE_LOGARITHMIC,
									SampleRate: outputSampleRate,
								},
							},
						},
					)

					if err != nil {
						panic(err)
					}

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
