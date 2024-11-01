# realworld-gin

![logo](logo.png)

![coverage](https://raw.githubusercontent.com/ndy2/realworld-gin/badges/.badges/main/coverage.svg)

- 90% of the code is written by [Github Copilot](https://github.com/features/copilot)... I just modified it a bit.
- see https://realworld-docs.netlify.app/specifications/backend/endpoints/

**Built with**

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![ChatGPT](https://img.shields.io/badge/chatGPT-74aa9c?style=for-the-badge&logo=openai&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-000000?style=for-the-badge&logo=go&logoColor=white)
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

- [ ] Get Profile
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

- [ ] Authentication
    - [x] Required
    - [ ] Optional
- [ ] Json Marshal/Unmarshal
    - [x] on success (single response)
    - [ ] on success (list response)
    - [x] on error (Error handling)
- [ ] Configuration
- [x] Database (MySQL)
    - [ ] Setup
    - [ ] Connection
- [ ] Testing
    - [ ] Unit Tests Samples
        - [ ] routes
        - [ ] logic
        - [ ] repo
    - [ ] Github Actions - Run tests on push
        - [ ] Unit Tests
        - [ ] Integration Test (newman/docker-compose)
- [ ] Documentation
- [x] Logging