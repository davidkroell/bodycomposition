package bodycomposition

import (
	"encoding/binary"
	"fmt"
	"github.com/abrander/garmin-connect"
	"github.com/tormoder/fit"
	"io"
	"time"
)

type BodyComposition struct {
	TimeStamp         time.Time
	Weight            float64
	PercentFat        float64
	PercentHydration  float64
	PercentBone       float64
	PercentMuscle     float64
	VisceralFatRating float64
	PhysiqueRating    float64
	MetabolicAge      float64
}

func (bc BodyComposition) writeFitFile(writer io.Writer) error {
	weightfile := fit.WeightFile{
		UserProfile: nil,
		WeightScales: []*fit.WeightScaleMsg{
			{
				Timestamp:         bc.TimeStamp,
				Weight:            fit.Weight(bc.Weight * 100),
				PercentFat:        uint16(bc.PercentFat * 100),
				PercentHydration:  uint16(bc.PercentHydration * 100),
				BoneMass:          uint16(bc.Weight * bc.PercentBone),
				MuscleMass:        uint16(bc.Weight * bc.PercentMuscle),
				VisceralFatRating: uint8(bc.VisceralFatRating),
				PhysiqueRating:    uint8(bc.PhysiqueRating),
				MetabolicAge:      uint8(bc.MetabolicAge),
			},
		},
	}

	fitfile := fit.File{
		FileId: struct {
			Type         fit.FileType
			Manufacturer fit.Manufacturer
			Product      uint16
			SerialNumber uint32
			TimeCreated  time.Time
			Number       uint16
			ProductName  string
		}{Type: fit.FileTypeWeight, Manufacturer: fit.ManufacturerTanita},
		Header: struct {
			Size            byte
			ProtocolVersion byte
			ProfileVersion  uint16
			DataSize        uint32
			DataType        [4]byte
			CRC             uint16
		}{Size: 14, ProtocolVersion: 16, ProfileVersion: 2092, DataType: [4]byte{46, 70, 73, 84}},
	}

	err := fitfile.SetWeight(&weightfile)
	if err != nil {
		panic(err)
	}

	return fit.Encode(writer, &fitfile, binary.BigEndian)
}

func (bc BodyComposition) uploadFitFile(reader io.Reader, email string, password string) bool {
	client := connect.NewClient(connect.Credentials(email, password))

	_, err := client.ImportActivity(reader, connect.ActivityFormatFIT)
	if err != nil {
		if err == connect.ErrWrongCredentials {
			fmt.Println("Authentication failed")
		} else {
			panic(err)
		}
	}

	return true
}

func (bc BodyComposition) UploadWeight(email, password string) {
	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()

		err := bc.writeFitFile(writer)
		if err != nil {
			panic(err)
		}
	}()

	bc.uploadFitFile(reader, email, password)
}

func NewBodyComposition(weight, percentFat, percentHydration, percentBone, percentMuscle, visceralFatRating, physiqueRating, metabolicAge float64, timestamp int64) BodyComposition {
	ts := time.Now()
	if timestamp != -1 {
		ts = time.Unix(timestamp, 0)
	}

	return BodyComposition{
		TimeStamp:         ts,
		Weight:            weight,
		PercentFat:        percentFat,
		PercentHydration:  percentHydration,
		PercentBone:       percentBone,
		PercentMuscle:     percentMuscle,
		VisceralFatRating: visceralFatRating,
		PhysiqueRating:    physiqueRating,
		MetabolicAge:      metabolicAge,
	}
}
