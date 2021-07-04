# es-intentmgmt

Example of using the Dialogflow ES API to export and import the Training Phrases used in Intents for external review.

See [Releases](https://github.com/ghchinoy/es-intentmgmt/releases) for Win, Linux, OS X binaries.

*Note*: Export is the only functioning command.

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

Given Dialogflow ES Agent, such as the [COVID-19 RRVA Agent](https://github.com/GoogleCloudPlatform/covid19-rapid-response-demo), the name of the GCP Project, "ghctest003", and an appropriate service account JSON key in the env var as above, a series of csv files will be exported, ready for viewing either directly or in Google Sheets.

```
$ export GOOGLE_APPLICATION_CREDENTIALS=<my file path to the service account.json>
$ es-intentmgmt --project ghctest003
2021/07/04 12:44:33 es intent management
2021/07/04 12:44:33 exporting all intents
2021/07/04 12:44:34 getting 'coronavirus.symptoms' (projects/ghctest003/locations/global/agent/intents/04cb6e8b-4680-4f22-ad2d-58cbe55f7100) ...
2021/07/04 12:44:34 training phrases found: 68
2021/07/04 12:44:34 written to coronavirus.symptoms_en.csv
...
2021/07/04 12:44:34 written to coronavirus.confirmed_cases_en.csv
2021/07/04 12:44:34 getting 'support.thank' (projects/ghctest003/locations/global/agent/intents/f8b85fb5-c2f7-4024-b1c0-9a6dd03b524a) ...
2021/07/04 12:44:34 training phrases found: 14
2021/07/04 12:44:34 written to support.thank_en.csv
2021/07/04 12:44:34 Intents 69, Training phrases 1277
2021/07/04 12:44:34 export complete

$ # list the csv files and sizes
$ find . -printf '%s\t%f\n' | grep csv | sort -n | tail
1371    coronavirus.stock_up_en.csv
1383    coronavirus.outdoor_activities_en.csv
1649    coronavirus.spread_en.csv
1666    coronavirus.treatment_en.csv
1682    coronavirus.death_en.csv
2002    coronavirus.testing_en.csv
2434    coronavirus.protect_en.csv
2627    coronavirus.symptoms_en.csv
4203    coronavirus.confirmed_cases_en.csv
4339    support.sick_en.csv

$ # Examine one of the Intent Training phrase export files
$ head coronavirus.symptoms_en.csv 
language code,training phrase
en,What about losing sense of taste
en,Is losing sense of smell a symptom of Covid?
en,is red eyes a symptom
en,what are the symptoms?
en,is chest pain a symptom
en,What are the signs of someone who caught Coronavirus?
en,And what are the symptoms?
en,what happens if you get it
en,what happens if you get the virus
```