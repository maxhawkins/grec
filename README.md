# grec: Google Recorder extractor

grec extracts the embedded transcripts from Google Recorder m4a files.

The [Google Recorder app](https://play.google.com/store/apps/details?id=com.google.android.apps.recorder&hl=en_US) is an audio recording app from Google available on Google Pixel phones. It records audio voice notes and generates time-aligned transcripts using speech recognition.

When audio is exported from Recorder as an m4a file, the generated transcripts are embedded in a metadata track. This tool exports those embedded transcripts as JSON.

### Usage

    grec audio.m4a

Output:

```js
{
  "transcript": [
    {
      "word": "I",
      "formatted": "\nI",
      "startMs": 5660,
      "endMs": 5900
    },
    {
      "word": "end",
      "startMs": 5900,
      "endMs": 6380
    },
    // ...
```

### Installation

Download the binary for your platform:

* [macOS 64-bit](https://github.com/maxhawkins/grec/releases/download/v1.0.0/grec-macos-amd64.tar.gz)
* [Linux 64-bit](https://github.com/maxhawkins/grec/releases/download/v1.0.0/grec-linux-amd64.tar.gz)
* [Windows 64-bit](https://github.com/maxhawkins/grec/releases/download/v1.0.0/grec-windows-amd64.zip)

And run in the command line:

```bash
tar xvf grec.tar.gz
chmod +x ./grec
./grec audio.m4a
```

**Or, compile from source**

Go 1.14 or later is required.

```bash
go get -u github.com/maxhawkins/grec/cmd/grec
```
