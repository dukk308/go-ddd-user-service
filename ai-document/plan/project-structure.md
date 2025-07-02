## Technology Stack

- **Framework:** Nestjs
- **Data Storage:** Typeorm with Postgres replication, Caching using redis sentinel
- **Authentication:** JWT (JSON Web Tokens)
- **Validation:** Class validator and custom validator in `src/decorators`
- **API Documentation:** Swagger

## Project Structure (Modular, Command-Query Separation Approach)

```
/src
  /modules/                      # Domain-specific modules
    /{module}/                   # Each domain module (user, configuration, etc.)
      /infrastructure/           # Infrastructure layer
        /controllers/            # HTTP API controllers
          *.controller.ts        # Controller with route definitions
        /persistence/            # Data access layer
          *.repository.ts        # abstract repository implementation
          /relational/           # Data access layer with relational database
            /entities/           # Typeorm entities
            /mappers/            # Mapper from data layer to bussiness layer
            /repositories/       # Data persistence repository implemenent
            *-persistence.module.ts # Module DI of relational persistence
      /domain/                   # Business layer
        /models/                  # Domain models layer
          *.ts                    # Model definitions and types
        /service/                 # Business logic: Service layer
          *.service.ts            # Service logic implement
        /commands/                # Command handlers layer: CQRS layer
          *.ts
        /queries/                 # Query handlers layer: CQRS layer
          *.ts
        /dtos/                    # Data object transfer from external or to external: DTO layer
          *.dto.ts
        /errors/                  # Bussiness error defination layer: Error layer
          *.{exception,error}.ts
      *.module.ts                # Module entry point and dependency injection
```

## Module description

**_auth_**: Contain authentication logic using jwt
**_configuration_**: Management config like json, theme, key-value
**_domain_**: Management domain like abc.com
**_domain-config_**: Management config per domain
**_permission_**: Management app permission
