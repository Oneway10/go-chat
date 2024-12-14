include "base.thrift"

struct LoginRequest {
    1: optional string name (api.query="name")
    2: optional string password
}

struct LoginResponse {
    1: optional string name
    2: optional string message

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

service UserService {
    LoginResponse Login(1: LoginRequest loginRequest) (api.get="/login")
    RegisterResponse Register(1: RegisterRequest registerRequest) (api.post="/register")
}