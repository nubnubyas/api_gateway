namespace go userAuth

struct UserProfile {
    1: required string firstName,
    2: required string lastName,
    3: optional i32 age,
    4: optional string email,
}

service AuthService {
    string authenticate(1: string username, 2: string password),
    bool authorize(1: string token),
}

service UserProfileService {
    UserProfile getProfile(1: string username),
    bool updateProfile(1: string username, 2: UserProfile profile),
}