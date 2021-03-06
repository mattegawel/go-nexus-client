package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoleRead(t *testing.T) {
	client := getTestClient()
	role, err := client.RoleRead("nx-admin")

	assert.Nil(t, err)
	assert.NotNil(t, role)
	assert.Equal(t, role.ID, "nx-admin")
	assert.Equal(t, role.Name, "nx-admin")
	assert.Equal(t, 1, len(role.Privileges))
	assert.Equal(t, "nx-all", role.Privileges[0])
	assert.Equal(t, 0, len(role.Roles))

}

func TestRoleCreateReadUpdateDelete(t *testing.T) {
	client := getTestClient()
	testRole := testRole("test-role", "test-role-name", "test-role-description", []string{"nx-all"}, []string{"nx-admin"})

	// Create
	err := client.RoleCreate(*testRole)
	assert.Nil(t, err)

	if err != nil {
		// Read
		createdRole, err := client.RoleRead(testRole.ID)
		assert.Nil(t, err)
		assert.NotNil(t, createdRole)
		assert.Equal(t, testRole.ID, createdRole.ID)
		assert.Equal(t, testRole.Name, createdRole.Name)
		assert.Equal(t, testRole.Description, createdRole.Description)
		assert.Equal(t, len(testRole.Privileges), len(createdRole.Privileges))
		assert.Equal(t, len(testRole.Roles), len(createdRole.Roles))

		if createdRole != nil {
			createdRole.Description = "changed"
			createdRole.Name = "changed"
			createdRole.Privileges = []string{"nx-repository-view-*-*-*"}
			createdRole.Roles = []string{"nx-anonymous"}

			// Update
			err = client.RoleUpdate(createdRole.ID, *createdRole)
			assert.Nil(t, err)

			updatedRole, err := client.RoleRead(createdRole.ID)
			assert.Nil(t, err)
			assert.NotNil(t, updatedRole)
			assert.Equal(t, "changed", updatedRole.Description)
			assert.Equal(t, "changed", updatedRole.Name)
			assert.Equal(t, []string{"nx-repository-view-*-*-*"}, updatedRole.Privileges)
			assert.Equal(t, []string{"nx-anonymous"}, updatedRole.Roles)
		}

		// Delete
		err = client.RoleDelete(createdRole.ID)
		assert.Nil(t, err)

		role, err := client.RoleRead(createdRole.ID)
		assert.Nil(t, err)
		assert.Nil(t, role)
	}
}

func testRole(id, name, description string, privileges []string, roles []string) *Role {
	return &Role{
		ID:          id,
		Name:        name,
		Description: description,
		Privileges:  privileges,
		Roles:       roles,
	}
}
