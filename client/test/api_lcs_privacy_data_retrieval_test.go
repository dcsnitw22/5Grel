/*
Nudm_SDM

Testing LCSPrivacyDataRetrievalAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_LCSPrivacyDataRetrievalAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LCSPrivacyDataRetrievalAPIService GetLcsPrivacyData", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ueId string

		resp, httpRes, err := apiClient.LCSPrivacyDataRetrievalAPI.GetLcsPrivacyData(context.Background(), ueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
