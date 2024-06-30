# TOKEN-VERIFIER

This service aims to provide two features:
1. Verify the token's signature using an RSA public key
2. Verify the "aud" claim of the token

## Pre-requisites

A Public Key obtained from the Private Key used to sign the token

## Installation

To compile this project:
1. At the project root, run "go build token-verifier"
    1. This will build the runnable "token-verifier"
2. Run "token-verifier"


## Use

This service is meant to be used by all USS to verify if the Access token they are receiving is valid or not

### Env

Configuration settings are set in the .env file

Some key configurations are:
- RSA_PUBLIC_KEY_FILE is the path to the RSA public key file

