// go telegram bot to check remote server 
// to install telebot, run: go get github.com/tucnak/telebot

package main 

import (
    "time"
    "github.com/tucnak/telebot"
    "log"
    "os/exec"
)

func cmd(command string)string{
    var (
        cmdOut []byte
        err     error
    )
    cmdOut, err = exec.Command(command).Output(); if err != nil{
        log.Fatal(err)
    }
    return string(cmdOut)
}


func main() {
    bot, err := telebot.NewBot(botToken)
    if err != nil {
        return
    }
    messages := make(chan telebot.Message)
    bot.Listen(messages, 1*time.Second)
    for message := range messages {
        if message.Text == "stat" {
            uptime := cmd("uptime")
            bot.SendMessage(message.Chat, uptime, nil)
        }
        if message.Text == "dir"{
            ls := cmd("ls")
            bot.SendMessage(message.Chat, ls, nil)
        }
    }
}