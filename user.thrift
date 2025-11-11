// 命名空间：生成 Go 代码时的包路径（根据实际项目调整）
namespace go user

// 登录请求参数（对应 UserLoginRequestDTO）
struct UserLoginRequest {
    1: optional i64 userLoginId,  // 可选：用户登录ID（对应 DTO 的 UserLoginId）
    2: required string userName,  // 必选：用户名（对应 DTO 的 UserName，非空）
    3: required string password   // 必选：密码（对应 DTO 的 Password，非空）
}

// 用户响应信息（对应 UserResponseDTO）
struct UserResponse {
    1: optional i64 userLoginId,  // 用户登录ID
    2: optional string userName,  // 用户名
    3: optional string password,  // 密码（注意：实际场景可能不返回密码，此处仅按DTO映射）
    4: optional i64 userInfoId,   // 用户信息ID
    5: optional string realName,  // 真实姓名
    6: optional string phone,     // 手机号
    7: optional string email,     // 邮箱
    8: optional i8 gender,        // 性别（0：未知，1：男，2：女）
    9: optional string birthday,  // 出生日期（日期字符串，如 "2000-01-01"）
    10: optional string createTime, // 创建时间（时间字符串，如 "2023-01-01 12:00:00"）
    11: optional string updateTime  // 更新时间（同上）
}

// 登录业务异常（覆盖服务端可能返回的错误，如参数为空、密码错误等）
exception LoginException {
    1: required string message,  // 错误详情（如 "用户名不能为空"）
    2: optional i32 code         // 错误码（可选，用于客户端快速判断错误类型）
}

// 登录服务接口（对应 Login 接口）
service UserService {
    // 登录方法：传入请求参数，返回用户信息；失败时抛出 LoginException
    UserResponse login(1: UserLoginRequest request) throws (1: LoginException ex)
}