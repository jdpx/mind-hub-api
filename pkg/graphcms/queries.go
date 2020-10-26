package graphcms

// GetAllCoursesQuery ...
const GetAllCoursesQuery = `{
  courses {
    id
    title
    description
  }
}`

// GetCourseByID ...
const GetCourseByID = `
  query Course($id: ID) {
      course(where: { id: $id }) {
          id
          title
          description
      }
  }
`

// GetCourseSessionsQuery ...
const GetCourseSessionsQuery = `{
  sessions {
    id
    title
    description
  }
}`

// GetSessionsByCourseID ...
const GetSessionsByCourseID = `query sessions($id: ID){
  sessions(where: { course: { id: $id } }) {
    id
    title
    description
  }
}
`

// GetSessionByID ...
const GetSessionByID = `query Session($id: ID) {
    session(where: { id: $id }) {
        id
        title
        description

        steps {
            id
            title
            description
            type
            videoUrl
            audioUrl
            question
        }

        course {
            id
            title
        }
    }
}`
