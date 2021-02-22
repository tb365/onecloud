// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aws

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudtrail"

	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
	"yunion.io/x/pkg/errors"

	"yunion.io/x/onecloud/pkg/cloudprovider"
)

type SEventResource struct {
	// The name of the resource referenced by the event returned. These are user-created
	// names whose values will depend on the environment. For example, the resource
	// name might be "auto-scaling-test-group" for an Auto Scaling Group or "i-1234567"
	// for an EC2 Instance.
	ResourceName string `type:"string"`

	// The type of a resource referenced by the event returned. When the resource
	// type cannot be determined, null is returned. Some examples of resource types
	// are: Instance for EC2, Trail for CloudTrail, DBInstance for RDS, and AccessKey
	// for IAM. To learn more about how to look up and filter events by the resource
	// types supported for a service, see Filtering CloudTrail Events (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/view-cloudtrail-events-console.html#filtering-cloudtrail-events).
	ResourceType string `type:"string"`
}

type SEvent struct {
	// The AWS access key ID that was used to sign the request. If the request was
	// made with temporary security credentials, this is the access key ID of the
	// temporary credentials.
	AccessKeyId string `type:"string"`

	// A JSON string that contains a representation of the event returned.
	CloudTrailEvent string `type:"string"`

	// The CloudTrail ID of the event returned.
	EventId string `type:"string"`

	// The name of the event returned.
	EventName string `type:"string"`

	// The AWS service that the request was made to.
	EventSource string `type:"string"`

	// The date and time of the event returned.
	EventTime time.Time `type:"timestamp"`

	// Information about whether the event is a write event or a read event.
	ReadOnly string `type:"string"`

	// A list of resources referenced by the event returned.
	Resources []SEventResource `type:"list"`

	// A user name or role name of the requester that called the API in the event
	// returned.
	Username string `type:"string"`
}

func (self *SEvent) GetName() string {
	return self.EventName
}

func (self *SEvent) GetService() string {
	return self.EventSource
}

func (self *SEvent) GetAction() string {
	return self.EventName
}

func (self *SEvent) GetResourceType() string {
	return self.EventSource
}

func (self *SEvent) GetRequestId() string {
	return self.EventId
}

func (self *SEvent) GetRequest() jsonutils.JSONObject {
	obj, _ := jsonutils.Parse([]byte(self.CloudTrailEvent))
	return obj
}

func (self *SEvent) GetAccount() string {
	return fmt.Sprintf("%s(%s)", self.AccessKeyId, self.Username)
}

func (self *SEvent) IsSuccess() bool {
	return !strings.Contains(self.CloudTrailEvent, "errorMessage")
}

func (self *SEvent) GetCreatedAt() time.Time {
	return self.EventTime
}

func (self *SAwsClient) LookupEvents(start, end time.Time) ([]SEvent, error) {
	s, err := self.getAwsCloudtrailSession()
	if err != nil {
		return nil, errors.Wrapf(err, "getAwsCloudtrailSession")
	}
	client := cloudtrail.New(s)
	input := &cloudtrail.LookupEventsInput{}
	if !start.IsZero() {
		input.SetStartTime(start)
	}
	if !end.IsZero() {
		input.SetEndTime(end)
	}
	events := []SEvent{}
	nextToken := ""
	for {
		if len(nextToken) > 0 {
			input.SetNextToken(nextToken)
		}
		output, err := client.LookupEvents(input)
		if err != nil {
			return nil, errors.Wrapf(err, "LookupEvents")
		}
		for i := range output.Events {
			err := FillZero(output.Events[i])
			if err != nil {
				return nil, errors.Wrapf(err, "FillZero")
			}
			event := SEvent{}
			err = jsonutils.Update(&event, jsonutils.Marshal(output.Events[i]))
			if err != nil {
				return nil, errors.Wrapf(err, "jsonutils.Update")
			}
			events = append(events, event)
		}
		nextToken = ""
		if output.NextToken != nil {
			nextToken = *output.NextToken
		}
		if len(nextToken) == 0 {
			break
		}
	}
	return events, nil
}

func (self *SRegion) GetICloudEvents(start time.Time, end time.Time, withReadEvent bool) ([]cloudprovider.ICloudEvent, error) {
	var events []SEvent
	var err error
	for {
		events, err = self.client.LookupEvents(start, end)
		if err != nil {
			if strings.Contains(err.Error(), "ThrottlingException") {
				log.Warningf("LookupEvents ThrottlingException, try after 3 seconds")
				time.Sleep(time.Second * 3)
				continue
			}
			return nil, errors.Wrapf(err, "LookupEvents(%s, %s)", start, end)
		}
		break
	}
	ret := []cloudprovider.ICloudEvent{}
	for i := range events {
		if !withReadEvent && events[i].ReadOnly == "true" {
			continue
		}
		ret = append(ret, &events[i])
	}
	return ret, nil
}
