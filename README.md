# es-intentmanagement

Example of using the Dialogflow ES API to export and import Intents for external review.

Note: Export is the only functioning command.

```
Usage of es-intentmgmt
  -action string
        import | export (default "export")
  -language string
        language code (for multilingual Agents) (default "en")
  -location string
        ES Agent location (default "global")
  -project string
        GCP Project ID
```

## Requirements

* Google service account with Dialogflow API Admin Role as env var `GOOGLE_APPLICATION_CREDENTIALS`
* `project` flag is required, set to GCP Project ID

For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.


## Example

Given a GCP Project named "ghctest003" with a Dialogflow ES Agent, and an appropriate service account JSON key in the env var as above:

```
$ es-intentmgmt --project ghctest003
2021/07/02 14:23:23 es intent management
2021/07/02 14:23:23 exporting all intents
2021/07/02 14:23:23 Getting 'Default Welcome Intent' (projects/ghctest003/locations/global/agent/intents/1b705f10-9db5-4ce4-b5f4-ab808ec932ea) ...
2021/07/02 14:23:23 Training Phrases: 16
2021/07/02 14:23:23 Getting 'Default Fallback Intent' (projects/ghctest003/locations/global/agent/intents/902e4e9b-47d1-4d15-a7c2-a261ff40b69a) ...
2021/07/02 14:23:23 No training phrases for 'Default Fallback Intent'
2021/07/02 14:23:23 Getting 'Talk to an agent' (projects/ghctest003/locations/global/agent/intents/c763df69-5fae-4414-8072-926cd57f1ecf) ...
2021/07/02 14:23:23 Training Phrases: 12
2021/07/02 14:23:23 Getting 'dealer-locator' (projects/ghctest003/locations/global/agent/intents/e964eb8f-591e-42df-8c48-2cffae813d2a) ...
2021/07/02 14:23:23 Training Phrases: 3
$ ls -1a *.csv
Default_Welcome_Intent_en.csv
Talk_to_an_agent_en.csv
dealer-locator_en.csv
```