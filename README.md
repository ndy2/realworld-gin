# realworld-gin

![logo](resources/logo.png)

![coverage](https://raw.githubusercontent.com/ndy2/realworld-gin/badges/.badges/main/coverage.svg)

- 90% of the code is written by [Github Copilot](https://github.com/features/copilot)... I just modified it a bit.
- see https://realworld-docs.netlify.app/specifications/backend/endpoints/

**Built with**

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Copilot](https://img.shields.io/badge/Copilot-000000?style=for-the-badge&logo=github&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=for-the-badge&logo=mysql&logoColor=white)
![Zap](https://img.shields.io/badge/Zap-00ADD8?style=for-the-badge&logo=go&logoColor=white)

## Auth/User

- [x] Authentication
    - [x] Endpoint
    - [x] JWT
- [x] Register
    - [x] Endpoint
    - [x] Insert Profile to DB
- [x] Get current user
- [x] Update

## Profile

- [x] Get Profile
    - [x] Current User Profile (authenticated)
    - [x] Other User Profile (authenticated)
        - [x] check Followed by current user 
    - [x] arbitrary User Profile (unauthenticated)
- [x] Update User (Profile)

### Follow

- [ ] Follow User
- [ ] Unfollow User

## Article

- [ ] List Articles
- [ ] Feed Articles
- [ ] Get Article
- [ ] Create Article
- [ ] Update Article
- [ ] Delete Article

## Comments

- [ ] Add Comments to an Article
- [ ] Get Comments from an Article
- [ ] Delete Comment

## Favorites

- [ ] Favorite Article
- [ ] Unfavorite Article

## Tags

- [ ] Get Tags

## Common

- [x] Authentication
    - [x] Required
    - [x] Optional
- [ ] Json Marshal/Unmarshal
    - [x] on success (single response)
    - [ ] on success (list response)
    - [x] on error (Error handling)
- [ ] Input Validation
- [x] Package structure - I'm satisfied with the current structure!
- [x] Configuration
    - [x] internal/config
    - [ ] ~~more sophisticated configuration~~ - SKIP FOR NOW
- [x] Concurrent Logic with Goroutine (w/ ErrGroup)
- [ ] Infrastructure (Database) 
    - [x] Database Setup (w/ MySQL, Docker)
    - [x] ~~Connection Pool~~ - Go has built-in connection pooling
    - [ ] ~~Transaction~~ - SKIP FOR NOW
    - [ ] Use SQLX
- [ ] Testing
    - [x] Unit Tests Samples (w/ go-mock, go-cmp)
        - [x] routes
        - [x] logic
        - [x] repo (w/ sqlmock)
    - [x] Github Actions - Run tests on push
        - [x] Unit Tests
        - [ ] Integration Test (w/ newman/docker-compose)
- [ ] ~~Documentation~~ - SKIP FOR NOW
  - [ ] about Project (w/ mkdocs-material)
  - [ ] ~~about API (w/ Swagger)~~ - SKIP FOR NOW
- [x] Logging (w/ Zap)

## Scripts

### Create go mock for `xxxapp#Logic` interfaces

```bash
packages=("auth" "profile" "user")
for pkg in "${packages[@]}"; do
    mockgen -source=internal/${pkg}/app/logic.go -destination=internal/${pkg}/app/logic_mock.go -package=app
done
```
