# discord-webhooks
package for sending messages &amp; embeds with a discord webhook

```embed := webhook.NewEmbed()
embed.Title = "This is an Embed!"
embed.TitleURL = "https://github.com/Squeakalus/discord-webhooks"
embed.Description = "This is an Embeds Description!"
embed.SetTimestamp()
embed.SetColour("0000FF")
embed.SetThumbnail("Thumbnail URL")
embed.SetAuthor("Author", "Author URL", "Author Thumbnail")
embed.SetFooter("Footer", "Footer Image URL")
embed.AddField("Field Title One", "Field Value One", false)
embed.AddField("Field Title Two", "Field Value Two", false)

err := embed.Send(
  "https://discordapp.com/api/webhooks/XXX/XXX",
  "Text message outside of Embed",
  "Webhook Profile Title",
  "Webhook Profile Image URL")

if err != nil {
  fmt.Println("err")
} else {
 fmt.Println("Sent Embed!")
}
```
