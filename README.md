# Telegram bot with ChatGPT

This is a Telegram bot that uses ChatGPT to generate responses to messages.

## How to use

1. Create a Telegram bot using [@BotFather](https://t.me/BotFather)
2. Apply for a OpenAI API key [here](https://platform.openai.com/)
3. Run the bot (preferably as a Docker container) with the following environment variables:
- `OPEANAI_APIKEY`: Your OpenAI API key
- `TELEGRAM_TOKEN`: Your Telegram bot token

## How to run as a Docker container

Here's an Docker compose example:

```yaml
version: '3.5'
services:
  chatgptbot:
    image: h00s/chatgpt-telegram-bot
    container_name: chatgptbot
    stop_grace_period: 15s
    restart: unless-stopped
    environment:
      - TELEGRAM_TOKEN=123456:abc
      - CHATGPT_APIKEY=sk-123456
```

## Commands and usage

- `/start`: Start the bot
- `/hello`: Test the bot if it's working
- `/new`: Start a new conversation (chatgpt will forget the previous conversation)

Any message sent to the bot will be sent to chatgpt and the response will be sent back to the user.