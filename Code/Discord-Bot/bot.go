package bot

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/akhil/discord-ping/config"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHendler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")

}

func messageHendler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "$trein" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hier word de trein informatie weergegeven")
	}

	if m.Content == "$fontys" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "opgeven dankzij fontys:, https://fontys.nl/Studeren/Uitschrijven.htm")
	}

	if m.Content == "$pjotr" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "De otter is niet meer te vinden in zijn natuurlijke habitat. Hij lijkt van de aardbodem te zijn verdwenen....")
	}
}

func connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Thomas:<welkom2022>@cluster0.sxdwxub.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}
