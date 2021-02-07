package service

type Service struct {
	Session        SessionServicer
	Step           StepServicer
	Course         CourseServicer
	CourseProgress CourseProgressServicer
	CourseNote     CourseNoteServicer
	StepProgress   StepProgressServicer
	StepNote       StepNoteServicer
	Timemap        TimemapServicer
}

// Option ...
type Option func(*Service)

func New(opts ...Option) *Service {
	s := &Service{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// WithCourse ...
func WithCourse(c CourseServicer) func(*Service) {
	return func(r *Service) {
		r.Course = c
	}
}

// WithSession ...
func WithSession(c SessionServicer) func(*Service) {
	return func(r *Service) {
		r.Session = c
	}
}

// WithStep ...
func WithStep(c StepServicer) func(*Service) {
	return func(r *Service) {
		r.Step = c
	}
}

// WithCourseProgress ...
func WithCourseProgress(c CourseProgressServicer) func(*Service) {
	return func(r *Service) {
		r.CourseProgress = c
	}
}

// WithStepNote ...

// WithCourseNote ...
func WithCourseNote(c CourseNoteServicer) func(*Service) {
	return func(r *Service) {
		r.CourseNote = c
	}
}

// WithStepProgress ...
func WithStepProgress(c StepProgressServicer) func(*Service) {
	return func(r *Service) {
		r.StepProgress = c
	}
}

// WithStepNote ...
func WithStepNote(c StepNoteServicer) func(*Service) {
	return func(r *Service) {
		r.StepNote = c
	}
}

// WithTimemap ...
func WithTimemap(c TimemapServicer) func(*Service) {
	return func(r *Service) {
		r.Timemap = c
	}
}
