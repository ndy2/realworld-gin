# realworld-gin

- see https://realworld-docs.netlify.app/specifications/backend/endpoints/

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
    - [ ] on error (Error handling)
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