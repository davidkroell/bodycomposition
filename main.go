package bodycomposition

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"

	connect "github.com/abrander/garmin-connect"
	"github.com/tormoder/fit"
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
	CaloriesActiveMet float64
}

func (bc BodyComposition) writeFitFile(writer io.Writer) error {
	fitfile, err := fit.NewFile(fit.FileTypeWeight, fit.NewHeader(fit.V20, true))
	if err != nil {
		return err
	}

	weight, err := fitfile.Weight()
	if err != nil {
		return err
	}

	weight.WeightScales = []*fit.WeightScaleMsg{
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
			ActiveMet:         uint16(bc.CaloriesActiveMet),
		},
	}

	return fit.Encode(writer, fitfile, binary.BigEndian)
}

func (bc BodyComposition) uploadFitFile(reader io.Reader, email string, password string) bool {
	client := connect.NewClient(connect.Credentials(email, password))

	_, err := client.ImportActivity(reader, connect.ActivityFormatFIT)
	if err != nil {
		fmt.Println("Error uploading file to Garmin Connect: ", err.Error())
		return false
	}

	return true
}

// UploadWeight uploads the bodycomposition data to garmin connect
func (bc BodyComposition) UploadWeight(email, password string) bool {
	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()

		err := bc.writeFitFile(writer)
		if err != nil {
			panic(err)
		}
	}()

	return bc.uploadFitFile(reader, email, password)
}

// NewBodyComposition creates a new bodycomposition type
func NewBodyComposition(weight, percentFat, percentHydration, percentBone, percentMuscle, visceralFatRating, physiqueRating, metabolicAge, caloriesActiveMet float64, timestamp int64) BodyComposition {
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
		CaloriesActiveMet: caloriesActiveMet,
	}
}
