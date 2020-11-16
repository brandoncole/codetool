# CodeTool

This Golang CLI is intended to eventually become a suite of tools that makes it much more efficient to work with AWS CodeCommit, CodeBuild and CodePipeline.

## Usage

This tool assumes the AWS environment is configured in the current shell so it does not explicitly accept any credentials.  The recommended way is to use the AWS Config approach.

- See: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html.

### Use AWS Config

```bash
export AWS_PROFILE=<profile>
```

### Use AWS Credentials

```bash
export AWS_ACCESS_KEY_ID=<placeholder>
export AWS_SECRET_ACCESS_KEY=<placeholder>
export AWS_SESSION_TOKEN=<placeholder>
export AWS_DEFAULT_REGION=<placeholder>
```

## Cheat Sheet

```bash
# Display how much time is building vs. waiting for AWS
./codetool codebuild analyze
Time		Percent	Phase
28h56m59s	100	    TOTAL TIME
16h57m22s	58	    PROVISIONING
5h53m8s		20	    QUEUED
4h8m27s		14	    BUILD
38m34s		2	    POST_BUILD
23m25s		1	    INSTALL
15m22s		0	    DOWNLOAD_SOURCE
14m48s		0	    FINALIZING
11m13s		0	    PRE_BUILD
0s		    0	    SUBMITTED
0s		    0	    UPLOAD_ARTIFACTS
Total Builds: 353
```

## Roadmap

- Support a configuration file that developers can share
- Show a list of all PRs across multiple accounts
- Show a list of all PRs across multiple regions