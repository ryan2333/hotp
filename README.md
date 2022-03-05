# hotp
#### build
```
GOOS=linux go build -o otp .
```

#### get secret
- use stream or charles get secret

#### get token
```
chmod +x otp
./otp -secret xxx
```
