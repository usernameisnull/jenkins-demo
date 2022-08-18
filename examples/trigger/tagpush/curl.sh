curl -v  -H 'X-GitLab-Token: 1234567' -H 'X-Gitlab-Event: Tag Push Hook' -H 'Content-Type: application/json' --data-binary "@tag-pushevents.json" http://10.6.223.2:32718
