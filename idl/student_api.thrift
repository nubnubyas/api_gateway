namespace go api

struct QueryStudentRequest {
    1: string Num (api.query="num", api.vd="$!='0'; msg: 'num cannot be 0'");
}

struct QueryStudentResponse {
    1: string Num;
    2: string Name;
    3: string Gender;
    4: string Msg;
}

struct InsertStudentRequest {
    1: string Num (api.form="num");
    2: string Name (api.form="name");
    3: string Gender (api.form="gender");
}

struct InsertStudentResponse {
    1: bool Ok;
    2: string Msg;
}

service StudentApi{
    QueryStudentResponse queryStudent(1: QueryStudentRequest req) (api.get="student/query");
    InsertStudentResponse insertStudent(1: InsertStudentRequest req) (api.post="student/insert");
}