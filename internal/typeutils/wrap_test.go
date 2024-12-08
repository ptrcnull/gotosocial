// GoToSocial
// Copyright (C) GoToSocial Authors admin@gotosocial.org
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package typeutils_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/superseriousbusiness/gotosocial/internal/ap"
	"github.com/superseriousbusiness/gotosocial/internal/typeutils"
)

type WrapTestSuite struct {
	TypeUtilsTestSuite
}

func (suite *WrapTestSuite) TestWrapNoteInCreateIRIOnly() {
	testStatus := suite.testStatuses["local_account_1_status_1"]

	note, err := suite.typeconverter.StatusToAS(context.Background(), testStatus)
	suite.NoError(err)

	create := typeutils.WrapStatusableInCreate(note, true)
	suite.NoError(err)
	suite.NotNil(create)

	createI, err := ap.Serialize(create)
	suite.NoError(err)

	bytes, err := json.MarshalIndent(createI, "", "  ")
	suite.NoError(err)

	suite.Equal(`{
  "@context": "https://www.w3.org/ns/activitystreams",
  "actor": "http://localhost:8080/users/the_mighty_zork",
  "cc": "http://localhost:8080/users/the_mighty_zork/followers",
  "id": "http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/activity#Create",
  "object": "http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY",
  "published": "2021-10-20T12:40:37+02:00",
  "to": "https://www.w3.org/ns/activitystreams#Public",
  "type": "Create"
}`, string(bytes))
}

func (suite *WrapTestSuite) TestWrapNoteInCreate() {
	testStatus := suite.testStatuses["local_account_1_status_1"]

	note, err := suite.typeconverter.StatusToAS(context.Background(), testStatus)
	suite.NoError(err)

	create := typeutils.WrapStatusableInCreate(note, false)
	suite.NoError(err)
	suite.NotNil(create)

	createI, err := ap.Serialize(create)
	suite.NoError(err)

	bytes, err := json.MarshalIndent(createI, "", "  ")
	suite.NoError(err)

	suite.Equal(`{
  "@context": [
    "https://gotosocial.org/ns",
    "https://www.w3.org/ns/activitystreams",
    {
      "sensitive": "as:sensitive"
    }
  ],
  "actor": "http://localhost:8080/users/the_mighty_zork",
  "cc": "http://localhost:8080/users/the_mighty_zork/followers",
  "id": "http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/activity#Create",
  "object": {
    "attachment": [],
    "attributedTo": "http://localhost:8080/users/the_mighty_zork",
    "cc": "http://localhost:8080/users/the_mighty_zork/followers",
    "content": "hello everyone!",
    "contentMap": {
      "en": "hello everyone!"
    },
    "id": "http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY",
    "interactionPolicy": {
      "canAnnounce": {
        "always": [
          "https://www.w3.org/ns/activitystreams#Public"
        ],
        "approvalRequired": []
      },
      "canLike": {
        "always": [
          "https://www.w3.org/ns/activitystreams#Public"
        ],
        "approvalRequired": []
      },
      "canReply": {
        "always": [
          "https://www.w3.org/ns/activitystreams#Public"
        ],
        "approvalRequired": []
      }
    },
    "published": "2021-10-20T12:40:37+02:00",
    "replies": {
      "first": {
        "id": "http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies?page=true",
        "next": "http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies?only_other_accounts=false\u0026page=true",
        "partOf": "http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies",
        "type": "CollectionPage"
      },
      "id": "http://localhost:8080/users/the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY/replies",
      "type": "Collection"
    },
    "sensitive": true,
    "summary": "introduction post",
    "tag": [],
    "to": "https://www.w3.org/ns/activitystreams#Public",
    "type": "Note",
    "updated": "2021-10-20T12:40:37+02:00",
    "url": "http://localhost:8080/@the_mighty_zork/statuses/01F8MHAMCHF6Y650WCRSCP4WMY"
  },
  "published": "2021-10-20T12:40:37+02:00",
  "to": "https://www.w3.org/ns/activitystreams#Public",
  "type": "Create"
}`, string(bytes))
}

func TestWrapTestSuite(t *testing.T) {
	suite.Run(t, new(WrapTestSuite))
}
