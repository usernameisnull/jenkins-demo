#curl -v -H 'X-GitLab-Token: 1234567' -H 'X-Gitlab-Event: Push Hook' -H 'Content-Type: application/json' --data-binary "@pushevents.json" http://10.6.233.2:30168


curl -v  -H 'X-GitLab-Token: 1234567' -H 'X-Gitlab-Event: Push Hook' -H 'Content-Type: application/json' --data-binary "@pushevents.json" http://10.6.223.2:30168
