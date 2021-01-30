package store_test

const (
	stepNoteTableName = "step_note"
)

// func TestStepNotehandlerGet(t *testing.T) {
// 	cID := fake.CharactersN(10)
// 	uID := fake.CharactersN(10)
// 	stepNote := builder.NewStepNoteBuilder().WithStepID(cID).WithUserID(uID).Build()

// 	testCases := []struct {
// 		desc               string
// 		userID             string
// 		stepID             string
// 		clientExpectations func(client *storemocks.MockStorer)

// 		expectedStepNote *store.StepNote
// 		expectedErr      error
// 	}{
// 		{
// 			desc:   "given a note is returned from store, note is returned",
// 			userID: uID,
// 			stepID: cID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				params := map[string]string{
// 					"stepID": cID,
// 					"userID": uID,
// 				}

// 				client.EXPECT().Get(gomock.Any(), stepNoteTableName, params, gomock.Any()).SetArg(3, stepNote)
// 			},

// 			expectedStepNote: &stepNote,
// 		},
// 		{
// 			desc:   "given a NotFound error is returned from the store, nil map is returned",
// 			userID: uID,
// 			stepID: cID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				params := map[string]string{
// 					"stepID": cID,
// 					"userID": uID,
// 				}

// 				client.EXPECT().Get(gomock.Any(), stepNoteTableName, params, gomock.Any()).Return(store.ErrNotFound)
// 			},

// 			expectedStepNote: nil,
// 		},
// 		{
// 			desc:   "given a generic error is returned from the store, error returned",
// 			userID: uID,
// 			stepID: cID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				params := map[string]string{
// 					"stepID": cID,
// 					"userID": uID,
// 				}

// 				client.EXPECT().Get(gomock.Any(), stepNoteTableName, params, gomock.Any()).Return(fmt.Errorf("error occurred"))
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

// 			resolver := store.NewStepNoteHandler(clientMock)
// 			ctx := context.Background()

// 			n, err := resolver.Get(ctx, tt.stepID, tt.userID)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.Equal(t, tt.expectedStepNote, n)
// 			}
// 		})
// 	}
// }

// func TestStepNotehandlerCreate(t *testing.T) {
// 	stepNote := builder.NewStepNoteBuilder().WithID("").Build()

// 	testCases := []struct {
// 		desc               string
// 		clientExpectations func(client *storemocks.MockStorer)

// 		expectedStepNote *store.StepNote
// 		expectedErr      error
// 	}{
// 		{
// 			desc: "given a note is returned from store, note is returned",
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				client.EXPECT().Put(gomock.Any(), stepNoteTableName, gomock.Any()).Return(nil)
// 			},

// 			expectedStepNote: &stepNote,
// 		},
// 		{
// 			desc: "given a generic error is returned from the store, error returned",
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				client.EXPECT().Put(gomock.Any(), stepNoteTableName, gomock.Any()).Return(fmt.Errorf("error occurred"))
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

// 			resolver := store.NewStepNoteHandler(clientMock)
// 			ctx := context.Background()

// 			n, err := resolver.Create(ctx, stepNote)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.NotEqual(t, tt.expectedStepNote.ID, n.ID)
// 				assert.Equal(t, tt.expectedStepNote.StepID, n.StepID)
// 				assert.Equal(t, tt.expectedStepNote.UserID, n.UserID)
// 				assert.Equal(t, tt.expectedStepNote.Value, n.Value)
// 			}
// 		})
// 	}
// }

// func TestStepNotehandlerUpdate(t *testing.T) {
// 	stepNote := builder.NewStepNoteBuilder().Build()

// 	testCases := []struct {
// 		desc               string
// 		clientExpectations func(client *storemocks.MockStorer)

// 		expectedStepNote *store.StepNote
// 		expectedErr      error
// 	}{
// 		{
// 			desc: "given a note is returned from store, note is returned",
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				keys := map[string]string{
// 					"stepID": stepNote.StepID,
// 					"userID": stepNote.UserID,
// 				}

// 				exp := "set info.value = :value"

// 				client.EXPECT().Update(gomock.Any(), stepNoteTableName, keys, exp, gomock.Any(), gomock.Any()).SetArg(5, stepNote)
// 			},

// 			expectedStepNote: &stepNote,
// 		},
// 		{
// 			desc: "given a generic error is returned from the store, error returned",
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				keys := map[string]string{
// 					"stepID": stepNote.StepID,
// 					"userID": stepNote.UserID,
// 				}

// 				exp := "set info.value = :value"

// 				client.EXPECT().Update(gomock.Any(), stepNoteTableName, keys, exp, gomock.Any(), gomock.Any()).Return(fmt.Errorf("error occurred"))
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

// 			resolver := store.NewStepNoteHandler(clientMock)
// 			ctx := context.Background()

// 			n, err := resolver.Update(ctx, stepNote)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.Equal(t, tt.expectedStepNote.ID, n.ID)
// 				assert.Equal(t, tt.expectedStepNote.StepID, n.StepID)
// 				assert.Equal(t, tt.expectedStepNote.UserID, n.UserID)
// 				assert.Equal(t, tt.expectedStepNote.Value, n.Value)
// 			}
// 		})
// 	}
// }
