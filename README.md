Automated door opener for Twilio
================================

If you live in a safe area that also requires a code to get in, and you don't want to answer the phone to let people in, let them come in, this will just simply play a DTMF dial tone (9 or 6), and let them in, and then it will send you a SMS alert to let you know someone is at the door and on their way to your home.

Disclaimer: Not responsible for misuse or abuse from others or yourself. Use at your own risk.

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
