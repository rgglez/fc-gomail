# fc-gomail

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause) ![GitHub release (latest by date)](https://img.shields.io/github/downloads/rgglez/fc-gomail/latest/total)

## Summary

This event handler sends an email using the [Gomail](https://github.com/go-gomail/gomail) module. It's written in [Go](https://go.dev/) for the Aliyun FC FaaS service, and is intended to be triggered by Aliyun MNS message service, where a message must be published by another program. Use cases could be a notification service or a login confirmation.

The message must have this structure, in JSON:

```
{
   "From":"<from@address>",
   "To":"<to@address>",
   "Subject":"The subject of the message",
   "Body":"The body of the message"
}
```javascript

