[![Go Report Card](https://goreportcard.com/badge/github.com/pablodz/sopro)](https://goreportcard.com/report/github.com/pablodz/sopro)

# SoPro (next generation SOund PROcessing)

Sox is a great tool, but it's not easy to use. SoPro is a next generation sound processing tool that is easy to use and easy to extend. By now only audio files can be converted to other formats, but in the future more features will be added, like video processing, etc.

```
                     ┌─────────────────┐
 raw data    ───────►│                 ├────────►   returns raw data in other format
                     │                 │
                     │                 │
 websocket   ───────►│                 ├────────►   returns raw data in other formats
                     │                 │
                     │    SOPRO-CORE   │
 chunked data───────►│                 ├────────►   returns chunked processed data
                     │                 │
                     │                 │
 gRPC        ───────►│                 ├────────►   returns grpc chunked data
                     │                 │
                     └─────────────────┘

Examples:

- ulaw -> wav pcm
- ulaw -> wav pcm normalized (on the fly)
```

## Installation

```bash
go get -v github.com/pablodz/sopro
```

## Methods planned to be implemented

- [x] Chunked
- [x] Full memory
- [ ] Batch
- [ ] Streaming

## Examples

Check [./examples](./examples/) folder

## Roadmap

- [ ] CLI (sox-friendly)
- [ ] GUI (in another repo)
- [ ] Microservice (in another repo)
  - [ ] HTTP
  - [ ] Websocket
  - [ ] gRPC
- [x] Audio file conversion
  - [ ] Format conversion [Work in progress...](docs/format_table.md)
  - [ ] Bitrate conversion
  - [ ] Channels conversion
  - [ ] Resampling
    - [ ] Downsampling
    - [ ] Upsampling
    - [ ] Interpolation
      - [ ] Linear
      - [ ] Cubic
      - [ ] B-spline
      - [ ] Polynomial
      - [ ] FIR filter 
      - [ ] IIR filter
      - [ ] LAG filter
      - [ ] Sinc
      - [ ] Spline
      - [ ] Windowed sinc
      - [ ] Windowed sinc (fast)
      - [ ] Neural network (NEW)
- [ ] Effects 
  - [ ] Tone/filter effects
    - [ ] allpass: RBJ all-pass biquad IIR filter
    - [ ] bandpass: RBJ band-pass biquad IIR filter
    - [ ] bandreject: RBJ band-reject biquad IIR filter
    - [ ] band: SPKit resonator band-pass IIR filter
    - [ ] bass: Tone control: RBJ shelving biquad IIR filter
    - [ ] equalizer: RBJ peaking equalisation biquad IIR filter
    - [ ] firfit+: FFT convolution FIR filter using given freq. response 
    - [ ] highpass: High-pass filter: Single pole or RBJ biquad IIR
    - [ ] hilbert: Hilbert transform filter (90 degrees phase shift)
    - [ ] lowpass: Low-pass filter: single pole or RBJ biquad IIR
    - [ ] sinc: Sinc-windowed low/high-pass/band-pass/reject FIR
    - [ ] treble: Tone control: RBJ shelving biquad IIR filter
  - [ ] Production effects
    - [ ] chorus: Make a single instrument sound like many
    - [ ] delay: Delay one or more channels
    - [ ] echo: Add an echo
    - [ ] echos: Add a sequence of echos
    - [ ] flanger: Stereo flanger
    - [ ] overdrive: Non-linear distortion
    - [ ] phaser: Phase shifter
    - [ ] repeat: Loop the audio a number of times
    - [ ] reverb: Add reverberation
    - [ ] reverse: Reverse the audio 
    - [ ] tremolo: Sinusoidal volume modulation
  - [ ] Volume/level effects
    - [ ] compand: Signal level compression/expansion/limiting
    - [ ] contrast: Phase contrast volume enhancement
    - [ ] dcshift: Apply or remove DC offset
    - [ ] fade: Apply a fade-in and/or fade-out to the audio
    - [ ] gain: Apply gain or attenuation; normalise/equalise/balance/headroom
    - [ ] loudness: Gain control with ISO 226 loudness compensation
    - [ ] mcompand: Multi-band compression/expansion/limiting
    - [ ] norm: Normalise to 0dB (or other)
    - [ ] vol: Adjust audio volume
  - [ ] Editing effects
    - [ ] pad: Pad (usually) the ends of the audio with silence
    - [ ] silence: Remove portions of silence from the audio
    - [ ] splice: Perform the equivalent of a cross-faded tape splice
    - [ ] trim: Cuts portions out of the audio
    - [ ] vad: Voice activity detector
  - [ ] Mixing effects
    - [ ] channels: Auto mix or duplicate to change number of channels
    - [ ] divide+: Divide sample values by those in the 1st channel 
    - [ ] remix: Produce arbitrarily mixed output channels
    - [ ] swap: Swap stereo channels
  - [ ] Pitch/tempo effects
    - [ ] bend: Bend pitch at given times without changing tempo
    - [ ] pitch: Adjust pitch (= key) without changing tempo
    - [ ] speed: Adjust pitch & tempo together
    - [ ] stretch: Adjust tempo without changing pitch (simple alg.)
    - [ ] tempo: Adjust tempo without changing pitch (WSOLA alg.)
  - [ ] Mastering effects
    - [ ] dither: Add dither noise to increase quantisation SNR
    - [ ] rate: Change audio sampling rate
  - [ ] Specialised filters/mixers
    - [ ] deemph: ISO 908 CD de-emphasis (shelving) IIR filter
    - [ ] earwax: Process CD audio to best effect for headphone use
    - [ ] noisered: Filter out noise from the audio
    - [ ] oops: Out Of Phase Stereo (or `Karaoke') effect
    - [ ] riaa: RIAA vinyl playback equalisation
  - [ ] Miscellaneous
    - [ ] ladspa: Apply LADSPA plug-in effects e.g. CMT (Computer Music Toolkit)
    - [ ] synth: Synthesise/modulate audio tones or noise signals
    - [ ] newfile: Create a new output file when an effects chain ends.
    - [ ] restart: Restart 1st effects chain when multiple chains exist.
  - [ ] Low-level signal processing effects
    - [ ] biquad: 2nd-order IIR filter using externally provided coefficients
    - [ ] downsample: Reduce sample rate by discarding samples
    - [ ] fir: FFT convolution FIR filter using externally provided coefficients
    - [ ] upsample: Increase sample rate by zero stuffing
- [ ] Mixer
  - [ ] Two channel mixer
  - [ ] Multi channel mixer
- [ ] Debugger
  - [x] Audio ascii graph in terminal
  - [x] Headers viewer in terminal
  - [ ] spectrogram: graph signal level vs. frequency & time
  - [ ] stat: Enumerate audio peak & RMS levels, approx. freq., etc.
  - [ ] stats: Multichannel aware `stat'


## Authors

All authors that contributed to this project are listed in the [AUTHORS](AUTHORS) file.

## License

SoPro is licensed under the [MIT License](LICENSE).