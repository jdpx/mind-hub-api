package service_test

import (
	"testing"

	"github.com/jdpx/mind-hub-api/pkg/graphcms"
	"github.com/jdpx/mind-hub-api/pkg/graphcms/builder"
	"github.com/jdpx/mind-hub-api/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestCourseFromCMS(t *testing.T) {
	gc := builder.NewCourseBuilder().Build()

	testCases := []struct {
		desc   string
		course graphcms.Course

		expectedCourse service.Course
	}{
		{
			desc:   "given a valid course, course converted",
			course: gc,

			expectedCourse: service.Course{
				ID:          gc.ID,
				Title:       gc.Title,
				Description: gc.Description,
				Sessions:    []*service.Session{},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			course := service.CourseFromCMS(&tt.course)

			assert.Equal(t, &tt.expectedCourse, course)
		})
	}
}

func TestCoursesFromCMS(t *testing.T) {
	gc := builder.NewCourseBuilder().Build()

	testCases := []struct {
		desc    string
		courses []*graphcms.Course

		expectedCourse []*service.Course
	}{
		{
			desc:    "given a valid course, course converted",
			courses: []*graphcms.Course{&gc},

			expectedCourse: []*service.Course{
				{
					ID:          gc.ID,
					Title:       gc.Title,
					Description: gc.Description,
					Sessions:    []*service.Session{},
				},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			course := service.CoursesFromCMS(tt.courses)

			assert.Equal(t, tt.expectedCourse, course)
		})
	}
}

func TestSessionFromCMS(t *testing.T) {
	gs := builder.NewSessionBuilder().Build()

	gsT := builder.NewStepBuilder().Build()
	gsWithSteps := builder.NewSessionBuilder().WithSteps(&gsT).Build()

	gc := builder.NewCourseBuilder().Build()
	gsWithCourse := builder.NewSessionBuilder().WithCourse(&gc).Build()

	testCases := []struct {
		desc    string
		session graphcms.Session

		expectedSession service.Session
	}{
		{
			desc:    "given a valid session without steps or course, session converted",
			session: gs,

			expectedSession: service.Session{
				ID:          gs.ID,
				Title:       gs.Title,
				Description: gs.Description,
				Steps:       []*service.Step{},
			},
		},
		{
			desc:    "given a valid session with steps, session with steps converted",
			session: gsWithSteps,

			expectedSession: service.Session{
				ID:          gsWithSteps.ID,
				Title:       gsWithSteps.Title,
				Description: gsWithSteps.Description,
				Steps: []*service.Step{
					{
						ID:          gsT.ID,
						Title:       gsT.Title,
						Description: gsT.Description,
						Type:        gsT.Type,
						VideoURL:    gsT.VideoURL,
						Question:    gsT.Question,
					},
				},
			},
		},
		{
			desc:    "given a valid session with a course, session with course converted",
			session: gsWithCourse,

			expectedSession: service.Session{
				ID:          gsWithCourse.ID,
				Title:       gsWithCourse.Title,
				Description: gsWithCourse.Description,
				Steps:       []*service.Step{},
				Course: &service.Course{
					ID:          gc.ID,
					Title:       gc.Title,
					Description: gc.Description,
					Sessions:    []*service.Session{},
				},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			session := service.SessionFromCMS(&tt.session)

			assert.Equal(t, &tt.expectedSession, session)
		})
	}
}

func TestSessionsFromCMS(t *testing.T) {
	gs := builder.NewSessionBuilder().Build()

	testCases := []struct {
		desc     string
		sessions []*graphcms.Session

		expectedSession []*service.Session
	}{
		{
			desc:     "given a valid session, session converted",
			sessions: []*graphcms.Session{&gs},

			expectedSession: []*service.Session{
				{
					ID:          gs.ID,
					Title:       gs.Title,
					Description: gs.Description,
					Steps:       []*service.Step{},
				},
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			sessions := service.SessionsFromCMS(tt.sessions)

			assert.Equal(t, tt.expectedSession, sessions)
		})
	}
}

func TestStepFromCMS(t *testing.T) {
	gs := builder.NewStepBuilder().Build()

	testCases := []struct {
		desc string
		step graphcms.Step

		expectedStep service.Step
	}{
		{
			desc: "given a valid step, step converted",
			step: gs,

			expectedStep: service.Step{
				ID:          gs.ID,
				Title:       gs.Title,
				Description: gs.Description,
				Type:        gs.Type,
				VideoURL:    gs.VideoURL,
				Question:    gs.Question,
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			step := service.StepFromCMS(&tt.step)

			assert.Equal(t, &tt.expectedStep, step)
		})
	}
}
