deploy:
	gcloud functions deploy "jess-chat" \
		--runtime go116 \
		--project "jessapi" \
		--allow-unauthenticated \
		--trigger-http \
		--entry-point="Say" \
		--service-account="jessapi@jessapi.iam.gserviceaccount.com"