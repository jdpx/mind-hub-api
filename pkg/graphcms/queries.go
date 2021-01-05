package graphcms

const getAllCoursesQuery = `{
  courses {
    id
    title
    description
  }
}`

const getCourseByID = `
  query Course($id: ID) {
      course(where: { id: $id }) {
          id
          title
          description
      }
  }`

const getSessionsByCourseID = `query sessions($id: ID){
  sessions(where: { course: { id: $id } }) {
    id
    title
    description

    steps {
      id
      title
    }
  }
}`

const getSessionByID = `query Session($id: ID) {
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
            description
        }
    }
}`

const getStepByID = `query Step($id: ID) {
    step(where: { id: $id }) {
      id
      title
      description
      type
      videoUrl
      audioUrl
      question
    }
}`
