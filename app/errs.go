package app

type Err struct {
    Message string 
    Code int
}

func New(message string, code int) Err {
    return Err{
        Message: message, 
        Code: code,
    }
}

func Convert(err error) Err {
    if err == nil {
        return Err{
            Message: "Catch Error but no message in error.",
            Code: 999,
        }
    }
    return Err{
        Message: err.Error(),
        Code: 500,
    }
}


var (
    ErrUserNotFound = New("User not found", 404)
)