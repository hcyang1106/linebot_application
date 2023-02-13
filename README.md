# CINNOX interview practice

# Introduction
Build an application that can
1. Save the messages that line users sent to the bot
2. Let the bot send messages to line users
3. Get the history messages the user has sent to the bot   


#  How to run
**Required**
- Docker

The container contains MongoDB

**Make**
- Makefile

The makefile helps you create and run the container with MongoDB

**Config**
- config.yml

You should modify the config.yml first

# Existing API
- Post /history
Save the user message sent to linebot
- Get /history
Get the history messages of a specific user
- Post /message
Send a message from a linebot to a specific user
