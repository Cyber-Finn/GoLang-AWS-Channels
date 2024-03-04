# GoLang-AWS-Channels
A very basic demo of how to use Go Channels to pass data between different parts of your go routine. This is exceptionally important in AWS, as it allows you to pass channels/objects between different lambdas and packages, etc.

## Why use Go for AWS?
Go is a natively supported language for AWS, and therefore operates faster than languages such as C/C++, because there is basically no "cold start" in the native environment.

## Isn't Go no longer supported?
While it's true that the first version of Go is no longer supported by the AWS SDK (Go v1), the second (and latest) version is still supported, and will likely continue being supported, as important cloud tools - like Terraform (by HashiCorp) - are actually written (at least partially) in Go.

## Usage:
1. Clone and change whatever you need
2. Upload the new code to your organisation's private repository
3. Create an import within the relevant go function you're writing, to link that package (created in Step 2) to your current project, to use the methods within.
