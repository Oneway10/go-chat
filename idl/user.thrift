include "base.thrift"

struct LoginRequest {
    1: optional string name
    2: optional string password
}

struct LoginResponse {
    1: optional UserInfo userInfo
    2: optional string token

    255: optional base.BaseResp BaseResp
}

struct RegisterRequest {
    1: required string name
    2: required string password
    3: required string confirmPassword
}

struct RegisterResponse {
    1: optional string token

    255: optional base.BaseResp BaseResp
}

struct GetUserInfoRequest {
    1: required i64 id
}

struct GetUserInfoResponse {
    1: optional UserInfo userInfo

    255: optional base.BaseResp BaseResp
}

struct UserInfo {
    1: optional i64 id
    2: optional string name
    3: optional string avatar
    4: optional string phone
    5: optional string email
    6: optional string descrition
}

service UserService {
//    LoginResponse Login(1: LoginRequest req) (api.get="/login")
    RegisterResponse Register(1: RegisterRequest req) (api.post="/register")
    GetUserInfoResponse GetUserInfo(1: GetUserInfoRequest req) (api.get="/auth/user/get")
}