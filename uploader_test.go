// +build integration

package bodycomposition

import (
	"os"
	"testing"
)

const garminEmailEnvKey = "GARMIN_EMAIL"
const garminPasswordEnvKey = "GARMIN_PASSWORD"

// TestUploadWeightToGarminIntegration uploading a weight-in to garmin connect
func TestUploadWeightToGarminIntegration(t *testing.T) {

	email := os.Getenv(garminEmailEnvKey)
	passWord := os.Getenv(garminPasswordEnvKey)

	if email == "" {
		t.Fatalf("Environment variable %s was not set. This is required.", garminEmailEnvKey)
	}

	if passWord == "" {
		t.Fatalf("Environment variable %s was not set. This is required.", garminPasswordEnvKey)
	}

	var bc = NewBodyComposition(80, 14.4, 55.2, 37, 2.98, 45.5, 55, 21, 5, 23, 2250, 12.6, -1)

	err := Upload(email, passWord, bc)

	if err != nil {
		t.Fatalf("Error uploading weight to Garmin Connect: %s", err.Error())
	}
}
