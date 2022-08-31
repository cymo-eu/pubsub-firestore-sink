# pubsub-firestore-sink
Containerised application that can sink data from pubSub into Firestore

## Configuration
Takes environment variables to configure source and desination
### Environment Variables
| Variable | Description    |
| :---:   | :---: | 
| PROJECT_ID | GCP Project ID where the subscription resides   |
| FIRESTORE_PROJECT_ID | GCP project ID where the Firestore instance resides | 
| SUBSCRIPTION | Name of the subscription that will be used for pulling the records from pubsub. Assumes it has been created in advance |
| FIRESTORE_COLLECTION | Name of the filestore collection where the records will be stored |
### Credentials
Initially created for deployment on a GCP runtime, takes the platform default credentials or Workload Identity credentials.
Might add functionality later to also take explicitly configured credentials.
