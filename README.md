# fc-gomail

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause) ![GitHub all releases](https://img.shields.io/github/downloads/rgglez/fc-gomail/total) ![GitHub issues](https://img.shields.io/github/issues/rgglez/fc-gomail) ![GitHub commit activity](https://img.shields.io/github/commit-activity/y/rgglez/fc-gomail)

This event handler sends an email using the [Gomail](https://github.com/go-gomail/gomail) module. It's written in [Go](https://go.dev/) for the Aliyun FC FaaS service, and is intended to be triggered by Aliyun MNS message service, where a message must be published by another program. Use cases could be a notification service or a login confirmation.

It's intended to run from a [custom Docker container](https://www.alibabacloud.com/help/en/function-compute/latest/create-a-function) sourced from Aliyun Container Service, so it implements a webserver on its own.

The message sent to Aliyun MNS must have this structure, in JSON:

```
{
   "From":"<from@email.address>",
   "To":"<to@email.address>",
   "Subject":"The subject of the message",
   "Body":"The body of the message"
}
```

