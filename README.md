# Bookdrop
**Send local document to kindle using resend!**

# Pre-Reqs
This program uses Resend to email documents to a kindle (or other e-reader) of your choice.
* A free account is required
* A domain that you can prove ownership over in order to send external emails
* An api key

https://resend.com/overview


# Installation and Setup
Run this command to download and install `bookdrop`
`go get -u https://github.com/GianniBYoung/bookdrop`

If you do not have `go` installed on your system pre compiled binaries can soon be found in releases

## Configuration
* On first run the wizard will guide you through the initial setup and save a yaml configuration in `~/.config/bookdrop.yml`
* The api key can be specified in the following ways:
    * By setting the environmental variable `RESEND_API_KEY` (Preferred)
        * Can be set securely using the `pass program` - `export RESEND_API_KEY=$(pass RESEND_API_KEY)`
    * By setting the field `apiKey` in the yaml file (insecure)

# Usage
`bookdrop <path to file>`
