namespace go api

struct QueryStudentRequest {
    1: string Id (api.query="id", api.vd="$!='0'; msg: 'id cannot be 0'");
}

struct QueryStudentResponse {
    1: string Id;
    2: string Name;
    3: string Gender;
    4: string Major;
    5: string Msg;
}

struct InsertStudentRequest {
    1: string Id (api.form="id");
    2: string Name (api.form="name");
    3: string Major (api.form="major");
    4: string Gender (api.form="gender");
}

struct InsertStudentResponse {
    1: bool Ok;
    2: string Msg;
}

service StudentApi{
    QueryStudentResponse queryStudent(1: QueryStudentRequest req) (api.get="/studentApi/query");
    InsertStudentResponse insertStudent(1: InsertStudentRequest req) (api.post="/studentApi/insert");
}