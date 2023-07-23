namespace go grader

// struct Student {
//   1: string name,
//   2: i32 id,
//   3: string major
//   4: string gender
// }

// struct Exam {
//   1: i32 id,
//   2: string name,
//   3: i32 credits,
//   4: list<i32> scores
// }

struct InsertGradeRequest {
  1: i32 student_id,
  2: list<string> grades
}

struct InsertGradeResponse {
    1: bool Ok;
    2: string Msg;
}

struct GetCapRequest {
  1: i32 student_id,
}

struct GetCapResponse {
  1: i32 id, 
  2: string name,
  3: string major,
  4: string gender,
  5: double Cap
}

service UniversityGrades {
//   void addStudent(1: Student student),
// post req to get input all the grades for a student 
  GetCapResponse getGrades(1: GetCapRequest req) (api.get="/UniversityGrades/query");
  InsertGradeResponse insertGrades(1: InsertGradeRequest req) (api.post="/UniversityGrades/insert");
  // void addExam(1: Exam exam),
  // void addGrade(1: Grade grade),
  // list<Grade> getGradesForStudent(1: i32 student_id),
  // list<Grade> getGradesForExam(1: i32 exam_id),
  // double getAverageGradeForStudent(1: i32 student_id),
  // double getAverageGradeForExam(1: i32 exam_id)
}