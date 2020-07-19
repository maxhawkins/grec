package grec

import (
	"encoding/binary"

	"github.com/maxhawkins/grec/internal/mp4"
	"github.com/maxhawkins/grec/internal/pb"
	"google.golang.org/protobuf/proto"
)

// Word is a time-aligned word from the transcript.
type Word struct {
	Word      string `json:"word"`
	Formatted string `json:"formatted,omitempty"`
	StartMs   int    `json:"startMs"`
	EndMs     int    `json:"endMs"`
}

// Transcript is a voice transcription for an audio file.
type Transcript struct {
	Transcript []Word `json:"transcript"`
}

// DecodeTranscript decodes the binary-encoded transcript protobuf
// embedded in a Google Recoder m4a.
func DecodeTranscript(data []byte) (Transcript, error) {
	var msg []byte
	var tran Transcript
	for len(data) > 0 {
		size := binary.BigEndian.Uint32(data)
		msg, data = data[4:4+size], data[4+size:]

		var result pb.Result
		if err := proto.Unmarshal(msg, &result); err != nil {
			return Transcript{}, err
		}

		for _, word := range result.Words {
			tran.Transcript = append(tran.Transcript, Word{
				Word:      word.Word,
				Formatted: word.Formatted,
				StartMs:   int(word.StartMs),
				EndMs:     int(word.EndMs),
			})
		}
	}

	return tran, nil
}

// ParseFile reads the transcript data from an m4a file saved by Google Recorder.
func ParseFile(filename string) (Transcript, error) {
	bin, err := mp4.ReadTrack(filename, 0)
	if err != nil {
		return Transcript{}, err
	}

	tran, err := DecodeTranscript(bin)
	if err != nil {
		return Transcript{}, err
	}

	return tran, nil
}
