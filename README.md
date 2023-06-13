# Api-Contract
A clear, easy to use and fully automated api is the basic for a microservice architecture.

[Protobuf](https://protobuf.dev/) offers a way to define an api with the following benefits:
- It is binary.
- It has clear rules for backward compatibility and breaking changes.
- It also generated very small byte representation against a hard-defined contract.

For remote procedure calls the usage of [Grpc](https://grpc.io/) is preferred.

# Motivation
Having one unique version number for a state of the API makes it easier to use. You only have to look at the version numbers to find out what you need. The version number has to follow the [semantic versioning](https://semver.org/)

Fetching this version must be fully automated.

# How to do feature branches
1. Increase the version number in the corresponding pipeline
2. Develop your API and check it in if you are ready. A prerelease version of your API is built then.
3. Use this API to verify that it is usable. 'API first' does not mean that you have to release the api before you evaluate it for the first time.
4. Generate all code that needs to be in the repository. Currently, C# is built using the pipeline and Golang has to be checked in.
5. Merge the code back to the master using a pull request.

# Scope of the API projects
Each first level folder (protos/kinnekode/<nameOfProject> is a project.

Each of the projects in this repository has its own lifecycle and version number.

There is a special project called 'protobuf' in the repository. 'protobuf' is used in other projects.

## Why one repository for all?
1. To compile using protoc you need the proto files in a directory. Otherwise, it does not compile.
2. To include the generated code you have three options. 
  1. Include the generated file of an included package to your new package. That may cause conflicts if you use api from different projects in your code. I have not tested that.
  2. You include the package as a reference (works at least for C# and for Go). As a downside, you have to modify the package in different feature branches. You also have to release the version of the used package first to use it in the other package.
  3. You include nothing and the user of the API has to find out which additional package in which version is needed.
3. The projects (at least 'protobuf') also include implemented helpers to deal with the types. To include this code the api contract or the user of the api contract has to reference the package.

### Conclusion
The easiest way to use the api contract project is to include the package inside packages.

The files will stay in one repository as all options to copy files between different repositories are too complicated.
