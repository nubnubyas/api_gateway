namespace go grader

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
  GetCapResponse getGrades(1: GetCapRequest req) (api.get="/UniversityGrades/query");
  InsertGradeResponse insertGrades(1: InsertGradeRequest req) (api.post="/UniversityGrades/insert");
}