Automated door opener for Twilio
================================

## Configure

* Open main.go
* Reset all constant variables
* Save file

## Build

#### Test locally
   ```
    $ goapp serve [LOCAL_FOLDER]
   ```

#### Deploy
   ```
    $ goapp deploy -oauth -application [APP_ENGINE_APP_ID_HERE] [LOCAL_FOLDER]
   ```

## Setup Twilio

* Sign up for Twilio
* Get/buy a phone number
* For Voice, point the Request URL to your deployed app. Eg. http://[APP_ENGINE_APP_ID].appspot.com
