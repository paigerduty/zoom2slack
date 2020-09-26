# zoom2slack

## Purpose
To automatically post zoom meeting recordings to a designated Slack channel

## Plan
1. ~~Set up cron to log something & prove out~~
1. ~~Add OTel launcher~~
1. Make LS_ACCESS_TOKEN optional param
1. Change "invocations" attr -> counter metric
1. Add workflow links [zoom status, slack status]
1. Split out configuration into sep file/env vars
1. Why is span getting named "start cron"?
1. Zoom API - authenticate
1. Zoom API - configure meeting to listen for
1. Zoom API - fetch meeting recording
1. Slack API - authenticate
1. Slack API - how to post to channel
1. Confirm dev tier details
1. Write graceful shutdown to log if in progress jobs were happening
1. [stretch] periodically GCal check cal invite in case it moved, use that date+1 day to set the cron

# Prereqs
* Get a Lightstep Access Token. [Sign up for free dev tier](https://app.lightstep.com/signup?signup_source=footer)
* `export LS_ACCESS_TOKEN="XXXXXX"`

# Relevant Docs
* [cron lib godoc](https://godoc.org/github.com/robfig/cron)
* [OTel launcher](https://github.com/lightstep/otel-launcher-go)
* [OTel tracer](LINK)
