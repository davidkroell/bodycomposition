package bodycomposition

import (
	"encoding/binary"
	"io"
	"time"

	"github.com/tormoder/fit"
)

// BodyComposition is the data container struct for managing the body measurements
type BodyComposition struct {
	TimeStamp         time.Time
	Weight            float64
	PercentFat        float64
	PercentHydration  float64
	BoneMass          float64
	MuscleMass        float64
	VisceralFatRating float64
	PhysiqueRating    float64
	MetabolicAge      float64
	CaloriesActiveMet float64
	BodyMassIndex     float64
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
			BoneMass:          uint16(bc.BoneMass * 100),
			MuscleMass:        uint16(bc.MuscleMass * 100),
			VisceralFatRating: uint8(bc.VisceralFatRating),
			PhysiqueRating:    uint8(bc.PhysiqueRating),
			MetabolicAge:      uint8(bc.MetabolicAge),
			ActiveMet:         uint16(bc.CaloriesActiveMet),
			Bmi:               uint16(bc.BodyMassIndex * 10),
		},
	}

	return fit.Encode(writer, fitfile, binary.BigEndian)
}

// NewBodyComposition creates a new BodyComposition instance
func NewBodyComposition(weight, percentFat, percentHydration, boneMass, muscleMass, visceralFatRating, physiqueRating, metabolicAge, caloriesActiveMet, bmi float64, timestamp int64) BodyComposition {
	ts := time.Now()
	if timestamp != -1 {
		ts = time.Unix(timestamp, 0)
	}

	return BodyComposition{
		TimeStamp:         ts,
		Weight:            weight,
		PercentFat:        percentFat,
		PercentHydration:  percentHydration,
		BoneMass:          boneMass,
		MuscleMass:        muscleMass,
		VisceralFatRating: visceralFatRating,
		PhysiqueRating:    physiqueRating,
		MetabolicAge:      metabolicAge,
		CaloriesActiveMet: caloriesActiveMet,
		BodyMassIndex:     bmi,
	}
}
