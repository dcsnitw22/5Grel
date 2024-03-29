/*
Nudm_SDM

Testing A2XSubscriptionDataRetrievalAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_A2XSubscriptionDataRetrievalAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test A2XSubscriptionDataRetrievalAPIService GetA2xData", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var supi string

		resp, httpRes, err := apiClient.A2XSubscriptionDataRetrievalAPI.GetA2xData(context.Background(), supi).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
