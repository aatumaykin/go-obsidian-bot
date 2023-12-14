# go-obsidian-bot

Бот для сохранения отправленных в телеграм сообщений в Obsidian

## Конфигурирование

```yaml
# config.yaml

telegram:
  bot_token: "токен бота"
  user: 123             # ID пользователя телеграм
  need_remove: false    # нужно ли удалять сообщение после обработки
  need_reply: true      # нужно ли отвечать сообщением для подтверждения обработки сообщения

obsidian:
  root: "/Obsidian"     # путь к каталогу Obsidian
  note_path: "Inbox"    # название каталога внутри Obsidian, в который нужно сохранять сообщения
```

## Build

```shell
make build
```

либо `go build` с нужными параметрами

## Run

```shell
obsidian-bot 
```

## Пример service-файла

```
# /lib/systemd/system/obsidian-bot.service
[Unit]
Description=Obsidian Bot
After=syslog.target
After=network.target

[Service]
RestartSec=2s
Type=simple
User=root
Group=root
WorkingDirectory=/etc/obsidian-bot
ExecStart=/usr/local/bin/obsidian-bot run --config /etc/obsidian-bot/config.yaml
Restart=always

[Install]
WantedBy=multi-user.target
```