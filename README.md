# VerificationCodeService
A velocity service which creates and validates verification codes

## Functions
- Create -- creates a jwt with a userID
- Validate -- validates jwt and returns the userID

## Workings
When a user registers, a verification code is generated for their userID and this will be sent using EmailService to user's email. Then user will use that verification code and enter it into velocity and the auth server will use this service to validate that verification code and mark that user as account status of Active. As of now this is only used to verify user's email but in futute can be used to do different things.