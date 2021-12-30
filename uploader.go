package bodycomposition

import (
	connect "github.com/abrander/garmin-connect"
	"io"
)

func uploadFitFile(reader io.Reader, email string, password string) error {
	client := connect.NewClient(connect.Credentials(email, password))

	_, err := client.ImportActivity(reader, connect.ActivityFormatFIT)

	return err
}

// Upload uploads the bodycomposition data to garmin connect
func Upload(email, password string, bc BodyComposition) error {
	reader, writer := io.Pipe()

	go func() {
		defer writer.Close()

		err := bc.writeFitFile(writer)
		if err != nil {
			panic(err)
		}
	}()

	return uploadFitFile(reader, email, password)
}
