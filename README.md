# OPS-CHALLENGE-G


**endpoints:**


/token: returns a token.

/health: returns 200 if healthy

/metrics: returns some basic information about the app.


## **required to deploy**: 

- [AWS CLI V2](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html). and configured creds
- [docker](https://docs.docker.com/engine/install/).
- access to an IAM role in my admin group *(shrug)*

## **requiredto work on src**: 

- [Go](https://golang.org/).


# infra:

/cfn cloudformation templates for hosting infra to be deployed outside of pipeline

/bin scripts with 0** prefix are non-pipeline scripts to deploy infra in /cfn 

bin/ 1** prefix for deploying/publishing


# deployment:


deployment is automated via github actions on pushes to this repo

**how:**

- I have configured an IAM user and administrator access group.

- 'configure AWS credentials' pipeline step allows github actions agents to assume a user in admin group mentioned above.



# manual deploy:


## step 0:

**check AWS resources exist outlined in /cfn (eg, ECR exists), can check CF stacks, query the outputs of them, run the scripts, whatever**

if not:

find bin/ scripts prefixed with 0. (eg 00-spin-up-ecr.sh)

to deploy cloud formation stacks, access to aws account or appropriate IAM user is needed and obviously won't be outlined in this repo because it's my own personal acc.But maybe I create an IAM user for you and add you to my admin group. 


then:

/bin 0**.sh scripts (scripts with 0 prefix)

these 0 steps are ordered, eg 00*.sh, 01*.sh etc.

if resources exist / once resource stacks are up:


## step 1:

deploy --> bin/1***.sh

run the scripts prefixed with 1 in order (e.g: 10*.sh before 11*.sh)

**these scripts can run 'localy' and in the pipeline. I have opted to put most of the work, for instance for building and pushing the images - getting commit sha, grabbing the ecr-name, rather than using github actions envs/ and output features - in the rare case I am unable to run the pipeline, or for whatever reason need to run all of these steps myself.**