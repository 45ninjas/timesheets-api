# Error
An error object comprises of the following fields.

- **Title** *string*
    The title of the error message, usually a extremely short description of what was being done for example: "updating shift"

- **Message** *string*
    The actual message of the error. On very rare occasions (like parsing json from users) this can be the exception message.

- **Reference** *string*
    Where you can find documentation about this error.