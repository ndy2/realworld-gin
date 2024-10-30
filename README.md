# realworld-gin

## Auth/User

- [ ] Authentication
   - [x] Endpoint
   - [ ] JWT
   - [ ] Insert Profile to DB (using goroutine?)
- [ ] Register
    - [x] Endpoint
    - [ ] Profile
- [ ] Get current user
- [ ] Update

## Profile

- [ ] Get
- [ ] Update

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

- [ ] Json Marshal/Unmarshal
  - [x] on success (single response)
  - [ ] on success (list response)
  - [ ] on error (Error handling)
- [ ] Configuration
- [x] Database (MySQL)
- [ ] Testing
    - [ ] Unit Tests Samples
        - [ ] routes
        - [ ] logic
        - [ ] repo
    - [ ] Github Actions - Run tests on push 
        - [ ] Unit Tests
        - [ ] Integration Test (newman/docker-compose)
- [ ] Documentation
