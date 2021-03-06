package gluecf

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/panther-labs/panther/tools/cfngen"
)

func TestDatabase(t *testing.T) {
	dbName := "db1"
	resources := map[string]interface{}{
		dbName: NewDatabase("12345", dbName, "Test db"),
	}

	cfTemplate := cfngen.NewTemplate("Test template", nil, resources, nil)
	cf, err := cfTemplate.CloudFormation()
	require.NoError(t, err)

	const expectedFile = "testdata/db.template.json"
	// uncomment to write new expected file
	// require.NoError(t, ioutil.WriteFile(expectedFile, cf, 0644))

	expected, err := ioutil.ReadFile(expectedFile)
	require.NoError(t, err)
	assert.JSONEq(t, string(expected), string(cf))
}
