package store_test

// const (
// 	courseNoteTableName = "course_note"
// )

// func TestCourseNotehandlerGet(t *testing.T) {
// 	cID := fake.CharactersN(10)
// 	uID := fake.CharactersN(10)
// 	courseNote := builder.NewCourseNoteBuilder().WithCourseID(cID).WithUserID(uID).Build()

// 	testCases := []struct {
// 		desc               string
// 		userID             string
// 		courseID           string
// 		clientExpectations func(client *storemocks.MockStorer)

// 		expectedCourseNote *store.CourseNote
// 		expectedErr        error
// 	}{
// 		{
// 			desc:     "given a courseNote is returned from store, note is returned",
// 			userID:   uID,
// 			courseID: cID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				params := map[string]string{
// 					"courseID": cID,
// 					"userID":   uID,
// 				}

// 				client.EXPECT().Get(gomock.Any(), courseNoteTableName, params, gomock.Any()).SetArg(3, courseNote)
// 			},

// 			expectedCourseNote: &courseNote,
// 		},
// 		{
// 			desc:     "given a NotFound error is returned from the store, nil map is returned",
// 			userID:   uID,
// 			courseID: cID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				params := map[string]string{
// 					"courseID": cID,
// 					"userID":   uID,
// 				}

// 				client.EXPECT().Get(gomock.Any(), courseNoteTableName, params, gomock.Any()).Return(store.ErrNotFound)
// 			},

// 			expectedCourseNote: nil,
// 		},
// 		{
// 			desc:     "given a generic error is returned from the store, error returned",
// 			userID:   uID,
// 			courseID: cID,
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				params := map[string]string{
// 					"courseID": cID,
// 					"userID":   uID,
// 				}

// 				client.EXPECT().Get(gomock.Any(), courseNoteTableName, params, gomock.Any()).Return(fmt.Errorf("error occurred"))
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

// 			resolver := store.NewCourseNoteHandler(clientMock)
// 			ctx := context.Background()

// 			n, err := resolver.Get(ctx, tt.courseID, tt.userID)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.Equal(t, tt.expectedCourseNote, n)
// 			}
// 		})
// 	}
// }

// func TestCourseNotehandlerCreate(t *testing.T) {
// 	courseNote := builder.NewCourseNoteBuilder().WithID("").Build()

// 	testCases := []struct {
// 		desc               string
// 		clientExpectations func(client *storemocks.MockStorer)

// 		expectedCourseNote *store.CourseNote
// 		expectedErr        error
// 	}{
// 		{
// 			desc: "given a note is returned from store, note is returned",
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				client.EXPECT().Put(gomock.Any(), courseNoteTableName, gomock.Any()).Return(nil)
// 			},

// 			expectedCourseNote: &courseNote,
// 		},
// 		{
// 			desc: "given a generic error is returned from the store, error returned",
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				client.EXPECT().Put(gomock.Any(), courseNoteTableName, gomock.Any()).Return(fmt.Errorf("error occurred"))
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

// 			resolver := store.NewCourseNoteHandler(clientMock)
// 			ctx := context.Background()

// 			n, err := resolver.Create(ctx, courseNote)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.NotEqual(t, tt.expectedCourseNote.ID, n.ID)
// 				assert.Equal(t, tt.expectedCourseNote.CourseID, n.CourseID)
// 				assert.Equal(t, tt.expectedCourseNote.UserID, n.UserID)
// 				assert.Equal(t, tt.expectedCourseNote.Value, n.Value)
// 			}
// 		})
// 	}
// }

// func TestCourseNotehandlerUpdate(t *testing.T) {
// 	courseNote := builder.NewCourseNoteBuilder().Build()

// 	testCases := []struct {
// 		desc               string
// 		clientExpectations func(client *storemocks.MockStorer)

// 		expectedCourseNote *store.CourseNote
// 		expectedErr        error
// 	}{
// 		{
// 			desc: "given a note is returned from store, note is returned",
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				keys := map[string]string{
// 					"courseID": courseNote.CourseID,
// 					"userID":   courseNote.UserID,
// 				}

// 				exp := "set info.value = :value"

// 				client.EXPECT().Update(gomock.Any(), courseNoteTableName, keys, exp, gomock.Any(), gomock.Any()).SetArg(5, courseNote)
// 			},

// 			expectedCourseNote: &courseNote,
// 		},
// 		{
// 			desc: "given a generic error is returned from the store, error returned",
// 			clientExpectations: func(client *storemocks.MockStorer) {
// 				keys := map[string]string{
// 					"courseID": courseNote.CourseID,
// 					"userID":   courseNote.UserID,
// 				}

// 				exp := "set info.value = :value"

// 				client.EXPECT().Update(gomock.Any(), courseNoteTableName, keys, exp, gomock.Any(), gomock.Any()).Return(fmt.Errorf("error occurred"))
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

// 			resolver := store.NewCourseNoteHandler(clientMock)
// 			ctx := context.Background()

// 			n, err := resolver.Update(ctx, courseNote)

// 			if tt.expectedErr != nil {
// 				assert.EqualError(t, err, tt.expectedErr.Error())
// 			} else {
// 				assert.Nil(t, err)
// 				assert.Equal(t, tt.expectedCourseNote.ID, n.ID)
// 				assert.Equal(t, tt.expectedCourseNote.CourseID, n.CourseID)
// 				assert.Equal(t, tt.expectedCourseNote.UserID, n.UserID)
// 				assert.Equal(t, tt.expectedCourseNote.Value, n.Value)
// 			}
// 		})
// 	}
// }
