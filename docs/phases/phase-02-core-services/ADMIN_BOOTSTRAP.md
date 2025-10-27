# Super Admin Bootstrap Strategy

**Document Version**: 1.0  
**Last Updated**: 2025-10-27  
**Status**: Implementation Required

---

## Overview

This document outlines the security strategy for creating the initial super admin user in the Angidi e-commerce platform. Proper admin bootstrap is critical for security - it must be secure, auditable, and prevent unauthorized admin access.

---

## Security Requirements

1. **No Default Admin Credentials**: Never ship with hardcoded admin credentials
2. **One-Time Bootstrap**: Admin creation should only work once (when no admin exists)
3. **Environment-Based Secrets**: Use environment variables for initial admin credentials
4. **Strong Password Requirements**: Enforce password complexity for admin accounts
5. **Audit Logging**: Log all admin creation attempts
6. **Secure Transport**: Admin credentials must never be transmitted in plain text
7. **Time-Limited**: Bootstrap endpoint should be disabled after first admin is created

---

## Recommended Approach: Environment-Based Bootstrap

### Strategy

The system will create an initial admin user on first startup using environment variables. This is secure because:

1. **No API Exposure**: Admin is created automatically, no public endpoint
2. **Environment Control**: Only someone with server access can set credentials
3. **One-Time Operation**: Only creates admin if none exists
4. **Production Ready**: Works in containerized/cloud environments

### Implementation

#### Step 1: Environment Variables

```bash
# Required only for initial bootstrap
ADMIN_EMAIL=admin@yourdomain.com
ADMIN_PASSWORD=ChangeThisSecurePassword123!
ADMIN_NAME=System Administrator
```

#### Step 2: Bootstrap Function

```go
// internal/user/bootstrap.go

package user

import (
    "context"
    "os"
    
    "go.uber.org/zap"
)

// BootstrapAdmin creates the initial admin user if no admin exists
func (s *service) BootstrapAdmin(ctx context.Context) error {
    // Check if any admin user exists
    // Note: This requires adding a method to check for admin users
    hasAdmin, err := s.repo.HasAdmin(ctx)
    if err != nil {
        s.logger.Error("Failed to check for existing admin", zap.Error(err))
        return err
    }
    
    if hasAdmin {
        s.logger.Info("Admin user already exists, skipping bootstrap")
        return nil
    }
    
    // Get admin credentials from environment
    email := os.Getenv("ADMIN_EMAIL")
    password := os.Getenv("ADMIN_PASSWORD")
    name := os.Getenv("ADMIN_NAME")
    
    // If no credentials provided, skip bootstrap (development mode)
    if email == "" || password == "" {
        s.logger.Warn("No admin credentials provided in environment, skipping bootstrap")
        return nil
    }
    
    // Validate password strength
    if len(password) < 12 {
        s.logger.Error("Admin password does not meet minimum requirements (12 characters)")
        return errors.New("admin password too weak")
    }
    
    // Create admin user
    s.logger.Info("Creating initial admin user", zap.String("email", email))
    
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
    if err != nil {
        s.logger.Error("Failed to hash admin password", zap.Error(err))
        return err
    }
    
    admin := &User{
        ID:           uuid.New().String(),
        Email:        email,
        PasswordHash: string(hashedPassword),
        Name:         name,
        Role:         "admin",
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
    }
    
    if err := s.repo.Create(ctx, admin); err != nil {
        s.logger.Error("Failed to create admin user", zap.Error(err))
        return err
    }
    
    s.logger.Info("Initial admin user created successfully", 
        zap.String("admin_id", admin.ID),
        zap.String("email", email))
    
    // Clear environment variables for security
    os.Unsetenv("ADMIN_PASSWORD")
    
    return nil
}
```

#### Step 3: Repository Extension

```go
// Add to internal/user/repository.go

// HasAdmin checks if any admin user exists
func (r *InMemoryRepository) HasAdmin(ctx context.Context) (bool, error) {
    r.mutex.RLock()
    defer r.mutex.RUnlock()
    
    for _, user := range r.users {
        if user.Role == "admin" {
            return true, nil
        }
    }
    
    return false, nil
}
```

#### Step 4: Main.go Integration

```go
// In cmd/api/main.go, after initializing services:

// Bootstrap admin user if needed
if err := userService.BootstrapAdmin(context.Background()); err != nil {
    zapLogger.Fatal("Failed to bootstrap admin user", zap.Error(err))
}
```

---

## Alternative Approach: CLI Command (More Secure)

### Strategy

Provide a CLI command to create the admin user separately from the main application.

### Implementation

```go
// cmd/admin-create/main.go

package main

import (
    "context"
    "flag"
    "fmt"
    "log"
    
    "golang.org/x/term"
    // ... imports
)

func main() {
    email := flag.String("email", "", "Admin email address")
    name := flag.String("name", "Administrator", "Admin name")
    flag.Parse()
    
    if *email == "" {
        log.Fatal("Email is required")
    }
    
    // Prompt for password securely
    fmt.Print("Enter admin password: ")
    password, err := term.ReadPassword(0)
    if err != nil {
        log.Fatal("Failed to read password:", err)
    }
    fmt.Println()
    
    fmt.Print("Confirm password: ")
    confirm, err := term.ReadPassword(0)
    if err != nil {
        log.Fatal("Failed to read confirmation:", err)
    }
    fmt.Println()
    
    if string(password) != string(confirm) {
        log.Fatal("Passwords do not match")
    }
    
    // Create admin user...
    // (Implementation details)
    
    fmt.Println("Admin user created successfully")
}
```

**Usage**:
```bash
go run cmd/admin-create/main.go -email admin@example.com -name "Super Admin"
# Prompts for password securely
```

---

## Security Considerations

### ✅ What We Do Right

1. **No Default Credentials**: No hardcoded admin/admin
2. **Environment-Based**: Credentials come from secure environment
3. **One-Time Only**: Bootstrap only runs if no admin exists
4. **Password Hashing**: Always hash with bcrypt
5. **Audit Logging**: Log all admin operations
6. **Credential Cleanup**: Clear password from environment after use

### ⚠️ Security Warnings

1. **Environment Variable Exposure**: 
   - Environment variables can be visible to other processes
   - Use secrets management in production (Vault, AWS Secrets Manager)
   
2. **Container Logs**:
   - Never log the actual password
   - Be careful with environment variable logging

3. **Initial Password Change**:
   - Force admin to change password on first login (future enhancement)
   - Implement password expiration policies

### 🔒 Production Best Practices

1. **Use Secrets Management**:
   ```bash
   # Instead of plain env vars, use:
   ADMIN_PASSWORD=$(vault kv get -field=password secret/admin)
   ```

2. **Rotate Credentials Immediately**:
   - After bootstrap, change admin password through the UI
   - Revoke any tokens from bootstrap process

3. **Multi-Factor Authentication**:
   - Implement MFA for admin accounts (Phase 5+)

4. **Audit Trail**:
   - Log all admin actions with timestamps
   - Monitor for suspicious admin activity

---

## Comparison of Approaches

| Approach | Security | Ease of Use | Production Ready | Docker Friendly |
|----------|----------|-------------|------------------|-----------------|
| **Environment Bootstrap** | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| **CLI Command** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ |
| **Database Migration** | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| **Web UI (First Run)** | ⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐ |

**Recommendation**: Use **Environment Bootstrap** for Phase 2, with plans to add CLI command in Phase 3.

---

## Implementation Checklist

- [ ] Add `Role` field validation to ensure only "user" or "admin"
- [ ] Implement `HasAdmin()` method in repository
- [ ] Create `BootstrapAdmin()` method in user service
- [ ] Add admin bootstrap call in main.go
- [ ] Create environment variable documentation
- [ ] Add tests for admin bootstrap logic
- [ ] Document admin password requirements (min 12 chars)
- [ ] Add logging for bootstrap operations
- [ ] Implement password change endpoint (future)
- [ ] Add MFA for admin accounts (Phase 5+)

---

## Testing Strategy

### Unit Tests

```go
func TestBootstrapAdmin(t *testing.T) {
    tests := []struct {
        name          string
        existingAdmin bool
        envEmail      string
        envPassword   string
        wantErr       bool
    }{
        {
            name:          "creates admin when none exists",
            existingAdmin: false,
            envEmail:      "admin@test.com",
            envPassword:   "SecurePassword123!",
            wantErr:       false,
        },
        {
            name:          "skips when admin exists",
            existingAdmin: true,
            envEmail:      "admin@test.com",
            envPassword:   "SecurePassword123!",
            wantErr:       false,
        },
        {
            name:          "fails with weak password",
            existingAdmin: false,
            envEmail:      "admin@test.com",
            envPassword:   "weak",
            wantErr:       true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation...
        })
    }
}
```

---

## Migration Path

### Phase 2 (Current)
- Implement environment-based bootstrap
- Document in README and .env.example

### Phase 3 (Database)
- Move to database-backed persistence
- Add migration script for admin creation

### Phase 4 (Enhanced Security)
- Add CLI tool for admin management
- Implement password change requirement

### Phase 5 (Production Hardening)
- Add MFA for admin accounts
- Implement secrets management integration
- Add admin action audit log

---

## Documentation Updates Required

1. **README.md**: Add admin bootstrap section
2. **.env.example**: Add admin environment variables with warnings
3. **SECURITY.md**: Document admin security best practices
4. **DEPLOYMENT.md**: Add production deployment checklist

---

## Example: Production Deployment

### Docker Compose

```yaml
version: '3.8'
services:
  api:
    image: angidi-api:latest
    environment:
      # Admin bootstrap (only for initial setup)
      - ADMIN_EMAIL=admin@yourcompany.com
      - ADMIN_PASSWORD=${ADMIN_PASSWORD}  # From secrets
      - ADMIN_NAME=System Administrator
      # JWT secrets
      - JWT_SECRET=${JWT_SECRET}
    secrets:
      - admin_password
      - jwt_secret

secrets:
  admin_password:
    external: true
  jwt_secret:
    external: true
```

### Kubernetes

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: angidi-api
spec:
  containers:
  - name: api
    image: angidi-api:latest
    env:
    - name: ADMIN_EMAIL
      value: "admin@yourcompany.com"
    - name: ADMIN_PASSWORD
      valueFrom:
        secretKeyRef:
          name: admin-credentials
          key: password
    - name: ADMIN_NAME
      value: "System Administrator"
```

---

## Conclusion

**Recommended Implementation**:
1. Use environment-based bootstrap for Phase 2
2. Check for admin existence before creating
3. Enforce strong password requirements (min 12 characters)
4. Log all bootstrap operations
5. Clear sensitive environment variables after use
6. Document in README with security warnings

This approach balances security, usability, and production readiness while being Docker/Kubernetes friendly.

---

## References

- [OWASP Authentication Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html)
- [12-Factor App: Config](https://12factor.net/config)
- [Docker Secrets Management](https://docs.docker.com/engine/swarm/secrets/)
- [Kubernetes Secrets](https://kubernetes.io/docs/concepts/configuration/secret/)
