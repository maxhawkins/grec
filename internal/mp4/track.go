package mp4

import (
	"fmt"
	"io"
)

// ReadTrack dumps all the samples for the given track ID
func ReadTrack(filename string, trackID int) ([]byte, error) {
	atom, err := Open(filename)
	if err != nil {
		return nil, err
	}

	if len(atom.moov.traks) <= trackID {
		return nil, fmt.Errorf("unknown track id %d", trackID)
	}
	trak := atom.moov.traks[trackID]

	var all []byte

	sampleID := 0
	for i := 0; i < len(trak.chunks); i++ {
		sampleOffset := trak.chunks[i].offset
		for j := 0; j < int(trak.chunks[i].sample_count); j++ {
			sampleSize := trak.samples[sampleID].size

			_, err := atom.File.Seek(int64(sampleOffset), 0)
			if err != nil {
				return nil, err
			}
			dat := make([]byte, int(sampleSize))
			if _, err := io.ReadFull(atom.File, dat); err != nil {
				return nil, err
			}
			all = append(all, dat...)

			sampleOffset += sampleSize
			sampleID++
		}
	}

	return all, nil
}
