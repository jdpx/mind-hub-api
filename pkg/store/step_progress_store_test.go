package store_test

// const (
// 	stepProgressTableName = "step_progress"
// )

// func TestStepProgresshandlerGet(t *testing.T) {
// 	sID := fake.CharactersN(10)
// 	uID := fake.CharactersN(10)
// 	stepProgress := builder.NewStepProgressBuilder().WithStepID(sID).WithUserID(uID).Build()

// 	testCases := []struct {
// 		desc               string
// 		userID             string
// 		stepID             string
// 		clientExpectations func(client *storemocks.MockStorer)

// 		expectedStepProgress *store.StepProgress
// 		expectedErr          error
// 	}{
// 		{
// 			desc:   "given a stepProgress is returned from store, progress is returned",
// 			userID: uID,
// 			stepID: sID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				params := map[string]string{
// 					"stepID": sID,
// 					"userID": uID,
// 				}

// 				client.EXPECT().Get(gomock.Any(), stepProgressTableName, params, gomock.Any()).SetArg(3, stepProgress)
// 			},

// 			expectedStepProgress: &stepProgress,
// 		},
// 		{
// 			desc:   "given a NotFound error is returned from the store, nil map is returned",
// 			userID: uID,
// 			stepID: sID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				params := map[string]string{
// 					"stepID": sID,
// 					"userID": uID,
// 				}

// 				client.EXPECT().Get(gomock.Any(), stepProgressTableName, params, gomock.Any()).Return(store.ErrNotFound)
// 			},

// 			expectedStepProgress: nil,
// 		},
// 		{
// 			desc:   "given a generic error is returned from the store, error returned",
// 			userID: uID,
// 			stepID: sID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				params := map[string]string{
// 					"stepID": sID,
// 					"userID": uID,
// 				}

// 				client.EXPECT().Get(gomock.Any(), stepProgressTableName, params, gomock.Any()).Return(fmt.Errorf("error occurred"))
// 			},

// 			expectedErr: fmt.Errorf("error occurred"),
// 		},
// 	}
// 	for _, tt := range testCases {
// 		t.Run(tt.desc, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			clientMock := storemocks.NewMockStorer(ctrl)

// 			if tt.clientExpectations != nil {
// 				tt.clientExpectations(clientMock)
// 			}

// 			resolver := store.NewStepProgressHandler(clientMock)
// 			ctx := context.Background()

// 			n, err := resolver.Get(ctx, tt.stepID, tt.userID)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.Equal(t, tt.expectedStepProgress, n)
// 			}
// 		})
// 	}
// }

// func TestStepProgresshandlerStart(t *testing.T) {
// 	sID := fake.CharactersN(10)
// 	uID := fake.CharactersN(10)
// 	stepProgress := builder.NewStepProgressBuilder().WithStepID(sID).WithUserID(uID).Build()

// 	testCases := []struct {
// 		desc               string
// 		userID             string
// 		stepID             string
// 		clientExpectations func(client *storemocks.MockStorer)

// 		expectedStepProgress *store.StepProgress
// 		expectedErr          error
// 	}{
// 		{
// 			desc:   "given a progress is returned from store, started progress is returned",
// 			userID: uID,
// 			stepID: sID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				client.EXPECT().Put(gomock.Any(), stepProgressTableName, gomock.Any()).Return(nil)
// 			},

// 			expectedStepProgress: &stepProgress,
// 		},
// 		{
// 			desc:   "given a generic error is returned from the store, error returned",
// 			userID: uID,
// 			stepID: sID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				client.EXPECT().Put(gomock.Any(), stepProgressTableName, gomock.Any()).Return(fmt.Errorf("error occurred"))
// 			},

// 			expectedErr: fmt.Errorf("error occurred"),
// 		},
// 	}
// 	for _, tt := range testCases {
// 		t.Run(tt.desc, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			clientMock := storemocks.NewMockStorer(ctrl)

// 			if tt.clientExpectations != nil {
// 				tt.clientExpectations(clientMock)
// 			}

// 			resolver := store.NewStepProgressHandler(clientMock)
// 			ctx := context.Background()

// 			n, err := resolver.Start(ctx, tt.stepID, tt.userID)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.NotEqual(t, tt.expectedStepProgress.ID, n.ID)
// 				assert.Equal(t, tt.expectedStepProgress.StepID, n.StepID)
// 				assert.Equal(t, tt.expectedStepProgress.UserID, n.UserID)
// 				assert.Equal(t, store.STATUS_STARTED, n.State)
// 			}
// 		})
// 	}
// }

// func TestStepProgresshandlerCompleted(t *testing.T) {
// 	sID := fake.CharactersN(10)
// 	uID := fake.CharactersN(10)
// 	stepProgress := builder.NewStepProgressBuilder().WithStepID(sID).WithUserID(uID).Build()

// 	testCases := []struct {
// 		desc               string
// 		userID             string
// 		stepID             string
// 		clientExpectations func(client *storemocks.MockStorer)

// 		expectedStepProgress *store.StepProgress
// 		expectedErr          error
// 	}{
// 		{
// 			desc:   "given a progress is returned from store, started progress is returned",
// 			userID: uID,
// 			stepID: sID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				keys := map[string]string{
// 					"stepID": stepProgress.StepID,
// 					"userID": stepProgress.UserID,
// 				}

// 				expression := "SET dateCompleted = :dateCompleted, progressState = :progressState"

// 				completedProgress := builder.NewStepProgressBuilder().WithStepID(sID).WithUserID(uID).Completed().Build()
// 				client.EXPECT().Update(gomock.Any(), stepProgressTableName, keys, expression, gomock.Any(), gomock.Any()).SetArg(5, completedProgress)
// 			},

// 			expectedStepProgress: &stepProgress,
// 		},
// 		{
// 			desc:   "given a generic error is returned from the store, error returned",
// 			userID: uID,
// 			stepID: sID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				keys := map[string]string{
// 					"stepID": stepProgress.StepID,
// 					"userID": stepProgress.UserID,
// 				}

// 				expression := "SET dateCompleted = :dateCompleted, progressState = :progressState"

// 				client.EXPECT().Update(gomock.Any(), stepProgressTableName, keys, expression, gomock.Any(), gomock.Any()).Return(fmt.Errorf("error occurred"))
// 			},

// 			expectedErr: fmt.Errorf("error occurred"),
// 		},
// 	}
// 	for _, tt := range testCases {
// 		t.Run(tt.desc, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			clientMock := storemocks.NewMockStorer(ctrl)

// 			if tt.clientExpectations != nil {
// 				tt.clientExpectations(clientMock)
// 			}

// 			resolver := store.NewStepProgressHandler(clientMock)
// 			ctx := context.Background()

// 			n, err := resolver.Complete(ctx, tt.stepID, tt.userID)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.NotEqual(t, tt.expectedStepProgress.ID, n.ID)
// 				assert.Equal(t, tt.expectedStepProgress.StepID, n.StepID)
// 				assert.Equal(t, tt.expectedStepProgress.UserID, n.UserID)
// 				assert.Equal(t, store.STATUS_COMPLETED, n.State)
// 			}
// 		})
// 	}
// }
