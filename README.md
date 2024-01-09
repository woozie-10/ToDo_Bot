# ToDo Bot

## Overview

The **ToDo Bot** project is a Telegram bot written in Go that helps users manage their to-do lists. This bot allows users to add, retrieve, update, and delete tasks using various commands. It leverages the [Telegram Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api) for interacting with Telegram and a MySQL database for task management.

## Features

- Users can interact with the bot to perform the following tasks:
    - Retrieve a list of existing tasks.
    - Get a task by ID.
    - Add a new task to the list.
    - Update an existing task by ID.
    - Delete a task by ID.
    - Clear all tasks from the list.
- The bot communicates with a MySQL database to store and manage tasks.

## Getting Started

To run the ToDo Bot project, follow these steps:

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/woozie-10/ToDo_Bot.git
2. Get your token using [BotFather](https://t.me/botfather) and paste it into the config/config.yml on the second line.
 
3. Launch the bot using the command:

```
make build && make run
```

## Usage

To use the ToDo Bot, follow these steps:

1. Start a chat with the bot on Telegram.

2. Use the following commands to interact with the bot:


    /getTasks: Retrieve a list of existing tasks.

    /getTask <task_id>: Get a task by its ID

    /addTask <task_description>: Add a new task to the list.

    /clearTasks: Clear all tasks from the list.

    /delTask <task_id>: Delete a task by its ID.

    /updTask <task_id> <new_task_description>: Update an existing task by its ID.
The bot will respond with the appropriate messages based on the command and perform the corresponding task management operations.

## Dependencies

- [Telegram Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api): Used for interacting with Telegram and sending messages.
- [MySQL](https://www.mysql.com/): Used to store and manage task data.
