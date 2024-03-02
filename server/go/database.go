package openapi

func Database() map[string]SessionManagementSubscriptionData {
	// Create a map with string keys (ueid) and SdmSubscription values
	subscriptions := make(map[string]SessionManagementSubscriptionData)

	// Define and initialize SdmSubscription objects
	createdSubscription1 := SessionManagementSubscriptionData{
		SingleNssai: Snssai{
			Sst: 1,
			Sd:  "sd1",
		},
		InternalGroupIds:               []string{"internalGroup1", "internalGroup2"},
		SharedVnGroupDataIds:           map[string]string{"groupId1": "sharedDataId1", "groupId2": "sharedDataId2"},
		SharedDnnConfigurationsId:      "sharedDnnConfigurationsId1",
		OdbPacketServices:              &OdbPacketServices{ /* Fill OdbPacketServices fields here */ },
		TraceData:                      &TraceData{ /* Fill TraceData fields here */ },
		SharedTraceDataId:              "sharedTraceDataId1",
		ExpectedUeBehavioursList:       map[string]ExpectedUeBehaviourData{"dnn1": ExpectedUeBehaviourData{ /* Fill ExpectedUeBehaviourData fields here */ }},
		SuggestedPacketNumDlList:       map[string]SuggestedPacketNumDl{"dnn1": SuggestedPacketNumDl{ /* Fill SuggestedPacketNumDl fields here */ }},
		Var3gppChargingCharacteristics: "3gppChargingCharacteristics1",
		SupportedFeatures:              "supportedFeatures1",
	}

	createdSubscription2 := SessionManagementSubscriptionData{
		SingleNssai: Snssai{
			Sst: 2,
			Sd:  "sd2",
		},
		InternalGroupIds:               []string{"internalGroup3", "internalGroup4"},
		SharedVnGroupDataIds:           map[string]string{"groupId3": "sharedDataId3", "groupId4": "sharedDataId4"},
		SharedDnnConfigurationsId:      "sharedDnnConfigurationsId2",
		OdbPacketServices:              &OdbPacketServices{ /* Fill OdbPacketServices fields here */ },
		TraceData:                      &TraceData{ /* Fill TraceData fields here */ },
		SharedTraceDataId:              "sharedTraceDataId2",
		ExpectedUeBehavioursList:       map[string]ExpectedUeBehaviourData{"dnn3": ExpectedUeBehaviourData{ /* Fill ExpectedUeBehaviourData fields here */ }},
		SuggestedPacketNumDlList:       map[string]SuggestedPacketNumDl{"dnn3": SuggestedPacketNumDl{ /* Fill SuggestedPacketNumDl fields here */ }},
		Var3gppChargingCharacteristics: "3gppChargingCharacteristics2",
		SupportedFeatures:              "supportedFeatures2",
	}

	createdSubscription3 := SessionManagementSubscriptionData{
		SingleNssai: Snssai{
			Sst: 3,
			Sd:  "sd3",
		},
		InternalGroupIds:               []string{"internalGroup5", "internalGroup6"},
		SharedVnGroupDataIds:           map[string]string{"groupId5": "sharedDataId5", "groupId6": "sharedDataId6"},
		SharedDnnConfigurationsId:      "sharedDnnConfigurationsId3",
		OdbPacketServices:              &OdbPacketServices{ /* Fill OdbPacketServices fields here */ },
		TraceData:                      &TraceData{ /* Fill TraceData fields here */ },
		SharedTraceDataId:              "sharedTraceDataId3",
		ExpectedUeBehavioursList:       map[string]ExpectedUeBehaviourData{"dnn5": ExpectedUeBehaviourData{ /* Fill ExpectedUeBehaviourData fields here */ }},
		SuggestedPacketNumDlList:       map[string]SuggestedPacketNumDl{"dnn5": SuggestedPacketNumDl{ /* Fill SuggestedPacketNumDl fields here */ }},
		Var3gppChargingCharacteristics: "3gppChargingCharacteristics3",
		SupportedFeatures:              "supportedFeatures3",
	}

	// Populate the map with the createdSubscription objects
	subscriptions["erb32"] = createdSubscription1
	subscriptions["sup12"] = createdSubscription2
	subscriptions["dsv21"] = createdSubscription3

	return subscriptions
}

// database retrieves SessionManagementSubscriptionData from Redis or another data source.
// If the data is not found in Redis, it can be fetched from another source (e.g., a database),
// and then stored in Redis for subsequent requests. This function showcases storing and retrieving
// data from Redis and converting the retrieved data to SessionManagementSubscriptionData struct.
// Adjust the logic based on your specific requirements.

// func database() map[string]SessionManagementSubscriptionData {
// 	// Create a Redis client
// 	redisClient := redisClient()
// 	defer redisClient.Close()

// 	// Example key for demonstration purposes
// 	key := "erb32"

// 	// Try to fetch data from Redis
// 	val, err := redisClient.Get(ctx, key).Result()
// 	if err == nil {
// 		// Data found in Redis, convert val to your SessionManagementSubscriptionData struct
// 		var subscriptionData SessionManagementSubscriptionData
// 		if err := json.Unmarshal([]byte(val), &subscriptionData); err != nil {
// 			panic(err)
// 		}

// 		// Populate the map with the retrieved SessionManagementSubscriptionData
// 		subscriptions := make(map[string]SessionManagementSubscriptionData)
// 		subscriptions[key] = subscriptionData

// 		return subscriptions
// 	}

// 	// Data not found in Redis, simulate fetching from another source (e.g., a database)
// 	// Your logic to fetch data from another source goes here

// 	// Example: Simulate fetching data from another source
// 	createdSubscription := SessionManagementSubscriptionData{
// 		SingleNssai: Snssai{
// 			Sst: 1,
// 			Sd:  "sd1",
// 		},
// 		// ... (Fill other fields accordingly)
// 	}

// 	// Populate the map with the createdSubscription object
// 	subscriptions := make(map[string]SessionManagementSubscriptionData)
// 	subscriptions[key] = createdSubscription

// 	// Store the fetched data in Redis for subsequent requests
// 	jsonData, err := json.Marshal(createdSubscription)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := redisClient.Set(ctx, key, string(jsonData), 0).Err(); err != nil {
// 		panic(err)
// 	}

// 	return subscriptions
// }
