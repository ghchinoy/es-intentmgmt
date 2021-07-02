// Copyright 2010 Google LLC
// Author: ghchinoy
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

var (
	action       string
	agentID      string
	location     string
	projectID    string
	languagecode string
)

func init() {
	flag.StringVar(&action, "action", "export", "import | export")
	flag.StringVar(&agentID, "agent", "", "ES Agent ID")
	flag.StringVar(&location, "location", "global", "ES Agent location")
	flag.StringVar(&projectID, "project", "", "GCP Project ID")
	flag.StringVar(&languagecode, "language", "en", "language code (for multilingual Agents)")
	flag.Parse()
}

func main() {
	log.Println("es intent management")

	if action == "" {
		log.Println("action flag required")
		os.Exit(1)
	}

	switch action {
	case "import":
		log.Println("importing intent")
	case "export":
		log.Println("exporting all intents")
		err := exportIntents()
		if err != nil {
			log.Printf("unable to connect to ES: %v", err)
			os.Exit(1)
		}
	default:
		log.Println("action flag must be either 'export' or 'import'")
		os.Exit(1)
	}
}

// exportIntents exports the Intents from an Agent into individual CSV files
func exportIntents() error {
	ctx := context.Background()
	var apiEndpoint string
	if location == "global" {
		apiEndpoint = "dialogflow.googleapis.com:443"
	} else {
		apiEndpoint = fmt.Sprintf("%s-dialogflow.googleapis.com:443", location)
	}
	c, err := dialogflow.NewIntentsClient(ctx, option.WithEndpoint(apiEndpoint))
	if err != nil {
		return err
	}
	defer c.Close()

	parent := fmt.Sprintf("projects/%s/locations/%s/agent", projectID, location)
	req := &dialogflowpb.ListIntentsRequest{
		Parent:       parent,
		LanguageCode: languagecode,
		IntentView:   1,
	}
	it := c.ListIntents(ctx, req)
	for {
		intent, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		if !strings.HasPrefix(intent.DisplayName, "Knowledge.") {
			log.Printf("Getting '%s' (%s) ...", intent.DisplayName, intent.Name)
			tps := intent.GetTrainingPhrases()
			if len(tps) > 0 {
				log.Printf("Training Phrases: %d", len(tps))

				var records [][]string
				records = append(records, []string{"language code", "training phrase"})

				for _, tp := range tps {
					var phraseparts []string
					for _, s := range tp.Parts {
						phraseparts = append(phraseparts, s.Text)
					}
					phrase := strings.Join(phraseparts, "")
					records = append(records, []string{languagecode, phrase})
				}

				filepath := fmt.Sprintf("%s_%s.csv", strings.Replace(intent.DisplayName, " ", "_", -1), languagecode)
				destination, err := os.Create(filepath)
				if err != nil {
					fmt.Println("os.Create:", err)
					return err
				}
				defer destination.Close()

				w := csv.NewWriter(destination)
				w.WriteAll(records)
				if err := w.Error(); err != nil {
					log.Println("error writing csv:", err)
				}
			} else {
				log.Printf("No training phrases for '%s'", intent.DisplayName)
			}
		}

		/*
			jsonbytes, _ := json.MarshalIndent(&intent, "", "  ")
			fmt.Printf("%s\n", jsonbytes)
		*/

		/*
			filepath, count, err := getIntent(ctx, intent.DisplayName, intent.Name)
			if err != nil {
				log.Printf("Unable to retrieve intent: %v", err)
			} else {
				log.Printf("Wrote %d intents to %s.", count, filepath)
			}
		*/

	}

	return nil
}
