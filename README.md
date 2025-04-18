# How To Run
##### 1. Install Go depedency
##### 2. Using Docker
- Adjust .env POSTGRES_HOST to 'postgres'
```
docker-compose up --build
```
##### 3. Manual Run
- Adjust .env POSTGRES_HOST to 'localhost'
- Make sure postgres is running in local
```
go run cmd/http/main.go
```

# Feature
##### 1. User can mocked log in by input valid email and role
- Session will be stored in cookie
- User role
-- internal
-- fieldOfficer
-- investor

##### 2. User can view list of all users for testing purpose
##### 3. User can download uploaded file for testing purpose
##### 4. As a fieldOfficer
- Get list of borrowers
- Create borrower
- Delete borrower
- Propose a loan

##### 5. As an internal
- Approve a loan
- Disburse a loan
- Upload proof of picture
- Upload borrower agreement letter

##### 6. As an investor
- Invest to a loan

##### 7. Get list of loans
##### 8. Get loan detail

# Happy Flow
1. Login as fieldOfficer
-- email must be in valid format
-- email can be dummy but try to use real email when login as investor (to actually receive agreement letter)
-- login using same email multiple time with different role will update the role
2. Create a borrower
3. Propose a loan
4. Login as internal
5. Upload proof of picture
6. Approve a loan
7. Login as investor
8. Invest to a loan
9. (Optional) Login as another investor
10. (Optional) Invest to a loan
11. If loan principal amount is equal to invested amount, status will be updated to invested and email will be sent to all investors
12. Login as internal
13. Upload borrower agreement letter
14. Disburse a loan