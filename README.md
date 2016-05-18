# hippie

A simple, but extensible webhooks-based [HipChat][hipchat] bot.


## Usage

1. Download, and install `hippie`:

    ```shell
    go get -u github.com/claudebot/hippie
    ```

2. Initiate the `hippie`:

    ```shell
    hippie -token=<hipchat token> -room=<room name> -url=<url to hosted bot>
    ```

Use [`ngrok`][ngrok] if you want to host the bot locally without the hassle of port forwarding. You can then use the `ngrok` URL for the bot (the default port is `8080`).


## Scripting

By default, it is shipped with a few basic scripts; including support for the infamous [Giphy][giphy]. See [`scripts/`](scripts/) for more information.

---

The bot's design was inspired by the following packages / projects:

- https://golang.org/pkg/database/sql/
- https://github.com/apex/apex/
- https://github.com/poptip/buster

`hippie` HipChat ChatOps implementation is arguably more effective, and reliable than that of [`lucille`][lucille], and its [Golang][golang] counterpart, [`buster`][buster]. It relies on webhooks rather than interval-based long-polling of the chat history endpoint. As such, it is able to respond to messages a lot quicker without expending unnecessary resources.


[hipchat]: https://www.hipchat.com/
[giphy]: http://giphy.com/
[ngrok]: https://ngrok.com/
[golang]: https://golang.org/
[lucille]: https://github.com/rhaining/lucille
[buster]: https://github.com/poptip/buster
