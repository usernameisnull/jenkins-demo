curl -v  -H 'X-GitLab-Token: 1234567' -H 'X-Gitlab-Event: Merge Request Hook' -H 'Content-Type: application/json' --data-binary "@mergeevents.json" http://10.6.223.2:30661
