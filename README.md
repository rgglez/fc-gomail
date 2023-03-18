# fc-gomail

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause) ![GitHub all releases](https://img.shields.io/github/downloads/rgglez/fc-gomail/total) ![GitHub issues](https://img.shields.io/github/issues/rgglez/fc-gomail) ![GitHub commit activity](https://img.shields.io/github/commit-activity/y/rgglez/fc-gomail)

This event handler sends an email using the [Gomail](https://github.com/go-gomail/gomail) module, in response of a trigger originated from the Aliyun MNS message service, where a message must be published by another program. It's written in [Go](https://go.dev/) for the Aliyun FC FaaS service. Use cases could be a notification service or a login confirmation.

It's intended to run from a [custom Docker container](https://www.alibabacloud.com/help/en/function-compute/latest/create-a-function) sourced from Aliyun Container Service, so it implements a webserver on its own.

The message sent to Aliyun MNS must have this structure, in JSON:

```javascript
{
   "From":"<from@email.address>",
   "To":"<to@email.address>",
   "Subject":"The subject of the message",
   "Body":"The body of the message"
}
```

## Configuration

The configuration, in the case of the code in this repository, is based in a YAML file, which is loaded into a variable map. The structure of the relevant part of the YAML must be like this:

```yaml
SMTP:
  HOST: smtp.yourserver.com
  PORT: 2525
  LOGIN: <your login>
  PASSWORD: <your password>
```

## Notes

* A sample *Dockerfile* in provided, which you might adjust to your use case.
* A sample *Makefile* to compile the Go code is provided, which you might adjust to your use case.
* Aliyun refers to Alibaba Cloud Services, but this code can be easily adapted to run in AWS Lambda, for instance.

## License

Copyright (c) 2023, Rodolfo González González.

Read the LICENSE file.
