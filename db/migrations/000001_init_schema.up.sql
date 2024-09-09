-- Enable the uuid-ossp extension for UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the Accounts table
CREATE TABLE "Accounts" (
                            "Id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                            "Username" VARCHAR(255) NOT NULL UNIQUE,
                            "PasswordHash" VARCHAR(255) NOT NULL,
                            "CreatedAt" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            "UpdatedAt" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Create the Roles table
CREATE TABLE "Roles" (
                         "Id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         "Name" VARCHAR(255) NOT NULL UNIQUE,
                         "CreatedAt" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create the Permissions table
CREATE TABLE "Permissions" (
                               "Id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                               "Name" VARCHAR(255) NOT NULL UNIQUE,
                               "Description" TEXT,
                               "CreatedAt" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create the UsersRoles table
CREATE TABLE "AccountsRoles" (
                              "Id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                              "AccountId" UUID NOT NULL,
                              "RoleId" UUID NOT NULL,
                              "CreatedAt" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              FOREIGN KEY ("AccountId") REFERENCES "Accounts" ("Id") ON DELETE CASCADE,
                              FOREIGN KEY ("RoleId") REFERENCES "Roles" ("Id") ON DELETE CASCADE
);

-- Create the RolesPermissions table
CREATE TABLE "RolesPermissions" (
                                    "Id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                    "RoleId" UUID NOT NULL,
                                    "PermissionId" UUID NOT NULL,
                                    "CreatedAt" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    FOREIGN KEY ("RoleId") REFERENCES "Roles" ("Id") ON DELETE CASCADE,
                                    FOREIGN KEY ("PermissionId") REFERENCES "Permissions" ("Id") ON DELETE CASCADE
);
